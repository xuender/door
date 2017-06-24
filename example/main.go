package main

import (
	"fmt"
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
	d.PUT("nick", nick)

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

func html(c echo.Context) error {
	return c.HTML(http.StatusOK, `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>聊天室</title>
  <base href="/">
  <meta name="viewport" content="width=device-width, initial-scale=1">
<link href="static/styles.d41d8cd98f00b204e980.bundle.css" rel="stylesheet"/></head>
<body>
  <app-root></app-root>
<script type="text/javascript" src="static/inline.893a24a4fa5604695bc9.bundle.js"></script><script type="text/javascript" src="static/polyfills.5ca19fb28cd74641de8d.bundle.js"></script><script type="text/javascript" src="static/vendor.75eee41300e46be93317.bundle.js"></script><script type="text/javascript" src="static/main.d06763307dcbc3a1afa5.bundle.js"></script></body>
</html>`)
}
