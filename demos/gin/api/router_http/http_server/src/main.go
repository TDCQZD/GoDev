package main

import (
	"time"

	"github.com/gin-gonic/gin"

	"net/http"
)

func main() {
	router := gin.Default()

	// 1.默认服务器
	router.Run()
	// 2.HTTP 服务器
	// 2.1 http.ListenAndServe()
	http.ListenAndServe(":8080", router)

	// 2.2 自定义 HTTP 服务器的配置
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	// 3. HTTP 服务器替换方案 想无缝重启 sever.go
}
