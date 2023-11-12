package eventdtos

import "github.com/datrine/domains"

type OrderCreatedDTO struct {
	ID              string `validate:"required"`
	CustomerId      string `validate:"required"`
	Status          string ` validate:"required"`
	Total           string
	Currency        string
	Cart            []domains.Cart
	ShippingAddress domains.Address
}

type ShippingAddressUpdatedDTO struct {
	OrderId string
	domains.Address
}

type OrderUpdatedDTO struct {
	OrderId string
	Fields  UpdateFieldsDTO
	domains.Address
}

type UpdateFieldsDTO struct {
	OrderId string
	domains.Address
}

type CartUpdatedDTO struct {
	OrderId string
	Fields  UpdateFieldsDTO
	domains.Address
}

type CartFieldsUpdatedDTO struct {
	ID    string
	Items []domains.Item
}

type CartCreatedDTO struct {
	ID      string
	OrderId string
	Items   []domains.Item
}

type CartDeletedDTO struct {
	OrderId string
	CartId  string
	Cart    domains.Cart
}

type ItemAddedToCartDTO struct {
	OrderId string
	CartId  string
	Item    domains.Item
}

type ItemEditedInCartDTO struct {
	OrderId  string
	CartId   string
	ItemId   string
	Quantity uint
}

type ItemRemovedFromCartDTO struct {
	OrderId  string
	CartId   string
	ItemId   string
	Quantity uint
}
