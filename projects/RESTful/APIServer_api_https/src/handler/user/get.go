package user

import (
	"go_project_code/APIServer/RESTful/APIServer_api_https/src/handler"
	"go_project_code/APIServer/RESTful/APIServer_api_https/src/model"
	"go_project_code/APIServer/RESTful/APIServer_api_https/src/pkg/errno"

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
