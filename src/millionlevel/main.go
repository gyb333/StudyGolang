package millionlevel

import (
				"os"
	"strconv"
	"time"
	"fmt"
	"runtime"
)


type Payload struct {
	Num int
}
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
	JobQueue = make(chan Job, 10)
	dispatcher := NewDispatcher(int(MaxWorker))
	dispatcher.Run()
	time.Sleep(1 * time.Second)
	go addQueue()
	time.Sleep(1000 * time.Second)

}

func addQueue() {
	for i := 0; i < 100; i++ {
		// 新建一个任务
		payLoad := Payload{Num: i}
		work := Job{Payload: payLoad}
		// 任务放入任务队列channal
		JobQueue <- work
		fmt.Println("JobQueue <- work", i)
		fmt.Println("当前协程数:", runtime.NumGoroutine())
		time.Sleep(100 * time.Millisecond)
	}
}