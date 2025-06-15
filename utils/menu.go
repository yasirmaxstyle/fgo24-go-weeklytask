package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func (cli *CLI) displayMenu(category MenuCategory) {
	pagination := NewPagination(len(category.Items))

	for {
		cli.clearScreen()
		cli.displayItemsPage(category, pagination)
		cli.displayNavigationOptions()

		if !cli.scanner.Scan() {
			break
		}

		input := strings.ToLower(strings.TrimSpace(cli.scanner.Text()))

		switch input {
		case "n", "next":
			if !pagination.NextPage() {
				fmt.Println("You're already on the last page!")
			}
		case "p", "prev", "previous":
			if !pagination.PreviousPage() {
				fmt.Println("You're already on the first page!")
			}
		case "s", "select":
			return
		case "0", "back":
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func (cli *CLI) displayItemsPage(category MenuCategory, pagination *Pagination) {
	cli.displayPaginationInfo(pagination.CurrentPage, pagination.TotalItems, pagination.ItemsPerPage)

	currentItems := pagination.GetCurrentPageItems(category.Items)
	startIdx := pagination.GetStartIndex()

	for i, item := range currentItems {
		cli.displayMenuItem(item, true, startIdx+i)
	}
}

func (cli *CLI) browseMenu() {
	for {
		cli.displayCategories()
		cli.scanner.Scan()

		categoryChoice, err := strconv.Atoi(cli.scanner.Text())
		if err != nil {
			continue
		}

		if categoryChoice == 0 {
			break // Back to main menu
		}

		if categoryChoice >= 1 && categoryChoice <= len(cli.menu.MenuCategories) {
			for {
				cli.displayCategoryItems(categoryChoice - 1)
				cli.scanner.Scan()

				itemChoice, err := strconv.Atoi(cli.scanner.Text())
				if err != nil {
					continue
				}

				category := cli.menu.MenuCategories[categoryChoice-1]
				if itemChoice == 0 {
					break // Back to categories
				}

				if itemChoice >= 1 && itemChoice <= len(category.Items) {
					cli.addToCart(category.Items[itemChoice-1])
				}
			}
		}
	}
}
