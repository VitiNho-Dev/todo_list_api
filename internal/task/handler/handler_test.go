package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	m "todo_list_api/internal/task/mocks"
	"todo_list_api/pkg/models"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	mockService := new(m.MockService)
	handler := NewHandler(mockService)

	t.Run("should make the request and return a bad request error when creating the task", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer([]byte("invalid json")))

		rr := httptest.NewRecorder()
		handler.CreateTask(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "Invalid request payload\n", rr.Body.String())
	})

	t.Run("should make the request and return a internal server error when creating the task", func(t *testing.T) {
		task := &models.Task{}

		mockService.On("CreateTask", task).Return(errors.New("status internal server error"))

		body, _ := json.Marshal(task)
		req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.CreateTask(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("must make the request and if the task was created return a created status", func(t *testing.T) {
		task := &models.Task{
			Title:       "New Task",
			Description: "New Description",
			Status:      "Pending",
		}

		mockService.On("CreateTask", task).Return(nil)

		body, _ := json.Marshal(task)
		req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		handler.CreateTask(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		mockService.AssertExpectations(t)
	})
}

func TestDeleteTask(t *testing.T) {
	mockService := new(m.MockService)
	handler := NewHandler(mockService)

	t.Run("should validate the ID and if the ID is empty return 400", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/tasks/", nil)

		rr := httptest.NewRecorder()
		handler.DeleteTask(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "ID cannot be empty\n", rr.Body.String())
	})

	t.Run("should validate the ID and if the ID is invalid return 400", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/tasks/invalid", nil)
		rr := httptest.NewRecorder()

		req.SetPathValue("id", "invalid")

		handler.DeleteTask(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "the id is invalid\n", rr.Body.String())
	})

	t.Run("should return 500 if service returns error", func(t *testing.T) {
		taskID := 1
		req, _ := http.NewRequest("DELETE", "/tasks/"+strconv.Itoa(taskID), nil)
		rr := httptest.NewRecorder()

		req.SetPathValue("id", strconv.Itoa(taskID))

		mockService.On("DeleteTask", int64(taskID)).Return(errors.New("service error")).Once()

		handler.DeleteTask(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, "service error\n", rr.Body.String())
		mockService.AssertExpectations(t)
	})

	t.Run("should return 204 if task is deleted successfully", func(t *testing.T) {
		taskID := 1
		req, _ := http.NewRequest("DELETE", "/tasks/"+strconv.Itoa(taskID), nil)
		rr := httptest.NewRecorder()

		req.SetPathValue("id", strconv.Itoa(taskID))

		mockService.On("DeleteTask", int64(taskID)).Return(nil).Once()

		handler.DeleteTask(rr, req)

		assert.Equal(t, http.StatusNoContent, rr.Code)
		mockService.AssertExpectations(t)
	})
}

func TestGetTask(t *testing.T) {
	mockService := new(m.MockService)
	handler := NewHandler(mockService)

	t.Run("should validate the ID and if the ID is empty return 400", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/tasks/", nil)

		rr := httptest.NewRecorder()
		handler.GetTask(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "ID cannot be empty\n", rr.Body.String())
	})

	t.Run("should validate the ID and if the ID is invalid return 400", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/tasks/invalid", nil)
		rr := httptest.NewRecorder()

		req.SetPathValue("id", "invalid")

		handler.GetTask(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		assert.Equal(t, "the id is invalid\n", rr.Body.String())
	})

	t.Run("should return 404 if task not found", func(t *testing.T) {
		taskID := 1
		req, _ := http.NewRequest("GET", "/tasks/"+strconv.Itoa(taskID), nil)
		rr := httptest.NewRecorder()

		req.SetPathValue("id", strconv.Itoa(taskID))

		mockService.On("GetTask", int64(taskID)).Return((*models.Task)(nil), nil).Once()

		handler.GetTask(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Equal(t, "task not found\n", rr.Body.String())
		mockService.AssertExpectations(t)
	})
}
