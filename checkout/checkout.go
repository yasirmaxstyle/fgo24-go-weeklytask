package checkout

import (
	"fmt"
	"go-cli/models"
	"strconv"
	"sync"
	"time"
)

type CheckoutProcessor struct{}

func NewCheckoutProcessor() *CheckoutProcessor {
	return &CheckoutProcessor{}
}

func (cp *CheckoutProcessor) ProcessCheckout(cart []models.OrderItem) error {
	if len(cart) == 0 {
		return fmt.Errorf("cart is empty")
	}

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
	return nil
}
