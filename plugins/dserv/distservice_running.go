package dserv

import (
	"context"
	"errors"
	"fmt"
	"git.golaxy.org/framework/plugins/broker"
	"git.golaxy.org/framework/plugins/discovery"
	"git.golaxy.org/framework/plugins/log"
	"time"
)

func (d *_DistService) mainLoop(serviceNode *discovery.Service, subs []broker.ISubscriber) {
	defer d.wg.Done()

	log.Infof(d.servCtx, "service %q node %q started", d.servCtx.GetName(), d.servCtx.GetId())

	if d.options.RefreshTTL {
		ticker := time.NewTicker(d.options.TTL / 2)
		defer ticker.Stop()

	loop:
		for {
			select {
			case <-ticker.C:
				// 刷新服务节点
				if err := d.registry.Register(d.ctx, serviceNode, d.options.TTL); err != nil {
					log.Errorf(d.servCtx, "refresh service %q node %q failed, %s", d.servCtx.GetName(), d.servCtx.GetId(), err)
					continue
				}

				log.Debugf(d.servCtx, "refresh service %q node %q success", d.servCtx.GetName(), d.servCtx.GetId())

			case <-d.ctx.Done():
				break loop
			}
		}
	} else {
		<-d.ctx.Done()
	}

	// 取消注册服务节点
	if err := d.registry.Deregister(context.Background(), serviceNode); err != nil {
		log.Errorf(d.servCtx, "deregister service %q node %q failed, %s", d.servCtx.GetName(), d.servCtx.GetId(), err)
	}

	// 取消订阅topic
	for _, sub := range subs {
		<-sub.Unsubscribe()
	}

	d.broker.Flush(context.Background())

	log.Infof(d.servCtx, "service %q node %q stopped", d.servCtx.GetName(), d.servCtx.GetId())
}

func (d *_DistService) watchingService() {
	defer d.wg.Done()

	log.Debug(d.servCtx, "watching service changes started")

retry:
	var watcher discovery.IWatcher
	var err error
	retryInterval := 3 * time.Second

	select {
	case <-d.ctx.Done():
		goto end
	default:
	}

	// 监控服务节点变化
	watcher, err = d.registry.Watch(d.ctx, "")
	if err != nil {
		log.Errorf(d.servCtx, "watching service changes failed, %s, retry it", err)
		time.Sleep(retryInterval)
		goto retry
	}

	for {
		e, err := watcher.Next()
		if err != nil {
			if errors.Is(err, discovery.ErrTerminated) {
				time.Sleep(retryInterval)
				goto retry
			}

			log.Errorf(d.servCtx, "watching service changes failed, %s, retry it", err)
			<-watcher.Terminate()
			time.Sleep(retryInterval)
			goto retry
		}

		switch e.Type {
		case discovery.Delete:
			for _, node := range e.Service.Nodes {
				d.deduplication.Remove(node.Address)
			}
		}
	}

end:
	if watcher != nil {
		<-watcher.Terminate()
	}

	log.Debug(d.servCtx, "watching service changes stopped")
}

func (d *_DistService) handleEvent(e broker.IEvent) error {
	mp, err := d.decoder.DecodeBytes(e.Message())
	if err != nil {
		return err
	}

	// 最少一次交付模式，需要消息去重
	if d.broker.GetDeliveryReliability() == broker.AtLeastOnce {
		if !d.deduplication.Validate(mp.Head.Src, mp.Head.Seq) {
			return fmt.Errorf("gap: discard duplicate msg-packet, head:%+v", mp.Head)
		}
	}

	var errs []error

	interrupt := func(err, _ error) bool {
		if err != nil {
			errs = append(errs, err)
		}
		return false
	}

	// 回调监控器
	d.msgWatchers.AutoRLock(func(watchers *[]*_MsgWatcher) {
		for i := range *watchers {
			(*watchers)[i].handler.Exec(interrupt, e.Topic(), mp)
		}
	})

	// 回调处理器
	d.options.RecvMsgHandler.Exec(interrupt, e.Topic(), mp)

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}
