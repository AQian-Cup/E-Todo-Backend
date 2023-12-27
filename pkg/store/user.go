package store

import (
	"e-todo-backend/pkg/api/user"
	"e-todo-backend/pkg/db"
	"e-todo-backend/pkg/model"
	"github.com/jinzhu/copier"
)

type UserStore struct {
}

func (u *UserStore) Create(r user.RegisterRequest) error {
	m := &model.User{}
	_ = copier.Copy(m, r)
	return db.DB.Create(m).Error
}

func (u *UserStore) Read(r user.LoginRequest) (*model.User, error) {
	m := &model.User{}
	if err := db.DB.Where("name = ?", r.Name).First(m).Error; err != nil {
		return nil, err
	}
	return m, nil
}
