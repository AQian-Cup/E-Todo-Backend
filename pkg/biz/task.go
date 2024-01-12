package biz

import (
	"e-todo-backend/pkg/api/task"
	"e-todo-backend/pkg/store"
	"time"
)

type TaskBiz struct {
}

func (t *TaskBiz) Create(r *task.CreateRequest, userId uint) (*task.CreateResponse, error) {
	s := &store.TaskStore{}
	m, err := s.Create(r, userId)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (t *TaskBiz) Edit(r *task.EditRequest, userId uint) (*task.EditResponse, error) {
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

func (t *TaskBiz) Read(r *task.ReadRequest, userId uint) (*task.ReadResponse, error) {
	s := &store.TaskStore{}
	m, err := s.ReadById(r, userId)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (t *TaskBiz) ReadList(r *task.ReadRequest, userId uint) (*[]task.ReadResponse, error) {
	s := &store.TaskStore{}
	if r.Year != 0 || r.Month != 0 {
		if r.Year == 0 {
			r.Year = time.Now().Year()
		}
		if r.Month == 0 {
			r.Month = int(time.Now().Month())
		}
		t := time.Date(r.Year, time.Month(r.Month), 1, 0, 0, 0, 0, time.UTC)
		startTimestamp := t.Unix()
		endTimestamp := t.AddDate(0, 1, 0).Unix()
		m, err := s.ReadByTimestampScope(startTimestamp, endTimestamp, userId)
		if err != nil {
			return nil, err
		}
		return m, err
	}
	m, err := s.ReadAll(userId)
	if err != nil {
		return nil, err
	}
	return m, err
}
