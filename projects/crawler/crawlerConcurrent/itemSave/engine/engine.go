package engine

import (
	"GoDev/projects/crawler/crawlerConcurrent/itemSave/fetcher"
	"log"
)

// 任务队列 获取到的种子，添加到一个切片requests中，然后循环处理每个任务。
func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		//	提取出worker,提高效率
		/*
			log.Printf("Fetching %s", r.Url)
			body, err := fetcher.Fetch(r.Url)
			if err != nil {
				log.Printf("Fetcher: error fetching url %s %v", r.Url, err)
				continue
			}
			parseResult := r.ParserFunc(body)
		*/
		parseResult, err := worker1(r)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func worker1(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)

	body, err := fetcher.Fetch(r.Url)

	if err != nil {
		log.Printf("Fetcher: error fetching url %s %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
