package door

import (
	"testing"

	"github.com/gorilla/websocket"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDoor(t *testing.T) {
	Convey("Door", t, func() {
		d := New()
		Convey("OPEN", func() {
			d.OPEN(func(c Context) error { return nil })
			So(len(d.router.routes[MethodEnum_OPEN]), ShouldEqual, 1)
		})
		Convey("CLOSE", func() {
			d.CLOSE(func(c Context) error { return nil })
			So(len(d.router.routes[MethodEnum_CLOSE]), ShouldEqual, 1)
		})
		Convey("GET", func() {
			d.GET(func(c Context) error { return nil }, "path")
			So(len(d.router.routes[MethodEnum_GET]), ShouldEqual, 1)
		})
		Convey("POST", func() {
			d.POST(func(c Context) error { return nil }, "path")
			So(len(d.router.routes[MethodEnum_POST]), ShouldEqual, 1)
		})
		Convey("PUT", func() {
			d.PUT(func(c Context) error { return nil }, "path")
			So(len(d.router.routes[MethodEnum_PUT]), ShouldEqual, 1)
		})
		Convey("DELETE", func() {
			d.DELETE(func(c Context) error { return nil }, "path")
			So(len(d.router.routes[MethodEnum_DELETE]), ShouldEqual, 1)
		})
		Convey("Close", func() {
			d.conns = make(map[uint32]*websocket.Conn, 1)
			d.conns[1] = nil
			So(len(d.conns), ShouldEqual, 1)
			d.Close(1)
			So(len(d.conns), ShouldEqual, 0)
		})
	})
}
