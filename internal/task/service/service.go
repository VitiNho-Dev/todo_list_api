package service

import (
	r "todo_list_api/internal/task/repository"
	"todo_list_api/pkg/models"
	"todo_list_api/pkg/utils"
)

type Service interface {
	CreateTask(task *models.Task) error
	DeleteTask(id int64) error
	GetTask(id int64) (*models.Task, error)
	ListTasks() ([]*models.Task, error)
	UpdateTask(task *models.Task) error
}

type TaskService struct {
	repo r.Repository
}

func NewTaskService(repo r.Repository) Service {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task *models.Task) error {
	if task.Title == "" {
		return utils.ErrEmptyTitle
	}

	if task.Status == "" {
		return utils.ErrEmptyStatus
	}

	err := utils.ValidateStatus(task.Status)
	if err != nil {
		return err
	}

	return s.repo.CreateTask(task)
}

func (s *TaskService) DeleteTask(id int64) error {
	if id < 0 {
		return utils.ErrInvalidId
	}

	return s.repo.DeleteTask(id)
}

func (s *TaskService) GetTask(id int64) (*models.Task, error) {
	if id < 0 {
		return nil, utils.ErrInvalidId
	}

	task, err := s.repo.GetTask(id)
	if err != nil {
		return nil, err
	}

	if task == nil {
		return nil, utils.ErrTaskNotFound
	}

	return task, nil
}

func (s *TaskService) ListTasks() ([]*models.Task, error) {
	tasks, err := s.repo.ListTasks()
	if err != nil {
		return nil, err
	}

	if tasks == nil {
		return []*models.Task{}, nil
	}

	return tasks, nil
}

func (s *TaskService) UpdateTask(task *models.Task) error {
	if task.Title == "" {
		return utils.ErrEmptyTitle
	}

	if task.Status == "" {
		return utils.ErrEmptyStatus
	}

	err := utils.ValidateStatus(task.Status)
	if err != nil {
		return err
	}

	return s.repo.UpdateTask(task)
}
