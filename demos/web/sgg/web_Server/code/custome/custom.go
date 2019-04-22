package utils

import (
	"time"
	"fmt"
	"net/http"
)
/*使用自己创建的多路复用器*/

//创建处理器
type MyHandler struct{}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "通过详细配置服务器的信息来处理请求！")
}

func main() {

	myHandler := MyHandler{}


	//创建Server结构，并详细配置里面的字段
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler,
		ReadTimeout: 2 * time.Second,
	}


	server.ListenAndServe()

}