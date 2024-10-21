package gitlab

import "github.com/go-git/go-git/v5/plumbing/object"

type Interface interface {
	DeleteRepo(path string) error
	LoadRepo(repoUrl string, path string) *object.Tree
}
