package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的没有任何中间件的路由
	r := gin.New()

	// 全局中间件
	// Logger 中间件将写日志到 gin.DefaultWriter ,即使你设置 GIN_MODE=release 。
	// 默认 gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery 中间件从任何 panic 恢复，如果出现 panic，它会写一个 500 错误。
	r.Use(gin.Recovery())

	// 每个路由的中间件, 你能添加任意数量的中间件
	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// 授权组
	// authorized := r.Group("/", AuthRequired())
	// 也可以这样:
	authorized := r.Group("/")
	// 每个组的中间件！ 在这个实例中，我们只需要在 "authorized" 组中
	// 使用自定义创建的 AuthRequired() 中间件
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	// 监听并服务于 0.0.0.0:8080
	r.Run(":8080")
}
