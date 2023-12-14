package biz

import (
	"e-todo-backend/pkg/api/user"
	"e-todo-backend/pkg/errno"
	"e-todo-backend/pkg/response"
	"e-todo-backend/pkg/store"
	"github.com/gin-gonic/gin"
)

type UserBiz struct {
}

func (U *UserBiz) Register(c *gin.Context, r user.CreateRequest) error {
	s := &store.UserStore{}
	if err := s.Register(r); err != nil {
		response.Write(c, errno.InternalServerError)
		return err
	}
	return nil
}
