package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string
	Description string
	UserId      uint
	Type        string
	Level       uint
	Timestamp   int64
}
