package task

import (
	"errors"
	"testing"
	m "todo_list_api/internal/task/mocks"
	"todo_list_api/pkg/models"
	"todo_list_api/pkg/utils"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	mockRepo := new(m.MockRepository)
	svc := NewTaskService(mockRepo)

	t.Run("should return error if title is empty", func(t *testing.T) {
		task := &models.Task{Status: "Pending"}
		err := svc.CreateTask(task)
		assert.ErrorIs(t, err, utils.ErrEmptyTitle)
	})

	t.Run("should return error if status is empty", func(t *testing.T) {
		task := &models.Task{Title: "New Task"}
		err := svc.CreateTask(task)
		assert.ErrorIs(t, err, utils.ErrEmptyStatus)
	})

	t.Run("should return error if status is invalid", func(t *testing.T) {
		task := &models.Task{Title: "New Task", Status: "Unknown"}
		err := svc.CreateTask(task)
		assert.ErrorIs(t, err, utils.ErrInvalidStatus)
	})

	t.Run("should create task if data is valid", func(t *testing.T) {
		task := &models.Task{Title: "New Task", Status: "Pending"}
		mockRepo.On("CreateTask", task).Return(nil).Once()

		err := svc.CreateTask(task)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if repository fails", func(t *testing.T) {
		task := &models.Task{Title: "New Task", Status: "Pending"}
		mockRepo.On("CreateTask", task).Return(errors.New("repository error")).Once()

		err := svc.CreateTask(task)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteTask(t *testing.T) {
	mockRepo := new(m.MockRepository)
	svc := NewTaskService(mockRepo)

	taskID := int64(1)

	t.Run("should a error if id is invalid", func(t *testing.T) {
		err := svc.DeleteTask(-1)
		assert.ErrorIs(t, err, utils.ErrInvalidId)
	})

	t.Run("must delete a task from the task table", func(t *testing.T) {
		mockRepo.On("DeleteTask", taskID).Return(nil).Once()

		err := svc.DeleteTask(taskID)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if repository fails", func(t *testing.T) {
		mockRepo.On("DeleteTask", taskID).Return(errors.New("repository error")).Once()

		err := svc.DeleteTask(taskID)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}

func TestGetTask(t *testing.T) {
	mockRepo := new(m.MockRepository)
	svc := NewTaskService(mockRepo)

	taskID := int64(1)

	t.Run("should return error if ID is invalid", func(t *testing.T) {
		_, err := svc.GetTask(-1)
		assert.ErrorIs(t, err, utils.ErrInvalidId)
	})

	t.Run("should return a task if ID is valid", func(t *testing.T) {
		mockTask := &models.Task{
			ID:    taskID,
			Title: "Valid Task",
		}

		mockRepo.On("GetTask", taskID).Return(mockTask, nil).Once()

		task, err := svc.GetTask(taskID)
		assert.NoError(t, err)
		assert.NotNil(t, task)
		assert.Equal(t, mockTask, task)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if task not found", func(t *testing.T) {
		mockRepo.On("GetTask", taskID).Return((*models.Task)(nil), nil).Once()

		task, err := svc.GetTask(taskID)
		assert.ErrorIs(t, err, utils.ErrTaskNotFound)
		assert.Nil(t, task)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if repository fails", func(t *testing.T) {
		mockRepo.On("GetTask", taskID).Return((*models.Task)(nil), errors.New("repository error")).Once()

		task, err := svc.GetTask(taskID)
		assert.Error(t, err)
		assert.Nil(t, task)
		mockRepo.AssertExpectations(t)
	})
}

func TestListTasks(t *testing.T) {
	mockRepo := new(m.MockRepository)
	svc := NewTaskService(mockRepo)

	t.Run("should return a task list", func(t *testing.T) {
		mockTask := []*models.Task{
			{ID: 1, Title: "Task 1"},
			{ID: 2, Title: "Task 2"},
		}

		mockRepo.On("ListTasks").Return(mockTask, nil).Once()

		tasks, err := svc.ListTasks()
		assert.NoError(t, err)
		assert.NotNil(t, tasks)
		assert.Equal(t, 2, len(tasks))
		assert.Equal(t, mockTask, tasks)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if repository fails", func(t *testing.T) {
		mockRepo.On("ListTasks").Return(([]*models.Task)(nil), errors.New("repository error")).Once()

		tasks, err := svc.ListTasks()
		assert.Error(t, err)
		assert.Nil(t, tasks)
		mockRepo.AssertExpectations(t)
	})
}

func TestUpdateTask(t *testing.T) {
	mockRepo := new(m.MockRepository)
	svc := NewTaskService(mockRepo)

	t.Run("should return error if title is empty", func(t *testing.T) {
		task := &models.Task{Status: "Pending"}
		err := svc.CreateTask(task)
		assert.ErrorIs(t, err, utils.ErrEmptyTitle)
	})

	t.Run("should return error if status is empty", func(t *testing.T) {
		task := &models.Task{Title: "New Task"}
		err := svc.CreateTask(task)
		assert.ErrorIs(t, err, utils.ErrEmptyStatus)
	})

	t.Run("should return error if status is invalid", func(t *testing.T) {
		task := &models.Task{Title: "New Task", Status: "Unknown"}
		err := svc.CreateTask(task)
		assert.ErrorIs(t, err, utils.ErrInvalidStatus)
	})

	task := &models.Task{Title: "New Task", Status: "Pending"}

	t.Run("should update the task if data is valid", func(t *testing.T) {
		mockRepo.On("UpdateTask", task).Return(nil).Once()

		err := svc.UpdateTask(task)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("should return error if repository fails", func(t *testing.T) {
		mockRepo.On("UpdateTask", task).Return(errors.New("repository error")).Once()

		err := svc.UpdateTask(task)
		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})
}
