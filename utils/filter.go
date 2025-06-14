package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func (cli *CLI) filterMenu() {
	filter := cli.getFilterOptions()
	if filter == nil {
		return
	}

	var filteredItems []MenuItem
	for _, category := range cli.menu.MenuCategories {
		// Check if category is selected
		categorySelected := len(filter.Categories) == 0
		for _, selectedCat := range filter.Categories {
			if selectedCat == category.ID {
				categorySelected = true
				break
			}
		}

		if !categorySelected {
			continue
		}

		for _, item := range category.Items {
			if cli.matchesFilter(item, *filter) {
				filteredItems = append(filteredItems, item)
			}
		}
	}

	cli.clearScreen()
	cli.displayHeader()
	fmt.Println("FILTERED RESULTS")
	fmt.Printf("Found %d items\n\n", len(filteredItems))

	if len(filteredItems) == 0 {
		fmt.Println("No items match your filter criteria")
		cli.waitForEnter()
		return
	}

	cli.displayMenuItem(filteredItems, true)
	fmt.Println("0. Back to Main Menu")
	fmt.Print("\nSelect item to add to cart (or back): ")

	cli.scanner.Scan()
	choice, err := strconv.Atoi(cli.scanner.Text())
	if err != nil || choice < 1 {
		return
	}

	if choice == 0 {
		return
	}

	if choice <= len(filteredItems) {
		cli.addToCart(filteredItems[choice-1])
	}
}

// Get filter options using checkbox-style interface
func (cli *CLI) getFilterOptions() *Filter {
	cli.clearScreen()
	cli.displayHeader()

	fmt.Println("FILTER OPTIONS")
	fmt.Println("Select categories (enter numbers separated by commas, or press Enter for all):")

	for i, category := range cli.menu.MenuCategories {
		fmt.Printf("%d. %s\n", i+1, category.Name)
	}

	fmt.Print("\nCategories: ")
	cli.scanner.Scan()
	categoryInput := strings.TrimSpace(cli.scanner.Text())

	var selectedCategories []string
	if categoryInput != "" {
		parts := strings.Split(categoryInput, ",")
		for _, part := range parts {
			index, err := strconv.Atoi(strings.TrimSpace(part))
			if err == nil && index >= 1 && index <= len(cli.menu.MenuCategories) {
				selectedCategories = append(selectedCategories, cli.menu.MenuCategories[index-1].ID)
			}
		}
	}

	fmt.Println("\nTemperature options:")
	fmt.Println("1. Hot items")
	fmt.Println("2. Iced items")
	fmt.Print("Select temperature (1, 2, or both separated by comma): ")
	cli.scanner.Scan()
	tempInput := strings.TrimSpace(cli.scanner.Text())

	hot := strings.Contains(tempInput, "1")
	iced := strings.Contains(tempInput, "2")
	if tempInput == "" {
		hot, iced = true, true
	}

	return &Filter{
		Categories: selectedCategories,
		Hot:        hot,
		Iced:       iced,
		Available:  true,
	}
}

// Check if item matches filter
func (cli *CLI) matchesFilter(item MenuItem, filter Filter) bool {
	if filter.Available && !item.Available {
		return false
	}

	if !filter.Hot && !filter.Iced {
		return true
	}

	if filter.Hot && filter.Iced {
		return (item.Hot != nil && *item.Hot) || (item.Iced != nil && *item.Iced)
	}

	if filter.Hot && (item.Hot == nil || !*item.Hot) {
		return false
	}

	if filter.Iced && (item.Iced == nil || !*item.Iced) {
		return false
	}

	return true
}
