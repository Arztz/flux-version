package gitlab

import (
	"flux-version/internals/infrastructure/gitlab"
	"fmt"
	"os"
)

func (r *GitlabRepository) LoadRepo(gitlab *gitlab.Gitlab) *gitlab.Gitlab {
	return gitlab.LoadRepo()

	//repo, err := git.PlainClone(clonePath, false, &git.CloneOptions{
	//	Auth: &http.BasicAuth{
	//		Username: "oauth", // yes, this can be anything except an empty string
	//		Password: config.Configura},
	//
	//	URL:      repoURL,
	//	Progress: os.Stdout,
	//})
}

func DeleteRepo(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf("failed to delete repository at %s: %w", path, err)
	}
	return nil
}
