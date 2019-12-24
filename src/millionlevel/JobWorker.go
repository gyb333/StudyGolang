package millionlevel

import (
			"encoding/json"
	"io"
	"net/http"
	"fmt"
	"time"
)

//待执行的工作
type Job struct {
	Payload Payload
}







//任务channal
var JobQueue chan Job

//执行任务的工作者单元
type Worker struct {
	WorkerPool chan chan Job //工作者池--每个元素是一个工作者的私有任务channal
	JobChannel chan Job      //每个工作者单元包含一个任务管道 用于获取任务
	quit       chan bool     //退出信号
	no         int           //编号
}

//创建一个新工作者单元
func NewWorker(workerPool chan chan Job, no int) Worker {
	fmt.Println("创建一个新工作者单元")
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
		no:         no,
	}
}

//循环  监听任务和结束信号
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel
			fmt.Println("w.WorkerPool <- w.JobChannel", w)
			select {
			case job := <-w.JobChannel:
				if err := job.Payload.UploadToS3(); err != nil {
					//log.Errorf("Error uploading to S3: %s", err.Error())
				}

				fmt.Println("job := <-w.JobChannel")
				// 收到任务
				fmt.Println(job)
				time.Sleep(100 * time.Second)
			case <-w.quit:
				// 收到退出信号
				return
			}
		}
	}()
}



// 停止信号
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}



func payloadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Read the body into a string for json decoding
	var content = &PayloadCollection{}
	err := json.NewDecoder(io.LimitReader(r.Body, MaxLength)).Decode(&content)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Go through each payload and queue items individually to be posted to S3
	for _, payload := range content.Payloads {

		// let's create a job with the payload
		work := Job{Payload: payload}

		// Push the work onto the queue.
		//JobQueue <- work
		//我们可以根据压测结果设置合适的并发数从而保证系统能够尽可能的发挥自己的能力，同时保证不会因为任务量太大而崩溃
		Limit(work)
	}

	w.WriteHeader(http.StatusOK)
}

//用于控制并发处理的协程数
var DispatchNumControl = make(chan bool, 10000)

func Limit(work Job) bool {
	select {
	case <-time.After(time.Millisecond * 100):
		fmt.Println("我很忙")
		return false
	case DispatchNumControl <- true:
		// 任务放入任务队列channal
		JobQueue <- work
		return true
	}
}
