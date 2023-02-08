package scheduler

import "crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(
	c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(
	w chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
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
// 		s.workerChan <- r
// }
