package gitlab

import (
	"bufio"
	"flux-version/types"
	"fmt"
	"github.com/go-git/go-git/v5/plumbing/object"
	"os"
	"path/filepath"
	"regexp"
)

func (s *Service) ReadFile(p string) map[string][]types.Service {
	tagPattern := regexp.MustCompile(s.config.TagPattern)
	versionPattern := regexp.MustCompile(s.config.VersionPattern)
	var currentCategory map[string][]types.Service
	fmt.Println("Project: ", p)
	repo := fmt.Sprintf("%s%s-flux.git", s.config.RepoURL, p)
	err := s.gitlabRepo.LoadRepo(repo).Files().ForEach(func(f *object.File) error {
		// Open each file for reading
		if filepath.Base(f.Name) == "patch.yaml" { //search patch.yaml
			filePath := filepath.Join(s.config.ClonePath, f.Name)
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
					currentCategory = s.InsertServices(f.Name, tagMatch, currentCategory)
				}
				if versionMatch := versionPattern.FindStringSubmatch(line); versionMatch != nil {
					// Print the version value found
					//fmt.Printf("Found version: '%s' in %s\n", versionMatch[1], f.Name)
					currentCategory = s.InsertServices(f.Name, versionMatch, currentCategory)
				}
			}
			if err := scanner.Err(); err != nil {
				return err
			}
		}
		return nil

	})
	if err != nil {
		panic(err)
	}
	s.gitlabRepo.DeleteRepo(s.config.ClonePath)
	return currentCategory
}
