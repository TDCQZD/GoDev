package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

func main() {
	// c := colly.NewCollector()
	c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))

	// 在请求之前调用
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	// 如果在请求期间发生错误，则调用
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	// 收到回复后调用
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})
	// 如果收到的内容是HTML，则在OnResponse之后立即调用
	c.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
		fmt.Println("First column of a table row:", e.Text)
	})
	// 如果收到的内容是HTML或XML，则在OnHTML之后立即调用
	c.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println(e.Text)
	})
	//在OnXML回调之后调用
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("http://go-colly.org/")
}
