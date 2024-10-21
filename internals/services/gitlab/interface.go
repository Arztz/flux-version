package gitlab

import (
	"flux-version/types"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type Interface interface {
	MergeService(services []types.Service) []types.Service
	InsertServices(path string, match []string, c map[string][]types.Service) map[string][]types.Service
	ReadFile(tree *object.Tree, p string) map[string][]types.Service
	GenerateJSON(p types.Project, c map[string][]types.Service) (types.Project, error)
	InitLoad() map[string]*object.Tree
}
