package main

import (
	"github.com/XiaoHuYi/zolo/api"
	h "github.com/XiaoHuYi/zolo/handler"
)

func main() {
	s := api.Newschedule()

	taskName := "test1"
	// taskID := "123456"                                                             // 任务ID必须唯一，用UUID算法生成吧
	task := s.GoTask(taskName, "/home/docker_images/sysdig/log/sysdig.json", 0, 0) // 生成task对象
	s.RegisterHandler(&h.HandlerDemo{}, task)                                      // 注册任务
	s.PutTask(task)                                                                // 添加任务到执行队列

	task2 := s.GoTask("task2", "/home/beancluster_log/sysdig/log.json", 0, 0)
	s.RegisterHandler(&h.HandlerDemo{}, task2) // 注册任务
	s.PutTask(task2)

	go s.Start() // 启动schedule

	select {}
}
