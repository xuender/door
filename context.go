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
