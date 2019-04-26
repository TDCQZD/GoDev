package persist

import (
	"GoDev/projects/crawler/crawlerDistributed/engine"
	"GoDev/projects/crawler/crawlerDistributed/persist"

	"github.com/olivere/elastic"
)

type ItemSaveService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaveService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, item, s.Index)
	if err == nil {
		*result = "ok"
	}
	return err
}
