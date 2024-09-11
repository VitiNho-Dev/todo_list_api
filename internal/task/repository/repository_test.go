package repository_test

import (
	"errors"
	"testing"
	"time"
	"todo_list_api/internal/task/repository"
	"todo_list_api/pkg/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewTaskRepository(db)

	t.Run("must validate the query and if query is valid, create a new task in the tasks table", func(t *testing.T) {
		task := &models.Task{
			Title:       "Test Task",
			Description: "Test Description",
			Status:      "Pending",
		}

		const query = "INSERT INTO tasks \\(title, description, status, created_at, updated_at\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5\\) RETURNING id"

		mock.ExpectQuery(query).
			WithArgs(
				task.Title,
				task.Description,
				task.Status,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
			).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		err = repo.CreateTask(task)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), task.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("should validate the query and return an error if the query is invalid", func(t *testing.T) {
		task := &models.Task{
			Title:       "Test Task",
			Description: "Test Description",
			Status:      "Pending",
		}

		const query = "INSERT INTO tasks (title, description, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id"

		mock.ExpectQuery(query).
			WithArgs(
				task.Title,
				task.Description,
				task.Status,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
			).
			WillReturnError(errors.New("query invalid"))

		err = repo.CreateTask(task)
		assert.Error(t, err)
		assert.Error(t, mock.ExpectationsWereMet())
	})
}

func TestGetTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewTaskRepository(db)

	t.Run("should validate the query and if the query is valid, return a task by id", func(t *testing.T) {
		columns := []string{"id", "title", "description", "status", "created_at", "updated_at"}

		expectedTask := &models.Task{
			ID:          1,
			Title:       "Test Task",
			Description: "Test Description",
			Status:      "Pending",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		const query = "SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id = \\$1"

		mock.ExpectQuery(query).
			WithArgs(expectedTask.ID).
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(
					expectedTask.ID,
					expectedTask.Title,
					expectedTask.Description,
					expectedTask.Status,
					expectedTask.CreatedAt,
					expectedTask.UpdatedAt,
				),
			)

		task, err := repo.GetTask(expectedTask.ID)
		assert.NoError(t, err)
		assert.Equal(t, expectedTask, task)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("should validate the query and return a error if the query is invalid", func(t *testing.T) {
		expectedTask := &models.Task{ID: 1}

		const query = "SELECT id, title, description, status, created_at, updated_at FROM tasks WHERE id = $1"

		mock.ExpectQuery(query).
			WithArgs(expectedTask.ID).
			WillReturnError(errors.New("query invalid"))

		_, err := repo.GetTask(expectedTask.ID)
		assert.Error(t, err)
		assert.Error(t, mock.ExpectationsWereMet())
	})
}

func TestUpdateTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewTaskRepository(db)

	t.Run("must be a valid query and if the query is valid, return an updated task", func(t *testing.T) {
		taskUpdated := &models.Task{
			ID:          1,
			Title:       "Test Task",
			Description: "Test Description",
			Status:      "Pending",
			CreatedAt:   time.Now().Add(-24 * time.Hour),
			UpdatedAt:   time.Now(),
		}

		const query = "UPDATE tasks SET title = \\$2, description = \\$3, status = \\$4, created_at = \\$5, updated_at = \\$6 WHERE id = \\$1"

		mock.ExpectExec(query).
			WithArgs(
				taskUpdated.ID,
				taskUpdated.Title,
				taskUpdated.Description,
				taskUpdated.Status,
				taskUpdated.CreatedAt,
				sqlmock.AnyArg(),
			).WillReturnResult(sqlmock.NewResult(0, 1))

		err = repo.UpdateTask(taskUpdated)
		assert.NoError(t, err)
		assert.WithinDuration(t, time.Now(), taskUpdated.UpdatedAt, time.Second)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("should validate the query and return a error if the query is invalid", func(t *testing.T) {
		taskUpdated := &models.Task{
			ID:          1,
			Title:       "Test Task",
			Description: "Test Description",
			Status:      "Pending",
			CreatedAt:   time.Now().Add(-24 * time.Hour),
			UpdatedAt:   time.Now(),
		}

		const query = "UPDATE tasks SET title = $2, description = $3, status = $4, created_at = $5, updated_at = $6 WHERE id = $1"

		mock.ExpectExec(query).
			WithArgs(
				taskUpdated.ID,
				taskUpdated.Title,
				taskUpdated.Description,
				taskUpdated.Status,
				taskUpdated.CreatedAt,
				sqlmock.AnyArg(),
			).WillReturnError(errors.New("query invalid"))

		err = repo.UpdateTask(taskUpdated)
		assert.Error(t, err)
		assert.Error(t, mock.ExpectationsWereMet())
	})
}

func TestDeleteTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewTaskRepository(db)

	t.Run("must validate the query and if the query is valid, delete the task from the tasks table", func(t *testing.T) {
		const id = 1

		const query = "DELETE FROM tasks WHERE id = \\$1"

		mock.ExpectExec(query).
			WithArgs(id).
			WillReturnResult(sqlmock.NewResult(0, 1))

		err := repo.DeleteTask(id)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("should validate the query and return a error if the query is invalid", func(t *testing.T) {
		const id = 1

		const query = "DELETE FROM tasks WHERE id = $1"

		mock.ExpectExec(query).
			WithArgs(id).
			WillReturnError(errors.New("query invalid"))

		err := repo.DeleteTask(id)
		assert.Error(t, err)
		assert.Error(t, mock.ExpectationsWereMet())
	})
}

func TestListTasks(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := repository.NewTaskRepository(db)

	t.Run("must validate the query and if the query is valid, return a list of tasks from the tasks table", func(t *testing.T) {
		columns := []string{"id", "title", "description", "status", "created_at", "updated_at"}

		task := models.Task{
			ID:          1,
			Title:       "Task",
			Description: "Description",
			Status:      "Pending",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		const query = "SELECT id, title, description, status, created_at, updated_at FROM tasks"

		mock.ExpectQuery(query).
			WithArgs().
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(
					task.ID,
					task.Title,
					task.Description,
					task.Status,
					task.CreatedAt,
					task.UpdatedAt,
				),
			)

		taskResult, err := repo.ListTasks()
		assert.NoError(t, err)
		assert.Equal(t, int64(1), taskResult[0].ID)
		assert.Equal(t, "Task", taskResult[0].Title)
		assert.Equal(t, "Description", taskResult[0].Description)
		assert.Equal(t, "Pending", taskResult[0].Status)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("should validate the query and return a error if the query is invalid", func(t *testing.T) {
		columns := []string{"id", "title", "description", "status", "created_at", "updated_at"}

		task := models.Task{
			ID:          1,
			Title:       "Task",
			Description: "Description",
			Status:      "Pending",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		const query = "SELECT id, title, description, status, created_at, updated_at FROM "

		mock.ExpectQuery(query).
			WithArgs().
			WillReturnRows(sqlmock.NewRows(columns).
				AddRow(
					task.ID,
					task.Title,
					task.Description,
					task.Status,
					task.CreatedAt,
					task.UpdatedAt,
				),
			)

		taskResult, err := repo.ListTasks()
		assert.NoError(t, err)
		assert.Equal(t, int64(1), taskResult[0].ID)
		assert.Equal(t, "Task", taskResult[0].Title)
		assert.Equal(t, "Description", taskResult[0].Description)
		assert.Equal(t, "Pending", taskResult[0].Status)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
