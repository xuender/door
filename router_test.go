package door

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRouter(t *testing.T) {
	Convey("Router", t, func() {
		r := NewRouter()
		Convey("Add", func() {
			So(r.Find(MethodEnum_GET, "test"), ShouldEqual, NotFoundHandler)
			h := func(c Context) error { return nil }
			r.Add(MethodEnum_GET, "test", h)
			So(r.Find(MethodEnum_GET, "test"), ShouldEqual, h)
			So(r.Find(MethodEnum_POST, "test"), ShouldEqual, NotFoundHandler)
		})
	})
}
