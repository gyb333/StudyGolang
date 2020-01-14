package main

import (
				"os"
	"strconv"
	"time"
	"fmt"
	"runtime"
	. "millionlevel/job"
	. "millionlevel/core"
	"millionlevel/queue"
)



var (
	MaxWorker ,_= strconv.ParseInt(os.Getenv("MAX_WORKERS"),10,64)
	MaxQueue  = os.Getenv("MAX_QUEUE")
)

/*
简单粗暴起协程处理耗时任务导致的系统不可控性
建立任务队列。golang提供了线程安全的任务队列实现方式–带缓冲的channal。但是这样只是延后了请求的爆发。
要解决这一问题，必须控制协程的数量。如何控制协程的数量？Job/Worker模式！
 */
func main(){
	JobQueue = make(chan Job, 100)

	go addQueue()
	//go SendData()
	time.Sleep(100 * time.Second)


}

func SendData()  {
	engine :=&queue.Engine{
		Scheduler:&queue.QueuedScheduler{},
		WorkerCount:1200000,
	}
	engine.Run()
	for i := 0; i < 30000000; i++ {
		payLoad := Payload{Num: i}
		job := Job{Payload: payLoad}
		// 任务放入任务队列channal
		engine.Scheduler.Submit(job)
		fmt.Printf("正在发生任务i:%d,当前协程数:%d\n",i+1, runtime.NumGoroutine())
		if (i+1)%10000==0{
			time.Sleep(1*time.Millisecond)
		}
	}
}

func addQueue() {
	dispatcher := NewDispatcher(1200000)
	dispatcher.Run()

	for i := 0; i < 30000000; i++ {
		// 新建一个任务
		payLoad := Payload{Num: i}
		job := Job{Payload: payLoad}
		// 任务放入任务队列channal
		JobQueue <- job
		//fmt.Println("JobQueue <- work", i)
		fmt.Printf("正在发生任务i:%d,当前协程数:%d\n",i+1, runtime.NumGoroutine())
		if (i+1)%10000==0{
			time.Sleep(100*time.Microsecond)
		}

	}
}