package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerFunc_Get(c *gin.Context) {
	c.String(http.StatusOK, "Get Success GetHandlerFunc!")
}
func HandlerFunc_Get_Parameters(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}
func HandlerFunc_Get_Parameters1(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + "  " + action
	c.String(http.StatusOK, message)
}

func HandlerFunc_Post(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")

	fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
}
func HandlerFunc_Put(c *gin.Context) {
	c.String(http.StatusOK, "Get Success GetHandlerFunc!")
}
func HandlerFunc_Delete(c *gin.Context) {
	c.String(http.StatusOK, "Get Success GetHandlerFunc!")
}
func HandlerFunc_Patch(c *gin.Context) {
	c.String(http.StatusOK, "Get Success GetHandlerFunc!")
}
func HandlerFunc_Head(c *gin.Context) {
	c.String(http.StatusOK, "Get Success GetHandlerFunc!")
}
func HandlerFunc_Options(c *gin.Context) {
	c.String(http.StatusOK, "Get Success GetHandlerFunc!")
}

func main() {

	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	// 使用默认中间件创建一个 gin 路由：
	// 日志与恢复中间件（无崩溃）。
	router := gin.Default()

	/*GET*/
	router.GET("/someGet", func(c *gin.Context) {
		c.String(http.StatusOK, "Get Success!")
	})
	router.GET("/get", HandlerFunc_Get)
	/*跨包访问 有Bug：No find fies*/
	// router.GET("/getp", HandlerFunc_Geting_)

	// path 中的参数
	router.GET("/user/:name", HandlerFunc_Get_Parameters) // 这个处理器将去匹配 /user/john ， 但是它不会去匹配 /user 说明：/user/中国 中文不会乱码
	// http://127.0.0.1:8080/user/中国 			Hello 中国
	router.GET("/user/:name/*action", HandlerFunc_Get_Parameters1) //这个处理器将会匹配 /user/john 并且也会匹配 /user/john/send,如果没有其他的路由匹配 /user/john,它将重定向到 /user/john/
	// http://127.0.0.1:8080/user/中国/北京  	中国  /北京

	// 查询字符串参数
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		// 这个是 c.Request.URL.Query().Get("lastname") 的快捷方式。
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	// http://127.0.0.1:8080/welcome  Hello Guest 

	/*Post*/
	router.POST("/somePost", HandlerFunc_Post)

	// router.PUT("/somePut", HandlerFunc_Put)
	// router.DELETE("/someDelete", HandlerFunc_Delete)
	// router.PATCH("/somePatch", HandlerFunc_Patch)
	// router.HEAD("/someHead", HandlerFunc_Head)
	// router.OPTIONS("/someOptions", HandlerFunc_Options)

	// 默认情况下，它使用：8080，除非定义了 PORT 环境变量。
	router.Run()
	// router.Run(":3000") 硬编码端口

}
