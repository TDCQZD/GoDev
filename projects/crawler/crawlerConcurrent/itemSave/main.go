package main

import (
	"GoDev/projects/crawler/crawlerConcurrent/itemSave/engine"
	"GoDev/projects/crawler/crawlerConcurrent/itemSave/parser"
	"GoDev/projects/crawler/crawlerConcurrent/itemSave/persist"
	"GoDev/projects/crawler/crawlerConcurrent/itemSave/scheduler"
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

	// 并发模式-并发调度器配置
	e := engine.ConcurrentEngine{
		// Scheduler: &scheduler.SimpleScheduler{}, // 简单的并发调度器
		Scheduler:   &scheduler.QueuedScheduler{}, // 队列的并发调度器
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
}
