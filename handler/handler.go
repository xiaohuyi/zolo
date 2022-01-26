package handler

import (
	"fmt"
)

// 需要实现Handler 接口
type HandlerDemo struct {
}

func (h *HandlerDemo) HandlerBefore() interface{} {
	//fmt.Println("start read log")
	return nil
}

func (h *HandlerDemo) Handler(line string, opts ...interface{}) interface{} {
	fmt.Println(line)
	return nil
}

func (h *HandlerDemo) HandlerAfter() interface{} {
	//fmt.Println("end...")
	return nil
}

func (h *HandlerDemo) StartHandler(line string) {
	h.HandlerBefore()
	h.Handler(line)
	h.HandlerAfter()
}
