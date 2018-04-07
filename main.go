package main

import (
	"crawler/zhenai/parser"
	"crawler/types"
	"crawler/engine"
)

func main() {
	seeds := types.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	}
	engine.Run(seeds)
}
