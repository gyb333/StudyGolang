package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"strconv"
	"fmt"
	)

type Executer interface {
	Execute() error
}




//Job任务
type Job struct {
	Tasks []Executer
}

func NewJob(taskId int,cap int) *Job{
	return &Job{
		Tasks:make([]Executer,0,cap),
	}
}

type Worker struct {
	WorkerId string
	JobChan  chan Job
	WorkChan chan chan Job
	Finished chan bool
}

func NewWorker(WorkChan chan chan Job,Id string) *Worker{
	return &Worker{
		WorkerId: Id,
		JobChan:  make(chan Job),
		WorkChan: WorkChan,
		Finished: make(chan bool),
	}
}

func (w *Worker) Start(){
	go func(){
		for {
			w.WorkChan <-w.JobChan
			log.Printf("把[%s]的工作台添加到工作台队列中，当前工作台队列长度：%d\n", w.WorkerId, len(w.WorkChan))
			select {
			case wJob:=<-w.JobChan:
				for _,task :=range wJob.Tasks{
					 task.Execute()
				}
			case bFinished :=<-w.Finished:
				if true==bFinished{
					log.Printf("%s 结束工作！\n", w.WorkerId)
					return
				}
			}
		}
	}()
}

func (w *Worker) Stop(){
	go func(){
		w.Finished<-true
	}()
}


type Dispatcher struct {
	DispatcherId string         //流水线ID
	MaxWorkers   int            //流水线上的员工(Worker)最大数量
	Workers      []*Worker      //流水线上所有员工(Worker)对象集合
	Closed       chan bool      //流水线工作状态通道
	EndDispatch  chan os.Signal //流水线停止工作信号
	JobQueue     chan Job       //流水线上的所有代加工产品(Job)队列通道
	WorkQueue    chan chan Job  //流水线上的所有操作台队列通道
}

func NewDispatcher(maxWorkers, maxQueue int) *Dispatcher {
	Closed := make(chan bool)
	EndDispatch := make(chan os.Signal)
	JobQueue := make(chan Job, maxQueue)
	WorkQueue := make(chan chan Job, maxWorkers)
	signal.Notify(EndDispatch, syscall.SIGINT, syscall.SIGTERM)
	return &Dispatcher{
		DispatcherId: "调度者",
		MaxWorkers:   maxWorkers,
		Closed:       Closed,
		EndDispatch:  EndDispatch,
		JobQueue:     JobQueue,
		WorkQueue:    WorkQueue,
	}
}

func (d *Dispatcher) Run() {
	// 开始运行
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(d.WorkQueue, fmt.Sprintf("work-%s", strconv.Itoa(i)))
		d.Workers = append(d.Workers, worker)
		//开始工作
		worker.Start()
	}
	//监控
	go d.Dispatch()
}

func (d *Dispatcher) Dispatch() {
FLAG:
	for {
		select {
		case endDispatch := <-d.EndDispatch:
			log.Printf("流水线关闭命令[%v]已发出...\n", endDispatch)
			close(d.JobQueue)
		case wJob, Ok := <-d.JobQueue:
			if true == Ok {
				log.Println("从流水线获取一个待加工产品(Job)")
				go func(wJob Job) {
					//获取未分配待加工产品的工作台
					Workbench := <-d.WorkQueue
					//将待加工产品(Job)放入工作台进行加工
					Workbench <- wJob
				}(wJob)
			} else {
				for _, w := range d.Workers {
					w.Stop()
				}
				d.Closed <- true
				break FLAG
			}
		}
	}
}

type WorkFlow struct {
	Dispatch *Dispatcher
}

func (wf *WorkFlow) StartWorkFlow(maxWorkers, maxQueue int) {
	//初始化一个调度器(流水线)，并指定该流水线上的员工(Worker)和待加工产品(Job)的最大数量
	wf.Dispatch = NewDispatcher(maxWorkers, maxQueue)
	//启动流水线
	wf.Dispatch.Run()
}

func (wf *WorkFlow) AddJob(wJob Job) {
	//向流水线中放入待加工产品(Job)
	wf.Dispatch.JobQueue <- wJob
}

func (wf *WorkFlow) CloseWorkFlow() {
	closed := <-wf.Dispatch.Closed
	if true == closed {
		log.Println("调度器(流水线)已关闭.")
	}
}

type Task struct {
	TaskId int
	data int
}
func NewTask(TaskId,data int) *Task{
	return &Task{
		TaskId:TaskId,
		data:data,
	}
}
func (t *Task) Execute() error {
	log.Printf("任务%d-%08d", t.TaskId,t.data)
	return nil
}


func main() {
	var wf WorkFlow
	//初始化并启动工作
	wf.StartWorkFlow(2, 4)
	for i := 0; i < 100; i++ {
		wJob := NewJob(i+1,2)
		wJob.Tasks =append(wJob.Tasks, NewTask(1,i+1))
		wJob.Tasks =append(wJob.Tasks, NewTask( 2,i+1))

		//添加工作
		wf.AddJob(*wJob)
		//time.Sleep(time.Millisecond * 10)
	}
	wf.CloseWorkFlow()




}
