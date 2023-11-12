package utils

import "github.com/datrine/domains"

type ErrResponse struct {
	Message string
}

type CreateOrderOkResponse struct {
	Message string
	Data    CreateOrderOkData
}

type CreateOrderOkData struct {
	OrderId string `json:"order_id"`
}

type CreateCartOkResponse struct {
	Message string
	Data    CreateCartOkData
}
type CreateCartOkData struct {
	CartId string `json:"cart_id"`
}

type DeleteCartOkResponse struct {
	Message string
	Data    DeleteCartOkData
}

type DeleteCartOkData struct {
	CartId string `json:"cart_id"`
}

type UpdateShippingAddressOkResponse struct {
	Message string
	Data    UpdateShippingAddressOkData
}

type UpdateShippingAddressOkData struct {
	UpdateId string `json:"update_id"`
}

type GetOrderOkResponse struct {
	Message string
	Data    GetOrderOkData
}

type GetOrderOkData struct {
	*domains.Order
}

type GetOrdersOkResponse struct {
	Message string
	Data    interface{}
}

type GetOrdersOkData struct {
	*domains.Order
}

type AddItemToCartOkResponse struct {
	Message string
	Data    AddItemToCartOkData
}

type AddItemToCartOkData struct {
	ItemId string `json:"item_id"`
}

type RemoveItemFromCartOkResponse struct {
	Message string
	Data    RemoveItemFromCartOkData
}

type RemoveItemFromCartOkData struct {
	ItemId string `json:"item_id"`
}
