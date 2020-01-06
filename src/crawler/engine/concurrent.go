package engine

import (
	"log"
	"crawler/fetcher"


)

type ConCurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

func (e ConCurrentEngine) Run(seeds ...Request){
	//worker公用一个in，out
	//in:=make(chan Request)
	out :=make(chan ParseResult)
	for i:=0;i<e.WorkerCount;i++{
		createWorker(in,out)
	}

	for _,request :=range seeds{
		//e.Scheduler.Submit(in,r)
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
			//e.Scheduler.Submit(in,request)
			in <-request
		}
	}
}

//创建worker
func createWorker(in chan Request, out chan ParseResult) {
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