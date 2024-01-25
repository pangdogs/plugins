package redis_discovery

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"git.golaxy.org/core"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/util/option"
	"git.golaxy.org/core/util/types"
	"git.golaxy.org/framework/plugins/discovery"
	"git.golaxy.org/framework/plugins/log"
	"git.golaxy.org/framework/plugins/util/concurrent"
	hash "github.com/mitchellh/hashstructure/v2"
	"github.com/redis/go-redis/v9"
	"sort"
	"strings"
	"time"
)

// NewRegistry 创建registry插件，可以配合cache registry将数据缓存本地，提高查询效率
func NewRegistry(settings ...option.Setting[RegistryOptions]) discovery.IRegistry {
	return &_Registry{
		options:   option.Make(Option{}.Default(), settings...),
		registers: concurrent.MakeLockedMap[string, uint64](0),
	}
}

type _Registry struct {
	options   RegistryOptions
	servCtx   service.Context
	client    *redis.Client
	registers concurrent.LockedMap[string, uint64]
}

// InitSP 初始化服务插件
func (r *_Registry) InitSP(ctx service.Context) {
	log.Infof(ctx, "init plugin <%s>:[%s]", plugin.Name, types.AnyFullName(*r))

	r.servCtx = ctx

	if r.options.RedisClient == nil {
		r.client = redis.NewClient(r.configure())
	} else {
		r.client = r.options.RedisClient
	}

	_, err := r.client.Ping(r.servCtx).Result()
	if err != nil {
		log.Panicf(r.servCtx, "ping redis %q failed, %v", r.client, err)
	}

	_, err = r.client.ConfigSet(r.servCtx, "notify-keyspace-events", "KEA").Result()
	if err != nil {
		log.Panicf(r.servCtx, "redis %q enable notify-keyspace-events failed, %v", r.client, err)
	}
}

// ShutSP 关闭服务插件
func (r *_Registry) ShutSP(ctx service.Context) {
	log.Infof(ctx, "shut plugin <%s>:[%s]", plugin.Name, types.AnyFullName(*r))

	if r.options.RedisClient == nil {
		if r.client != nil {
			r.client.Close()
		}
	}
}

// Register 注册服务
func (r *_Registry) Register(ctx context.Context, service *discovery.Service, ttl time.Duration) error {
	if ctx == nil {
		ctx = context.Background()
	}

	if service == nil {
		return fmt.Errorf("%w: %w: serivce is nil", discovery.ErrRegistry, core.ErrArgs)
	}

	if len(service.Nodes) <= 0 {
		return fmt.Errorf("%w: require at least one node", discovery.ErrRegistry)
	}

	var errs []error

	for i := range service.Nodes {
		node := &service.Nodes[i]

		if err := r.registerNode(ctx, service, node, ttl); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", node.Id, err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("%w: %w", discovery.ErrRegistry, errors.Join(errs...))
	}

	return nil
}

// Deregister 取消注册服务
func (r *_Registry) Deregister(ctx context.Context, service *discovery.Service) error {
	if ctx == nil {
		ctx = context.Background()
	}

	if service == nil {
		return fmt.Errorf("%w: %w: serivce is nil", discovery.ErrRegistry, core.ErrArgs)
	}

	if len(service.Nodes) <= 0 {
		return fmt.Errorf("%w: require at least one node", discovery.ErrRegistry)
	}

	var errs []error

	for i := range service.Nodes {
		node := &service.Nodes[i]

		if err := r.deregisterNode(ctx, service, node); err != nil {
			errs = append(errs, fmt.Errorf("%s: %w", node.Id, err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("%w: %w", discovery.ErrRegistry, errors.Join(errs...))
	}

	return nil
}

// GetServiceNode 查询服务节点
func (r *_Registry) GetServiceNode(ctx context.Context, serviceName, nodeId string) (*discovery.Service, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if serviceName == "" || nodeId == "" {
		return nil, discovery.ErrNotFound
	}

	nodeVal, err := r.client.Get(ctx, getNodePath(r.options.KeyPrefix, serviceName, nodeId)).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, discovery.ErrNotFound
		}
		return nil, fmt.Errorf("%w: %w", discovery.ErrRegistry, err)
	}

	return decodeService(nodeVal)
}

// GetService 查询服务
func (r *_Registry) GetService(ctx context.Context, serviceName string) ([]discovery.Service, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if serviceName == "" {
		return nil, discovery.ErrNotFound
	}

	nodeKeys, err := r.client.Keys(ctx, getServicePath(r.options.KeyPrefix, serviceName)).Result()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", discovery.ErrRegistry, err)
	}

	if len(nodeKeys) <= 0 {
		return nil, discovery.ErrNotFound
	}

	nodeVals, err := r.client.MGet(ctx, nodeKeys...).Result()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", discovery.ErrRegistry, err)
	}

	serviceMap := map[string]*discovery.Service{}

	for _, v := range nodeVals {
		service, err := decodeService([]byte(v.(string)))
		if err != nil {
			log.Errorf(r.servCtx, "decode service %q failed, %s", v, err)
			continue
		}

		s, ok := serviceMap[service.Version]
		if !ok {
			serviceMap[s.Version] = service
			continue
		}

		s.Nodes = append(s.Nodes, service.Nodes...)
	}

	services := make([]discovery.Service, 0, len(serviceMap))
	for _, service := range serviceMap {
		services = append(services, *service)
	}

	// sort the services
	sort.Slice(services, func(i, j int) bool {
		return services[i].Version < services[j].Version
	})

	return services, nil
}

