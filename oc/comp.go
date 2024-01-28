package oc

import (
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/service"
)

// ComponentBehavior 组件行为，需要在开发新组件时，匿名嵌入至组件结构体中
type ComponentBehavior struct {
	ec.ComponentBehavior
}

// GetRtCtx 获取运行时上下文
func (c *ComponentBehavior) GetRtCtx() runtime.Context {
	return runtime.Current(c)
}

// GetServCtx 获取服务上下文
func (c *ComponentBehavior) GetServCtx() service.Context {
	return service.Current(c)
}
