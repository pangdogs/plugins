package gtp

import (
	"errors"
	"fmt"
	"git.golaxy.org/core"
	"git.golaxy.org/framework/util/concurrent"
	"reflect"
)

var (
	ErrNotRegistered = errors.New("gtp: msg not registered") // 消息未注册
)

// IMsgCreator 消息对象构建器接口
type IMsgCreator interface {
	// Declare 注册消息
	Declare(msg Msg)
	// Undeclare 取消注册消息
	Undeclare(msgId MsgId)
	// New 创建消息指针
	New(msgId MsgId) (Msg, error)
}

var msgCreator = NewMsgCreator()

// DefaultMsgCreator 默认消息对象构建器
func DefaultMsgCreator() IMsgCreator {
	return msgCreator
}

func init() {
	DefaultMsgCreator().Declare(&MsgHello{})
	DefaultMsgCreator().Declare(&MsgECDHESecretKeyExchange{})
	DefaultMsgCreator().Declare(&MsgChangeCipherSpec{})
	DefaultMsgCreator().Declare(&MsgAuth{})
	DefaultMsgCreator().Declare(&MsgContinue{})
	DefaultMsgCreator().Declare(&MsgFinished{})
	DefaultMsgCreator().Declare(&MsgRst{})
	DefaultMsgCreator().Declare(&MsgHeartbeat{})
	DefaultMsgCreator().Declare(&MsgSyncTime{})
	DefaultMsgCreator().Declare(&MsgPayload{})
}

// NewMsgCreator 创建消息对象构建器
func NewMsgCreator() IMsgCreator {
	return &_MsgCreator{
		msgTypeMap: concurrent.MakeLockedMap[MsgId, reflect.Type](0),
	}
}

// _MsgCreator 消息对象构建器
type _MsgCreator struct {
	msgTypeMap concurrent.LockedMap[MsgId, reflect.Type]
}

// Declare 注册消息
func (c *_MsgCreator) Declare(msg Msg) {
	if msg == nil {
		panic(fmt.Errorf("%w: msg is nil", core.ErrArgs))
	}

	c.msgTypeMap.Insert(msg.MsgId(), reflect.TypeOf(msg).Elem())
}

// Undeclare 取消注册消息
func (c *_MsgCreator) Undeclare(msgId MsgId) {
	c.msgTypeMap.Delete(msgId)
}

// New 创建消息指针
func (c *_MsgCreator) New(msgId MsgId) (Msg, error) {
	rtype, ok := c.msgTypeMap.Get(msgId)
	if !ok {
		return nil, ErrNotRegistered
	}
	return reflect.New(rtype).Interface().(Msg), nil
}
