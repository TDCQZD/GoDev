package view

import (
	"GoDev/projects/crawler/crawlerConcurrent/concurrent/engine"
	"GoDev/projects/crawler/crawlerConcurrent/concurrent/frontend/model"
	model2 "GoDev/projects/crawler/crawlerConcurrent/concurrent/model"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	//template := template.Must(template.ParseFiles("template.html"))

	view := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")

	page := model.SearchResult{}

	page.Hits = 123
	//page.Start = 0

	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1214814888",
		Type: "zhenai",
		Id:   "1214814888",
		Payload: model2.Profile{
			Name:       "林YY",
			Marriage:   "未婚",
			Age:        "26岁",
			Xingzuo:    "魔羯座(12.22-01.19)",
			Height:     "165cm",
			Weight:     "50kg",
			Income:     "月收入:5-8千",
			Occupation: "职业技术教师",
			Education:  "高中及以下",
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	//err = template.Execute(out, page)
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
