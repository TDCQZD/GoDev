package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	res, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {

		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})

}

func main() {

	// ExampleScrape()
	newDocment()
	// select_ID()
	// select_class()

}
func newDocment() {
	doc, err := goquery.NewDocument("http://studygolang.com/topics")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".topics .topic").Each(func(i int, contentSelection *goquery.Selection) {
		title := contentSelection.Find(".title a").Text()
		log.Println("第", i+1, "个帖子的标题：", title)
	})
}

// 基于HTML Element 元素的选择器
func select_Element() {
	html := `<body>

				<div id="div1">DIV1</div>
				<div>DIV2</div>
				<span>SPAN</span>

			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("div").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}

// ID 选择器
func select_ID() {
	html := `<body>

				<div id="div1">DIV1</div>
				<div>DIV2</div>
				<span>SPAN</span>

			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("#div1").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())

	})
	// Element ID 选择器
	dom.Find("div#div1").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())

	})
}

// Class选择器
func select_class() {
	html := `<body>

				<div id="div1">DIV1</div>
				<div class="name">DIV2</div>
				<span>SPAN</span>

			</body>
			`

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find(".name").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
	// Element Class 选择器
	dom.Find("div.name").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}
