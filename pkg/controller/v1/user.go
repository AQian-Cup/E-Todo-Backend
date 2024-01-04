package v1

import (
	"crypto/ecdsa"
	"e-todo-backend/pkg/api/user"
	"e-todo-backend/pkg/biz"
	"e-todo-backend/pkg/errno"
	"e-todo-backend/pkg/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
}

func (u *UserController) Register(c *gin.Context) {
	r := &user.RegisterRequest{}
	if err := c.ShouldBindJSON(r); err != nil {
		response.Write(c, errno.BindError)
		return
	}
	b := &biz.UserBiz{}
	if err := b.Register(r); err != nil {
		response.Write(c, errno.InternalServerError)
		return
	}
	response.Write(c, errno.OK)
}

func (u *UserController) Login(c *gin.Context) {
	r := &user.LoginRequest{}
	if err := c.ShouldBindJSON(r); err != nil {
		response.Write(c, errno.BindError)
		return
	}
	b := &biz.UserBiz{}
	if ts, err := b.Login(r, u.PrivateKey); err != nil {
		response.Write(c, errno.InternalServerError)
		return
	} else {
		response.Write(c, &response.Response{
			HTTP: http.StatusOK,
			Result: response.OkResult{
				"token": ts,
			},
		})
		return
	}
}
