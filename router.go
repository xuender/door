package door

import "strings"

// Router 路由器.
type Router struct {
	routes map[MethodEnum]map[string]HandlerFunc
}

// Add 增加路由.
func (router *Router) Add(h HandlerFunc, method MethodEnum, paths ...string) {
	path := strings.Join(paths, "/")
	pathMap, ok := router.routes[method]
	if !ok {
		pathMap = make(map[string]HandlerFunc)
		router.routes[method] = pathMap
	}
	pathMap[path] = h
}

// Find 路由查找.
func (router *Router) Find(method MethodEnum, paths ...string) HandlerFunc {
	path := strings.Join(paths, "/")
	if pathMap, methodOk := router.routes[method]; methodOk {
		if handlerFunc, pathOk := pathMap[path]; pathOk {
			return handlerFunc
		}
	}
	return NotFoundHandler
}

// Finds 路由全部查找.
func (router *Router) Finds(method MethodEnum) []HandlerFunc {
	ret := make([]HandlerFunc, 0)
	if pathMap, methodOk := router.routes[method]; methodOk {
		for _, hf := range pathMap {
			ret = append(ret, hf)
		}
	}
	return ret
}

// NewRouter 新建路由器.
func NewRouter() *Router {
	return &Router{
		routes: make(map[MethodEnum]map[string]HandlerFunc, 0),
	}
}
