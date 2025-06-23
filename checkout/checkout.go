package checkout

import (
	"fmt"
	"go-cli/models"
	"strconv"
	"time"
)

type CheckoutProcessor struct {
	notificationCh chan string
	statusCh       chan string
}

func NewCheckoutProcessor() *CheckoutProcessor {
	return &CheckoutProcessor{
		notificationCh: make(chan string, 10), // Buffered channel for notifications
		statusCh:       make(chan string, 5),  // Buffered channel for status updates
	}
}

func (cp *CheckoutProcessor) ProcessCheckout(cart []models.OrderItem) error {
	if len(cart) == 0 {
		return fmt.Errorf("cart is empty")
	}

	orderID := "CF" + strconv.FormatInt(time.Now().Unix(), 10)

	// Start background processes using goroutines
	go func() {
		cp.processOrder(orderID)
	}()
	go cp.sendNotifications()
	go func() {
		cp.updateOrderStatus()
	}()

	// Wait for completion signals from channels
	cp.waitForCompletion()

	return nil
}

// processOrder runs in a separate goroutine
func (cp *CheckoutProcessor) processOrder(orderID string) {
	cp.notificationCh <- fmt.Sprintf("Order %s confirmed!", orderID)
	cp.statusCh <- "CONFIRMED"

	time.Sleep(1 * time.Second)
	cp.statusCh <- "PREPARING"
	cp.notificationCh <- "Kitchen is preparing your order..."

	time.Sleep(2 * time.Second)
	cp.statusCh <- "READY"
	cp.notificationCh <- "Your order is ready for pickup!"

	time.Sleep(1 * time.Second)
	cp.statusCh <- "COMPLETED"
	cp.notificationCh <- "Thank you for your order!"

	// Signal completion
	close(cp.notificationCh)
	close(cp.statusCh)
}

// sendNotifications runs in a separate goroutine
func (cp *CheckoutProcessor) sendNotifications() {
	for notification := range cp.notificationCh {
		fmt.Printf("%s\n", notification)
		time.Sleep(500 * time.Millisecond) // Simulate notification delay
	}
}

// updateOrderStatus runs in a separate goroutine
func (cp *CheckoutProcessor) updateOrderStatus() {
	for status := range cp.statusCh {
		switch status {
		case "CONFIRMED":
			fmt.Printf("Order Status: %s\n", status)
		case "PREPARING":
			fmt.Printf("Order Status: %s\n", status)
		case "READY":
			fmt.Printf("Order Status: %s\n", status)
		case "COMPLETED":
			fmt.Printf("Order Status: %s\n", status)
		}
		time.Sleep(300 * time.Millisecond)
	}
}

// waits for all goroutines to finish
func (cp *CheckoutProcessor) waitForCompletion() {
	time.Sleep(5 * time.Second)
}
