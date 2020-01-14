package queue

import (
	. "millionlevel/job"

	"time"
	"fmt"
)

type Engine struct {
	Scheduler   Scheduler
	WorkerCount int
	//ItemChan chan Item
	//Worker Worker
}

func (e *Engine) Run() {
	e.Scheduler.Run()
	for i := 1; i < e.WorkerCount+1; i++ {
		e.CreateWorker(i,e.Scheduler.WorkerChan(), e.Scheduler)
	}

}

func (e *Engine) CreateWorker(no int,in chan Job, ready ReadyNotifier) {
	go func() {
		for {
			//需要让scheduler知道已经就绪了
			ready.WorkerReady(in)
			job := <-in
			//err := e.Worker.Work(job) //Work(request)
			//if err != nil {
			//	continue
			//}
			fmt.Printf("Worker NO: %d 正在执行任务:%v\n", no, job.Payload.Num)
			// 收到任务
			time.Sleep(10 * time.Second)
			//out <- result
		}
	}()
}

