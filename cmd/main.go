package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/Prizze/GitHub-monitor/docs"
	delivery "github.com/Prizze/GitHub-monitor/gh-monitor/delivery/http"
	usecase "github.com/Prizze/GitHub-monitor/gh-monitor/usecase"
	github "github.com/Prizze/GitHub-monitor/infrastructure/github"
	"github.com/Prizze/GitHub-monitor/internal/middleware"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// GitHubAPI
	client := &http.Client{Timeout: 10 * time.Second}
	githubAPI := github.NewGitHubAPI(client)
	// Usecase
	usecase := usecase.NewGHMonitorUsecase(githubAPI)
	go usecase.InitFetching()
	// Delivery
	delivery := delivery.NewGHMonitorDelivery(usecase)

	// Роутер
	r := mux.NewRouter()

	// Rate limiting
	rl := middleware.NewRateLimiter()
	r.Use(rl.RateLimitingMiddleware)
	// API
	r.HandleFunc("/top", delivery.GetTop).Methods(http.MethodGet)

	// Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Запуск сервера
	fmt.Println("Сервис запущен на localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server not started: ", err.Error())
		return
	}
}
