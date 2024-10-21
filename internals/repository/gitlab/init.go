package gitlab

import (
	"flux-version/internals/config"
)

type Repository struct {
	config config.Configuration
}

func NewRepository(config config.Configuration) Interface {
	return &Repository{
		config: config,
	}
}
