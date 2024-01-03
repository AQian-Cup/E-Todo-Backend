package biz

import (
	"e-todo-backend/pkg/api/task"
	"e-todo-backend/pkg/model"
	"e-todo-backend/pkg/store"
)

type TaskBiz struct {
}

func (t *TaskBiz) Create(r *task.CreateRequest, userId uint) (*model.Task, error) {
	s := &store.TaskStore{}
	m, err := s.Create(r, userId)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (t *TaskBiz) Edit(r task.CreateRequest) {

}

func (t *TaskBiz) Delete(r task.CreateRequest) {

}

func (t *TaskBiz) Read(r task.CreateRequest) {

}
