package service

import (
	"github.com/LordSunchips/intelligent-task-scheduler/backend/internal/model"
)

// TaskService provides methods for task management
type TaskService struct {
	Tasks []model.Task
}

// NewTaskService creates a new TaskService
func NewTaskService() *TaskService {
	return &TaskService{
		Tasks: make([]model.Task, 0),
	}
}

// AddTask adds a new task to the task list
func (ts *TaskService) AddTask(name string) *model.Task {
	task := &model.Task{
		ID: "unique-id",
		Name: name,
		Status: "Pending",
	}
	ts.Tasks = append(ts.Tasks, *task)
	return task
}

// GetTasks returns the list of tasks
func (ts *TaskService) GetTasks() []model.Task {
	return ts.Tasks
}