package main

import (
	_"fmt"
	"errors"
	"go_project_code/APIServer/RESTful/APIServer_config/src/config"
	"go_project_code/APIServer/RESTful/APIServer_config/src/router"

	"log"
	"net/http"
	"time"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {

	pflag.Parse()

	// 初始化配置文件
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	/*热更新测试代码*/
	// for {
	// 	fmt.Println(viper.GetString("runmode"))
	// 	time.Sleep(4 * time.Second)
	// }

	// 设置 gin 的运行模式.
	gin.SetMode(viper.GetString("runmode"))

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

	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

/*API 服务器健康状态自检流程：
1、在启动 HTTP 端口前 go 一个 pingServer 协程
2、启动 HTTP 端口后，该协程不断地 ping /sd/health 路径
3、如果失败次数超过一定次数，则终止 HTTP 服务器进程。
*/

// ping http服务器以确保路由器正常工作。
func pingServer() error {

	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		//1秒后继续下一次ping。
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
