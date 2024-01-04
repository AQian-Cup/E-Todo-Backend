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

func (t *TaskStore) Read(r *task.ReadRequest, userId uint) (*model.Task, error) {
	m := &model.Task{
		Id:     r.Id,
		UserId: userId,
	}
	return m, db.DB.First(m).Error
}

func (t *TaskStore) Update(r *task.EditRequest, userId uint) (*model.Task, error) {
	m, err := t.Read(&task.ReadRequest{Id: r.Id}, userId)
	if err != nil {
		return nil, err
	}
	err = copier.CopyWithOption(m, r, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return nil, err
	}
	return m, db.DB.Save(m).Error
}

func (t *TaskStore) DeleteById(r *task.DeleteRequest, userId uint) error {
	m := &model.Task{
		Id:     r.Id,
		UserId: userId,
	}
	return db.DB.Delete(m).Error
}

func (t *TaskStore) DeleteAll(userId uint) error {
	m := &model.Task{
		UserId: userId,
	}
	return db.DB.Where("user_id = ?", userId).Delete(m).Error
}
