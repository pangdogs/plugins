package framework

import (
	"git.golaxy.org/core"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/service"
	"github.com/spf13/viper"
	"sync"
)

// IRuntimeInstantiation 运行时实例化接口
type IRuntimeInstantiation interface {
	Instantiation() runtime.Context
}

// RuntimeInstance 运行时实例
type RuntimeInstance struct {
	runtime.ContextBehavior
}

// GetStartupConf 获取启动参数配置
func (instance *RuntimeInstance) GetStartupConf() *viper.Viper {
	v, _ := instance.GetMemKVs().Load("startup.conf")
	if v == nil {
		panic("service memory startup.conf not existed")
	}
	return v.(*viper.Viper)
}

// GetMemKVs 获取服务内存KV数据库
func (instance *RuntimeInstance) GetMemKVs() *sync.Map {
	memKVs, _ := service.Current(instance).Value("mem_kvs").(*sync.Map)
	if memKVs == nil {
		panic("service memory not existed")
	}
	return memKVs
}

// CreateEntity 创建实体
func (instance *RuntimeInstance) CreateEntity() core.EntityCreator {
	return core.CreateEntity(instance)
}