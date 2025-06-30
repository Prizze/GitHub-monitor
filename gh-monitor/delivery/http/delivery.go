package http

import (
	"encoding/json"
	"net/http"

	"github.com/Prizze/GitHub-monitor/gh-monitor/domain"
	"github.com/Prizze/GitHub-monitor/pkg/validate"
	"github.com/gorilla/schema"
)

type GHMonitorDelivery struct {
	usecase ghMonitorUsecase
}

func NewGHMonitorDelivery(usecase ghMonitorUsecase) *GHMonitorDelivery {
	return &GHMonitorDelivery{
		usecase: usecase,
	}
}

// GetTop godoc
// @Summary Получение статистики репозиториев по языку
// @Descripiton Возвращает список репозиториев с наибольшим количеством звезд по языку
// @Tags monitor
// @Produce json
// @Param lang query string true "Выбор языка"
// @Param n query int false "Лимит репозиториев: от 1 до 10, по умолчанию - 5"
// @Success 200 {object} domain.APIResponseDTO
// @Failure 500 {string} string "Внутренняя ошибка"
// @Router /top [get]
func (h *GHMonitorDelivery) GetTop(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	decoder := schema.NewDecoder()

	// Считывание query параметров
	var params domain.QueryParamsRequest
	if err := decoder.Decode(&params, r.Form); err != nil {
		http.Error(w, "Invalid query parameters", http.StatusBadRequest)
		return
	}

	// Количество репозиториев в списке по умолчанию
	if params.Number == 0 {
		params.Number = domain.DefaultReposNumber
	}

	// Валидация
	validator := validate.NewValidator()
	if err := validator.Validator.Struct(params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Получение статистики по языку
	langStat, err := h.usecase.GetLanguageStatistic(params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(langStat); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
