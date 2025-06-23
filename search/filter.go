package search

import (
	"go-cli/models"
	"strconv"
	"strings"
)

func GetFilterOptions(menu *models.MenuData, categoryInput string) *models.Filter {
	var selectedCategories []string

	if categoryInput != "" {
		parts := strings.Split(categoryInput, ",")
		for _, part := range parts {
			index, err := strconv.Atoi(strings.TrimSpace(part))
			if err == nil && index >= 1 && index <= len(menu.MenuCategories) {
				selectedCategories = append(selectedCategories, menu.MenuCategories[index-1].ID)
			}
		}
	} else {
		// Select all categories if none specified
		for _, category := range menu.MenuCategories {
			selectedCategories = append(selectedCategories, category.ID)
		}
	}

	return &models.Filter{
		Categories: selectedCategories,
		Available:  true,
	}
}
