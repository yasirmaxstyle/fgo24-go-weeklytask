package cmd

import (
	"bufio"
	"fmt"
	"go-cli/api"
	"go-cli/cart"
	"go-cli/checkout"
	"go-cli/models"
	"go-cli/search"
	"go-cli/ui"
	"os"
	"strconv"
	"strings"
)

type CLI struct {
	menu            *models.MenuData
	scanner         *bufio.Scanner
	menuService     models.MenuService
	cartService     models.CartService
	displayService  models.DisplayService
	searchService   models.SearchService
	checkoutService models.CheckoutService
}

func NewCLI() *CLI {
	return &CLI{
		scanner:         bufio.NewScanner(os.Stdin),
		menuService:     api.NewAPIMenuService(),
		cartService:     cart.NewCart(),
		displayService:  ui.NewDisplay(),
		searchService:   search.NewSearchEngine(),
		checkoutService: checkout.NewCheckoutProcessor(),
	}
}

func (cli *CLI) FetchMenuData() error {
	menuData, err := cli.menuService.FetchMenu()
	if err != nil {
		return err
	}
	cli.menu = menuData
	return nil
}

func (cli *CLI) Run() {
	for {
		cli.displayService.ClearScreen()
		cartCount := len(cli.cartService.GetItems())
		cli.displayService.DisplayMainMenu(cli.menu, cartCount)
		cli.scanner.Scan()

		choice, err := strconv.Atoi(cli.scanner.Text())
		if err != nil {
			continue
		}

		switch choice {
		case 1: // Browse Menu
			cli.browseMenu()
		case 2: // Search Menu
			cli.searchMenu()
		case 3: // Filter Menu
			cli.filterMenu()
		case 4: // View Cart
			cli.viewCart()
		case 5: // Checkout
			cli.processCheckout()
		case 0: // Exit
			fmt.Println("👋 Thank you for visiting Kopi Maxstyle!")
			return
		default:
			fmt.Println("Invalid choice")
			cli.displayService.WaitForEnter()
			cli.scanner.Scan()
		}
	}
}

func (cli *CLI) browseMenu() {
	for {
		cli.displayService.DisplayCategories(cli.menu)
		cli.scanner.Scan()

		categoryChoice, err := strconv.Atoi(cli.scanner.Text())
		if err != nil {
			continue
		}

		if categoryChoice == 0 {
			break // Back to main menu
		}

		if categoryChoice >= 1 && categoryChoice <= len(cli.menu.MenuCategories) {
			cli.browseCategory(categoryChoice - 1)
		}
	}
}

func (cli *CLI) browseCategory(categoryIndex int) {
	category := cli.menu.MenuCategories[categoryIndex]
	pagination := ui.NewPagination(len(category.Items))

	for {
		cli.displayService.ClearScreen()
		cli.displayService.DisplayHeader(cli.menu)

		fmt.Printf("📋 %s\n", strings.ToUpper(category.Name))
		fmt.Printf("Page %d of %d\n", pagination.CurrentPage+1, pagination.GetTotalPages())
		fmt.Println(strings.Repeat("-", 60))

		// Display current page items
		currentItems := pagination.GetCurrentPageItems(category.Items)
		startIdx := pagination.GetStartIndex()

		for i, item := range currentItems {
			cli.displayService.DisplayMenuItem(item, true, startIdx+i)
		}

		// Display navigation options
		fmt.Println(strings.Repeat("-", 60))

		// Show pagination controls
		if pagination.CurrentPage > 0 {
			fmt.Println("p. Previous Page")
		}
		if pagination.CurrentPage < pagination.GetTotalPages()-1 {
			fmt.Println("n. Next Page")
		}

		fmt.Println("0. Back to Categories")
		fmt.Print("\nSelect item number to add to cart, 'n' for next, 'p' for previous, or '0' to go back: ")

		cli.scanner.Scan()
		input := strings.TrimSpace(cli.scanner.Text())

		if input == "0" {
			break // Back to categories
		}

		if input == "n" && pagination.NextPage() {
			continue
		}

		if input == "p" && pagination.PreviousPage() {
			continue
		}

		// Try to parse as item selection
		choice, err := strconv.Atoi(input)
		if err != nil {
			continue
		}

		// Adjust choice to be relative to current page
		relativeChoice := choice - startIdx - 1
		if relativeChoice >= 0 && relativeChoice < len(currentItems) {
			cli.addToCart(currentItems[relativeChoice])
		}
	}
}

