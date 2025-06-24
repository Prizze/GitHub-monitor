package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Prizze/GitHub-monitor/gh-monitor/domain"
	"github.com/Prizze/GitHub-monitor/gh-monitor/domain/errors"
)

type GitHubAPI struct {
	client *http.Client
}

func NewGitHubAPI(client *http.Client) *GitHubAPI {
	return &GitHubAPI{
		client: client,
	}
}

func (f *GitHubAPI) FetchLanguageStatistic(language string) (*domain.APIResponseDTO, error) {
	// Создаем запрос
	req, err := http.NewRequest(http.MethodGet, domain.BaseGitHubApiURL, nil)
	if err != nil {
		return nil, errors.ErrFailedCreateRequest
	}

	// Формируем URL параметры
	q := req.URL.Query()
	q.Add("q", fmt.Sprintf("language:%s", language))
	q.Add("sort", "stars")
	q.Add("order", "desc")
	req.URL.RawQuery = q.Encode()

	// Делаем запрос
	resp, err := f.client.Do(req)
	if err != nil {
		return nil, errors.ErrFailedRequest
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitHub API error: %s", string(body))
	}

	// Декодирование json
	var apiResponse domain.APIResponseDTO
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %s", err.Error())
	}

	return &apiResponse, nil
}
