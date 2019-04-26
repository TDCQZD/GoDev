package model

import "GoDev/projects/crawler/crawlerConcurrent/concurrent/engine"

type SearchResult struct {
	Hits  int64
	Start int
	Items []engine.Item
}
