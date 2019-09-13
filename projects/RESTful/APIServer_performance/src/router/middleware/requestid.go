package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get("X-Request-Id")

		if requestId == "" {
			u4, _ := uuid.NewV4()
			requestId = u4.String()
		}
		c.Set("X-Request-Id", requestId)

		c.Writer.Header().Set("X-Request-Id", requestId)
		c.Next()
	}
}
