package router

import (
	"go_project_code/APIServer/RESTful/APIServer_log/src/router/middleware"
	"go_project_code/APIServer/RESTful/APIserver_log/src/handler/sd"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	// middlewares
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	//404 Handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// API健康检查处理程序
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)

	}

	return g
}
