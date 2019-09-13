package main

import (
	"GoDev/projects/chat_gin/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home 首页
func Home() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	}
}

// HttpConn HTTP连接方式：websocket 轮询 刷新
func HttpConn() gin.HandlerFunc {
	return func(c *gin.Context) {
		httpConn := c.Query("httpconn")
		user := c.Query("user")

		switch httpConn {
		case "refresh":
			c.Redirect(http.StatusMovedPermanently, "/refresh?user="+user)
		case "longpolling":
			c.Redirect(http.StatusMovedPermanently, "/longpolling/room?user="+user)
		case "websocket":
			c.Redirect(http.StatusMovedPermanently, "/websocket/room?user="+user)
		default:
			c.Redirect(http.StatusMovedPermanently, "/websocket/room?user="+user)

		}
	}
}

func main() {
	r := gin.Default()
	r.GET("/", Home())
	r.GET("/httpconn", HttpConn())

	//设置路由
	routes.WebSocket(r)
	routes.LongPolling(r)
	routes.Refresh(r)

	// 设置静态资源
	r.LoadHTMLGlob("./templates/*")
	r.Static("/static", "./static")

	r.Run()
}
