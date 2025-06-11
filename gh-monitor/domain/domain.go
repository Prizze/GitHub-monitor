package domain

import "time"

// Структура ответа от API GitHub
type APIResponseDTO struct {
	TotalCount uint32          `json:"total_count"`
	Items      []GitHubRepoDTO `json:"items"`
}

// Структура с ифнормацией об одном репозитории
type GitHubRepoDTO struct {
	FullName        string    `json:"full_name"`
	Url             string    `json:"url"`
	CreatedAt       time.Time `json:"created_at"`
	Language        string    `json:"language"`
	StargazersCount uint32    `json:"stargazers_count"`
}
