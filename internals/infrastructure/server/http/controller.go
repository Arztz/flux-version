package http

import (
	"flux-version/internals/controllers/gitlab"
	"flux-version/internals/controllers/healthcheck"
)

type Controller struct {
	gitlab      *gitlab.Controller
	healthcheck *healthcheck.Controller
}

func NewController(gitlab *gitlab.Controller, healthcheck *healthcheck.Controller) *Controller {
	return &Controller{
		gitlab:      gitlab,
		healthcheck: healthcheck,
	}
}
