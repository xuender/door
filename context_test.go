package door

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestContext(t *testing.T) {
	Convey("Context", t, func() {
		d := New()
		d.open(nil, 1)
		c := Context{
			door: d,
			num:  1,
		}
		Convey("attribute", func() {
			c.PutAttribute("test", 13)
			v, _ := c.GetAttribute("test")
			So(v, ShouldEqual, 13)
		})
	})
}
