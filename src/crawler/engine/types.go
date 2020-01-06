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
	Submit(in chan Request,request Request)
	//ConfigureMasterWorkerChan(chan Request)
}
