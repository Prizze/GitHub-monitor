package github

import "net/http"

type GitHubAPI struct {
	client *http.Client
}

func NewGitHubAPI(client *http.Client) *GitHubAPI {
	return &GitHubAPI{
		client: client,
	}
}