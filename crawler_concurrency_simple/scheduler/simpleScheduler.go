package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(
	c chan engine.Request) {
	s.workerChan = c
}

// type 2
func (s *SimpleScheduler) Submit(
	r engine.Request) {
	go func() {
		s.workerChan <- r
	}()
}

// // type 1
// func (s *SimpleScheduler) Submit(
// 	r engine.Request) {
// 	s.workerChan <- r
// }
