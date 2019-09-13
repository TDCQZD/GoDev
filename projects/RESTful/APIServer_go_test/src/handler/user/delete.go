package user

import (
	"go_project_code/APIServer/RESTful/APIServer_startup_script/src/handler"
	"go_project_code/APIServer/RESTful/APIServer_startup_script/src/model"
	"go_project_code/APIServer/RESTful/APIServer_startup_script/src/pkg/errno"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	err := model.DeleteUser(uint64(userId))
	if err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}