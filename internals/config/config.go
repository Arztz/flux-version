package config

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type Configuration struct {
	TagPattern        string   `env:"TAG_PATTERN" envDefault:"tag:\\s*(\\S+)"`
	VersionPattern    string   `env:"VERSION_PATTERN" envDefault:"version:\\s*(\\S+)"`
	GitlabToken       string   `env:"GITLAB_TOKEN" envDefault:""`
	RepoURL           string   `env:"REPO_URL" envDefault:"https://git.robodev.co/robowealth/operation/fluxcd/"`
	ClonePath         string   `env:"REPO_URL" envDefault:"./repo"`
	HttpPort          string   `env:"HTTP_PORT" envDefault:"3001"`
	HTTPServerTimeout int      `env:"HTTP_SERVER_TIMEOUT" envDefault:"5"`
	ProjectList       []string `env:"PROJECT_LIST" envSeparator:"," envDefault:"roa,finvest,fundii,odini"`
}

func NewConfiguration() Configuration {
	config := Configuration{}

	if err := env.Parse(&config); err != nil {
		log.Errorf("%+v\n", err)
	}

	return config
}
