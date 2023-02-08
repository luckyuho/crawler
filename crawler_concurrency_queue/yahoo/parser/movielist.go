package parser

import (
	"crawler/engine"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseMovieList(
	contents string) engine.ParseResult {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(contents))
	if err != nil {
		log.Fatalln(err)
	}

	result := engine.ParseResult{}
	dom.Find("div[class=release_info_text]>div[class=release_movie_name]>a").Each(func(i int, selection *goquery.Selection) {
		name := strings.Replace(selection.Text(), " ", "", -1)
		url, _ := selection.Attr("href")

		result.Items = append(
			result.Items, name)
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        url,
				ParserFunc: ParseMovieDetail,
			})
	})

	return result
}
