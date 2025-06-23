package models

// MenuService handles menu data operations
type MenuService interface {
	FetchMenu() (*MenuData, error)
}

// CartService handles cart operations
type CartService interface {
	AddItem(item MenuItem, quantity int) error
	GetItems() []OrderItem
	ClearCart()
	GetTotal() int
}

// DisplayService handles UI display operations
type DisplayService interface {
	DisplayHeader(menu *MenuData)
	DisplayMainMenu(menu *MenuData, cartCount int)
	DisplayCategories(menu *MenuData)
	DisplayCategoryItems(category MenuCategory, categoryIndex int)
	DisplayMenuItem(item MenuItem, numbered bool, idx int)
	FormatPrice(price int) string
	ClearScreen()
	WaitForEnter()
}

// SearchService handles search and filter operations
type SearchService interface {
	SearchItems(menu *MenuData, searchTerm string) []MenuItem
	FilterItems(menu *MenuData, filter Filter) []MenuItem
}

// CheckoutService handles checkout operations
type CheckoutService interface {
	ProcessCheckout(cart []OrderItem) error
}
