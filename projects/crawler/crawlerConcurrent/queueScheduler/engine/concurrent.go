package engine

import "log"

// Scheduler接口
type Scheduler interface {
	// 简单调度器方法
	Submit(request Request)
	ConfigureMasterWorkerChan(chan Request)
	// 队列调度器增加方法
	WorkerReady(chan Request)
	Run()
}

type ConcurrentEngine struct {
	Scheduler   Scheduler //Sheduler
	WorkerCount int       //worker的数量
}

// 调度器
func (e *ConcurrentEngine) Run(seeds ...Request) {

	//worker公用一个in，out
	// in := make(chan Request)
	out := make(chan ParseResult)

	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		// createWorker(in, out) //创建worker
		createWorker(out, e.Scheduler)

	}

	//参数seeds的request，要分配任务
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	//从out中获取result，对于item就打印即可，对于request，就继续分配
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got %d  item : %v", itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

//创建worker
func createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			//需要让scheduler知道已经就绪了
			s.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

/*
// 队列调度器
func (e *ConcurrentEngine) Run(seeds ...Request) {

	//worker公用一个in，out
	//in := make(chan Request)
	out := make(chan ParseResult)

	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		//createWorker(in, out) //创建worker
		createWorker(out, e.Scheduler)
	}

	//参数seeds的request，要分配任务
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	//从out中获取result，对于item就打印即可，对于request，就继续分配
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got %d  item : %v", itemCount, item)
			// itemCount++
			go func() {
				e.ItemChan <- item
			}()
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

//创建worker
func createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			//需要让scheduler知道已经就绪了
			s.WorkerReady(in)
			request := <-in
			result, err := Worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

*/
