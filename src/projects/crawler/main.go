package main

import (
	"./engine"
	"./zhenai/parser"
	)

func main()  {
	// 1. SimpleEngine
	// 网络利用率：50-70k/s
	engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
