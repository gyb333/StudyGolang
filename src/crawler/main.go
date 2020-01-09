package main

import (
	"crawler/engine"
	"crawler/parser"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/config"

	"crawler/worker"
)



var url = config.Url

var cityUrl = config.CityUrl

var profileUrl = config.ProfileUrl
var hosts =[]int{
	9000,9001,9002,9003,
}

func main() {
	//Simple()
	ConCurrent()
}

func ConCurrent() {
	itemChan, err := persist.ItemPrint(config.ItemSaverPort)
	if err !=nil{
		return
	}
	engine.ConCurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:   itemChan ,
		//ItemChan:  persist.SimpleItemPrint(),// itemChan ,
		//Worker: engine.Worker{},
		Worker:worker.RpcWorker{
			WorkChan:worker.CreatWorkerPool(hosts),
		},
	}.Run(
		engine.Request{
			Url:    url,
			Parser: parser.CityListParser{},
		},
		//engine.Request{
		//	Url:    cityUrl,
		//	Parser: parser.CityParser{},
		//},
		//engine.Request{
		//	Url:    profileUrl,
		//	Parser: parser.ProfileParser{},
		//},
	)
}

func Simple() {
	engine.SimpleEngine{
		ItemChan: persist.SimpleItemPrint(),
	}.Run(
		engine.Request{
			Url:    url,
			Parser: parser.CityListParser{},
		},
		//engine.Request{
		//	Url:    cityUrl,
		//	Parser: parser.CityParser{},
		//},
		//engine.Request{
		//	Url:    profileUrl,
		//	Parser: parser.ProfileParser{},
		//},
	)
}
