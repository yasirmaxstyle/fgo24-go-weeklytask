package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Checkout
func (cli *CLI) checkout() {
	if len(cli.cart) == 0 {
		fmt.Println("Your cart is empty")
		cli.waitForEnter()
		return
	}

	cli.clearScreen()
	cli.displayHeader()

	fmt.Println("CHECKOUT")
	fmt.Println("Order Summary:")

	total := 0
	for _, orderItem := range cli.cart {
		itemTotal := orderItem.Item.Price * orderItem.Quantity
		total += itemTotal

		fmt.Printf("• %s x%d = %s\n",
			orderItem.Item.Name,
			orderItem.Quantity,
			cli.formatPrice(itemTotal))
	}

	fmt.Printf("\nTOTAL: %s\n\n", cli.formatPrice(total))

	fmt.Print("Confirm order? (y/n): ")
	cli.scanner.Scan()

	if strings.ToLower(strings.TrimSpace(cli.scanner.Text())) == "y" {
		fmt.Println("\nOrder confirmed!")
		fmt.Println("Order ID: CF" + strconv.FormatInt(time.Now().Unix(), 10))
		fmt.Println("⏱ Estimated preparation time: 10-15 minutes")
		fmt.Println("Thank you for your order!")

		cli.cart = make([]OrderItem, 0) // Clear cart
		cli.waitForEnter()
	}
}

// Format price to IDR
func (cli *CLI) formatPrice(price int) string {
	return fmt.Sprintf("Rp %s", cli.addCommas(price))
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
			result = append(result, ",")
		}
		result = append(result, string(digit))
	}

	return strings.Join(result, "")
}
