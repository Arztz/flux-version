package gitlab

import "flux-version/types"

type Interface interface {
	MergeService(services []types.Service) []types.Service
	InsertServices(path string, match []string, c map[string][]types.Service) map[string][]types.Service
	ReadFile(p string) map[string][]types.Service
	GenerateJSON(p types.Project, c map[string][]types.Service) (types.Project, error)
}
