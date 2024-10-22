package gitlab

import (
	"errors"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	log "github.com/sirupsen/logrus"
	"os"
)

func (r *Repository) CloneRepo(repoUrl string, p string) (*object.Tree, error) {
	rp := fmt.Sprintf("%s%s-flux.git", repoUrl, p)
	path := fmt.Sprintf("%s/%s", r.config.ClonePath, p)

	log.Printf("clone repo: %s", p)
	//Git Clone
	repo, err := git.PlainClone(path, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: "oauth", // yes, this can be anything except an empty string
			Password: r.config.GitlabToken},

		URL:      rp,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Println(err)
	}
	return getTree(repo)
}
func (r *Repository) PullRepo(repoUrl string, p string) (*object.Tree, error) {
	//rp := fmt.Sprintf("%s%s-flux.git", repoUrl, p)
	path := fmt.Sprintf("%s/%s", r.config.ClonePath, p)

	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	log.Printf("pull repo: %s", p)
	wk, _ := repo.Worktree()
	err = wk.Pull(&git.PullOptions{
		RemoteName:    "origin",
		ReferenceName: plumbing.ReferenceName("refs/heads/main"), // Replace 'main' with your branch
		Auth: &http.BasicAuth{
			Username: "oauth", // can be anything except an empty string
			Password: r.config.GitlabToken,
		},
	})
	if err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
		return nil, err
	}
	return getTree(repo)
}
func getTree(repo *git.Repository) (*object.Tree, error) {
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

	return tree, err
}
func (r *Repository) DeleteRepo(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf("failed to delete repository at %s: %w", path, err)
	}
	return nil
}
