package main

import (
	"net/http"

	delivery "github.com/Prizze/GitHub-monitor/gh-monitor/delivery/http"
	usecase "github.com/Prizze/GitHub-monitor/gh-monitor/usecase"
	"github.com/gorilla/mux"
)

func main() {
	// Usecase
	usecase := usecase.NewGHMonitorUsecase()
	// Delivery
	delivery := delivery.NewGHMonitorDelivery(usecase)

	// Роутер
	r := mux.NewRouter()

	r.HandleFunc("/top", delivery.GetTop).Methods(http.MethodGet)

	// Запуск сервера
	http.ListenAndServe(":8080", r)
}
