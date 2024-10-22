package gitlab

import (
	"flux-version/types"
	"strings"
)

func (s *Service) MergeService(services []types.Service) []types.Service {
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
			if existing.Prod == "" {
				existing.Prod = service.Prod
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
	//sort.Slice(result, func(i, j int) bool {
	//	return result[i].Name < result[j].Name
	//})
	return result
}

func (s *Service) insertServices(path string, tag string, c map[string][]types.Service) map[string][]types.Service {

	text := strings.Split(path, "/") //split path
	environment := text[1]           //nonprod
	serviceName := text[2]           //account-service
	if c == nil {
		c = make(map[string][]types.Service)
	}

	service := types.Service{Name: serviceName}

	switch environment {
	case "nonprod", "develop":
		service.NonProd = tag
	case "uat":
		service.UAT = tag
	case "prod":
		service.Prod = tag

	}

	c[text[0]] = append(c[text[0]], service)
	for k, v := range c {
		c[k] = s.MergeService(v)
	}

	return c
}

func (s *Service) GenerateJSON(p types.Project, c map[string][]types.Service) (types.Project, error) {

	var totalCategory []types.Category
	for name, services := range c {
		totalCategory = append(totalCategory, types.Category{Name: name, Service: services})
	}
	p.Category = totalCategory
	//projectJSON, err := json.MarshalIndent(p, "", "  ")
	//if err != nil {
	//	fmt.Println("Error marshalling category to JSON:", err)
	//	return nil, err
	//}
	return p, nil
}
