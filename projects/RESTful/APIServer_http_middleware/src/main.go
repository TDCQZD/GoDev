package main

import (
	"go_project_code/APIServer/RESTful/APIServer_http_middleware/src/router/middleware"
	"errors"
	"go_project_code/APIServer/RESTful/APIServer_http_middleware/src/config"
	"go_project_code/APIServer/RESTful/APIServer_http_middleware/src/model"
	"go_project_code/APIServer/RESTful/APIServer_http_middleware/src/router"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
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

	// init db
	model.DB.Init()
	defer model.DB.Close()

	// 设置 gin 的运行模式.
	gin.SetMode(viper.GetString("runmode"))

	//创建Gin引擎。
	g := gin.New()

	

	// 加载路由
	router.Load(
		// Cores.
		g,

		// Middlwares.
		middleware.Logging(),
		middleware.RequestId(),
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

/*API 服务器健康状态自检流程：
1、在启动 HTTP 端口前 go 一个 pingServer 协程
2、启动 HTTP 端口后，该协程不断地 ping /sd/health 路径
3、如果失败次数超过一定次数，则终止 HTTP 服务器进程。
*/

// ping http服务器以确保路由器正常工作。
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
