package gitlab

import (
	"flux-version/internals/config"
	"flux-version/internals/repository/gitlab"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Service struct {
	gitlabRepo gitlab.Interface
	config     config.Configuration
	tree       map[string]*object.Tree
}

func NewService(
	gitlabRepo gitlab.Interface,
	config config.Configuration,
) Interface {
	return &Service{
		gitlabRepo: gitlabRepo,
		config:     config,
	}
}
