package utils

import (
	"fmt"
	"net/http"
)

func handlerBody(w http.ResponseWriter, r *http.Request) {
	//获取请求体重内容的长度
	len := r.ContentLength
	//创建byte切片
	body := make([]byte, len)
	//将请求体中的内容读到body中
	r.Body.Read(body)
	//在浏览器中显示请求体中的内容
	fmt.Fprintln(w, "请求体中的内容有：", string(body))

}

func BodyDemo() {
	http.HandleFunc("/body", handlerBody)
	http.ListenAndServe(":8080", nil)
}
