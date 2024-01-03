package store

import (
	"e-todo-backend/pkg/api/task"
	"e-todo-backend/pkg/db"
	"e-todo-backend/pkg/model"
	"github.com/jinzhu/copier"
)

type TaskStore struct {
}

func (t *TaskStore) Create(r *task.CreateRequest, userId uint) (*model.Task, error) {
	m := &model.Task{}
	_ = copier.Copy(m, r)
	m.UserId = userId
	return m, db.DB.Create(m).Error
}

func (t *TaskStore) Read() {

}

func (t *TaskStore) Update(r *task.EditRequest, userId uint) {

}
