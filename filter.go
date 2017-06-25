package door

// Filter 过滤器.
type Filter struct {
	path     string
	executer Executer
}

// Executer 过滤器接口.
type Executer interface {
	Execute(*Context) error
}

// Inline 是单行过滤器.
type Inline struct {
	Handler func(*Context) error
}

// Execute runs the inlined handler.
func (filter *Inline) Execute(context *Context) error {
	return filter.Handler(context)
}
