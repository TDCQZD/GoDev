package main

import (
	"errors"
	"go_project_code/APIServer/RESTful/APIserver/src/router"

	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	//创建Gin引擎。
	g := gin.New()
	// gin 中间件
	middlewares := []gin.HandlerFunc{}
	// 加载路由
	router.Load(
		g,
		middlewares...,
	)
	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())
}

/*API 服务器健康状态自检流程：
1、在启动 HTTP 端口前 go 一个 pingServer 协程
2、启动 HTTP 端口后，该协程不断地 ping /sd/health 路径
3、如果失败次数超过一定次数，则终止 HTTP 服务器进程。
*/

// ping http服务器以确保路由器正常工作。
func pingServer() error {

	for i := 0; i < 10; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
