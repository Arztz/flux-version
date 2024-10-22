package gitlab

import (
	"flux-version/types"
)

type Interface interface {
	MergeService(services []types.Service) []types.Service
	ReadFile(p string) (map[string][]types.Service, error)
	GenerateJSON(p types.Project, c map[string][]types.Service) (types.Project, error)
	Init()
	Healthcheck() bool
}
