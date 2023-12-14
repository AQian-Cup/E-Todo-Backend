package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	Nickname string
	Tasks    []Task
}

func (u *User) BeforeCreate(_ *gorm.DB) error {
	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(password)
	return nil
}