func (cli *CLI) searchMenu() {
	cli.displayService.ClearScreen()
	cli.displayService.DisplayHeader(cli.menu)

	fmt.Println("🔍 SEARCH MENU")
	fmt.Print("Enter search term (name or description): ")
	cli.scanner.Scan()
	searchTerm := strings.TrimSpace(cli.scanner.Text())

	if searchTerm == "" {
		fmt.Println("Please enter a search term")
		cli.displayService.WaitForEnter()
		cli.scanner.Scan()
		return
	}

	foundItems := cli.searchService.SearchItems(cli.menu, searchTerm)

	if len(foundItems) == 0 {
		cli.displayService.ClearScreen()
		cli.displayService.DisplayHeader(cli.menu)
		fmt.Printf("🔍 No items found for: \"%s\"\n", searchTerm)
		cli.displayService.WaitForEnter()
		cli.scanner.Scan()
		return
	}

	cli.displaySearchResultsWithPagination(foundItems, fmt.Sprintf("🔍 Search results for: \"%s\"", searchTerm))
}

func (cli *CLI) filterMenu() {
	cli.displayService.ClearScreen()
	cli.displayService.DisplayHeader(cli.menu)

	fmt.Println("🔧 FILTER OPTIONS")
	fmt.Println("Select categories (enter numbers separated by commas, or press Enter for all):")

	for i, category := range cli.menu.MenuCategories {
		fmt.Printf("%d. %s\n", i+1, category.Name)
	}

	fmt.Print("\nCategories: ")
	cli.scanner.Scan()
	categoryInput := strings.TrimSpace(cli.scanner.Text())

	filter := search.GetFilterOptions(cli.menu, categoryInput)
	filteredItems := cli.searchService.FilterItems(cli.menu, *filter)

	if len(filteredItems) == 0 {
		cli.displayService.ClearScreen()
		cli.displayService.DisplayHeader(cli.menu)
		fmt.Println("🔧 No items match your filter criteria")
		cli.displayService.WaitForEnter()
		cli.scanner.Scan()
		return
	}

	cli.displaySearchResultsWithPagination(filteredItems, "🔧 Filtered Results")
}

func (cli *CLI) displaySearchResultsWithPagination(items []models.MenuItem, title string) {
	pagination := ui.NewPagination(len(items))

	for {
		cli.displayService.ClearScreen()
		cli.displayService.DisplayHeader(cli.menu)

		fmt.Printf("%s\n", title)
		fmt.Printf("Found %d items - Page %d of %d\n", len(items), pagination.CurrentPage+1, pagination.GetTotalPages())
		fmt.Println(strings.Repeat("-", 60))

		// Display current page items
		currentItems := pagination.GetCurrentPageItems(items)
		startIdx := pagination.GetStartIndex()

		for i, item := range currentItems {
			cli.displayService.DisplayMenuItem(item, true, startIdx+i)
		}

		// Display navigation options
		fmt.Println(strings.Repeat("-", 60))

		// Show pagination controls
		if pagination.CurrentPage > 0 {
			fmt.Println("p. Previous Page")
		}
		if pagination.CurrentPage < pagination.GetTotalPages()-1 {
			fmt.Println("n. Next Page")
		}

		fmt.Println("0. Back to Main Menu")
		fmt.Print("\nSelect item number to add to cart, 'n' for next, 'p' for previous, or '0' to go back: ")

		cli.scanner.Scan()
		input := strings.TrimSpace(cli.scanner.Text())

		if input == "0" {
			break // Back to main menu
		}

		if input == "n" && pagination.NextPage() {
			continue
		}

		if input == "p" && pagination.PreviousPage() {
			continue
		}

		// Try to parse as item selection
		choice, err := strconv.Atoi(input)
		if err != nil {
			continue
		}

		// Adjust choice to be relative to current page
		relativeChoice := choice - startIdx - 1
		if relativeChoice >= 0 && relativeChoice < len(currentItems) {
			cli.addToCart(currentItems[relativeChoice])
		}
	}
}

