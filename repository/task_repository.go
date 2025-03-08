package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	Create(task *model.Task) error
	GetAll(tasks *[]model.Task, userId uint) error
	GetByID(task *model.Task, userId uint, taskId uint) error
	Update(task *model.Task, userId uint, taskId uint) error
	Delete(userId uint, taskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) Create(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetAll(tasks *[]model.Task, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetByID(task *model.Task, userId uint, taskId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).First(task, taskId).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) Update(task *model.Task, userId uint, taskId uint) error {
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("user_id = ? AND id = ?", userId, taskId).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (tr *taskRepository) Delete(userId uint, taskId uint) error {
	if err := tr.db.Where("user_id = ? AND id = ?", userId, taskId).Delete(&model.Task{}).Error; err != nil {
		return err
	}
	return nil
}
