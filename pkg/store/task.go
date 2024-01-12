package store

import (
	"e-todo-backend/pkg/api/task"
	"e-todo-backend/pkg/db"
	"e-todo-backend/pkg/model"
	"github.com/jinzhu/copier"
)

type TaskStore struct {
}

func (t *TaskStore) Create(r *task.CreateRequest, userId uint) (*task.CreateResponse, error) {
	m := &model.Task{
		UserId: userId,
	}
	_ = copier.CopyWithOption(m, r, copier.Option{IgnoreEmpty: true})
	result := db.DB.Create(m)
	res := &task.CreateResponse{}
	_ = copier.Copy(res, m)
	return res, result.Error
}

func (t *TaskStore) ReadAll(userId uint) (*[]task.ReadResponse, error) {
	m := &[]model.Task{}
	res := &[]task.ReadResponse{}
	return res, db.DB.Model(m).Where("user_id = ?", userId).Find(res).Error
}

func (t *TaskStore) ReadByTimestamp(timestamp int64, userId uint) (*task.ReadResponse, error) {
	m := &model.Task{}
	res := &task.ReadResponse{}
	return res, db.DB.Model(m).Where("timestamp = ? AND user_id = ?", timestamp, userId).First(res).Error
}

func (t *TaskStore) ReadByTimestampScope(startTimestamp int64, endTimestamp int64, userId uint) (*[]task.ReadResponse, error) {
	m := &[]model.Task{}
	res := &[]task.ReadResponse{}
	return res, db.DB.Model(m).Where("timestamp >= ? AND timestamp < ? AND user_id = ?", startTimestamp, endTimestamp, userId).Find(res).Error
}

func (t *TaskStore) ReadById(r *task.ReadRequest, userId uint) (*task.ReadResponse, error) {
	m := &model.Task{
		Id:     r.Id,
		UserId: userId,
	}
	result := db.DB.First(m)
	res := &task.ReadResponse{}
	_ = copier.Copy(res, m)
	return res, result.Error
}

func (t *TaskStore) Update(r *task.EditRequest, userId uint) (*task.EditResponse, error) {
	m, err := t.ReadById(&task.ReadRequest{Id: r.Id}, userId)
	if err != nil {
		return nil, err
	}
	err = copier.CopyWithOption(m, r, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return nil, err
	}
	result := db.DB.Save(m)
	res := &task.EditResponse{}
	_ = copier.Copy(res, m)
	return res, result.Error
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
