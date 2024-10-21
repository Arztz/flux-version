package container

import (
	"flux-version/internals/config"
	gitlabController "flux-version/internals/controllers/gitlab"
	"flux-version/internals/infrastructure/server/http"
	"flux-version/internals/repository/gitlab"
	gitlabService "flux-version/internals/services/gitlab"
	"flux-version/internals/utils/logrus"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type Container struct {
	container *dig.Container
}

func (c *Container) Configure() error {

	servicesConstructors := []interface{}{
		config.NewConfiguration,
		http.NewServer,
		http.NewController,
		gitlab.NewRepository,
		gitlabService.NewService,
		gitlabController.NewController,
	}
	for _, service := range servicesConstructors {
		if err := c.container.Provide(service); err != nil {
			return err
		}
	}
	logrus.NewLog()
	return nil
}

func (c *Container) Start() error {
	log.Info("Start Container")

	if err := c.container.Invoke(func(h *http.Server) {
		h.Start()
	}); err != nil {
		log.Errorf("%s", err)

		return err
	}

	return nil
}

func NewContainer() (*Container, error) {
	d := dig.New()

	container := &Container{
		container: d,
	}

	if err := container.Configure(); err != nil {
		return nil, err
	}

	return container, nil
}
