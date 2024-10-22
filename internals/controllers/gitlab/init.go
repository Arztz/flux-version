package gitlab

import (
	"flux-version/internals/config"
	"flux-version/internals/services/gitlab"
)

type Controller struct {
	service gitlab.Interface
	config  config.Configuration
}

func NewController(service gitlab.Interface, config config.Configuration) *Controller {
	return &Controller{
		service: service,
		config:  config,
	}
}
