package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// 禁用控制台颜色
	// gin.DisableConsoleColor()

	// 使用默认中间件创建一个 gin 路由：
	// 日志与恢复中间件（无崩溃）。
	router := gin.Default()

	/*单文件*/
	// 为 multipart 表单设置一个较低的内存限制（默认是 32 MiB）
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		// 上传文件到指定的 dst 。
		// c.SaveUploadedFile(file, dst)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	/*多文件*/
	// 为 multipart 表单设置一个较低的内存限制（默认是 32 MiB）
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/uploadpart", func(c *gin.Context) {

		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["files"]

		for _, file := range files {
			log.Println(file.Filename)

			// 上传文件到指定的 dst.
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	// 默认情况下，它使用：8080，除非定义了 PORT 环境变量。
	router.Run()
	// router.Run(":3000") 硬编码端口

}
