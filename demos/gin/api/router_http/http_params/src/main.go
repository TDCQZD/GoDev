package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	// 使用默认中间件创建一个 gin 路由：
	// 日志与恢复中间件（无崩溃）。
	router := gin.Default()

	/* 参数处理：
	path 中的参数 通过Context的Param方法来获取
	URL 参数通过 DefaultQuery 或 Query 方法获取
	表单参数通过 PostForm 方法获取
	*/
	//  http://127.0.0.1:8080/user/中国 			Hello 中国
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	}) // 这个处理器将去匹配 /user/john ， 但是它不会去匹配 /user 说明：/user/中国 中文不会乱码

	// 查询字符串参数使用现有的底层 request 对象解析。
	// http://localhost:8080/welcome  Hello Guest
	// http://localhost:8080/welcome?firstname=Jane Hello Jane
	// http://localhost:8080/welcome?lastname=Doe Hello Guest Doe
	// http://localhost:8080/welcome?firstname=Jane&lastname=Doe Hello Jane Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest") //可设置默认值
		// 是 c.Request.URL.Query().Get("lastname") 的简写
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	// Multipart/Urlencoded 表单
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		c.String(http.StatusOK, "id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	// 默认情况下，它使用：8080，除非定义了 PORT 环境变量。
	router.Run()
	// router.Run(":3000") 硬编码端口

}
