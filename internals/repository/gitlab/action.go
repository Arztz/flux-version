package gitlab

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	log "github.com/sirupsen/logrus"
	"os"
)

func (r *Repository) LoadRepo(repoUrl string) *object.Tree {
	//Remove Folder Before fetch new one
	var err error
	err = r.DeleteRepo(r.config.ClonePath)
	if err != nil {
		log.Println(err)
	}

	//Git Clone
	repo, err := git.PlainClone(r.config.ClonePath, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: "oauth", // yes, this can be anything except an empty string
			Password: r.config.GitlabToken},

		URL:      repoUrl,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Println(err)
	}

	//Checkout to Head
	ref, err := repo.Head()
	if err != nil {
		log.Fatal(err)
	}
	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		log.Fatal(err)
	}

	//Get file to tree
	tree, err := commit.Tree()
	if err != nil {
		log.Fatal(err)
	}

	return tree
}

func (r *Repository) DeleteRepo(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf("failed to delete repository at %s: %w", path, err)
	}
	return nil
}
