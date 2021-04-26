package handler

import "github.com/hpcloud/tail"

type Handler interface {
	HandlerBefore() interface{}
	Handler(line *tail.Line, opts ...interface{}) interface{}
	HandlerAfter() interface{}
	StartHandler(line *tail.Line)
}
