package handler

import (
	"fmt"

	"github.com/hpcloud/tail"
)

type HandlerDemo struct {
}

func (h *HandlerDemo) HandlerBefore() interface{} {
	fmt.Println("start read log")
	return nil
}

func (h *HandlerDemo) Handler(line *tail.Line, opts ...interface{}) interface{} {
	// fmt.Println(line.Text)
	return nil
}

func (h *HandlerDemo) HandlerAfter() interface{} {
	fmt.Println("end...")
	return nil
}

func (h *HandlerDemo) StartHandler(line *tail.Line) {
	h.HandlerBefore()
	h.Handler(line)
	h.HandlerAfter()
}
