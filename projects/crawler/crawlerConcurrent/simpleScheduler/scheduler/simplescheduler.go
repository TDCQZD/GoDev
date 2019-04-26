package scheduler

import "GoDev/projects/crawler/crawlerConcurrent/simpleScheduler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// send request down to worker chan
	go func() { s.workerChan <- request }()
}

