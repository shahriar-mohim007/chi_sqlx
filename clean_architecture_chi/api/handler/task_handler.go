//API (Web Layer):

// This layer deals with HTTP requests and responses. We use chi for routing and connect it to the use cases.

// api/handler/task_handler.go
package handler

import (
	"clean-architecture/internal/domain"
	"clean-architecture/internal/usecase"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type TaskHandler struct {
	taskUsecase *usecase.TaskUsecase
}

func NewTaskHandler(uc *usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{taskUsecase: uc}
}

func (h *TaskHandler) GetTask(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")             // Get the id parameter from URL
	id, err := strconv.ParseInt(idParam, 10, 64) // Convert string to int64
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	task, err := h.taskUsecase.GetTask(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var task domain.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.taskUsecase.CreateTask(&task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
