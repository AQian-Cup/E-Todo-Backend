package store

import (
	"e-todo-backend/pkg/api/user"
	"e-todo-backend/pkg/db"
	"e-todo-backend/pkg/model"
	"github.com/jinzhu/copier"
)

type UserStore struct {
}

func (U *UserStore) Register(r user.CreateRequest) error {
	m := &model.User{}
	_ = copier.Copy(m, r)
	return db.DB.Create(m).Error
}
