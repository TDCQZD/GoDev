package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerFunc_Geting_(c *gin.Context) {
	c.String(http.StatusOK, "GetHandlerFunc Success !")
}
