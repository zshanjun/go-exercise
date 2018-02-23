package scheduler

import (
	"zshanjun/go-exercise/crawler/engine"
)

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigMasterWorkerChan(r chan engine.Request) {
	s.workerChan = r
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		s.workerChan <- request
	}()
}

