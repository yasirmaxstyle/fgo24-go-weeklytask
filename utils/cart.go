package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// Add item to cart
func (cli *CLI) addToCart(item MenuItem) {
	if !item.Available {
		fmt.Println("Sorry, this item is not available")
		cli.waitForEnter()
		return
	}

	fmt.Printf("\nAdding %s to cart\n", item.Name)
	fmt.Print("Quantity: ")
	cli.scanner.Scan()

	quantity, err := strconv.Atoi(cli.scanner.Text())
	if err != nil || quantity < 1 {
		fmt.Println("Invalid quantity")
		cli.waitForEnter()
		return
	}

	// Check if item already in cart
	for i, cartItem := range cli.cart {
		if cartItem.Item.ID == item.ID {
			cli.cart[i].Quantity += quantity
			fmt.Printf("✅ Updated %s quantity to %d\n", item.Name, cli.cart[i].Quantity)
			cli.waitForEnter()
			return
		}
	}

	// Add new item to cart
	cli.cart = append(cli.cart, OrderItem{
		Item:     item,
		Quantity: quantity,
	})

	fmt.Printf("✅ Added %d x %s to cart\n", quantity, item.Name)
	cli.waitForEnter()
}

// View cart
func (cli *CLI) viewCart() {
	cli.clearScreen()
	cli.displayHeader()

	fmt.Println("🛒 YOUR CART")

	if len(cli.cart) == 0 {
		fmt.Println("Your cart is empty")
		fmt.Println("\n0. Back to Main Menu")
		fmt.Print("\nPress Enter to continue: ")
		cli.scanner.Scan()
		return
	}

	total := 0
	for i, orderItem := range cli.cart {
		itemTotal := orderItem.Item.Price * orderItem.Quantity
		total += itemTotal

		fmt.Printf("%d. %s x%d\n", i+1, orderItem.Item.Name, orderItem.Quantity)
		fmt.Printf("   %s each = %s\n\n",
			cli.formatPrice(orderItem.Item.Price),
			cli.formatPrice(itemTotal))
	}

	fmt.Printf("💰 TOTAL: %s\n\n", cli.formatPrice(total))

	fmt.Println("99. Clear Cart")
	fmt.Println("\n0. Back to Main Menu")
	fmt.Print("\nSelect option: ")

	cli.scanner.Scan()
	choice, err := strconv.Atoi(cli.scanner.Text())
	if err != nil {
		return
	}

	switch choice {
	case 99: // Clear cart
		fmt.Print("Are you sure? (y/n): ")
		cli.scanner.Scan()
		if strings.ToLower(strings.TrimSpace(cli.scanner.Text())) == "y" {
			cli.cart = make([]OrderItem, 0)
			fmt.Println("Cart cleared!")
			cli.waitForEnter()
		}
	}
}
