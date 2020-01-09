package client

import (
	"net/rpc"
	"projects/crawler_distributed/worker"
	"projects/crawler_distributed/config"
	"projects/crawler_distributed/engine"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor{

	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		client := <-clientChan
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err!=nil{
			return engine.ParseResult{},err
		}
		return worker.DesrializeResult(sResult),nil
	}
}