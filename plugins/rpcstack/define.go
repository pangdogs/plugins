package rpcstack

import (
	"git.golaxy.org/core/define"
)

var (
	self      = define.RuntimePlugin(newRPCStack)
	Name      = self.Name
	Using     = self.Using
	Install   = self.Install
	Uninstall = self.Uninstall
)
