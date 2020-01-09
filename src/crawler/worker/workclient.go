package worker

import (

	"crawler/config"

	"fmt"


	"utils"
	"projects/crawler_distributed/worker"
	"projects/crawler_distributed/engine"
)

func CreateProcessor() (engine.Processor, error) {
	client, err := utils.NewClient(fmt.Sprintf(":%d", config.WorkerPort0))
	if err != nil {
		return nil, err
	}

	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult
		err := client.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResult), nil
	}, nil
}
