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

	router.HandleFunc("/api/createTask", taskController.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/api/task/{id}", taskController.GetTaskHandler).Methods("GET")
	router.HandleFunc("/api/updateTask/{id}", taskController.UpdateTaskHandler).Methods("PUT")
	router.HandleFunc("/api/deleteTask/{id}", taskController.DeleteTaskHandler).Methods("DELETE")
	router.HandleFunc("/api/tasks", taskController.GetTasksHandler).Methods("GET")

	http.ListenAndServe(":5001", router)
}