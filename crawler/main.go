package main

import (
	"zshanjun/go-exercise/crawler/engine"
	"zshanjun/go-exercise/crawler/zhenai/parser"
	"zshanjun/go-exercise/crawler/scheduler"
)

func main() {

	//engine.SimpleEngine{}.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun", ParserFunc: parser.ParseCityList})

	engine.ConcurrentEngine{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:100,
	}.Run(engine.Request{Url: "http://www.zhenai.com/zhenghun", ParserFunc: parser.ParseCityList})
}
