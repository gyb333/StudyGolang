package main

import (
	"crawler/engine"
	"crawler/parser"
	"crawler/scheduler"
)

var url = "https://www.zhenai.com/zhenghun"

var cityUrl = "http://www.zhenai.com/zhenghun/aba"

var profileUrl = "http://album.zhenai.com/u/1280064210"

func main() {
	engine.ConCurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 10,
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
