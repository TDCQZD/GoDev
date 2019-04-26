package persist

import (
	"GoDev/projects/crawler/crawlerConcurrent/elasticSearch/engine"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/olivere/elastic"
)

// 存储Item 存储爬取的数据
// func ItemSaver() chan interface{} {
func ItemSaver() chan engine.Item {
	// out := make(chan interface{})
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got %d  item : %v", itemCount, item)
			itemCount++

			err := save(item) //保存item
			if err != nil {
				log.Printf("Item Saver :error saving item %v : %v ", item, err)
			}
		}
	}()
	return out

}

/*
//保存item
func save(item interface{}) (id string, err error) {
	//关闭内网的sniff
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		// log.Println(err)
		return "", err
	}

	resp, err := client.Index(). //存储数据，可以添加或者修改，要看id是否存在
					Index("datint_profile").
					Type("zhenai").
					BodyJson(item).
					Do(context.Background())

	if err != nil {
		// log.Println(err)
		return "", err
	}

	fmt.Printf("%+v", resp) //格式化输出结构体对象的时候包含了字段名称
	return resp.Id, nil
}
*/
//保存item-添加Url和Id
func save(item engine.Item) error {
	//关闭内网的sniff
	fmt.Printf("save..%+v\n", item)
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		//log.Println(err)
		return err
	}

	if item.Type == "" {
		return errors.New("must supply Type ..")
	}

	indexService := client.Index(). //存储数据，可以添加或者修改，要看id是否存在
					Index("datint_profile").
		//Type("zhenai").
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.Do(context.Background())

	if err != nil {
		return err
	}

	//fmt.Printf("%+v",resp)//格式化输出结构体对象的时候包含了字段名称
	return nil
}
