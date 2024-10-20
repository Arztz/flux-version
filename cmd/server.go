package main

import (
	"bufio"
	"encoding/json"
	"flux-version/internals/config"
	"flux-version/internals/container"
	types "flux-version/types"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	var p *types.Project
	var currentCategory map[string][]types.Service
	var totalCategory []types.Category
	//var mergeCategory *types.Category
	//var currentProject *types.Project
	fmt.Println("starting....")
	server, err := container.NewContainer()
	if err := server.Start(); err != nil {
		log.Panic(err)
	}
	fmt.Println("started")
	appConfig := config.NewConfiguration()
	tagPattern := regexp.MustCompile(appConfig.TagPattern)
	versionPattern := regexp.MustCompile(appConfig.VersionPattern)
	fmt.Println("loaded config")

	p = &types.Project{Project: "fundii", Category: []types.Category{}}

	//Loop file
	fmt.Println("find file....")
	LoadRepo(appConfig).Files().ForEach(func(f *object.File) error {
		// Open each file for reading
		if filepath.Base(f.Name) == "patch.yaml" { //search patch.yaml
			filePath := filepath.Join(appConfig.ClonePath, f.Name)
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
					fmt.Printf("Found tag: '%s' in %s\n", tagMatch[1], f.Name)
					currentCategory = insertServices(f.Name, tagMatch, currentCategory)
				}
				if versionMatch := versionPattern.FindStringSubmatch(line); versionMatch != nil {
					// Print the version value found
					fmt.Printf("Found version: '%s' in %s\n", versionMatch[1], f.Name)
					currentCategory = insertServices(f.Name, versionMatch, currentCategory)
				}
			}
			//currentCategory = mergeCategory(currentCategory)
			//err = mergo.Merge(&totalCategory.Service, currentCategory, mergo.WithOverride)
			//array := []types.Category{*totalCategory}
			//if err != nil {
			//	return err
			//}
			//err = mergo.Merge(&p.Category, array, mergo.WithOverride)
			//
			//if err != nil {
			//	return err
			//}
			if err := scanner.Err(); err != nil {
				return err
			}
		}
		return nil

	})
	DeleteRepo(appConfig.ClonePath)
	fmt.Println("start merge project")
	// convert map service to Category
	for name, services := range currentCategory {
		totalCategory = append(totalCategory, types.Category{Name: name, Service: services})
	}
	p.Category = totalCategory
	projectJSON, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling category to JSON:", err)
		return
	}

	// Print the JSON output
	fmt.Println(string(projectJSON))
}

func LoadRepo(c config.Configuration) *object.Tree {
	//Remove Folder Before fetch new one
	var err error
	err = DeleteRepo(c.ClonePath)
	if err != nil {
		log.Println(err)
	}

	//Git Clone
	repo, err := git.PlainClone(c.ClonePath, false, &git.CloneOptions{
		Auth: &http.BasicAuth{
			Username: "oauth", // yes, this can be anything except an empty string
			Password: c.GitlabToken},

		URL:      c.RepoURL,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("cloned repo....")
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

func DeleteRepo(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf("failed to delete repository at %s: %w", path, err)
	}
	return nil
}

func mergeService(services []types.Service) []types.Service {
	merged := make(map[string]types.Service)

	for _, service := range services {
		if existing, ok := merged[service.Name]; ok {
			// Merge nonprod field if it is empty in the existing service
			if existing.NonProd == "" {
				existing.NonProd = service.NonProd
			}
			// Merge uat field if it is empty in the existing service
			if existing.UAT == "" {
				existing.UAT = service.UAT
			}
			merged[service.Name] = existing
		} else {
			// Add new service if it doesn't exist
			merged[service.Name] = service
		}
	}

	// Convert map back to slice
	result := make([]types.Service, 0, len(merged))
	for _, service := range merged {
		result = append(result, service)
	}

	return result
}

func insertServices(path string, match []string, c map[string][]types.Service) map[string][]types.Service {

	text := strings.Split(path, "/") //split path
	environment := text[1]           //nonprod
	serviceName := text[2]           //account-service
	if c == nil {
		c = make(map[string][]types.Service)
	}

	service := types.Service{Name: serviceName}

	if environment == "nonprod" || environment == "develop" {
		service.NonProd = match[1]
	} else if environment == "uat" {
		service.UAT = match[1]
	}

	c[text[0]] = append(c[text[0]], service)
	for k, v := range c {
		c[k] = mergeService(v)
	}

	return c
}
