package door

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFilter(t *testing.T) {
	Convey("Filter", t, func() {
		d := New()
		d.open(nil, 1)
		c := Context{
			door: d,
			num:  1,
		}
		Convey("AddFilter", func() {
			d.AddFilter(&Inline{Handler: func(c *Context) error {
				c.PutAttribute("f1", true)
				c.PutAttribute("num", 3)
				return c.Next()
			}}, "a")
			d.AddFilter(&Inline{Handler: func(c *Context) error {
				c.PutAttribute("f2", true)
				num, _ := c.GetAttribute("num")
				c.PutAttribute("num", 8*num.(int))
				return c.Next()
			}}, "a")
			// 忽略的过滤器
			d.AddFilter(&Inline{Handler: func(c *Context) error {
				c.PutAttribute("f2", false)
				num, _ := c.GetAttribute("num")
				c.PutAttribute("num", 8*num.(int))
				return c.Next()
			}}, "a", "c")
			So(len(d.filters), ShouldEqual, 3)
			c.Execute("a/b")
			So(c.Pass(), ShouldEqual, true)
			f1, _ := c.GetAttribute("f1")
			f2, _ := c.GetAttribute("f2")
			num, _ := c.GetAttribute("num")
			So(f1, ShouldEqual, true)
			So(f2, ShouldEqual, true)
			So(num, ShouldEqual, 24)
		})
	})
}
