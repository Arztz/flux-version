package gitlab

import (
	"flux-version/internals/config"
	"flux-version/internals/services/gitlab"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Controller struct {
	service gitlab.Interface
	config  config.Configuration
	repo    map[string]*object.Tree
}

func NewController(service gitlab.Interface, config config.Configuration) *Controller {

	var repo = service.InitLoad()
	return &Controller{
		service: service,
		config:  config,
		repo:    repo,
	}
}
