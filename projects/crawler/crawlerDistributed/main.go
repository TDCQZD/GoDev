package main

import (
	"GoDev/projects/crawler/crawlerConcurrent/concurrent/engine"
	"GoDev/projects/crawler/crawlerConcurrent/concurrent/parser"
	"GoDev/projects/crawler/crawlerConcurrent/concurrent/persist"
	"GoDev/projects/crawler/crawlerConcurrent/concurrent/scheduler"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	/*
		// 单任务engine
		engine.Run(engine.Request{
			Url:        url,
			ParserFunc: parser.ParseCityList,
		})
	*/
	// itemChan, err := persist.ItemSaver()
	itemChan, err := persist.ItemSaver("datint_profile")
	if err != nil {
		panic(err)
	}
	// 并发模式-并发调度器配置
	e := engine.ConcurrentEngine{
		// Scheduler: &scheduler.SimpleScheduler{}, // 简单的并发调度器
		Scheduler:   &scheduler.QueuedScheduler{}, // 队列的并发调度器
		WorkerCount: 100,
		// ItemChan:    persist.ItemSaver(),
		ItemChan: itemChan,
	}

	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
}
