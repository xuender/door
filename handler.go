package door

// HandlerFunc 定义 websocket 请求.
type HandlerFunc func(Context) error

// NotFoundHandler 没找到.
var NotFoundHandler = func(c Context) error {
	return ErrorNotFound
}
