package main

import (
	"utils"
	"log"
	. "fmt"
	"crawler/config"
	"crawler/worker"
)

func main()  {
		log.Fatal(utils.ServeRpc(Sprintf(":%d", config.WorkerPort0), worker.CrawlService{}))

}
