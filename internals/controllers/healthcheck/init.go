package healthcheck

import "flux-version/internals/services/gitlab"

type Controller struct {
	gitlab gitlab.Interface
}

func NewController(gitlab gitlab.Interface) *Controller {
	return &Controller{
		gitlab: gitlab,
	}
}
