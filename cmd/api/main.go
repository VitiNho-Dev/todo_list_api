package main

import (
	"log"
	"net/http"
	"todo_list_api/internal/db"
	"todo_list_api/internal/task/handler"
	"todo_list_api/internal/task/repository"
	"todo_list_api/internal/task/service"

	"github.com/gorilla/mux"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewHandler(taskService)

	r := mux.NewRouter()

	r.HandleFunc("/tasks", taskHandler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", taskHandler.GetTask).Methods("GET")
	r.HandleFunc("/tasks/{id}", taskHandler.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", taskHandler.DeleteTask).Methods("DELETE")
	r.HandleFunc("/tasks", taskHandler.ListTasks).Methods("GET")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
