package distributed

import (
	"golang.org/x/net/context"
	"kit.golaxy.org/plugins/broker"
	"kit.golaxy.org/plugins/log"
	"time"
)

func (d *_Distributed) run() {
	defer d.wg.Done()

	log.Infof(d.ctx, "start service %q node %q", d.ctx.GetName(), d.ctx.GetId())

	ticker := time.NewTicker(d.Options.RefreshInterval)
	defer ticker.Stop()

loop:
	for {
		select {
		case <-ticker.C:
			// 刷新服务节点
			err := d.registry.Register(d.ctx, d.service, d.Options.RefreshInterval*2)
			if err != nil {
				log.Errorf(d.ctx, "refresh service %q node %q failed, %s", d.ctx.GetName(), d.ctx.GetId(), err)
				continue
			}

			log.Infof(d.ctx, "refresh service %q node %q success", d.ctx.GetName(), d.ctx.GetId())

		case <-d.ctx.Done():
			break loop
		}
	}

	// 取消注册服务节点
	err := d.registry.Deregister(context.Background(), d.service)
	if err != nil {
		log.Errorf(d.ctx, "deregister service %q node %q failed, %s", d.ctx.GetName(), d.ctx.GetId(), err)
	}

	// 取消订阅topic
	for _, sub := range d.subs {
		<-sub.Unsubscribe()
	}

	log.Infof(d.ctx, "stop service %q node %q", d.ctx.GetName(), d.ctx.GetId())
}

// handleEvent 处理事件
func (d *_Distributed) handleEvent(e broker.Event) error {
	e.Message()

	e.Topic()

	return nil
}