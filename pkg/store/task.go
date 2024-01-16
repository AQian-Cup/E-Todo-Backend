package store

import (
	"e-todo-backend/pkg/api/task"
	"e-todo-backend/pkg/db"
	"e-todo-backend/pkg/model"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type TaskStore struct {
}

func (t *TaskStore) Create(r *task.CreateRequest, userId uint) (*map[string]interface{}, error) {
	m := &model.Task{
		UserId: userId,
	}
	_ = copier.CopyWithOption(m, r, copier.Option{IgnoreEmpty: true})
	result := db.DB.Create(m)
	res := &task.CreateResponse{}
	resMap := &map[string]interface{}{}
	return resMap, result.Model(m).First(res).Error
}

func (t *TaskStore) ReadAll(userId uint) (*[]map[string]interface{}, error) {
	m := &[]model.Task{}
	res := &[]task.ReadResponse{}
	resMap := &[]map[string]interface{}{}
	return resMap, db.DB.Model(m).Where("user_id = ?", userId).Find(res).Find(resMap).Error
}

func (t *TaskStore) ReadByTimestamp(timestamp int64, userId uint) (*[]map[string]interface{}, error) {
	m := &model.Task{}
	res := &task.ReadResponse{}
	resMap := &[]map[string]interface{}{}
	return resMap, db.DB.Model(m).Where("timestamp = ? AND user_id = ?", timestamp, userId).Find(res).Find(resMap).Error
}

func (t *TaskStore) ReadByTimestampScope(startTimestamp int64, endTimestamp int64, userId uint) (*[]map[string]interface{}, error) {
	m := &[]model.Task{}
	res := &[]task.ReadResponse{}
	resMap := &[]map[string]interface{}{}
	return resMap, db.DB.Model(m).Where("timestamp >= ? AND timestamp < ? AND user_id = ?", startTimestamp, endTimestamp, userId).Find(res).Find(resMap).Error
}

func (t *TaskStore) ReadById(r *task.ReadRequest, userId uint) (*map[string]interface{}, error) {
	m := &model.Task{
		Model:  gorm.Model{ID: r.Id},
		UserId: userId,
	}
	res := &task.ReadResponse{}
	resMap := &map[string]interface{}{}
	return resMap, db.DB.Model(m).First(res).First(resMap).Error
}

func (t *TaskStore) Update(r *task.EditRequest, userId uint) (*map[string]interface{}, error) {
	m := &model.Task{
		UserId: userId,
	}
	_ = copier.CopyWithOption(m, r, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	result := db.DB.Model(m).Updates(m)
	res := &task.EditResponse{}
	_ = copier.Copy(res, m)
	resMap := &map[string]interface{}{}
	return resMap, result.Model(res).First(resMap).Error
}

func (t *TaskStore) DeleteById(r *task.DeleteRequest, userId uint) error {
	m := &model.Task{
		Model:  gorm.Model{ID: r.Id},
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
