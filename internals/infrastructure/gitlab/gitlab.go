package gitlab

import "flux-version/internals/config"

type Gitlab struct {
	env config.Configuration
}

func NewGitLab(env config.Configuration) *Gitlab {
	return &Gitlab{env: env}
}
