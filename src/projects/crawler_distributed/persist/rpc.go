package persist

import (
	"gopkg.in/olivere/elastic.v5"
	"log"
	"projects/crawler_distributed/engine"

)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.",item)
	if err == nil{
		*result="ok"
	}else {
		log.Printf("Error saving item %v:%v",item,err)
	}
	return err
}
