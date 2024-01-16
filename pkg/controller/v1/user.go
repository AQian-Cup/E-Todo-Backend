package v1

import (
	"crypto/ecdsa"
	"e-todo-backend/pkg/api/user"
	"e-todo-backend/pkg/biz"
	"e-todo-backend/pkg/errno"
	"e-todo-backend/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
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
	} else {
		response.Write(c, errno.OK)
	}
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

func (u *UserController) ReadCurrent(c *gin.Context) {
	b := &biz.UserBiz{}
	originalUserId, _ := c.Get("userId")
	userId, _ := originalUserId.(uint)
	if m, err := b.ReadCurrent(userId); err != nil {
		response.Write(c, errno.InternalServerError)
		return
	} else {
		okResult := &response.OkResult{}
		_ = copier.Copy(okResult, m)
		response.Write(c, &response.Response{
			HTTP:   http.StatusOK,
			Result: okResult,
		})
	}
}
