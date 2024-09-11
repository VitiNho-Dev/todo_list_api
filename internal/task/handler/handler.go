package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	s "todo_list_api/internal/task/service"
	"todo_list_api/pkg/models"
	"todo_list_api/pkg/utils"
)

type Handler struct {
	service s.Service
}

func NewHandler(service s.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, utils.ErrInvalidPayload.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateTask(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, utils.ErrEmptyID.Error(), http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, utils.ErrInvalidId.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteTask(int64(ID)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, utils.ErrEmptyID.Error(), http.StatusBadRequest)
		return
	}

	ID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, utils.ErrInvalidId.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.service.GetTask(int64(ID))
	if task == nil {
		http.Error(w, utils.ErrTaskNotFound.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, utils.ErrFailedEncode.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.ListTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, utils.ErrFailedEncode.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, utils.ErrInvalidPayload.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateTask(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
