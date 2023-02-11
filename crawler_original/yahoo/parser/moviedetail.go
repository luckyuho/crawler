package parser

import (
	"crawler/engine"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseMovieDetail(
	contents string) engine.ParseResult {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(contents))
	if err != nil {
		log.Fatalln(err)
	}

	result := engine.ParseResult{}
	dom.Find("div[class=movie_intro_info_r]").Each(func(i int, selection *goquery.Selection) {
		name := strings.Replace(selection.Text(), " ", "", -1)
		name = organize_text(name)

		result.Items = append(
			result.Items, name)
		result.Requests = append(
			result.Requests, engine.Request{
				Url:        "",
				ParserFunc: engine.NilParser,
			})
	})

	return result
}

func organize_text(name string) string {
	special_char := "*"
	name = strings.Replace(name, "、", "", -1)
	name = strings.Replace(name, "\n", special_char, -1)
	name = strings.Replace(name, "****", special_char, -1)
	name = strings.Replace(name, "***", special_char, -1)
	name = strings.Replace(name, "**", special_char, -1)
	name = strings.Replace(name, "：*", ":", -1)
	// name = strings.Replace(name, "*", special_char+"\n", -1)
	name = strings.Replace(name, "*", "\n", -1)

	return name
}
