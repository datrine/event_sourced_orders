package domains

import "time"

type Order struct {
	ID              string
	CustomerId      string
	Status          string
	OrderDate       time.Time
	Total           float64
	Currency        string
	Carts           []Cart
	ShippingAddress Address
}

type Cart struct {
	ID    string
	Items []Item
}

type Item struct {
	ID        string
	Name      string
	Quantity  uint
	UnitPrice float64
}

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
	Country string
}
