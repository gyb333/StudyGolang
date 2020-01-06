package engine

import (
	"log"
	"crawler/fetcher"


)

type ConCurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
}

func (e ConCurrentEngine) Run(seeds ...Request){
	e.Scheduler.Run()
	out := make(chan ParseResult)
	for i := 0; i < e.WorkerCount; i++ {
		//createWorker(in, out) //创建worker
		createWorker(e.Scheduler.WorkerChan(),out, e.Scheduler)
	}
	//参数seeds的request，要分配任务
	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}

	//从out中获取result，对于item就打印即可，对于request，就继续分配
	for {
		result := <-out
		for _, item := range result.Items {
			go func() {
				e.ItemChan <- item
			}()
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request,out chan ParseResult,ready  ReadyNotifier) {
	//in := make(chan Request)
	go func() {
		for {
			//需要让scheduler知道已经就绪了
			ready.WorkerReady(in)
			request := <-in
			result, err := Work(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}


func (e ConCurrentEngine)SimpleRun(seeds ...Request)  {
	//worker公用一个in，out
	in:=make(chan Request)
	out :=make(chan ParseResult)
	for i:=0;i<e.WorkerCount;i++{
		createSimpleWorker(in,out)
	}
	for _,request :=range seeds{
		in <-request
	}
	itemCount := 0
	//从out中获取result，对于item就打印即可，对于request，就继续分配
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got %d  item : %v",itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			in <-request
		}
	}
}

//创建worker
func createSimpleWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			request := <-in
			result, err := Work(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func Work(request Request) (ParseResult,error) {
	log.Printf("Fetching %s",request.Url)
	body,err:=fetcher.Fetch(request.Url)
	if err!=nil{
		return ParseResult{},err
	}
	//log.Printf("%s",string(body))
	return request.Parser.Parse(body),nil

}