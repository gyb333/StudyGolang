package engine

type ParseResult struct {
	Requests []Request
	Items []interface{}
}

type Request struct {
	Url string
	Parser Parser
}

type Parser interface {
	Parse([]byte) ParseResult
}

type Enginer interface {
	Run(seeds ...Request)
}


//接口
type Scheduler interface {
	Submit(request Request)
	ReadyNotifier
	Run()
	WorkerChan() chan Request
}

//接口
type ReadyNotifier interface {
	WorkerReady(chan Request)
}