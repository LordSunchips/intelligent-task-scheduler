package controller

import (
	"encoding/json"
	"net/http"
	"github.com/LordSunchips/intelligent-task-scheduler/backend/internal/service"
)

// TaskController handles HTTP requests related to tasks
type TaskController struct {
	TaskService *service.TaskService
}

// NewTaskController creates a new instance of TaskController
func NewTaskController(service *service.TaskService) *TaskController {
	return &TaskController{ TaskService: service }
}

// AddTaskHandler handles HTTP requests to add a new task
func (tc *TaskController) AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Name string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task := tc.TaskService.AddTask(requestBody.Name)
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// GetTasksHandler handles HTTP requests to get all tasks
func (tc *TaskController) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := tc.TaskService.GetTasks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}