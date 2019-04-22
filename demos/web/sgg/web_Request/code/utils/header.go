package utils


import (
	"fmt"
	"net/http"
)

func handlerHeader(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "发送请求的请求头信息：", r.Header)
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息是：", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头中Accept-Encoding的属性值是：", r.Header.Get("Accept-Encoding"))
	
	
}

func HeaderDemo() {
	http.HandleFunc("/header", handlerHeader)
	http.ListenAndServe(":8080", nil)
}
