package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	someGroup := router.Group("/someGroup")
	{
		someGroup.GET("/someGet", func(c *gin.Context) {
			c.String(http.StatusOK, "Grouping routes ")
		})

	}
	/*
		// 简单组： v1
		v1 := router.Group("/v1")
		{
			v1.POST("/login", loginEndpoint)
			v1.POST("/submit", submitEndpoint)
			v1.POST("/read", readEndpoint)
		}

		// 简单组： v2
		v2 := router.Group("/v2")
		{
			v2.POST("/login", loginEndpoint)
			v2.POST("/submit", submitEndpoint)
			v2.POST("/read", readEndpoint)
		}
	*/
	router.Run()
}
