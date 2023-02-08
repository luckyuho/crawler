package main

import (
	"crawler/engine"
	"crawler/scheduler"
	"crawler/yahoo/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 3,
	}

	e.Run(engine.Request{
		Url:        "https://movies.yahoo.com.tw/movie_intheaters.html",
		ParserFunc: parser.ParseMovieList,
	})
}
