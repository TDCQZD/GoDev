package middleware

import (
	"go_project_code/APIServer/RESTful/APIServer_makefile/src/handler"
	"go_project_code/APIServer/RESTful/APIServer_makefile/src/pkg/errno"
	"go_project_code/APIServer/RESTful/APIServer_makefile/src/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
