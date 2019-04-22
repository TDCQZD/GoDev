package utils

import (
	"fmt"
	"net/http"
)

func handlerURL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "发送请求的请求地址是：", r.URL.Path)
	fmt.Fprintln(w, "发送请求的请求地址的查询字符串是：", r.URL.RawQuery)

}

func URLDemo() {
	http.HandleFunc("/url", handlerURL)
	http.ListenAndServe(":8080", nil)
}
