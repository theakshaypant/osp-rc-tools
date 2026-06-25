package github

import (
	"github.com/google/go-github/v88/github"
)

func NewClient(token string) (*github.Client, error) {
	return github.NewClient(github.WithAuthToken(token))
}
