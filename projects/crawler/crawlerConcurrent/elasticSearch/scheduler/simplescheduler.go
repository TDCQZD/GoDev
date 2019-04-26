package scheduler

import "GoDev/projects/crawler/crawlerConcurrent/elasticSearch/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// send request down to worker chan
	go func() { s.workerChan <- request }()
}
func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}
