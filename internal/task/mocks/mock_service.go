package task

import (
	"todo_list_api/pkg/models"

	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

// CreateTask implements task.Service.
func (m *MockService) CreateTask(task *models.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

// DeleteTask implements task.Service.
func (m *MockService) DeleteTask(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

// GetTask implements task.Service.
func (m *MockService) GetTask(id int64) (*models.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Task), args.Error(1)
}

// ListTasks implements task.Service.
func (m *MockService) ListTasks() ([]*models.Task, error) {
	args := m.Called()
	return args.Get(0).([]*models.Task), args.Error(1)
}

// UpdateTask implements task.Service.
func (m *MockService) UpdateTask(task *models.Task) error {
	args := m.Called(task)
	return args.Error(0)
}
