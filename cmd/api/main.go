package main

import (
	"log"
	"net/http"
	"todo_list_api/internal/db"
	"todo_list_api/internal/task/handler"
	"todo_list_api/internal/task/repository"
	"todo_list_api/internal/task/service"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}

	mux := http.NewServeMux()

	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewHandler(taskService)

	mux.HandleFunc("POST /tasks", taskHandler.CreateTask)
	mux.HandleFunc("GET /tasks/{id}", taskHandler.GetTask)
	mux.HandleFunc("PUT /tasks/{id}", taskHandler.UpdateTask)
	mux.HandleFunc("DELETE /tasks/{id}", taskHandler.DeleteTask)
	mux.HandleFunc("GET /tasks", taskHandler.ListTasks)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
