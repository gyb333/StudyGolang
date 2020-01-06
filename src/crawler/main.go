package main

import (
	"crawler/engine"
	"crawler/parser"
	"crawler/scheduler"
	"crawler/persist"
)

var url = "https://www.zhenai.com/zhenghun"

var cityUrl = "http://www.zhenai.com/zhenghun/aba"

var profileUrl = "http://album.zhenai.com/u/1280064210"

func main() {
	engine.ConCurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:persist.ItemSaver(),
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
