package http

import "flux-version/internals/controllers/gitlab"

type Controller struct {
	gitlab *gitlab.Controller
}

func NewController(gitlab *gitlab.Controller) *Controller {
	return &Controller{
		gitlab: gitlab,
	}
}
