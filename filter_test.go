package door

import (
	"fmt"
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
			d.AddFilter(&Inline{func(c *Context) error {
				fmt.Println(1)
				err := c.Next()
				fmt.Println(-1)
				return err
			}}, "a", "b")
			d.AddFilter(&Inline{func(c *Context) error {
				fmt.Println(2)
				err := c.Next()
				fmt.Println(-2)
				return err
			}}, "a", "b")
			So(len(d.filters), ShouldEqual, 1)
		})
		Convey("Execute", func() {
			c.Execute("a/b")
			So(c.Pass(), ShouldEqual, true)
		})
	})
}
