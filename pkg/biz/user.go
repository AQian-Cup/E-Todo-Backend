package biz

import (
	"e-todo-backend/pkg/api/user"
	"e-todo-backend/pkg/store"
)

type UserBiz struct {
}

func (U *UserBiz) Register(r user.CreateRequest) error {
	s := &store.UserStore{}
	if err := s.Register(r); err != nil {
		return err
	}
	return nil
}
