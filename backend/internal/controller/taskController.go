package controller

import (
	"encoding/json"
	"net/http"
	"github.com/LordSunchips/intelligent-task-scheduler/backend/internal/service"
	"github.com/LordSunchips/intelligent-task-scheduler/backend/internal/model"
	"github.com/gorilla/mux"
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

	createdTask := tc.TaskService.CreateTask(task.Name, task.Priority, task.Deadline, task.ResourceNeeded)

	// Respond with the created task
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTask)
}

// GetTaskHandler handles HTTP GET requests to get a task by ID
func (tc *TaskController) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]
	task, err := tc.TaskService.GetTask(taskID)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// UpdateTaskHandler handles HTTP PUT requests to update a task by ID
func (tc *TaskController) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]
	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	updatedTask, err := tc.TaskService.UpdateTask(taskID, task.Name, task.Priority, task.Deadline, task.ResourceNeeded)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	// Respond with the updated task
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

// DeleteTaskHandler handles HTTP DELETE requests to delete a task by ID
func (tc *TaskController) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]
	err := tc.TaskService.DeleteTask(taskID)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetTasksHandler handles HTTP requests to get all tasks
func (tc *TaskController) GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := tc.TaskService.GetTasks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}