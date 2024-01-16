package store

import (
	"e-todo-backend/pkg/api/user"
	"e-todo-backend/pkg/db"
	"e-todo-backend/pkg/model"
	"github.com/jinzhu/copier"
)

type UserStore struct {
}

func (u *UserStore) Create(r *user.RegisterRequest) error {
	m := &model.User{}
	_ = copier.Copy(m, r)
	return db.DB.Create(m).Error
}

func (u *UserStore) Read(r *user.LoginRequest) (*model.User, error) {
	m := &model.User{}
	if err := db.DB.Where("name = ?", r.Name).First(m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (u *UserStore) ReadByUserId(userId uint) (*map[string]interface{}, error) {
	m := &model.User{}
	res := &user.ReadResponse{}
	resMap := &map[string]interface{}{}
	return resMap, db.DB.Model(m).Where("ID = ?", userId).First(res).First(resMap).Error
}
