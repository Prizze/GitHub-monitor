package main

import (
	"fmt"
	"net/http"
	"time"

	delivery "github.com/Prizze/GitHub-monitor/gh-monitor/delivery/http"
	usecase "github.com/Prizze/GitHub-monitor/gh-monitor/usecase"
	github "github.com/Prizze/GitHub-monitor/infrastructure/github"
	"github.com/gorilla/mux"
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

	r.HandleFunc("/top", delivery.GetTop).Methods(http.MethodGet)

	// Запуск сервера
	fmt.Println("Сервис запущен на localhost:8080")
	http.ListenAndServe(":8080", r)
}
