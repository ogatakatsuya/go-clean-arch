package usecase

import (
	"errors"
	"go-rest-api/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTaskRepository struct {
	mock.Mock
}

func newMockTaskRepository() *MockTaskRepository {
	return &MockTaskRepository{}
}

func (mr *MockTaskRepository) Create(task *model.Task) error {
	args := mr.Called(task)
	return args.Error(0)
}

func (mr *MockTaskRepository) GetAll(tasks *[]model.Task, userId uint) error {
	args := mr.Called(tasks, userId)
	return args.Error(0)
}

func (mr *MockTaskRepository) GetByID(task *model.Task, userId uint, taskId uint) error {
	args := mr.Called(task, userId, taskId)
	return args.Error(0)
}

func (mr *MockTaskRepository) Update(task *model.Task, userId uint, taskId uint) error {
	args := mr.Called(task, userId, taskId)
	return args.Error(0)
}

func (mr *MockTaskRepository) Delete(userId uint, taskId uint) error {
	args := mr.Called(userId, taskId)
	return args.Error(0)
}

type MockTaskValidator struct {
	mock.Mock
}

func newMockTaskValidator() *MockTaskValidator {
	return &MockTaskValidator{}
}

func (mv *MockTaskValidator) TaskValidate(task model.Task) error {
	args := mv.Called(task)
	return args.Error(0)
}

func TestCreateTask_Success(t *testing.T) {
	mr := newMockTaskRepository()
	mv := newMockTaskValidator()
	mr.On("Create", mock.Anything).Return(nil)
	mv.On("TaskValidate", mock.Anything).Return(nil)

	tu := NewTaskUseCase(mr, mv)

	_, err := tu.CreateTask(model.Task{Title: "test"})
	assert.NoError(t, err)
	mr.AssertCalled(t, "Create", mock.Anything)
	mv.AssertCalled(t, "TaskValidate", mock.Anything)
}

func TestCreateTask_Repository_Failure(t *testing.T) {
	mr := newMockTaskRepository()
	mv := newMockTaskValidator()
	mr.On("Create", mock.Anything).Return(errors.New("error"))
	mv.On("TaskValidate", mock.Anything).Return(nil)

	tu := NewTaskUseCase(mr, mv)

	_, err := tu.CreateTask(model.Task{Title: "test"})
	assert.Error(t, err)
}

func TestCreateTask_Validator_Failure(t *testing.T) {
	mr := newMockTaskRepository()
	mv := newMockTaskValidator()
	mr.On("Create", mock.Anything).Return(nil)
	mv.On("TaskValidate", mock.Anything).Return(errors.New("error"))

	tu := NewTaskUseCase(mr, mv)

	_, err := tu.CreateTask(model.Task{Title: "test"})
	assert.Error(t, err)
}

func TestGetAllTasks_Success(t *testing.T) {
	mr := newMockTaskRepository()
	mv := newMockTaskValidator()
	mr.On("GetAll", mock.Anything, mock.Anything).Return(nil)

	tu := NewTaskUseCase(mr, mv)

	_, err := tu.GetAllTasks(1)
	assert.NoError(t, err)
	mr.AssertCalled(t, "GetAll", mock.Anything, mock.Anything)
}

func TestGetAllTasks_Repository_Failure(t *testing.T) {
	mr := newMockTaskRepository()
	mv := newMockTaskValidator()
	mr.On("GetAll", mock.Anything, mock.Anything).Return(errors.New("error"))

	tu := NewTaskUseCase(mr, mv)

	_, err := tu.GetAllTasks(1)
	assert.Error(t, err)
}

func TestGetTaskByID_Success(t *testing.T) {
	mr := newMockTaskRepository()
	mv := newMockTaskValidator()
	mr.On("GetByID", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	tu := NewTaskUseCase(mr, mv)

	_, err := tu.GetTaskByID(1, 1)
	assert.NoError(t, err)
	mr.AssertCalled(t, "GetByID", mock.Anything, mock.Anything, mock.Anything)
}

func TestGetTaskByID_Repository_Failure(t *testing.T) {
	mr := newMockTaskRepository()
	mv := newMockTaskValidator()
	mr.On("GetByID", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))

	tu := NewTaskUseCase(mr, mv)

	_, err := tu.GetTaskByID(1, 1)
	assert.Error(t, err)
}

func TestUpdateTask_Success(t *testing.T) {
	mr := newMockTaskRepository()
	mv := newMockTaskValidator()
	mr.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	mv.On("TaskValidate", mock.Anything).Return(nil)

	tu := NewTaskUseCase(mr, mv)

	_, err := tu.UpdateTask(1, 1, model.Task{Title: "test"})
	assert.NoError(t, err)
	mr.AssertCalled(t, "Update", mock.Anything, mock.Anything, mock.Anything)
	mv.AssertCalled(t, "TaskValidate", mock.Anything)
}

func TestUpdateTask_Respository_Failure(t *testing.T) {
	mr := newMockTaskRepository()
	mv := newMockTaskValidator()
	mr.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error"))
	mv.On("TaskValidate", mock.Anything).Return(nil)

	tu := NewTaskUseCase(mr, mv)

	_, err := tu.UpdateTask(1, 1, model.Task{Title: "test"})
	assert.Error(t, err)
}

func TestUpdateTask_Validator_Failure(t *testing.T) {
	mr := newMockTaskRepository()
	mv := newMockTaskValidator()
	mr.On("Update", mock.Anything).Return(nil)
	mv.On("TaskValidate", mock.Anything).Return(errors.New("error"))

	tu := NewTaskUseCase(mr, mv)

	_, err := tu.CreateTask(model.Task{Title: "test"})
	assert.Error(t, err)
}

func TestDeleteTask_Success(t *testing.T) {
	mr := newMockTaskRepository()
	mv := newMockTaskValidator()
	mr.On("Delete", mock.Anything, mock.Anything).Return(nil)

	tu := NewTaskUseCase(mr, mv)

	err := tu.DeleteTask(1, 1)
	assert.NoError(t, err)
	mr.AssertCalled(t, "Delete", mock.Anything, mock.Anything)
}

func TestDeleteTask_Repository_Failure(t *testing.T) {
	mr := newMockTaskRepository()
	mv := newMockTaskValidator()
	mr.On("Delete", mock.Anything, mock.Anything).Return(errors.New("error"))

	tu := NewTaskUseCase(mr, mv)

	err := tu.DeleteTask(1, 1)
	assert.Error(t, err)
}
