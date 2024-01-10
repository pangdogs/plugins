package registry

import (
	"kit.golaxy.org/golaxy/define"
)

var (
	plugin = define.DefineServicePluginInterface[IRegistry]()
	Name   = plugin.Name
	Using  = plugin.Using
)
