package engine

import (
	"crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		log.Printf("Fetching %s", r.Url)
		requests = requests[1:]

		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error "+
				"fetching url %s: %v",
				r.Url, err)
			continue
		}

		parseResult := r.ParserFunc(body)
		requests = append(requests,
			parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
