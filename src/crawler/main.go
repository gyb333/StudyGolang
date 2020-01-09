package main

import (
	"crawler/engine"
	"crawler/parser"
	"crawler/persist"
	"crawler/scheduler"
	"crawler/config"

)



var url = config.Url

var cityUrl = config.CityUrl

var profileUrl = config.ProfileUrl

func main() {
	//Simple()
	ConCurrent()
}

func ConCurrent() {
	//itemChan, err := persist.ItemPrint(":7788")
	//if err !=nil{
	//	return
	//}
	engine.ConCurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:  persist.SimpleItemPrint(),// itemChan ,
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
