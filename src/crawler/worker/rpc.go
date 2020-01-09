package worker

import "crawler/engine"

type CrawlService struct {
}

func (CrawlService) Process(req Request, result *ParseResult) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}

	engineRes, err := engine.Work(engineReq)
	if err != nil {
		return err
	}

	*result = SerializeParseResult(engineRes)
	return nil
}
