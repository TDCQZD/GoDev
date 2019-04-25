package main

import (
	"GoDev/projects/crawler/crawlerSingle/engine"
	"GoDev/projects/crawler/crawlerSingle/parser"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	engine.Run(engine.Request{
		Url:        url,
		ParserFunc: parser.ParseCityList,
	})

}
