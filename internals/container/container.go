package container

import (
	"flux-version/internals/config"
	"flux-version/internals/infrastructure/gitlab"
	"flux-version/internals/utils/logrus"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
	"net/http"
)

type Container struct {
	container *dig.Container
}

func (c *Container) Configure() error {

	servicesConstructors := []interface{}{
		config.NewConfiguration,
		http.NewServeMux,
		gitlab.NewGitLab,
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
	//if err := c.container.Invoke(func(h *httpServer.Server) {
	//	go func() {
	//		_ = h.Start()
	//	}()
	//
	//}); err != nil {
	//	log.Errorf("%s", err)
	//
	//	return err
	//}

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
