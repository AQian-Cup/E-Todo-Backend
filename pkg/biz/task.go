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

func (t *TaskBiz) Edit(r *task.EditRequest, userId uint) (*model.Task, error) {
	s := &store.TaskStore{}
	m, err := s.Update(r, userId)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (t *TaskBiz) Delete(r *task.DeleteRequest, userId uint) error {
	s := &store.TaskStore{}
	if r.Id != 0 {
		return s.DeleteById(r, userId)
	} else {
		return s.DeleteAll(userId)
	}
}

func (t *TaskBiz) Read(r *task.ReadRequest, userId uint) (*model.Task, error) {
	s := &store.TaskStore{}
	m, err := s.Read(r, userId)
	if err != nil {
		return nil, err
	}
	return m, nil
}
