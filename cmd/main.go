package main

import (
	"log"
	"n1ktarchik/go-final/internal/core/transport/server"
	"net/http"
	"os"

	tasks_repository "n1ktarchik/go-final/internal/features/tasks/repository"
	tasks_service "n1ktarchik/go-final/internal/features/tasks/service"
	tasks_transport_http "n1ktarchik/go-final/internal/features/tasks/transport/http"

	auth_service "n1ktarchik/go-final/internal/features/auth/service"
	auth_transport_http "n1ktarchik/go-final/internal/features/auth/transport/http"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not found")
	}

	serverPort := os.Getenv("TODO_PORT")
	if serverPort == "" {
		log.Fatal("server port not faund in .env file")
	}

	webDir := os.Getenv("WEB_DIR")
	if webDir == "" {
		log.Fatal("web directory not faund in .env file")
	}

	dbPath := os.Getenv("TODO_DBFILE")
	if dbPath == "" {
		log.Fatal("databse path not faund in .env file")
	}

	secretKey := os.Getenv("TODO_SECRET")
	if secretKey == "" {
		log.Fatal("secret key for jwt not faund in .env file")
	}

	password := os.Getenv("TODO_PASSWORD")

	TasksRepository, err := tasks_repository.NewTasksRepository(dbPath)
	if err != nil {
		log.Fatal("error to create repository")
	}
	TasksService := tasks_service.NewTasksService(TasksRepository)
	TasksHandler := tasks_transport_http.NewTasksTransport(TasksService)

	AuthService := auth_service.NewAuthService(secretKey, password)
	AuthHandler := auth_transport_http.NewAuthTransport(AuthService)

	server := server.NewHTTPServer()
	r := server.Router

	r.HandleFunc("/api/signin", AuthHandler.Login).Methods("POST")
	r.HandleFunc("/api/nextdate", TasksHandler.NextDate).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(AuthHandler.AuthMiddleware)

	api.HandleFunc("/task", TasksHandler.CreateTask).Methods("POST")
	api.HandleFunc("/tasks", TasksHandler.GetAllTasks).Methods("GET")
	api.HandleFunc("/task", TasksHandler.GetTaskByID).Methods("GET")
	api.HandleFunc("/task", TasksHandler.UpdateTask).Methods("PUT")
	api.HandleFunc("/task/done", TasksHandler.CompleteTask).Methods("POST")
	api.HandleFunc("/task", TasksHandler.DeleteTask).Methods("DELETE")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir(webDir)))

	server.Run(serverPort)

}
