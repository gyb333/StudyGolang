package server

import (
	"flag"
	"projects/crawler_distributed/rpcsupport"
	"projects/crawler_distributed/worker"
	"log"
	"fmt"
)

//命令行库，可以读取加入的参数
// go run worker.go --port=9000
var port = flag.Int("port",0,"the port for me to listen on")
func main() {
	flag.Parse()
	if *port == 0{
		log.Println("must specify a port")
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d",*port),worker.CrawlService{}))
}
