package main

import (
	"fmt"
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
	d.PUT("nick", nick)

	e := echo.New()
	e.GET("/ws", func(c echo.Context) error {
		return d.WebsocketHandler(c.Response().Writer, c.Request())
	})
	e.Static("/", "dist")
	e.Logger.Fatal(e.Start(":8888"))
}

func send(c door.Context) error {
	ca := &chat.Chat{}
	c.Unmarshal(ca)
	ca.Timestamp = time.Now().UnixNano()
	chats = append(chats, ca)
	if len(chats) > 20 {
		chats = chats[1:21]
	}
	for _, num := range c.Numbers() {
		c.Send(num, ca, door.MethodEnum_POST, "send")
	}
	return nil
}

func nick(c door.Context) error {
	ca := &chat.Chat{}
	c.Unmarshal(ca)
	ca.Timestamp = time.Now().UnixNano()
	ca.Context = fmt.Sprintf("欢迎 [ %s ] 进入聊天室！", ca.Nick)
	ca.Nick = "机器人"
	chats = append(chats, ca)
	if len(chats) > 20 {
		chats = chats[1:21]
	}
	for _, num := range c.Numbers() {
		c.Send(num, ca, door.MethodEnum_POST, "send")
	}
	return nil
}

func open(c door.Context) error {
	c.Revert(&chat.Chats{
		Chats: chats,
	}, door.MethodEnum_PUT, "send")
	ca := &chat.Chat{
		Nick:      "机器人",
		Context:   "你进入了一个聊天室",
		Timestamp: time.Now().UnixNano(),
	}
	c.Revert(ca, door.MethodEnum_POST, "send")
	return nil
}
