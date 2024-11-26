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

package rpcpcsr

import (
	"fmt"
	"git.golaxy.org/core"
	"git.golaxy.org/core/ec"
	"git.golaxy.org/core/extension"
	"git.golaxy.org/core/runtime"
	"git.golaxy.org/core/service"
	"git.golaxy.org/core/utils/async"
	"git.golaxy.org/core/utils/types"
	"git.golaxy.org/core/utils/uid"
	"git.golaxy.org/framework/net/gap/variant"
	"git.golaxy.org/framework/plugins/rpcstack"
	"reflect"
)

func CallService(svcCtx service.Context, cc rpcstack.CallChain, pluginName, method string, args variant.Array) (rets variant.Array, err error) {
	defer func() {
		if panicErr := types.Panic2Err(recover()); panicErr != nil {
			err = fmt.Errorf("%w: %w", core.ErrPanicked, panicErr)
		}
	}()

	var scriptRV reflect.Value

	if pluginName == "" {
		scriptRV = service.UnsafeContext(svcCtx).GetReflected()
	} else {
		ps, ok := svcCtx.GetPluginBundle().Get(pluginName)
		if !ok {
			return nil, ErrPluginNotFound
		}

		if ps.State() != extension.PluginState_Active {
			return nil, ErrPluginInactive
		}

		scriptRV = ps.Reflected()
	}

	methodRV := scriptRV.MethodByName(method)
	if !methodRV.IsValid() {
		return nil, ErrMethodNotFound
	}

	argsRV, err := parseArgs(methodRV, cc, args)
	if err != nil {
		return nil, err
	}

	return variant.MakeSerializedArray(methodRV.Call(argsRV))
}

func CallRuntime(svcCtx service.Context, cc rpcstack.CallChain, entityId uid.Id, pluginName, method string, args variant.Array) (asyncRet async.AsyncRet, err error) {
	defer func() {
		if panicErr := types.Panic2Err(recover()); panicErr != nil {
			err = fmt.Errorf("%w: %w", core.ErrPanicked, panicErr)
		}
	}()

	return svcCtx.Call(entityId, func(entity ec.Entity, _ ...any) async.Ret {
		var scriptRV reflect.Value

		if pluginName == "" {
			scriptRV = runtime.UnsafeContext(runtime.Current(entity)).GetReflected()
		} else {
			ps, ok := runtime.Current(entity).GetPluginBundle().Get(pluginName)
			if !ok {
				return async.MakeRet(nil, ErrPluginNotFound)
			}

			if ps.State() != extension.PluginState_Active {
				return async.MakeRet(nil, ErrPluginInactive)
			}

			scriptRV = ps.Reflected()
		}

		methodRV := scriptRV.MethodByName(method)
		if !methodRV.IsValid() {
			return async.MakeRet(nil, ErrMethodNotFound)
		}

		argsRV, err := parseArgs(methodRV, cc, args)
		if err != nil {
			return async.MakeRet(nil, err)
		}

		stack := rpcstack.Using(runtime.Current(entity))
		rpcstack.UnsafeRPCStack(stack).PushCallChain(cc)
		defer rpcstack.UnsafeRPCStack(stack).PopCallChain()

		return async.MakeRet(variant.MakeSerializedArray(methodRV.Call(argsRV)))
	}), nil
}

func CallEntity(svcCtx service.Context, cc rpcstack.CallChain, entityId uid.Id, component, method string, args variant.Array) (asyncRet async.AsyncRet, err error) {
	defer func() {
		if panicErr := types.Panic2Err(recover()); panicErr != nil {
			err = fmt.Errorf("%w: %w", core.ErrPanicked, panicErr)
		}
	}()

	return svcCtx.Call(entityId, func(entity ec.Entity, _ ...any) async.Ret {
		var scriptRV reflect.Value

		if component == "" {
			scriptRV = entity.GetReflected()
		} else {
			comp := entity.GetComponent(component)
			if comp == nil {
				return async.MakeRet(nil, ErrComponentNotFound)
			}
			scriptRV = comp.GetReflected()
		}

		methodRV := scriptRV.MethodByName(method)
		if !methodRV.IsValid() {
			return async.MakeRet(nil, ErrMethodNotFound)
		}

		argsRV, err := parseArgs(methodRV, cc, args)
		if err != nil {
			return async.MakeRet(nil, err)
		}

		stack := rpcstack.Using(runtime.Current(entity))
		rpcstack.UnsafeRPCStack(stack).PushCallChain(cc)
		defer rpcstack.UnsafeRPCStack(stack).PopCallChain()

		return async.MakeRet(variant.MakeSerializedArray(methodRV.Call(argsRV)))
	}), nil
}

var (
	callChainRT = reflect.TypeFor[rpcstack.CallChain]()
)

func parseArgs(methodRV reflect.Value, cc rpcstack.CallChain, args variant.Array) ([]reflect.Value, error) {
	methodRT := methodRV.Type()
	var argsRV []reflect.Value
	var argsPos int

	switch methodRT.NumIn() {
	case len(args) + 1:
		if !callChainRT.AssignableTo(methodRT.In(0)) {
			return nil, ErrMethodParameterTypeMismatch
		}
		argsRV = append(make([]reflect.Value, 0, len(args)+1), reflect.ValueOf(cc))
		argsPos = 1

	case len(args):
		argsRV = make([]reflect.Value, 0, len(args))
		argsPos = 0

	default:
		return nil, ErrMethodParameterCountMismatch
	}

	for i := range args {
		argRV, err := args[i].Convert(methodRT.In(argsPos + i))
		if err != nil {
			return nil, ErrMethodParameterTypeMismatch
		}
		argsRV = append(argsRV, argRV)
	}

	return argsRV, nil
}
