package api

import (
	"context"
	"fmt"
	com "mytail/common"
	"mytail/handler"
	"mytail/tailfile"
	"strconv"
	"time"

	"github.com/hpcloud/tail"
)

type schedule struct {
	tasks   chan *tailfile.TailTask
	taskMap map[string]*tailfile.TailTask // 任务ID与该任务的上下文对象的map，通过这个ctx控制协程的结束
}

func Newschedule() *schedule {
	return &schedule{
		tasks:   make(chan *tailfile.TailTask, 10), //  代表同时可以处理10个日志对象
		taskMap: make(map[string]*tailfile.TailTask),
	}
}

// registerHandler 给任务注册handler
func (s *schedule) RegisterHandler(handle handler.Handler, task *tailfile.TailTask) {
	task.Handler = handle
}

// 生成tail 任务
func (s *schedule) GoTask(taskName string, taskID string, path string, offset int64, whence int) *tailfile.TailTask {

	// 确认offset,如果找不到对应的record，则默认offset为0
	isexists, _ := com.PathExists(fmt.Sprintf("%s.record", taskID))
	if isexists {
		recordFilePath := fmt.Sprintf("%s.record", taskID)
		buffer := com.ReadFile(recordFilePath)
		tmp1, _ := strconv.Atoi(string(buffer))
		offset = int64(tmp1)
		fmt.Println(offset)
	}

	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: offset, Whence: whence},
		MustExist: false,
		Poll:      true,
	}
	ctx, cancel := context.WithCancel(context.Background()) // ctx 管理goroutine
	tailTask := tailfile.NewTailTask(ctx, cancel, path, config, taskName, taskID)
	s.taskMap[taskID] = tailTask
	return tailTask
}

// 推送任务到队列
func (s *schedule) PutTask(task *tailfile.TailTask) {
	s.tasks <- task
}

// 开始任务
func (s *schedule) Start() {
	for {
		select {
		case task := <-s.tasks:
			go task.TailFile() // 这里的协程数量得控制一下，用协程池吧
		default:
			time.Sleep(time.Second)
		}
	}
}

// 停止任务
func (s *schedule) Stop(taskIDs ...string) {
	// 如果没有传入taskIDs，则停止全部任务
	if len(taskIDs) == 0 {
		for _, task := range s.taskMap {
			task.Stop()
		}
		return
	}
	for _, taskID := range taskIDs {
		if task, ok := s.taskMap[taskID]; ok {
			task.Stop()
		}
	}
}
