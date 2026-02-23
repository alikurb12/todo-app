package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alikurb12/todo-app-go/internal/config"
	"github.com/alikurb12/todo-app-go/internal/handler"
	"github.com/alikurb12/todo-app-go/internal/repository"
	"github.com/alikurb12/todo-app-go/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg := config.Load()

	connString := "postgres://" + cfg.DBUser + ":" + cfg.DBPass + "@" + cfg.DBHost +
		":" + cfg.DBPort + "/" + cfg.DBName + "?sslmode=" + cfg.SSLMode
	dbpool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal("Unable to connect to database")
	}
	defer dbpool.Close()
	log.Println("Connected to database")

	taskRepo := repository.NewTaskRepository(dbpool)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	r := chi.NewRouter()

	//MIDDLEWARE
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.AllowContentType("application/json"))

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", taskHandler.GetAllTasks)
		r.Post("/", taskHandler.CreateTask)
		r.Get("/{id}", taskHandler.GetTaskById)
		r.Put("/{id}", taskHandler.UpdateTask)
		r.Delete("/{id}", taskHandler.DeleteTask)
	})

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Printf("Server starting on port: %s", cfg.Port)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Server failed to start: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shut down: ", err)
	}
	log.Println("Server exited")
}
