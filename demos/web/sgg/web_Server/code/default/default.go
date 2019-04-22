package utils

import (
	"fmt"
	"net/http"
)
/*使用默认的多路复用器*/

/* 2）使用处理器处理请求*/

//创建处理器
type MyHandler struct{}
//处理器方法
func (mh *MyHandler)ServeHTTP(respone http.ResponseWriter, requst *http.Request)  {
	fmt.Fprintln(respone, "正在通过处理器处理你的请求")
}

func WebHandler()  {
	myHandler := MyHandler{}
	//调用处理器
	http.Handle("/",&myHandler)
	http.ListenAndServe(":8080", nil)
}