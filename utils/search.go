package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func (cli *CLI) searchMenu() {
	cli.clearScreen()
	cli.displayHeader()

	fmt.Println("🔍 SEARCH MENU")
	fmt.Print("Enter search term (name or description): ")
	cli.scanner.Scan()
	searchTerm := strings.ToLower(strings.TrimSpace(cli.scanner.Text()))

	if searchTerm == "" {
		fmt.Println("Please enter a search term")
		cli.waitForEnter()
		return
	}

	var results []MenuItem
	for _, category := range cli.menu.MenuCategories {
		for _, item := range category.Items {
			if strings.Contains(strings.ToLower(item.Name), searchTerm) {
				results = append(results, item)
			}
		}
	}

	cli.clearScreen()
	cli.displayHeader()
	fmt.Printf("🔍 SEARCH RESULTS FOR: \"%s\"\n", searchTerm)
	fmt.Printf("Found %d items\n\n", len(results))

	if len(results) == 0 {
		fmt.Println("No items found matching your search")
		cli.waitForEnter()
		return
	}

	cli.displayMenuItem(results, true)
	fmt.Println("\n0. Back to Main Menu")
	fmt.Print("\nSelect item to add to cart (or back): ")

	cli.scanner.Scan()
	choice, err := strconv.Atoi(cli.scanner.Text())
	if err != nil || choice < 1 {
		return
	}

	if choice == len(results)+1 {
		return
	}

	if choice <= len(results) {
		cli.addToCart(results[choice-1])
	}
}
