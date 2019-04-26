package scheduler

import "GoDev/projects/crawler/crawlerConcurrent/queueScheduler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) Submit(request engine.Request) {
	s.requestChan <- request
}

func (s *QueuedScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
}

// 负责接收Request
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}
func (s *QueuedScheduler) Run() {
	// 传递Request的worker
	s.workerChan = make(chan chan engine.Request)
	// 传递Request
	s.requestChan = make(chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			}
			select {
			case r := <-s.requestChan: // 如果有request来了就加入到Request队列
				requestQ = append(requestQ, r)
			case w := <-s.workerChan: // 如果有worker来了就加入到Worker队列
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest: // 如果同时有requst和worker，我们就把这个request发给这个worker
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
