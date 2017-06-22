package door

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/xuender/goutils"
)

// Door golang 的 websocket 交互工具.
type Door struct {
	router   *Router
	upgrader websocket.Upgrader
	conns    map[uint32]*websocket.Conn
}

// Router 路由.
func (door *Door) Router() *Router {
	return door.router
}

// GET 设置获取功能.
func (door *Door) GET(path string, h HandlerFunc) {
	door.router.Add(MethodEnum_GET, path, h)
}

// POST 设置增加功能.
func (door *Door) POST(path string, h HandlerFunc) {
	door.router.Add(MethodEnum_POST, path, h)
}

// PUT 设置修改功能.
func (door *Door) PUT(path string, h HandlerFunc) {
	door.router.Add(MethodEnum_PUT, path, h)
}

// DELETE 设置删除功能.
func (door *Door) DELETE(path string, h HandlerFunc) {
	door.router.Add(MethodEnum_DELETE, path, h)
}

// WebsocketHandler Websocket 回调.
func (door *Door) WebsocketHandler(w http.ResponseWriter, r *http.Request) error {
	conn, err := door.upgrader.Upgrade(w, r, nil)
	goutils.CheckError(err)
	num := goutils.UniqueUint32()
	door.conns[num] = conn
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		if messageType == websocket.BinaryMessage {
			event := &Event{}
			if err = proto.Unmarshal(p, event); err != nil {
				continue
			}
			door.router.Find(MethodEnum(event.Method), event.Path)(Context{
				conn: conn,
				data: event.Data,
				num:  num,
				door: door,
			})
		}
	}
}

// New 新建Door.
func New() *Door {
	return &Door{
		router: NewRouter(),
		conns:  make(map[uint32]*websocket.Conn, 0),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}
