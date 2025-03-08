package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskByID(userId uint, taskId uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(userId uint, taskId uint, task model.Task) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

func NewTaskUseCase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUsecase {
	return &taskUsecase{tr, tv}
}

func (tu *taskUsecase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	var tasks []model.Task
	if err := tu.tr.GetAll(&tasks, userId); err != nil {
		return nil, err
	}

	var taskResponses []model.TaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, model.TaskResponse{
			ID:        task.ID,
			Title:     task.Title,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		})
	}
	return taskResponses, nil
}

func (tu *taskUsecase) GetTaskByID(userId uint, taskId uint) (model.TaskResponse, error) {
	task := model.Task{}
	if err := tu.tr.GetByID(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	return model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}, nil
}

func (tu *taskUsecase) CreateTask(task model.Task) (model.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}
	if err := tu.tr.Create(&task); err != nil {
		return model.TaskResponse{}, err
	}
	return model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}, nil
}

func (tu *taskUsecase) UpdateTask(userId uint, taskId uint, task model.Task) (model.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}
	if err := tu.tr.Update(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	return model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}, nil
}

func (tu *taskUsecase) DeleteTask(userId uint, taskId uint) error {
	return tu.tr.Delete(userId, taskId)
}
