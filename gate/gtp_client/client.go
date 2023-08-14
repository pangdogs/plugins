package gtp_client

import (
	"fmt"
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"kit.golaxy.org/plugins/gate"
	"kit.golaxy.org/plugins/transport"
	"kit.golaxy.org/plugins/transport/protocol"
	"sync"
)

type (
	RecvDataHandler  = func(client *Client, data []byte, sequenced bool) error         // 客户端接收的数据的处理器
	RecvEventHandler = func(client *Client, event protocol.Event[transport.Msg]) error // 客户端接收的自定义事件的处理器
)

// Client 客户端
type Client struct {
	context.Context
	cancel        context.CancelFunc
	mutex         sync.Mutex
	options       ClientOptions
	sessionId     string
	endpoint      string
	transceiver   protocol.Transceiver
	dispatcher    protocol.EventDispatcher
	trans         protocol.TransProtocol
	ctrl          protocol.CtrlProtocol
	sendDataChan  chan gate.SendData
	recvDataChan  chan gate.RecvData
	sendEventChan chan protocol.Event[transport.Msg]
	recvEventChan chan gate.RecvEvent
	logger        *zap.SugaredLogger
}

// String implements fmt.Stringer
func (c *Client) String() string {
	return fmt.Sprintf("{SessionId:%s AuthToken:%s Endpoint:%s}", c.GetSessionId(), c.GetToken(), c.GetEndpoint())
}

// GetSessionId 获取会话Id
func (c *Client) GetSessionId() string {
	return c.sessionId
}

// GetToken 获取token
func (c *Client) GetToken() string {
	return c.options.AuthToken
}

// GetEndpoint 获取服务器地址
func (c *Client) GetEndpoint() string {
	return c.endpoint
}

// SendData 发送数据
func (c *Client) SendData(data []byte, sequenced bool) error {
	return c.trans.SendData(data, sequenced)
}

// SendEvent 发送自定义事件
func (c *Client) SendEvent(event protocol.Event[transport.Msg]) error {
	return protocol.Retry{
		Transceiver: &c.transceiver,
		Times:       c.options.IORetryTimes,
	}.Send(c.transceiver.Send(event))
}

// SendDataChan 发送数据的channel
func (c *Client) SendDataChan() chan<- gate.SendData {
	if c.sendDataChan == nil {
		c.logger.Panic("send data channel size less equal 0, can't be used")
	}
	return c.sendDataChan
}

// RecvDataChan 接收数据的channel
func (c *Client) RecvDataChan() <-chan gate.RecvData {
	if c.recvDataChan == nil {
		c.logger.Panic("receive data channel size less equal 0, can't be used")
	}
	return c.recvDataChan
}

// SendEventChan 发送自定义事件的channel
func (c *Client) SendEventChan() chan<- protocol.Event[transport.Msg] {
	if c.sendEventChan == nil {
		c.logger.Panic("send event channel size less equal 0, can't be used")
	}
	return c.sendEventChan
}

// RecvEventChan 接收自定义事件的channel
func (c *Client) RecvEventChan() <-chan gate.RecvEvent {
	if c.recvEventChan == nil {
		c.logger.Panic("receive event channel size less equal 0, can't be used")
	}
	return c.recvEventChan
}

// Close 关闭
func (c *Client) Close() {
	c.cancel()
}