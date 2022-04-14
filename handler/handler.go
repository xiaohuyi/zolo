/*
 * @Author: xiaohuyi
 * @Date: 2022-04-14 14:26:54
 * @Description:
 */
package handler

import (
	"fmt"
	"os"
	"strconv"
)

// 需要实现Handler 接口
type HandlerDemo struct {
}

// 写入到文件，IO效率低，正式使用可以改为放到redis，再起一个线程定时去同步redis的offset到文件，也是为了持久化吧
func (h *HandlerDemo) RecordOffset(offset int64, taskID string) {
	f, err := os.OpenFile(fmt.Sprintf("%s.record", taskID), os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	f.Write([]byte(strconv.Itoa(int(offset))))
	f.Close()
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
