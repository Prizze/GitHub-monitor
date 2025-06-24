package domain

import "time"

// Структура запроса query параметров
type QueryParamsRequest struct {
	Language string `schema:"lang" validate:"lang,required"`
	Number   uint32 `schema:"n" validate:"gte=1,lte=10"`
}

// Количество репозиториев в списке по умолчанию
const DefaultReposNumber uint32 = 5

// Структура ответа от API GitHub
type APIResponseDTO struct {
	TotalCount uint32          `json:"total_count"`
	Items      []GitHubRepoDTO `json:"items"`
}

// Структура с информацией об одном репозитории
type GitHubRepoDTO struct {
	FullName        string    `json:"full_name"`
	Url             string    `json:"url"`
	CreatedAt       time.Time `json:"created_at"`
	Language        string    `json:"language"`
	StargazersCount uint32    `json:"stargazers_count"`
}

// Поддерживаемые языки
var AllowedLanguages = map[string]struct{}{
	"go":         {},
	"python":     {},
	"rust":       {},
	"c++":        {},
	"java":       {},
	"javascript": {},
	"swift":      {},
}

const BaseGitHubApiURL string = "https://api.github.com/search/repositories"
