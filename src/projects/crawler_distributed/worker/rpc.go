package worker

import "projects/crawler_distributed/engine"

type CrawlService struct{}

func (CrawlService) Process(req Request, result *ParseResult) error {
	engineReq,err := DesrializeRequest(req)
	if err!=nil{
		return err
	}

	engineResult, err := engine.Worker(engineReq)
	if err!=nil {
		return err
	}

	*result = SerializeResult(engineResult)
	return nil
}
