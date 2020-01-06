package engine

import (
	"log"
	"crawler/fetcher"
)

type SimpleEngine struct {

}

func (SimpleEngine) Run(seeds ...Request){
	var requests []Request
	requests = append(requests,seeds...)
	for len(requests)>0{
		request:=requests[0]
		requests=requests[1:]
		log.Printf("Fetching %s",request.Url)
		body,err:=fetcher.Fetch(request.Url)
		if err!=nil{
			log.Printf("Fetcher: error fetching url %s %v",request.Url,err)
			continue
		}
		//log.Printf("%s",string(body))
		parseResult:=request.Parser.Parse(body)
		requests=append(requests,parseResult.Requests...)
		for i,item:=range parseResult.Items{
			log.Printf("Got item %d %v",i,item)
		}
	}
}
