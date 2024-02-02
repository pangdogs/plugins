package dent

import "git.golaxy.org/core/define"

var (
	self      = define.DefineRuntimePlugin(newDistEntities)
	Name      = self.Name
	Using     = self.Using
	Install   = self.Install
	Uninstall = self.Uninstall
)
