package controller

import (
	"GoDev/projects/crawler/crawlerConcurrent/concurrent/frontend/model"
	"GoDev/projects/crawler/crawlerConcurrent/concurrent/frontend/view"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/olivere/elastic"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

//localhost:9527/search?q=女&from=10
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))

	from, err := strconv.Atoi(req.FormValue("from"))
	if err != nil {
		from = 0
	}
	//fmt.Fprintf(w, "q=%s, from=%d", q, from)

	fmt.Printf("q:%s, form:%d\n", q, from)
	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.view.Render(w, page)
	fmt.Println("....", err)
	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := h.client.Search("datint_profile").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}
	fmt.Println("resp-->", resp)
	result.Hits = resp.TotalHits()
	result.Start = from

	// result.Items = resp.Each(reflect.TypeOf(engine.Item{}))

	//for _, v := range result.Items {
	//  fmt.Printf("%+v\n", v)
	//}

	/*
	   if resp.Hits.TotalHits >0{
	       for _,hit:=range resp.Hits.Hits{

	           var item model2.Profile
	           err :=json.Unmarshal(*hit.Source,&item)
	           if err != nil{
	               panic(err)
	           }
	           fmt.Printf("%s\n",*hit.Source)
	           fmt.Printf("%+v\n",item)
	           result.Items = append(result.Items,item)
	       }
	   }
	*/

	/*
	   itemRaw :=resp.Each(reflect.TypeOf(engine.Item{}))
	   fmt.Println("len-->",len(itemRaw))
	   for _, v := range itemRaw {
	       item := v.(engine.Item)
	       //fmt.Printf("%+v\n", v)
	       result.Items = append(result.Items, item)
	   }

	*/

	return result, nil

}
func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}
