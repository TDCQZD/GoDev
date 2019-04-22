package main

import (
	"fmt"
	"net/http"
)

//创建处理器函数
func handler(respone http.ResponseWriter ,request *http.Request){
fmt.Fprintln(respone,"--",request.URL.Path,"--",request.URL)

}

func main()  {
	http.HandleFunc("/", handler)

	//创建路由
	http.ListenAndServe(":8080",nil)
}