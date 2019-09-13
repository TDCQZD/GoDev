package util

import (
	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
)

func GenShortId() (string, error) {
	return shortid.Generate()
}
func GetReqID(c *gin.Context) string {
	v, ok := c.Get("X-Request-Idquest")
	if !ok {
		return ""
	}
	if requestId, ok := v.(string); ok {
		return requestId
	}
	return ""
}
