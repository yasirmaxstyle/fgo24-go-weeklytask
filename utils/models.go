package utils

// shop description home
type CoffeeShopInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

type MenuItem struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       int     `json:"price"`
	Category    string  `json:"category"`
	Available   bool    `json:"available"`
	Rating      float32 `json:"rating"`
}

type MenuCategory struct {
	ID    string     `json:"id"`
	Name  string     `json:"name"`
	Items []MenuItem `json:"items"`
}

type HomeMenuItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type MenuData struct {
	CoffeeShop     CoffeeShopInfo `json:"coffeeshop"`
	HomeMenu       []HomeMenuItem `json:"home_menu"`
	MenuCategories []MenuCategory `json:"menu_categories"`
}

type Filter struct {
	Categories []string
	Available  bool
	Rating     float32
}

type OrderItem struct {
	Item     MenuItem
	Quantity int
}
