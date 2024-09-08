package task

import (
	"database/sql"
	"errors"
	"time"
	"todo_list_api/pkg/models"
)

type Repository interface {
	CreateTask(task *models.Task) error
	DeleteTask(id int64) error
	GetTask(id int64) (*models.Task, error)
	ListTasks() ([]*models.Task, error)
	UpdateTask(task *models.Task) error
}

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	const query = "INSERT INTO tasks (title, description, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	return r.db.QueryRow(
		query,
		task.Title,
		task.Description,
		task.Status,
		task.CreatedAt,
		task.UpdatedAt,
	).Scan(&task.ID)
}

func (r *TaskRepository) GetTask(id int64) (*models.Task, error) {
	const query = "SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id = $1"
	task := &models.Task{}
	if err := r.db.QueryRow(query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return task, nil
}

func (r *TaskRepository) UpdateTask(task *models.Task) error {
	const query = "UPDATE tasks SET title = $2, description = $3, status = $4, created_at = $5, updated_at = $6  WHERE id = $1"
	task.UpdatedAt = time.Now()
	if _, err := r.db.Exec(
		query,
		task.ID,
		task.Title,
		task.Description,
		task.Status,
		task.CreatedAt,
		task.UpdatedAt,
	); err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) DeleteTask(id int64) error {
	const query = "DELETE FROM tasks WHERE id = $1"
	if _, err := r.db.Exec(query, id); err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) ListTasks() ([]models.Task, error) {
	const query = "SELECT id, title, description, status, created_at, updated_at FROM tasks"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
