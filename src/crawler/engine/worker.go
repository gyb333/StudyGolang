package engine

import (
	"log"
	"crawler/fetcher"
)

func Work(request Request) (ParseResult,error) {
	log.Printf("Fetching %s",request.Url)
	body,err:=fetcher.Fetch(request.Url)
	if err!=nil{
		return ParseResult{},err
	}
	//log.Printf("%s",string(body))
	return request.Parser.Parse(body),nil

}
