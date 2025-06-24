package usecase

import "github.com/Prizze/GitHub-monitor/gh-monitor/domain"

type gitHubAPI interface {
	FetchLanguageStatistic(language string) (*domain.APIResponseDTO, error)
}
