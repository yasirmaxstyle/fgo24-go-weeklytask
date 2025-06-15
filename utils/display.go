package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func (cli *CLI) displayHeader() {
	fmt.Printf("\n" + strings.Repeat("=", 60) + "\n")
	fmt.Printf("☕ %s\n", cli.menu.CoffeeShop.Name)
	fmt.Printf("📍 %s\n", cli.menu.CoffeeShop.Location)
	fmt.Printf(strings.Repeat("=", 60) + "\n")
}

func (cli *CLI) displayMainMenu() {
	cli.displayHeader()

	fmt.Println("\n🏠 MAIN MENU")
	fmt.Println(strings.Repeat("-", 30))

	for i, item := range cli.menu.HomeMenu {
		if item.Title == "View Cart" {
			fmt.Printf("%d. %s (%d)\n", i+1, item.Title, len(cli.cart))
		} else {
			fmt.Printf("%d. %s\n", i+1, item.Title)
		}
	}

	fmt.Printf("\n0. Exit\n")
	fmt.Print("\nSelect an option: ")
}

func (cli *CLI) displayCategories() {
	cli.clearScreen()
	cli.displayHeader()

	fmt.Println("📋 MENU CATEGORIES")
	for i, category := range cli.menu.MenuCategories {
		fmt.Printf("%d. %s (%d items)\n", i+1, category.Name, len(category.Items))
	}
	fmt.Println("\n0. Back to Main Menu")
	fmt.Print("\nSelect category: ")
}

// Display items in a category
func (cli *CLI) displayCategoryItems(categoryIndex int) {
	cli.clearScreen()
	cli.displayHeader()

	category := cli.menu.MenuCategories[categoryIndex]
	fmt.Printf("%s\n", strings.ToUpper(category.Name))

	cli.displayMenu(category)

	fmt.Println("\n0. Back to Categories")
	fmt.Print("\nSelect item to add to cart (or back): ")
}

func (cli *CLI) displayPaginationInfo(currentPage, totalItems, itemsPerPage int) {
	totalPages := (totalItems + itemsPerPage - 1) / itemsPerPage
	startIdx := currentPage * itemsPerPage
	endIdx := min(startIdx+itemsPerPage, totalItems)

	fmt.Printf("\nPage %d of %d (Items %d-%d of %d)\n\n",
		currentPage+1, totalPages, startIdx+1, endIdx, totalItems)
}

func (cli *CLI) displayMenuItem(item MenuItem, numbered bool, idx int) {
	status := "Available"
	if !item.Available {
		status = "Out of stock"
	}

	numStr := ""
	if numbered {
		numStr = fmt.Sprintf("%d. ", idx+1)
	}

	fmt.Printf("\n" + strings.Repeat("=", 60) + "\n")
	fmt.Printf("%s%s ⭐ %.2f\nDesc: %s\nStock: %s\nPrice: %s\n", numStr, item.Name, item.Rating, item.Description, status, cli.formatPrice(item.Price))

}

func (cli *CLI) displayNavigationOptions() {
	fmt.Print("\nOptions: [n]ext, [p]revious, [s]elect, [0] back: ")
}

func (cli *CLI) clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Format price to IDR
func (cli *CLI) formatPrice(price int) string {
	return fmt.Sprintf("Rp. %s", cli.addCommas(price))
}

// Add commas to number
func (cli *CLI) addCommas(n int) string {
	str := strconv.Itoa(n)
	if len(str) <= 3 {
		return str
	}

	var result []string
	for i, digit := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result = append(result, ".")
		}
		result = append(result, string(digit))
	}

	return strings.Join(result, "")
}
