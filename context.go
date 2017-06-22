package door

import (
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

// Context 内容.
type Context struct {
	conn *websocket.Conn
	data []byte
}

// Unmarshal 解码proto.
func (context *Context) Unmarshal(pb proto.Message) error {
	return proto.Unmarshal(context.data, pb)
}

// Send 发送对象.
func (context *Context) Send(method MethodEnum, path string, pb proto.Message) error {
	pbBytes, pbErr := proto.Marshal(pb)
	if pbErr != nil {
		return pbErr
	}
	event := &Event{
		Method: method,
		Path:   path,
		Data:   pbBytes,
	}
	eventBytes, eventErr := proto.Marshal(event)
	if eventErr != nil {
		return eventErr
	}
	return context.conn.WriteMessage(websocket.BinaryMessage, eventBytes)
}
