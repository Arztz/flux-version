package gitlab

import (
	"flux-version/internals/infrastructure/gitlab"
)

type GitlabRepository struct {
	gitlab *gitlab.Gitlab
}

func NewRepository(gitlab *gitlab.Gitlab) Interface {
	return &GitlabRepository{
		gitlab: gitlab,
	}
}
