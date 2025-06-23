package cart

import (
	"fmt"
	"go-cli/models"
)

type Cart struct {
	items []models.OrderItem
}

func NewCart() *Cart {
	return &Cart{
		items: make([]models.OrderItem, 0),
	}
}

func (c *Cart) AddItem(item models.MenuItem, quantity int) error {
	if !item.Available {
		return fmt.Errorf("item %s is not available", item.Name)
	}

	if quantity < 1 {
		return fmt.Errorf("quantity must be at least 1")
	}

	// Check if item already in cart
	for i, cartItem := range c.items {
		if cartItem.Item.ID == item.ID {
			c.items[i].Quantity += quantity
			return nil
		}
	}

	// Add new item to cart
	c.items = append(c.items, models.OrderItem{
		Item:     item,
		Quantity: quantity,
	})

	return nil
}

func (c *Cart) GetItems() []models.OrderItem {
	return c.items
}

func (c *Cart) ClearCart() {
	c.items = make([]models.OrderItem, 0)
}

func (c *Cart) GetTotal() int {
	total := 0
	for _, orderItem := range c.items {
		total += orderItem.Item.Price * orderItem.Quantity
	}
	return total
}
