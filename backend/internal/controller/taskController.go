package controller

import (
	"encoding/json"
	"net/http"
	"github.com/LordSunchips/intelligent-task-scheduler/backend/internal/service"
	"github.com/LordSunchips/intelligent-task-scheduler/backend/internal/model"
)

// TaskController handles HTTP requests related to tasks
type TaskController struct {
	TaskService *service.TaskService
}

// NewTaskController creates a new instance of TaskController
func NewTaskController(service *service.TaskService) *TaskController {
	return &TaskController{ TaskService: service }
}

// CreateTaskHandler handles HTTP POST requests to create a new task
func (tc *TaskController) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	createdTask := tc.TaskService.CreateTask(task.Name)

	// Respond with the created task
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTask)
}

// GetTasksHandler handles HTTP requests to get all tasks
func (tc *TaskController) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := tc.TaskService.GetTasks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}