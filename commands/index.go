package commands

import "github.com/datrine/domains"

type Order struct {
	ID              string
	CustomerId      string
	Status          string
	Total           float64
	Currency        string
	Carts           []domains.Cart
	ShippingAddress domains.Address
}
