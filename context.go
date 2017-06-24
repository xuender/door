package door

import (
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

// Context 内容.
type Context struct {
	conn *websocket.Conn
	data []byte
	num  uint32
	door *Door
}

// Unmarshal 解码proto.
func (context *Context) Unmarshal(pb proto.Message) error {
	return proto.Unmarshal(context.data, pb)
}

// Num 当前Context的编号.
func (context *Context) Num() uint32 {
	return context.num
}

// Numbers 全部客户端编号.
func (context *Context) Numbers() (nums []uint32) {
	nums = make([]uint32, len(context.door.conns))
	i := 0
	for k := range context.door.conns {
		nums[i] = k
		i++
	}
	return
}

// Send 发送对象.
func (context *Context) Send(num uint32, pb proto.Message, method MethodEnum, path ...string) error {
	if conn, ok := context.door.conns[num]; ok {
		return context.send(num, conn, pb, method, path...)
	}
	return ErrorNotNum{Num: num}
}

// Revert 回复.
func (context *Context) Revert(pb proto.Message, method MethodEnum, path ...string) error {
	return context.send(context.num, context.conn, pb, method, path...)
}

func (context *Context) send(num uint32, conn *websocket.Conn, pb proto.Message, method MethodEnum, path ...string) error {
	pbBytes, pbErr := proto.Marshal(pb)
	if pbErr != nil {
		return pbErr
	}
	event := &Event{
		Method: method,
		Path:   strings.Join(path, "/"),
		Data:   pbBytes,
	}
	eventBytes, eventErr := proto.Marshal(event)
	if eventErr != nil {
		return eventErr
	}
	err := conn.WriteMessage(websocket.BinaryMessage, eventBytes)
	if err != nil {
		context.door.Close(num)
	}
	return err
}
