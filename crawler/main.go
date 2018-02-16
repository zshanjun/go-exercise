package main

import (
	"zshanjun/go-exercise/crawler/engine"
	"zshanjun/go-exercise/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun", ParserFunc: parser.ParseCityList})
}
