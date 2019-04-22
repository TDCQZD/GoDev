package utils

import (
	"fmt"
	"net/http"
)
/*使用默认的多路复用器*/

/* 1）使用处理器函数处理请求*/

//创建处理器函数
func handler(respone http.ResponseWriter, requst *http.Request)  {
	fmt.Fprintln(respone, "正在通过处理器函数处理你的请求")
}

func WebHandlerFunc()  {
	http.HandleFunc("/", handler)

	//创建路由
	http.ListenAndServe(":8080", nil)
}