package gate

import (
	"git.golaxy.org/core/define"
)

var (
	self      = define.DefineServicePlugin(newGate)
	Name      = self.Name
	Using     = self.Using
	Install   = self.Install
	Uninstall = self.Uninstall
)