package main

import (
	"GoDev/projects/crawler/crawlerConcurrent/simpleScheduler/engine"
	"GoDev/projects/crawler/crawlerConcurrent/simpleScheduler/parser"
	"GoDev/projects/crawler/crawlerConcurrent/simpleScheduler/scheduler"
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

	// 简单调度器
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})
}
