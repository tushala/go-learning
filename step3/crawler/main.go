package main

import (
	"go-learning/step3/crawler/engine"
	"go-learning/step3/crawler/zhenai/parser"
)


func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})


}
