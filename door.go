package door

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/xuender/goutils"
)

// Door is golang websocket router.
type Door struct {
	router   *Router
	upgrader websocket.Upgrader
	conns    map[uint32]*websocket.Conn
}

// OPEN 设置开启功能.
func (door *Door) OPEN(h HandlerFunc) {
	door.router.Add(MethodEnum_OPEN, "", h)
}

// CLOSE 设置关闭功能.
func (door *Door) CLOSE(h HandlerFunc) {
	door.router.Add(MethodEnum_CLOSE, "", h)
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

// WebsocketHandler Websocket handler.
func (door *Door) WebsocketHandler(w http.ResponseWriter, r *http.Request) error {
	conn, err := door.upgrader.Upgrade(w, r, nil)
	goutils.CheckError(err)
	num := goutils.UniqueUint32()
	defer door.Close(num)
	door.conns[num] = conn
	door.router.Find(MethodEnum_OPEN, "")(Context{
		conn: conn,
		num:  num,
		door: door,
	})
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

// Close connect by num.
func (door *Door) Close(num uint32) {
	if _, ok := door.conns[num]; ok {
		conn := door.conns[num]
		door.router.Find(MethodEnum_CLOSE, "")(Context{
			conn: conn,
			num:  num,
			door: door,
		})
		delete(door.conns, num)
		if conn != nil {
			conn.Close()
		}
	}
}

// New Door.
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
