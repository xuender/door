package door

import (
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/xuender/goutils"
)

// Door is golang websocket router.
type Door struct {
	router     *Router
	upgrader   websocket.Upgrader
	conns      map[uint32]*websocket.Conn
	attributes map[uint32]goutils.ChMap
	filters    map[string][]Executer
}

// OPEN 设置开启功能.
func (door *Door) OPEN(h HandlerFunc, paths ...string) {
	door.router.Add(h, MethodEnum_OPEN, paths...)
}

// CLOSE 设置关闭功能.
func (door *Door) CLOSE(h HandlerFunc, paths ...string) {
	door.router.Add(h, MethodEnum_CLOSE, paths...)
}

// GET 设置获取功能.
func (door *Door) GET(h HandlerFunc, paths ...string) {
	door.router.Add(h, MethodEnum_GET, paths...)
}

// POST 设置增加功能.
func (door *Door) POST(h HandlerFunc, paths ...string) {
	door.router.Add(h, MethodEnum_POST, paths...)
}

// PUT 设置修改功能.
func (door *Door) PUT(h HandlerFunc, paths ...string) {
	door.router.Add(h, MethodEnum_PUT, paths...)
}

// DELETE 设置删除功能.
func (door *Door) DELETE(h HandlerFunc, paths ...string) {
	door.router.Add(h, MethodEnum_DELETE, paths...)
}

// WebsocketHandler Websocket handler.
func (door *Door) WebsocketHandler(w http.ResponseWriter, r *http.Request) error {
	conn, err := door.upgrader.Upgrade(w, r, nil)
	goutils.CheckError(err)
	num := goutils.UniqueUint32()
	defer door.Close(num)
	door.open(conn, num)
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
			context := Context{
				Method: event.Method,
				conn:   conn,
				data:   event.Data,
				num:    num,
				door:   door,
			}
			context.Execute(event.Path)
			if context.Pass() {
				door.router.Find(MethodEnum(event.Method), event.Path)(context)
			}
		}
	}
}

func (door *Door) open(conn *websocket.Conn, num uint32) {
	door.conns[num] = conn
	door.attributes[num] = goutils.NewChMap()
	for _, handlerFunc := range door.router.Finds(MethodEnum_OPEN) {
		handlerFunc(Context{
			Method: MethodEnum_OPEN,
			conn:   conn,
			num:    num,
			door:   door,
		})
	}
}

// Close connect by num.
func (door *Door) Close(num uint32) {
	if conn, ok := door.conns[num]; ok {
		for _, handlerFunc := range door.router.Finds(MethodEnum_CLOSE) {
			handlerFunc(Context{
				Method: MethodEnum_CLOSE,
				conn:   conn,
				num:    num,
				door:   door,
			})
		}
		delete(door.conns, num)
		if conn != nil {
			conn.Close()
		}
	}
	if m, ok := door.attributes[num]; ok {
		m.Close()
		delete(door.attributes, num)
	}
}

// AddFilter 增加过滤器.
func (door *Door) AddFilter(filter Executer, paths ...string) {
	path := strings.Join(paths, "/")
	filters, ok := door.filters[path]
	if !ok {
		filters = make([]Executer, 0)
	}
	filters = append(filters, filter)
	door.filters[path] = filters
}

// New Door.
func New() *Door {
	return &Door{
		router:     NewRouter(),
		conns:      make(map[uint32]*websocket.Conn, 0),
		attributes: make(map[uint32]goutils.ChMap, 0),
		filters:    make(map[string][]Executer, 0),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}
