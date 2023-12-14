package v1

import (
	"e-todo-backend/pkg/api/user"
	"e-todo-backend/pkg/biz"
	"e-todo-backend/pkg/errno"
	"e-todo-backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (U *UserController) Register(c *gin.Context) {
	var r user.CreateRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		response.Write(c, errno.BindError)
		return
	}
	b := &biz.UserBiz{}
	if err := b.Register(c, r); err != nil {
		response.Write(c, errno.InternalServerError)
		return
	}
	response.Write(c, errno.OK)
}
