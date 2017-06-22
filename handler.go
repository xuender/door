package door

import (
	"errors"
)

// HandlerFunc 定义 websocket 请求.
type HandlerFunc func(Context) error

// ErrNotFound 没找到错误.
var ErrNotFound = errors.New("Not found")

// NotFoundHandler 没找到.
var NotFoundHandler = func(c Context) error {
	return ErrNotFound
}
