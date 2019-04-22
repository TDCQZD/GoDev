package utils

import (
	"fmt"
	"net/http"
)
/*使用自己创建的多路复用器*/

//创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	
	fmt.Fprintln(w, "通过自己创建的多路复用器处理请求!", r.URL.Path)
}

func WebCustom() {
	//创建多路复用器
	mux := http.NewServeMux()

	// http.HandleFunc("/", handler)
	mux.HandleFunc("/", handler)

	//创建路由
	// http.ListenAndServe(":8080", nil)
	http.ListenAndServe(":8080", mux)
}