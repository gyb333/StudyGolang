package engine

import (
	"log"
	"../fetcher"
	)

// 从Run中分离出worker，为了让多个worker并发执行
func (SimpleEngine) worker(r Request) (ParseResult, error) {
	//log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error "+"fetching url %s: %v",
			r.Url, err)
		return ParseResult{}, err
	}
	//fmt.Printf("%s \n", body)
	return r.ParserFunc(body, r.Url), nil
	//return NilParser(body),nil
}