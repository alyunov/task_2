package main

import (
	"net/http"
	"task_2/internal/database"
	"task_2/internal/handlers"
	"task_2/internal/taskService"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})

	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")

	router.HandleFunc("/api/patch", handler.PatchTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/delete", handler.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
