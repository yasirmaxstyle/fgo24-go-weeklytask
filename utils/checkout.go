package utils

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
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

		fmt.Printf("• %s x%d = Rp. %s\n",
			orderItem.Item.Name,
			orderItem.Quantity,
			cli.formatPrice(itemTotal))
	}

	fmt.Printf("\nTOTAL: %s\n\n", cli.formatPrice(total))

	fmt.Print("Confirm order? (y/n): ")
	cli.scanner.Scan()

	if strings.ToLower(strings.TrimSpace(cli.scanner.Text())) == "y" {
		var mu sync.Mutex
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("\nOrder confirmed!")
			fmt.Println("Order ID: CF" + strconv.FormatInt(time.Now().Unix(), 10))
			fmt.Print("Preparing your order. Please wait...")
			time.Sleep(3 * time.Second)
			mu.Lock()
			fmt.Println("\n\nThank you for your order!")
			mu.Unlock()
		}()
		wg.Wait()

		cli.cart = make([]OrderItem, 0) // Clear cart
		cli.waitForEnter()
	}
}
