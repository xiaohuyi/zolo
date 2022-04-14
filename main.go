/*
 * @Author: xiaohuyi
 * @Date: 2022-04-14 14:26:54
 * @Description:
 */
package main

import (
	"time"

	"github.com/xiaohuyi/zolo/api"
	h "github.com/xiaohuyi/zolo/handler"
)

func main() {
	s := api.Newschedule()

	taskName := "test1"
	// taskID := "123456"                                                             // 任务ID必须唯一，用UUID算法生成吧
	task := s.GoTask(taskName, "eventlog.xml", 0, 0) // 生成task对象
	s.RegisterHandler(&h.HandlerDemo{}, task)        // 注册任务
	s.PutTask(task)                                  // 添加任务到执行队列

	task2 := s.GoTask("task2", "eventlog.xml", 0, 0)
	s.RegisterHandler(&h.HandlerDemo{}, task2) // 注册任务
	s.PutTask(task2)

	go s.Start() // 启动schedule

	time.Sleep(5 * time.Second)

	s.Stop("task2") //  停止某个任务，传入taskid即可
	select {}
}
