package door

import (
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
)

// Context 内容.
type Context struct {
	Method MethodEnum
	door   *Door
	conn   *websocket.Conn
	data   []byte
	num    uint32
	filter int
	path   string
}

// Execute 执行过滤器.
func (context *Context) Execute(path string) error {
	i := 0
	for _, filter := range context.door.filters {
		if strings.HasPrefix(path, filter.path) {
			if i == context.filter {
				context.filter++
				context.path = path
				return filter.executer.Execute(context)
			}
			i++
		}
	}
	return nil
}

// Next 下一个过滤器
func (context *Context) Next() error {
	return context.Execute(context.path)
}

// Pass 是否通过过滤器.
func (context *Context) Pass() bool {
	i := 0
	for _, filter := range context.door.filters {
		if strings.HasPrefix(context.path, filter.path) {
			i++
		}
	}
	return context.filter == i
}

// Unmarshal 解码proto.
func (context *Context) Unmarshal(pb proto.Message) error {
	return proto.Unmarshal(context.data, pb)
}

// Marshal 编码proto.
func (context *Context) Marshal(pb proto.Message) (err error) {
	context.data, err = proto.Marshal(pb)
	return
}

// Num 当前Context的编号.
func (context *Context) Num() uint32 {
	return context.num
}

// PutAttribute 设置属性.
func (context *Context) PutAttribute(key string, value interface{}) {
	context.door.attributes[context.num].Put(key, value)
}

// GetAttribute 获取属性.
func (context *Context) GetAttribute(key string) (value interface{}, ok bool) {
	value, ok = context.door.attributes[context.num].Get(key)
	return
}

// DelAttribut 删除属性.
func (context *Context) DelAttribut(key string) {
	context.door.attributes[context.num].Remove(key)
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
func (context *Context) Send(num uint32, pb proto.Message, method MethodEnum, paths ...string) error {
	if conn, ok := context.door.conns[num]; ok {
		return context.send(num, conn, pb, method, paths...)
	}
	return ErrorNotNum{Num: num}
}

// Revert 回复.
func (context *Context) Revert(pb proto.Message, method MethodEnum, paths ...string) error {
	return context.send(context.num, context.conn, pb, method, paths...)
}

func (context *Context) send(num uint32, conn *websocket.Conn, pb proto.Message, method MethodEnum, paths ...string) error {
	event := &Event{
		Method: method,
		Path:   strings.Join(paths, "/"),
	}
	if pb != nil {
		pbBytes, pbErr := proto.Marshal(pb)
		if pbErr != nil {
			return pbErr
		}
		event.Data = pbBytes
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