// ListServices 查询所有服务
func (r *_Registry) ListServices(ctx context.Context) ([]discovery.Service, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	nodeKeys, err := r.client.Keys(ctx, r.options.KeyPrefix+"*").Result()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", discovery.ErrRegistry, err)
	}

	if len(nodeKeys) <= 0 {
		return nil, nil
	}

	nodeVals, err := r.client.MGet(ctx, nodeKeys...).Result()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", discovery.ErrRegistry, err)
	}

	versions := make(map[string]*discovery.Service)

	for _, v := range nodeVals {
		service, err := decodeService([]byte(v.(string)))
		if err != nil {
			log.Errorf(r.servCtx, "decode service %q failed, %s", v, err)
			continue
		}

		version := service.Name + ":" + service.Version

		s, ok := versions[version]
		if !ok {
			versions[version] = service
			continue
		}

		// append to service:version nodes
		s.Nodes = append(s.Nodes, service.Nodes...)
	}

	services := make([]discovery.Service, 0, len(versions))
	for _, service := range versions {
		services = append(services, *service)
	}

	// sort the services
	sort.Slice(services, func(i, j int) bool {
		if services[i].Name == services[j].Name {
			return services[i].Version < services[j].Version
		}
		return services[i].Name < services[j].Name
	})

	return services, nil
}

// Watch 获取服务监听器
func (r *_Registry) Watch(ctx context.Context, pattern string) (discovery.IWatcher, error) {
	return r.newWatcher(ctx, pattern)
}

func (r *_Registry) configure() *redis.Options {
	if r.options.RedisConfig != nil {
		return r.options.RedisConfig
	}

	if r.options.RedisURL != "" {
		conf, err := redis.ParseURL(r.options.RedisURL)
		if err != nil {
			log.Panicf(r.servCtx, "parse redis url %q failed, %s", r.options.RedisURL, err)
		}
		return conf
	}

	conf := &redis.Options{}
	conf.Username = r.options.CustUsername
	conf.Password = r.options.CustPassword
	conf.Addr = r.options.CustAddress
	conf.DB = r.options.CustDB

	return conf
}

func (r *_Registry) registerNode(ctx context.Context, service *discovery.Service, node *discovery.Node, ttl time.Duration) error {
	if service.Name == "" {
		return errors.New("service name can't empty")
	}

	if node.Id == "" {
		return errors.New("service node id can't empty")
	}

	if ttl < 0 {
		ttl = 0
	}

	hv, err := hash.Hash(node, hash.FormatV2, nil)
	if err != nil {
		return err
	}

	nodePath := getNodePath(r.options.KeyPrefix, service.Name, node.Id)
	var keepAlive bool

	if ttl.Seconds() > 0 {
		keepAlive, err = r.client.Expire(ctx, nodePath, ttl).Result()
		if err != nil {
			return err
		}
		log.Debugf(r.servCtx, "renewing existing service %q node %q with ttl %q, result %t", service.Name, node.Id, ttl, keepAlive)
	}

	rhv, ok := r.registers.Get(nodePath)
	if ok && rhv == hv && keepAlive {
		log.Debugf(r.servCtx, "service %q node %q unchanged skipping registration", service.Name, node.Id)
		return nil
	}

	serviceNode := service
	serviceNode.Nodes = []discovery.Node{*node}
	serviceNodeData := encodeService(serviceNode)

	log.Debugf(r.servCtx, "registering service %q node %q content %q with ttl %q", serviceNode.Name, node.Id, serviceNodeData, ttl)

	_, err = r.client.Set(ctx, nodePath, serviceNodeData, ttl).Result()
	if err != nil {
		return err
	}

	r.registers.Insert(nodePath, hv)

	log.Debugf(r.servCtx, "register service %q node %q success", serviceNode.Name, node.Id)

	return nil
}

func (r *_Registry) deregisterNode(ctx context.Context, service *discovery.Service, node *discovery.Node) error {
	log.Debugf(r.servCtx, "deregistering service %q node %q", service.Name, node.Id)

	nodePath := getNodePath(r.options.KeyPrefix, service.Name, node.Id)

	r.registers.Delete(nodePath)

	if _, err := r.client.Del(ctx, nodePath).Result(); err != nil {
		return err
	}

	log.Debugf(r.servCtx, "deregister service %q node %q success", service.Name, node.Id)

	return nil
}

func encodeService(s *discovery.Service) string {
	b, _ := json.Marshal(s)
	return string(b)
}

func decodeService(ds []byte) (*discovery.Service, error) {
	var s *discovery.Service

	if err := json.Unmarshal(ds, &s); err != nil {
		return nil, fmt.Errorf("%w: %w", discovery.ErrRegistry, err)
	}

	return s, nil
}

func getNodePath(prefix, s, id string) string {
	service := strings.ReplaceAll(s, ":", "-")
	node := strings.ReplaceAll(id, ":", "-")
	return fmt.Sprintf("%s%s:%s", prefix, service, node)
}

func getServicePath(prefix, s string) string {
	service := strings.ReplaceAll(s, ":", "-")
	return fmt.Sprintf("%s%s:*", prefix, service)
}
