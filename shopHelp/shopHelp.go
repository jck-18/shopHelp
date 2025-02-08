package shopHelp

// Item represents a product that can be added to the cart.
type Item struct {
	Name  string
	Price float64
}

// Cart represents a shopping cart containing multiple items.
type Cart struct {
	items []Item
}

// NewCart initializes and returns a new, empty Cart.
func NewCart() *Cart {
	return &Cart{
		items: []Item{},
	}
}

// AddItem adds a new item to the cart.
func (c *Cart) AddItem(name string, price float64) {
	newItem := Item{
		Name:  name,
		Price: price,
	}
	c.items = append(c.items, newItem)
}

// Total calculates the total price of all items in the cart.
func (c *Cart) Total() float64 {
	var total float64
	for _, item := range c.items {
		total += item.Price
	}
	return total
}

// Items returns a copy of all items in the cart.
func (c *Cart) Items() []Item {
	// Return a copy so external code can’t directly modify internal items.
	itemsCopy := make([]Item, len(c.items))
	copy(itemsCopy, c.items)
	return itemsCopy
}
