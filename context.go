package door

import (
	"errors"
	fmt "fmt"

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
func (context *Context) Send(num uint32, method MethodEnum, path string, pb proto.Message) error {
	if conn, ok := context.door.conns[num]; ok {
		return context.send(num, conn, method, path, pb)
	}
	return errors.New(fmt.Sprintf("错误的代号:%d", num))
}

// Revert 回复.
func (context *Context) Revert(method MethodEnum, path string, pb proto.Message) error {
	return context.send(context.num, context.conn, method, path, pb)
}

func (context *Context) send(num uint32, conn *websocket.Conn, method MethodEnum, path string, pb proto.Message) error {
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
	err := conn.WriteMessage(websocket.BinaryMessage, eventBytes)
	if err != nil {
		context.door.Close(num)
	}
	return err
}
