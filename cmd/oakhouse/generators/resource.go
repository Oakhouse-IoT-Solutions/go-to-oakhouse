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
	createdFiles = append(createdFiles, fmt.Sprintf("entity/%s.go", strings.ToLower(name)))

	if err := GenerateRepository(name); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("repository/%s_repo.go", strings.ToLower(name)))

	if err := GenerateService(name); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("service/%s_service.go", strings.ToLower(name)))

	if err := GenerateHandler(name); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("handler/%s_handler.go", strings.ToLower(name)))

	if err := GenerateDTO(name); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("dto/%s/create_%s_dto.go", strings.ToLower(name), strings.ToLower(name)))
	createdFiles = append(createdFiles, fmt.Sprintf("dto/%s/update_%s_dto.go", strings.ToLower(name), strings.ToLower(name)))
	createdFiles = append(createdFiles, fmt.Sprintf("dto/%s/get_%s_dto.go", strings.ToLower(name), strings.ToLower(name)))

	if err := GenerateRoute(name); err != nil {
		return nil, err
	}
	createdFiles = append(createdFiles, fmt.Sprintf("route/%s.go", strings.ToLower(name)))

	return createdFiles, nil
}