package handler

type Handler interface {
	HandlerBefore() interface{}                           // 正式handle之前的hook
	Handler(line string, opts ...interface{}) interface{} // handle逻辑
	HandlerAfter() interface{}                            // 正式handle之后的hook
	StartHandler(line string)
}
