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

func (u *UserBiz) Register(r *user.RegisterRequest) error {
	s := &store.UserStore{}
	if err := s.Create(r); err != nil {
		return err
	}
	return nil
}

func (u *UserBiz) Login(r *user.LoginRequest, key *ecdsa.PrivateKey) (string, error) {
	s := &store.UserStore{}
	m, err := s.Read(r)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(r.Password))
	if err != nil {
		return "", err
	}
	var ts string
	ts, err = jwt.Sign(m.ID, key)
	if err != nil {
		return "", err
	}
	return ts, nil
}
