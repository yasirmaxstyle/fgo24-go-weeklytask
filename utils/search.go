package utils

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
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

	var foundItems []MenuItem
	for _, category := range cli.menu.MenuCategories {
		for _, item := range category.Items {
			if strings.Contains(strings.ToLower(item.Name), searchTerm) ||
				strings.Contains(strings.ToLower(item.Description), searchTerm) {
				foundItems = append(foundItems, item)
			}
		}
	}

	// sort searched items by rating in descending order
	sort.Slice(foundItems, func(i, j int) bool {
		return foundItems[i].Rating > foundItems[j].Rating
	})

	var mu sync.Mutex
	var wg sync.WaitGroup

	cli.clearScreen()
	cli.displayHeader()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("🔍 SEARCHING RESULTS FOR: \"%s\" ...", searchTerm)
		time.Sleep(3 * time.Second)
		mu.Lock()
		fmt.Printf("Found %d items\n\n", len(foundItems))
		mu.Unlock()
	}()
	wg.Wait()

	if len(foundItems) == 0 {
		fmt.Println("No items found matching your search")
		cli.waitForEnter()
		return
	}

	searchCategory := MenuCategory{
		Name:  "Search Results",
		Items: foundItems,
	}

	if len(foundItems) >= ItemsPerPage {
		cli.displayMenu(searchCategory)
	} else {
		for idx, item := range foundItems {
			cli.displayMenuItem(item, true, idx)
		}
	}
	// cli.displayMenu(searchCategory)

	fmt.Println("\n0. Back to Main Menu")
	fmt.Print("\nSelect item to add to cart (or back): ")

	cli.scanner.Scan()
	choice, err := strconv.Atoi(cli.scanner.Text())
	if err != nil || choice < 1 {
		return
	}

	if choice == 0 {
		return
	}

	if choice <= len(foundItems) {
		cli.addToCart(foundItems[choice-1])
	}
}
