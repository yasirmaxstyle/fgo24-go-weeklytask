package ui

import (
	"fmt"
	"go-cli/models"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

type Display struct{}

func NewDisplay() *Display {
	return &Display{}
}

func (d *Display) DisplayHeader(menu *models.MenuData) {
	fmt.Printf("\n" + strings.Repeat("=", 60) + "\n")
	fmt.Printf("☕ %s\n", menu.CoffeeShop.Name)
	fmt.Printf("📍 %s\n", menu.CoffeeShop.Location)
	fmt.Printf(strings.Repeat("=", 60) + "\n")
}

func (d *Display) DisplayMainMenu(menu *models.MenuData, cartCount int) {
	d.DisplayHeader(menu)

	fmt.Println("\n🏠 MAIN MENU")
	fmt.Println(strings.Repeat("-", 30))

	for i, item := range menu.HomeMenu {
		if item.Title == "View Cart" {
			fmt.Printf("%d. %s (%d)\n", i+1, item.Title, cartCount)
		} else {
			fmt.Printf("%d. %s\n", i+1, item.Title)
		}
	}

	fmt.Printf("\n0. Exit\n")
	fmt.Print("\nSelect an option: ")
}

func (d *Display) DisplayCategories(menu *models.MenuData) {
	d.ClearScreen()
	d.DisplayHeader(menu)

	fmt.Println("📋 MENU CATEGORIES")
	for i, category := range menu.MenuCategories {
		fmt.Printf("%d. %s (%d items)\n", i+1, category.Name, len(category.Items))
	}
	fmt.Println("\n0. Back to Main Menu")
	fmt.Print("\nSelect category: ")
}

func (d *Display) DisplayCategoryItems(category models.MenuCategory, categoryIndex int) {
	d.ClearScreen()
	// Note: DisplayHeader needs menu data, this would need to be passed or restructured
	fmt.Printf("%s\n", strings.ToUpper(category.Name))

	fmt.Println("\n0. Back to Categories")
	fmt.Print("\nSelect item to add to cart (or back): ")
}

func (d *Display) DisplayMenuItem(item models.MenuItem, numbered bool, idx int) {
	status := "Available"
	if !item.Available {
		status = "Out of stock"
	}

	numStr := ""
	if numbered {
		numStr = fmt.Sprintf("%d. ", idx+1)
	}

	fmt.Printf("\n" + strings.Repeat("=", 60) + "\n")
	fmt.Printf("%s%s ⭐ %.2f\nDesc: %s\nStock: %s\nPrice: %s\n",
		numStr, item.Name, item.Rating, item.Description, status, d.FormatPrice(item.Price))
}

func (d *Display) FormatPrice(price int) string {
	return fmt.Sprintf("Rp. %s", d.addCommas(price))
}

func (d *Display) addCommas(n int) string {
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

func (d *Display) ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (d *Display) WaitForEnter() {
	fmt.Print("\nPress Enter to continue...")
}
