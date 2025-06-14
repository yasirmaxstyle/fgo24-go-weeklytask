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
		fmt.Printf("%d. %s\n", i+1, item.Title)
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
	fmt.Println("0. Back to Main Menu")
	fmt.Print("\nSelect category: ")
}

// Display items in a category
func (cli *CLI) displayCategoryItems(categoryIndex int) {
	cli.clearScreen()
	cli.displayHeader()

	category := cli.menu.MenuCategories[categoryIndex]
	fmt.Printf("%s\n", strings.ToUpper(category.Name))

	cli.displayMenuItem(category.Items, true)

	fmt.Println("0. Back to Categories")
	fmt.Print("\nSelect item to add to cart (or back): ")
}

func (cli *CLI) displayMenuItem(items []MenuItem, numbered bool) {
	for idx, item := range items {
		status := "✅"
		if !item.Available {
			status = "❌"
		}

		temp := ""
		if item.Hot != nil && item.Iced != nil {
			if *item.Hot && *item.Iced {
				temp = " 🔥❄️"
			} else if *item.Hot {
				temp = " 🔥"
			} else if *item.Iced {
				temp = " ❄️"
			}
		}

		numStr := ""
		if numbered {
			numStr = fmt.Sprintf("%d. ", idx+1)
		}

		fmt.Printf("%s %s%s%s\n", status, numStr, item.Name, temp)
		fmt.Printf("     Rp %s\n", FormatPrice(item.Price))
	}
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

func FormatPrice(price int) string {
	str := strconv.Itoa(price)
	n := len(str)
	if n <= 3 {
		return str
	}

	result := ""
	for i, digit := range str {
		if i > 0 && (n-i)%3 == 0 {
			result += ","
		}
		result += string(digit)
	}
	return result
}
