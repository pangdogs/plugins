package transport

import (
	"git.golaxy.org/framework/net/gtp"
)

// Event 消息事件
type Event[T gtp.MsgReader] struct {
	Flags gtp.Flags // 标志位
	Seq   uint32    // 消息序号
	Ack   uint32    // 应答序号
	Msg   T         // 消息
}

// Interface 泛化事件，转换为事件通用类型
func (e Event[T]) Interface() Event[gtp.MsgReader] {
	return Event[gtp.MsgReader]{
		Flags: e.Flags,
		Seq:   e.Seq,
		Ack:   e.Ack,
		Msg:   e.Msg,
	}
}

// EventT 特化事件，转换为事件具体类型
func EventT[T gtp.MsgReader](e Event[gtp.MsgReader]) Event[T] {
	ret := Event[T]{
		Flags: e.Flags,
		Seq:   e.Seq,
		Ack:   e.Ack,
	}

	if e.Msg == nil {
		return ret
	}

	msgPtr, ok := any(e.Msg).(*T)
	if ok {
		ret.Msg = *msgPtr
		return ret
	}

	msg, ok := any(e.Msg).(T)
	if ok {
		ret.Msg = msg
		return ret
	}

	panic("gtp: incorrect msg type")
}
