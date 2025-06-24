package usecase

import (
	"log"
	"sync"
	"time"

	"github.com/Prizze/GitHub-monitor/gh-monitor/domain/errors"

	"github.com/Prizze/GitHub-monitor/gh-monitor/domain"
)

var StatisticCache = map[string]*domain.APIResponseDTO{}

type GHMonitorUsecase struct {
	gitHubAPI gitHubAPI
}

func NewGHMonitorUsecase(gitHubAPI gitHubAPI) *GHMonitorUsecase {
	return &GHMonitorUsecase{
		gitHubAPI: gitHubAPI,
	}
}

func (uc *GHMonitorUsecase) GetLanguageStatistic(params domain.QueryParamsRequest) (*domain.APIResponseDTO, error) {
	langStat, ok := StatisticCache[params.Language]
	if !ok {
		return nil, errors.ErrLanguageUnexpected
	}

	if len(langStat.Items) > int(params.Number) {
		langStat.Items = langStat.Items[:params.Number]
	}

	return langStat, nil
}

func (uc *GHMonitorUsecase) InitFetching() {
	uc.fetchAllLanguages()

	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for {
		<-ticker.C
		uc.fetchAllLanguages()
	}
}

func (uc *GHMonitorUsecase) fetchAllLanguages() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	// Проход по всем поддерживаемым языкам
	for l := range domain.AllowedLanguages {
		wg.Add(1)
		// Горутина для обновления данных о языке
		go func(language string) {
			defer wg.Done()

			resp, err := uc.gitHubAPI.FetchLanguageStatistic(language)

			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				log.Println(err.Error())
				StatisticCache[language] = nil
			} else {
				StatisticCache[language] = resp
			}
		}(l)
	}

	wg.Wait()
}
