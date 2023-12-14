package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	UserID uint
	Date   int
	Detail string
	Type   string
	Level  int
}
