package user

import (
	"go_project_code/APIServer/RESTful/APIServer_api_https/src/handler"
	"go_project_code/APIServer/RESTful/APIServer_api_https/src/pkg/errno"
	"go_project_code/APIServer/RESTful/APIServer_api_https/src/service"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	handler.SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
