package main

import (
	"mytail/api"
	h "mytail/handler"
)

func main() {
	s := api.Newschedule()

	taskName := "test"
	taskID := "123456"
	task := s.GoTask(taskName, taskID, "/opt/mytail/test.txt", 0, 0)
	// handler := func() func(line *tail.Line) {
	// 	count := 0
	// 	return func(line *tail.Line) {
	// 		fmt.Println(line.Text)
	// 		count += 1
	// 	}
	// }()
	s.RegisterHandler(&h.HandlerDemo{}, task)
	s.PutTask(task)
	go s.Start()

	// time.Sleep(time.Second * 5)
	// s.Stop([]string{taskID}...)
	select {}
}
