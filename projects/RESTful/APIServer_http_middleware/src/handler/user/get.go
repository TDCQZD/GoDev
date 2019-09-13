package user

import (
	"go_project_code/APIServer/RESTful/APIServer_http_middleware/src/handler"
	"go_project_code/APIServer/RESTful/APIServer_http_middleware/src/model"
	"go_project_code/APIServer/RESTful/APIServer_http_middleware/src/pkg/errno"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	username := c.Param("username")
	user, err := model.GetUser(username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	handler.SendResponse(c, nil, user)
}
