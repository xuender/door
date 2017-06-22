package main

import (
	"log"
	"net/http"
	"time"

	"./chat"
	"github.com/labstack/echo"
	"github.com/xuender/door"
)

var chats []*chat.Chat

func main() {
	chats = make([]*chat.Chat, 0)
	d := door.New()
	d.OPEN(open)
	d.POST("send", send)
	log.Println("启动")

	e := echo.New()
	e.GET("/", html)
	e.GET("/ws", func(c echo.Context) error {
		return d.WebsocketHandler(c.Response().Writer, c.Request())
	})
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8888"))
}

func send(c door.Context) error {
	ca := &chat.Chat{}
	c.Unmarshal(ca)
	ca.Timestamp = time.Now().UnixNano()
	log.Printf("chat: %v\n", c)
	chats = append(chats, ca)
	if len(chats) > 20 {
		chats = chats[1:21]
	}
	for _, num := range c.Numbers() {
		c.Send(num, door.MethodEnum_POST, "send", ca)
	}
	return nil
}

func open(c door.Context) error {
	c.Revert(door.MethodEnum_PUT, "send", &chat.Chats{
		Chats: chats,
	})
	ca := &chat.Chat{
		Nick:      "机器人",
		Context:   "欢迎光临",
		Timestamp: time.Now().UnixNano(),
	}
	c.Revert(door.MethodEnum_POST, "send", ca)
	return nil
}

func html(c echo.Context) error {
	return c.HTML(http.StatusOK, `<html><body><script>
		var s = new WebSocket('ws://localhost:8888/ws');
		s.onopen = function(event) {
			// s.send('I am the client and I\'m listening!');
			s.onmessage = function(event) {
				console.log('Client received a message',event);
			};
			// 监听Socket的关闭
			s.onclose = function(event) {
				console.log('Client notified socket has closed',event);
			};
		};
		</script></body></html>`)
}