func (cli *CLI) addToCart(item models.MenuItem) {
	fmt.Printf("\nAdding %s to cart\n", item.Name)
	fmt.Print("Quantity: ")
	cli.scanner.Scan()

	quantity, err := strconv.Atoi(cli.scanner.Text())
	if err != nil || quantity < 1 {
		fmt.Println("Invalid quantity")
		cli.displayService.WaitForEnter()
		cli.scanner.Scan()
		return
	}

	err = cli.cartService.AddItem(item, quantity)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		cli.displayService.WaitForEnter()
		cli.scanner.Scan()
		return
	}

	fmt.Printf("✅ Added %d x %s to cart\n", quantity, item.Name)
	cli.displayService.WaitForEnter()
	cli.scanner.Scan()
}

func (cli *CLI) viewCart() {
	cli.displayService.ClearScreen()
	cli.displayService.DisplayHeader(cli.menu)

	fmt.Println("🛒 YOUR CART")
	cartItems := cli.cartService.GetItems()

	if len(cartItems) == 0 {
		fmt.Println("Your cart is empty")
		fmt.Println("\n0. Back to Main Menu")
		fmt.Print("\nPress Enter to continue: ")
		cli.scanner.Scan()
		return
	}

	for i, orderItem := range cartItems {
		itemTotal := orderItem.Item.Price * orderItem.Quantity
		fmt.Printf("%d. %s x%d\n", i+1, orderItem.Item.Name, orderItem.Quantity)
		fmt.Printf("   %s each = %s\n\n",
			cli.displayService.FormatPrice(orderItem.Item.Price),
			cli.displayService.FormatPrice(itemTotal))
	}

	fmt.Printf("💰 TOTAL: %s\n\n", cli.displayService.FormatPrice(cli.cartService.GetTotal()))

	fmt.Println("99. Clear Cart")
	fmt.Println("\n0. Back to Main Menu")
	fmt.Print("\nSelect option: ")

	cli.scanner.Scan()
	choice, err := strconv.Atoi(cli.scanner.Text())
	if err != nil {
		return
	}

	if choice == 99 {
		fmt.Print("Are you sure? (y/n): ")
		cli.scanner.Scan()
		if strings.ToLower(strings.TrimSpace(cli.scanner.Text())) == "y" {
			cli.cartService.ClearCart()
			fmt.Println("Cart cleared!")
			cli.displayService.WaitForEnter()
			cli.scanner.Scan()
		}
	}
}

func (cli *CLI) processCheckout() {
	cartItems := cli.cartService.GetItems()
	if len(cartItems) == 0 {
		fmt.Println("Your cart is empty")
		cli.displayService.WaitForEnter()
		cli.scanner.Scan()
		return
	}

	cli.displayService.ClearScreen()
	cli.displayService.DisplayHeader(cli.menu)

	fmt.Println("💳 CHECKOUT")
	fmt.Println("Order Summary:")

	for _, orderItem := range cartItems {
		itemTotal := orderItem.Item.Price * orderItem.Quantity
		fmt.Printf("• %s x%d = %s\n",
			orderItem.Item.Name,
			orderItem.Quantity,
			cli.displayService.FormatPrice(itemTotal))
	}

	fmt.Printf("\nTOTAL: %s\n\n", cli.displayService.FormatPrice(cli.cartService.GetTotal()))

	fmt.Print("Confirm order? (y/n): ")
	cli.scanner.Scan()

	if strings.ToLower(strings.TrimSpace(cli.scanner.Text())) == "y" {
		err := cli.checkoutService.ProcessCheckout(cartItems)
		if err != nil {
			fmt.Printf("Checkout failed: %s\n", err.Error())
		} else {
			cli.cartService.ClearCart() // Clear cart after successful checkout
		}
		cli.displayService.WaitForEnter()
		cli.scanner.Scan()
	}
}
