package gitlab

import (
	"bufio"
	"errors"
	"flux-version/types"
	"fmt"
	"github.com/go-git/go-git/v5/plumbing/object"
	"os"
	"path/filepath"
	"regexp"
)

func (s *Service) ReadFile(p string) (map[string][]types.Service, error) {
	var (
		tagPattern      = regexp.MustCompile(s.config.TagPattern)
		versionPattern  = regexp.MustCompile(s.config.VersionPattern)
		currentCategory map[string][]types.Service
		err             error
	)

	s.tree[p], err = s.gitlabRepo.PullRepo(s.config.RepoURL, p)
	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("%s/%s", s.config.ClonePath, p)
	tree, ok := s.tree[p]
	if !ok {
		return nil, errors.New("no tree")
	}
	err = tree.Files().ForEach(func(f *object.File) error {
		// Open each file for reading
		if filepath.Base(f.Name) == "patch.yaml" { //search patch.yaml
			filePath := filepath.Join(path, f.Name)
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			// Search for the word in the file
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()
				// Check if line contains 'version:'
				if tagMatch := tagPattern.FindStringSubmatch(line); tagMatch != nil {
					// Print the tag value found
					//fmt.Printf("Found tag: '%s' in %s\n", tagMatch[1], f.Name)
					currentCategory = s.insertServices(f.Name, tagMatch[1], currentCategory)
				}
				if versionMatch := versionPattern.FindStringSubmatch(line); versionMatch != nil {
					// Print the version value found
					//fmt.Printf("Found version: '%s' in %s\n", versionMatch[1], f.Name)
					currentCategory = s.insertServices(f.Name, versionMatch[1], currentCategory)
				}
			}
			if err := scanner.Err(); err != nil {
				return err
			}
		}
		return nil

	})
	if err != nil {
		return nil, err
	}
	//s.gitlabRepo.DeleteRepo(s.config.ClonePath)
	return currentCategory, err
}

func (s *Service) Init() {
	s.tree = make(map[string]*object.Tree)
	var err error
	for _, p := range s.config.ProjectList {
		s.tree[p], err = s.gitlabRepo.CloneRepo(s.config.RepoURL, p)
		if err != nil {
			panic(err)
		}
	}

}

func (s *Service) Healthcheck() bool {
	return len(s.config.ProjectList) == len(s.tree)
}
