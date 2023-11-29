package gap

import (
	"errors"
	"fmt"
	"kit.golaxy.org/golaxy"
	"reflect"
	"sync"
)

var (
	ErrNotRegistered = errors.New("gap.msg not registered") // 消息未注册
)

// IMsgCreator 消息对象构建器接口
type IMsgCreator interface {
	// Register 注册消息
	Register(msg Msg)
	// Deregister 取消注册消息
	Deregister(msgId MsgId)
	// Spawn 构建消息
	Spawn(msgId MsgId) (Msg, error)
}

var msgCreator = NewMsgCreator()

// DefaultMsgCreator 默认消息对象构建器
func DefaultMsgCreator() IMsgCreator {
	return msgCreator
}

func init() {
	DefaultMsgCreator().Register(&MsgRPCRequest{})
	DefaultMsgCreator().Register(&MsgRPCReply{})
	DefaultMsgCreator().Register(&MsgOneWayRPC{})
}

// NewMsgCreator 创建消息对象构建器
func NewMsgCreator() IMsgCreator {
	return &_MsgCreator{
		msgTypeMap: make(map[MsgId]reflect.Type),
	}
}

// _MsgCreator 消息对象构建器
type _MsgCreator struct {
	sync.RWMutex
	msgTypeMap map[MsgId]reflect.Type
}

// Register 注册消息
func (c *_MsgCreator) Register(msg Msg) {
	c.Lock()
	defer c.Unlock()

	if msg == nil {
		panic(fmt.Errorf("%w: msg is nil", golaxy.ErrArgs))
	}

	c.msgTypeMap[msg.MsgId()] = reflect.TypeOf(msg).Elem()
}

// Deregister 取消注册消息
func (c *_MsgCreator) Deregister(msgId MsgId) {
	c.Lock()
	defer c.Unlock()

	delete(c.msgTypeMap, msgId)
}

// Spawn 构建消息
func (c *_MsgCreator) Spawn(msgId MsgId) (Msg, error) {
	c.RLock()
	defer c.RUnlock()

	rtype, ok := c.msgTypeMap[msgId]
	if !ok {
		return nil, ErrNotRegistered
	}

	return reflect.New(rtype).Interface().(Msg), nil
}