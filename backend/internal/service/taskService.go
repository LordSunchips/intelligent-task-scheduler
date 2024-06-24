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

// GetTask returns a task by ID
func (ts *TaskService) GetTask(id string) (*model.Task, error) {
	for _, task := range ts.Tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, fmt.Errorf("Task not found")
}

// UpdateTask updates a task by ID

// DeleteTask deletes a task by ID
func (ts *TaskService) DeleteTask(id string) error {
	for i, task := range ts.Tasks {
		if task.ID == id {
			ts.Tasks = append(ts.Tasks[:i], ts.Tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Task not found")
}

// GetTasks returns the list of tasks
func (ts *TaskService) GetTasks() []model.Task {
	return ts.Tasks
}