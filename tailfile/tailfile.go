package tailfile

import (
	"context"
	"fmt"
	"log"
	"mytail/handler"
	"os"
	"strconv"
	"time"

	"github.com/hpcloud/tail"
)

type Handler func(line *tail.Line)

type TailTask struct {
	path     string
	ctx      context.Context
	cancel   context.CancelFunc
	config   tail.Config
	Handler  handler.Handler
	TaskName string
	TaskID   string
}

type TailConfig struct {
	Offset int64
	Whence int
}

func NewTailTask(ctx context.Context, cancel context.CancelFunc, path string, config tail.Config, taskName string, taskID string) *TailTask {
	return &TailTask{
		path:     path,
		config:   config,
		ctx:      ctx,
		TaskName: taskName,
		TaskID:   taskID,
		cancel:   cancel,
	}
}

func (t *TailTask) Stop() {
	t.cancel()
}

func (t *TailTask) TailFile() {
	var err error
	var tailobj *tail.Tail
	tailobj, err = tail.TailFile(t.path, t.config)
	defer tailobj.Cleanup()
	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		select {
		case <-t.ctx.Done(): // 等待上级通知
			fmt.Println("任务结束")
			return
		case msg, _ := <-tailobj.Lines: // 这里是非阻塞的，写到default就是阻塞的
			offset, err := tailobj.Tell()
			if err != nil {
				fmt.Println(err)
			}
			record_offset(offset, t.TaskID)
			go t.Handler.StartHandler(msg)
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func record_offset(offset int64, taskid string) {
	f, err := os.OpenFile(fmt.Sprintf("%s.record", taskid), os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	f.Write([]byte(strconv.Itoa(int(offset))))
	f.Close()
}
