package router

import (
	"go_project_code/APIServer/RESTful/APIServer_startup_script/src/handler/sd"
	"go_project_code/APIServer/RESTful/APIServer_startup_script/src/handler/user"
	"go_project_code/APIServer/RESTful/APIServer_startup_script/src/router/middleware"
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
	g.POST("/login", user.Login)
	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
	}

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
