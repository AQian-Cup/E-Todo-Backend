package db

import (
	"e-todo-backend/pkg/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgreSQLOptions struct {
	Host     string
	User     string
	Password string
	Database string
	Port     int
}

var DB *gorm.DB
var err error

func (o *PostgreSQLOptions) DB() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(o.DSN()), &gorm.Config{})
}

func (o *PostgreSQLOptions) DSN() string {
	return fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%d`, o.Host, o.User, o.Password, o.Database, o.Port)
}

func (o *PostgreSQLOptions) Init() error {
	DB, err = o.DB()
	if err != nil {
		return err
	}
	if err := DB.AutoMigrate(&model.User{}); err != nil {
		return err
	}
	return nil
}
