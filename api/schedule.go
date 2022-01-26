package api

import (
	"context"
	"fmt"
	"strconv"
	"time"

	com "github.com/XiaoHuYi/zolo/common"
	"github.com/XiaoHuYi/zolo/handler"
	"github.com/XiaoHuYi/zolo/tailfile"

	"github.com/hpcloud/tail"
)

const (
	MAXTASK int = 10
)

type schedule struct {
	tasks   chan *tailfile.TailTask
	taskMap map[string]*tailfile.TailTask // 任务ID与该任务的上下文对象的map，通过这个ctx控制协程的结束
}

func Newschedule() *schedule {
	return &schedule{
		tasks:   make(chan *tailfile.TailTask, MAXTASK), //  代表同时可以处理10个日志对象
		taskMap: make(map[string]*tailfile.TailTask),
	}
}

// registerHandler 给任务注册handler
func (s *schedule) RegisterHandler(handle handler.Handler, task *tailfile.TailTask) {
	task.Handler = handle
}

// 生成任务ID  具体可能考虑用文件路径的哈希值吧，形成 文件路径 与 ID 的一一对应
func (s *schedule) genTaskID(taskName string) string {
	return taskName
}

// 生成tail 任务
func (s *schedule) GoTask(taskName string, path string, offset int64, whence int) *tailfile.TailTask {

	// 基于taskName+时间戳 生成一个唯一的taskID
	taskID := s.genTaskID(taskName)

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
			go task.TailFile() // s.tasks最多容量为10，所以最多10个携程数量
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
