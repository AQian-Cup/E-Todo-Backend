package biz

import (
	"crypto/ecdsa"
	"e-todo-backend/pkg/api/user"
	"e-todo-backend/pkg/jwt"
	"e-todo-backend/pkg/store"
	"golang.org/x/crypto/bcrypt"
)

type UserBiz struct {
}

func (u *UserBiz) Register(r user.RegisterRequest) error {
	s := &store.UserStore{}
	if err := s.Create(r); err != nil {
		return err
	}
	return nil
}

func (u *UserBiz) Login(r user.LoginRequest, key *ecdsa.PrivateKey) (string, error) {
	s := &store.UserStore{}
	m, err := s.Read(r)
	if err != nil {
		println(err.Error(), 1)
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(r.Password))
	if err != nil {
		println(err.Error(), 2)
		return "", err
	}
	var ts string
	ts, err = jwt.Sign(r.Name, key)
	if err != nil {
		println(err.Error(), 3)
		return "", err
	}
	return ts, nil
}
