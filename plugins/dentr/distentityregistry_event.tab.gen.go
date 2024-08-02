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

// Code generated by eventc eventtab --name=distEntityRegistryEventTab; DO NOT EDIT.

package dentr

import (
	event "git.golaxy.org/core/event"
)

type IDistEntityRegistryEventTab interface {
	EventDistEntityOnline() event.IEvent
	EventDistEntityOffline() event.IEvent
}

var (
	_distEntityRegistryEventTabId = event.DeclareEventTabIdT[distEntityRegistryEventTab]()
	EventDistEntityOnlineId = _distEntityRegistryEventTabId + 0
	EventDistEntityOfflineId = _distEntityRegistryEventTabId + 1
)

type distEntityRegistryEventTab [2]event.Event

func (eventTab *distEntityRegistryEventTab) Init(autoRecover bool, reportError chan error, recursion event.EventRecursion) {
	(*eventTab)[0].Init(autoRecover, reportError, recursion)
	(*eventTab)[1].Init(autoRecover, reportError, recursion)
}

func (eventTab *distEntityRegistryEventTab) Get(id uint64) event.IEvent {
	if _distEntityRegistryEventTabId != id & 0xFFFFFFFF00000000 {
		return nil
	}
	pos := id & 0xFFFFFFFF
	if pos >= uint64(len(*eventTab)) {
		return nil
	}
	return &(*eventTab)[pos]
}

func (eventTab *distEntityRegistryEventTab) Open() {
	for i := range *eventTab {
		(*eventTab)[i].Open()
	}
}

func (eventTab *distEntityRegistryEventTab) Close() {
	for i := range *eventTab {
		(*eventTab)[i].Close()
	}
}

func (eventTab *distEntityRegistryEventTab) Clean() {
	for i := range *eventTab {
		(*eventTab)[i].Clean()
	}
}

func (eventTab *distEntityRegistryEventTab) EventDistEntityOnline() event.IEvent {
	return &(*eventTab)[0]
}

func (eventTab *distEntityRegistryEventTab) EventDistEntityOffline() event.IEvent {
	return &(*eventTab)[1]
}
