package router

import (
	"go_project_code/APIServer/RESTful/APIServer_config/src/handler/sd"
	"go_project_code/APIServer/RESTful/APIServer_config/src/router/middleware"

	"net/http"

	"github.com/gin-gonic/gin"
)

/*Load 路由加载器*/
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 设置 HTTP Header
	g.Use(gin.Recovery()) //恢复 API 服务器
	g.Use(middleware.Nocache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
