package main

import (
	"crawler/engine"
	"crawler/yahoo/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://movies.yahoo.com.tw/movie_intheaters.html",
		ParserFunc: parser.ParseMovieList,
	})
}
