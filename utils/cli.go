package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type CLI struct {
	// inputHandler *InputHandler
	menu        *MenuData
	scanner     *bufio.Scanner
	cart        []OrderItem
	currentView string
}

func NewCLI() *CLI {
	return &CLI{
		scanner:     bufio.NewScanner(os.Stdin),
		cart:        make([]OrderItem, 0),
		currentView: "main",
		// inputHandler: NewInputHandler(),
	}
}

func (cli *CLI) FetchMenuData() error {
	menuData, err := FetchMenuFromAPI()
	if err != nil {
		return err
	}
	cli.menu = menuData
	return nil
}

func (cli *CLI) Run() {
	for {
		cli.clearScreen()
		cli.displayMainMenu()
		cli.scanner.Scan()

		choice, err := strconv.Atoi(cli.scanner.Text())
		if err != nil {
			continue
		}

		switch choice {
		case 1: // Browse Menu
			for {
				cli.displayCategories()
				cli.scanner.Scan()

				categoryChoice, err := strconv.Atoi(cli.scanner.Text())
				if err != nil {
					continue
				}

				if categoryChoice == len(cli.menu.MenuCategories)+1 {
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
						if itemChoice == len(category.Items)+1 {
							break // Back to categories
						}

						if itemChoice >= 1 && itemChoice <= len(category.Items) {
							cli.addToCart(category.Items[itemChoice-1])
						}
					}
				}
			}

		case 2: // Search Menu
			cli.searchMenu()

		case 3: // Filter Menu
			cli.filterMenu()

		case 4: // View Cart
			cli.viewCart()

		case 5: // Checkout
			cli.checkout()

		case 0: // Exit
			fmt.Println("👋 Thank you for visiting Kopi Maxstyle!")
			return

		default:
			fmt.Println("Invalid choice")
			cli.waitForEnter()
		}
	}
}

// Wait for user input
func (cli *CLI) waitForEnter() {
	fmt.Print("\nPress Enter to continue...")
	cli.scanner.Scan()
}
