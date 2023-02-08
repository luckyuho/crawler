package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// in := make(chan Request) // type 1, 2
	out := make(chan ParseResult)
	// e.Scheduler.ConfigureMasterWorkerChan(in) // type 1, 2
	e.Scheduler.Run()

	//type 3
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	// // type 1, 2
	// for i := 0; i < e.WorkerCount; i++ {
	// 	createWorker(in, out)
	// }

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item %v", item)
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

// type 3
func createWorker(in chan Request,
	out chan ParseResult, ready ReadyNotifier,
) {
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

// // type 1, 2
// func createWorker(
// 	in chan Request, out chan ParseResult,
// ) {
// 	go func() {
// 		for {
// 			// tell scheduler i'm ready
// 			request := <-in
// 			result, err := worker(request)
// 			if err != nil {
// 				continue
// 			}
// 			out <- result
// 		}
// 	}()
// }
