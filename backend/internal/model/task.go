package model

type Task struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Status string `json:"status"`
	Priority int `json:"priority"`
	Deadline string `json:"deadline"`
	ResourceNeeded int `json:"resource_needed"`
}