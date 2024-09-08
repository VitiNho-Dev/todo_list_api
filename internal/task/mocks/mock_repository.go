package task

import (
	"todo_list_api/pkg/models"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

// CreateTask implements Repository.
func (m *MockRepository) CreateTask(task *models.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

// DeleteTask implements Repository.
func (m *MockRepository) DeleteTask(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

// GetTask implements Repository.
func (m *MockRepository) GetTask(id int64) (*models.Task, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Task), args.Error(1)
}

// ListTasks implements Repository.
func (m *MockRepository) ListTasks() ([]*models.Task, error) {
	args := m.Called()
	return args.Get(0).([]*models.Task), args.Error(1)
}

// UpdateTask implements Repository.
func (m *MockRepository) UpdateTask(task *models.Task) error {
	args := m.Called(task)
	return args.Error(0)
}
