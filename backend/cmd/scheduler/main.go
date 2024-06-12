package main

import (
	"net/http"
	"github.com/LordSunchips/intelligent-task-scheduler/backend/internal/controller"
	"github.com/LordSunchips/intelligent-task-scheduler/backend/internal/service"
	"github.com/gorilla/mux"
)

func main() {
	taskService := service.NewTaskService()
	taskController := controller.NewTaskController(taskService)
	router := mux.NewRouter()

	router.HandleFunc("/api/tasks", taskController.GetTasksHandler).Methods("GET")
	router.HandleFunc("api/tasks", taskController.AddTaskHandler).Methods("POST")

	http.ListenAndServe(":8080", router)
}