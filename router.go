package door

// Router 路由器.
type Router struct {
	routes map[MethodEnum]map[string]HandlerFunc
}

// Add 增加路由.
func (router *Router) Add(method MethodEnum, path string, h HandlerFunc) {
	pathMap, ok := router.routes[method]
	if !ok {
		pathMap = make(map[string]HandlerFunc)
		router.routes[method] = pathMap
	}
	pathMap[path] = h
}

// Find 路由查找.
func (router *Router) Find(method MethodEnum, path string) HandlerFunc {
	if pathMap, methodOk := router.routes[method]; methodOk {
		if handlerFunc, pathOk := pathMap[path]; pathOk {
			return handlerFunc
		}
	}
	return NotFoundHandler
}

// NewRouter 新建路由器.
func NewRouter() *Router {
	return &Router{
		routes: make(map[MethodEnum]map[string]HandlerFunc, 0),
	}
}
