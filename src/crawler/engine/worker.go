package engine

import (
	"log"
	"crawler/fetcher"
)

type Worker struct {

}

func (Worker)FetchRequest(request Request) (ParseResult,error) {
	return  FetchWork(request)

}


func FetchWork(request Request) (ParseResult,error) {
	log.Printf("Fetching %s",request.Url)
	body,err:=fetcher.Fetch(request.Url)
	if err!=nil{
		return ParseResult{},err
	}
	//log.Printf("%s",string(body))
	return request.Parser.Parse(body),nil

}