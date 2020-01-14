package core

import (
	"fmt"
	. "millionlevel/job"
	)

//调度中心
type Dispatcher struct {
	//工作者池
	WorkerPool chan chan Job
	//工作者数量
	MaxWorkers int
}

//创建调度中心
func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, MaxWorkers: maxWorkers}
}

//工作者池的初始化
func (d *Dispatcher) Run() {
	// starting n number of workers
	for i := 1; i < d.MaxWorkers+1; i++ {
		worker := NewWorker(d.WorkerPool, i)
		worker.Start()
	}
	go d.dispatch()
}

//调度
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			//fmt.Println("job := <-JobQueue:")
			//这个调度方法也是在不断的创建协程等待空闲的worker
			go func(job Job) {
				//等待空闲worker (任务多的时候会阻塞这里)
				worker := <-d.WorkerPool
				//fmt.Println("worker := <-d.WorkerPool", reflect.TypeOf(worker))
				// 将任务放到上述woker的私有任务channal中
				worker <- job
				fmt.Println("worker <- job")
			}(job)
		}
	}
}
