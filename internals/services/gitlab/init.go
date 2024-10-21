package gitlab

import (
	"flux-version/internals/config"
	"flux-version/internals/repository/gitlab"
)

type Service struct {
	gitlabRepo gitlab.Interface
	config     config.Configuration
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
