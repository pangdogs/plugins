// Code generated by 'yaegi extract git.golaxy.org/framework/plugins/discovery/cache_discovery'. DO NOT EDIT.

package fwlib

import (
	"git.golaxy.org/framework/plugins/discovery/cache_discovery"
	"reflect"
)

func init() {
	Symbols["git.golaxy.org/framework/plugins/discovery/cache_discovery/cache_discovery"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Install":   reflect.ValueOf(&cache_discovery.Install).Elem(),
		"Uninstall": reflect.ValueOf(&cache_discovery.Uninstall).Elem(),
		"With":      reflect.ValueOf(&cache_discovery.With).Elem(),

		// type definitions
		"RegistryOptions": reflect.ValueOf((*cache_discovery.RegistryOptions)(nil)),
	}
}
