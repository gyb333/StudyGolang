package main

import (
		"crawler/worker"
	. "fmt"
	"log"
	"utils"
	"flag"
	)

var port = flag.Int("port", 9000, "the port for me to listen on")
//go run worker.go --port=9000
func main() {
	flag.Parse()
	if *port == 0 {
		log.Println("must specify a port ... ")
		return
	}

	log.Fatal(utils.ServeRpc(Sprintf(":%d", *port),
		worker.CrawlService{}))


}
