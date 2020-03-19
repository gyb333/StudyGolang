package simple

/*
Goroutine作为轻量级执行流程，也不需要CPU调度器的切换
如果无休止的开辟Goroutine依然会出现高频率的调度Goroutine,那么依然会浪费很多上下文切换的资源,导致做无用功。
超大规模并发的场景下,不加限制大规模的goroutine可能造成内存暴涨,给机器带来极大的压力,
吞吐量下降和处理速度变慢还是其次,更危险的是可能使得程序crash崩溃
 */


import (
	"fmt"
	"time"
)

type Task struct {
	f func() error
}

func NewTask(f func() error) *Task{
	return &Task{
		f:f,
	}
}

func (t *Task) Execute(){
	t.f()
}


type Pool struct {
	EntryChannel chan *Task
	worker_num int
	JobsChannel chan *Task
}

func NewPool(cap int) *Pool{
	return &Pool{
		EntryChannel:make(chan *Task),
		worker_num:cap,
		JobsChannel: make(chan *Task),
	}

}

func (p *Pool) worker(worker_ID int){
	for task :=range p.JobsChannel{
		task.Execute()
		fmt.Println("worker_ID",worker_ID," 执行完毕任务！")
	}
}

func (p *Pool) Run()  {
	defer close(p.JobsChannel)
	defer close(p.EntryChannel)

	for i:=0;i<p.worker_num;i++{
		go p.worker(i)
	}

	for task :=range p.EntryChannel{
		p.JobsChannel<-task
	}



}

func main() {
	t :=NewTask(func() error{
		fmt.Println(time.Now())
		return nil
	})

	p :=NewPool(3)
	go func(){
		for{
			p.EntryChannel<-t
			//p.JobsChannel<- t
		}
	}()

	p.Run()
}
