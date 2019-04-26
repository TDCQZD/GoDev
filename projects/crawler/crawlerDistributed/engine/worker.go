package engine

import (
	"GoDev/projects/crawler/crawlerConcurrent/concurrent/fetcher"
	"log"
)

func Worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)

	body, err := fetcher.Fetch(r.Url)

	if err != nil {
		log.Printf("Fetcher: error fetching url %s %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
