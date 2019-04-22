package main

import (
	"github.com/gin-gonic/gin"
)

func mian() {
	// Default With the Logger and Recovery middleware already attached
	// r := gin.Default()

	// blank gin
	r := gin.New()
	r.Run()
}
