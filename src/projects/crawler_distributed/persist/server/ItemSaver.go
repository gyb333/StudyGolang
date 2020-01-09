package server

import (
	"flag"
	"gopkg.in/olivere/elastic.v5"
	"projects/crawler_distributed/persist"
	"projects/crawler_distributed/config"
	"projects/crawler_distributed/rpcsupport"
	"log"
	"fmt"
)

var port = flag.Int("port",0,"the port for me to listen on")
func main() {
	flag.Parse()
	if *port == 0{
		log.Println("must specify a port")
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d",*port), config.ElasticIndex))
	//Fatal，若有异常，则挂了。panic还有recover的机会

}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}