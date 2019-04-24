package main

import (
	"GoDev/projects/crawler/src/zhenai_Single/engine"
	"GoDev/projects/crawler/src/zhenai_Single/parser"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	engine.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})

}
