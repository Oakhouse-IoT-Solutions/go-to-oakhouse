package generators

import (
	"fmt"
	"strings"
)

// generateResource generates a complete REST resource including model, repository, service, handler, DTOs and routes.
// This provides a full CRUD implementation following clean architecture principles with proper separation of concerns.
// Returns a list of all created files for user feedback.
// 🚀 Proudly Created by Htet Waiyan From Oakhouse 🏡
func GenerateResource(name string, fields []string) ([]string, error) {
	var createdFiles []string

	if err := GenerateModel(name, fields); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("model/%s.go", strings.ToLower(name)))

	if err := GenerateRepository(name); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("repository/%s_repo.go", strings.ToLower(name)))

	// Generate service interface
	if err := GenerateServiceInterface(name); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("service/%s_interface.go", strings.ToLower(name)))

	// Generate service implementation
	if err := GenerateService(name, fields); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("service/%s_service.go", strings.ToLower(name)))

	if err := GenerateHandler(name); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("handler/%s_handler.go", strings.ToLower(name)))

	if err := GenerateDTO(name, fields); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("dto/%s/create_%s_dto.go", strings.ToLower(name), strings.ToLower(name)))
	createdFiles = append(createdFiles, fmt.Sprintf("dto/%s/update_%s_dto.go", strings.ToLower(name), strings.ToLower(name)))
	createdFiles = append(createdFiles, fmt.Sprintf("dto/%s/get_%s_dto.go", strings.ToLower(name), strings.ToLower(name)))

	// Add scope generation - this was missing!
	if err := GenerateScope(name, "filter"); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("scope/%s/filter.go", strings.ToLower(name)))

	if err := GenerateRoute(name); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("route/%s.go", strings.ToLower(name)))

	return createdFiles, nil
}
