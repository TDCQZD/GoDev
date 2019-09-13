package user

import (
	"fmt"
	"go_project_code/APIServer/RESTful/APIServer_http/src/handler"
	"go_project_code/APIServer/RESTful/APIServer_http/src/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

/*说明：
window:	handler.SendResponse(c, errno.ErrBind, nil)
Linux:	SendResponse(c, errno.ErrBind, nil)
*/

// Create creates a new user account.
func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	admin2 := c.Param("username")
	log.Infof("URL username: %s", admin2)

	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username is: [%s], password is [%s]", r.Username, r.Password)
	if r.Username == "" {
		handler.SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
		return
	}

	if r.Password == "" {
		handler.SendResponse(c, fmt.Errorf("password is empty"), nil)
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	handler.SendResponse(c, nil, rsp)
}
