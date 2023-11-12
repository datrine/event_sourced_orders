package dtos

import "github.com/datrine/domains"

type CreateOrderDTO struct {
	CustomerId      string `json:"customer_id" validate:"required"`
	Total           string
	Currency        string
	Carts           []domains.Cart
	ShippingAddress domains.Address
}

type UpdateShippingAddressDTO struct {
	OrderId string `json:"order_id" validate:"required"`
	Street  string `json:"street" validate:"required"`
	City    string `json:"city" validate:"required"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country" validate:"required"`
}

type GetOrdersDTO struct {
	Filters interface{}
}

type GetOrderDTO struct {
	OrderId string
}

type CreateCartDTO struct {
	OrderId string    `json:"order_id" validate:"required"`
	Items   []ItemDTO `json:"items" validate:"required,dive,min=1"`
}

type DeleteCartDTO struct {
	OrderId string `validate:"required"`
	CartId  string `validate:"required"`
}
type ItemDTO struct {
	Name      string  `json:"name" validate:"required"`
	Quantity  uint    `json:"quantity" validate:"required"`
	UnitPrice float64 `json:"unit_price" validate:"required"`
}

type AddItemToCartDTO struct {
	OrderId string  `validate:"required"`
	CartId  string  `validate:"required"`
	Item    ItemDTO `json:"item" validate:"required"`
}

type EditItemInCartDTO struct {
	OrderId  string `validate:"required"`
	CartId   string `validate:"required"`
	ItemId   string `validate:"required"`
	Quantity uint   `json:"quantity" validate:"required"`
}

type RemoveItemFromCartDTO struct {
	OrderId string `validate:"required"`
	CartId  string `validate:"required"`
	ItemId  string `validate:"required"`
}
