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

// GenerateTaskID generates a unique task ID
func (ts *TaskService) GenerateTaskID() string {
	// use the current time and hash it
	currentTime := time.Now()
	hash := md5.New()
	hash.Write([]byte(currentTime.String()))
	return hex.EncodeToString(hash.Sum(nil))
}

// CreateTask creates a new task
func (ts *TaskService) CreateTask(name string, priority int, deadline string, resource_needed int) *model.Task {
	task := &model.Task{
		ID: ts.GenerateTaskID(),
		Name: name,
		Status: "Pending",
		Priority: priority,
		Deadline: deadline,
		ResourceNeeded: resource_needed,
	}
	ts.Tasks = append(ts.Tasks, *task)
	return task
}

// 

// GetTasks returns the list of tasks
func (ts *TaskService) GetTasks() []model.Task {
	return ts.Tasks
}