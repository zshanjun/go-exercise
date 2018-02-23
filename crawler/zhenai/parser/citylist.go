package parser

import (
	"zshanjun/go-exercise/crawler/engine"
	"regexp"
	"log"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z]+)"[^>]*>([^<]*)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	results := engine.ParseResult{}
	itemCount := 0
	for _, m := range matches {
		itemCount++
		results.Items = append(results.Items, "City " + string(m[2]))
		results.Requests = append(results.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	log.Printf("City Count: %d", itemCount)
	return results
}
