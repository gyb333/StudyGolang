package main

import (
	"crawler/engine"
	"crawler/parser"
	"crawler/persist"
	"crawler/scheduler"
)

var url = "https://www.zhenai.com/zhenghun"

var cityUrl = "http://www.zhenai.com/zhenghun/aba"

var profileUrl = "http://album.zhenai.com/u/1280064210"

func main() {
	//Simple()
	ConCurrent()
}

func ConCurrent() {
	engine.ConCurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemSaver(),
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
		ItemChan: persist.ItemSaver(),
	}.Run(
		//engine.Request{
		//	Url:    url,
		//	Parser: parser.CityListParser{},
		//},
		engine.Request{
			Url:    cityUrl,
			Parser: parser.CityParser{},
		},
		//engine.Request{
		//	Url:    profileUrl,
		//	Parser: parser.ProfileParser{},
		//},
	)
}
