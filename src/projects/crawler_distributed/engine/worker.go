package engine

import (
	"log"
	"projects/crawler_distributed/fetcher"
)

func Worker(r Request) (ParseResult, error) {
	log.Printf("fetching %s",r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err!=nil{
		log.Printf("fetcher:error fetching url %s: %v", r.Url,err)
		return ParseResult{},err
	}
	return r.Parser.Parser(body,r.Url),nil
}
