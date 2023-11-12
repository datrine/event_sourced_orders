package queries

import (
	"github.com/datrine/domains"
	"github.com/datrine/projections"
)

func GetOrders(filters interface{}) (map[string]*domains.Order, error) {
	return projections.OrdersMap, nil
}

func GetOrderByIdQuery(orderId string) (*domains.Order, error) {
	order := projections.OrdersMap[orderId]
	return order, nil
}

type Order struct {
	ID              string
	CustomerId      string
	Status          string
	Total           float64
	Currency        string
	Carts           []domains.Cart
	ShippingAddress domains.Address
}
