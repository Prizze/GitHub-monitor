package http

import "github.com/Prizze/GitHub-monitor/gh-monitor/domain"

type ghMonitorUsecase interface {
	GetLanguageStatistic(domain.QueryParamsRequest) (*domain.APIResponseDTO, error)
}
