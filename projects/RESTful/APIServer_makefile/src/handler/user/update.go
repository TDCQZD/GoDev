package user

import (
	"go_project_code/APIServer/RESTful/APIServer_makefile/src/handler"
	"go_project_code/APIServer/RESTful/APIServer_makefile/src/model"
	"go_project_code/APIServer/RESTful/APIServer_makefile/src/pkg/errno"
	"go_project_code/APIServer/RESTful/APIServer_makefile/src/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	userId, _ := strconv.Atoi(c.Param("id"))

	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u.Id = uint64(userId)

	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	if err := u.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
