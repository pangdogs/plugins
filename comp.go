/*
 * This file is part of Golaxy Distributed Service Development Framework.
 *
 * Golaxy Distributed Service Development Framework is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 2.1 of the License, or
 * (at your option) any later version.
 *
 * Golaxy Distributed Service Development Framework is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with Golaxy Distributed Service Development Framework. If not, see <http://www.gnu.org/licenses/>.
 *
 * Copyright (c) 2024 pangdogs.
 */

package framework

import (
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/plugin"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/reinterpret"
)

// ComponentBehavior 组件行为，在开发新组件时，匿名嵌入至组件结构体中
type ComponentBehavior struct {
	ec.ComponentBehavior
}

// GetRuntime 获取运行时
func (c *ComponentBehavior) GetRuntime() IRuntimeInstance {
	return reinterpret.Cast[IRuntimeInstance](runtime.Current(c))
}

// GetService 获取服务
func (c *ComponentBehavior) GetService() IServiceInstance {
	return reinterpret.Cast[IServiceInstance](service.Current(c))
}

// GetPluginBundle 获取插件包
func (c *ComponentBehavior) GetPluginBundle() plugin.PluginBundle {
	return runtime.Current(c).GetPluginBundle()
}

// IsAlive 是否活跃
func (c *ComponentBehavior) IsAlive() bool {
	return c.GetState() <= ec.ComponentState_Alive
}
