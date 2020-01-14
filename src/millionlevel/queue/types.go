package queue

import . "millionlevel/job"

//接口
type Scheduler interface {
	Submit(job Job)
	ReadyNotifier
	Run()
	WorkerChan() chan Job
}

//接口
type ReadyNotifier interface {
	WorkerReady(chan Job)
}


type Worker interface {
	Work(Job) error
}