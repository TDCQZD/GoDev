package user

import (
	"go_project_code/APIServer/RESTful/APIServer_startup_script/src/handler"
	"go_project_code/APIServer/RESTful/APIServer_startup_script/src/model"
	"go_project_code/APIServer/RESTful/APIServer_startup_script/src/pkg/auth"
	"go_project_code/APIServer/RESTful/APIServer_startup_script/src/pkg/errno"
	"go_project_code/APIServer/RESTful/APIServer_startup_script/src/pkg/token"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	d, err := model.GetUser(u.Username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	if err := auth.Compare(d.Password, u.Password); err != nil {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	if err != nil {
		handler.SendResponse(c, errno.ErrToken, nil)
		return
	}

	handler.SendResponse(c, nil, model.Token{Token: t})

}
