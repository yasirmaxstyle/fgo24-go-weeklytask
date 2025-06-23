package search

import (
	"go-cli/models"
	"sort"
	"strings"
)

type SearchEngine struct{}

func NewSearchEngine() *SearchEngine {
	return &SearchEngine{}
}

func (se *SearchEngine) SearchItems(menu *models.MenuData, searchTerm string) []models.MenuItem {
	searchTerm = strings.ToLower(strings.TrimSpace(searchTerm))
	if searchTerm == "" {
		return []models.MenuItem{}
	}

	var foundItems []models.MenuItem
	for _, category := range menu.MenuCategories {
		for _, item := range category.Items {
			if strings.Contains(strings.ToLower(item.Name), searchTerm) ||
				strings.Contains(strings.ToLower(item.Description), searchTerm) {
				foundItems = append(foundItems, item)
			}
		}
	}

	// Sort by rating descending
	sort.Slice(foundItems, func(i, j int) bool {
		return foundItems[i].Rating > foundItems[j].Rating
	})

	return foundItems
}

func (se *SearchEngine) FilterItems(menu *models.MenuData, filter models.Filter) []models.MenuItem {
	var filteredItems []models.MenuItem

	for _, category := range menu.MenuCategories {
		// Check if category is selected
		categorySelected := contains(filter.Categories, category.ID)
		if !categorySelected {
			continue
		}

		for _, item := range category.Items {
			if se.matchesFilter(item, filter) {
				filteredItems = append(filteredItems, item)
			}
		}
	}

	// Sort by rating descending
	sort.Slice(filteredItems, func(i, j int) bool {
		return filteredItems[i].Rating > filteredItems[j].Rating
	})

	return filteredItems
}

func (se *SearchEngine) matchesFilter(item models.MenuItem, filter models.Filter) bool {
	if filter.Available && !item.Available {
		return false
	}
	return true
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
