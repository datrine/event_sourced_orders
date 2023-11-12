package queries

import (
	"github.com/datrine/http/apis/dtos"
)

func GetOrderQuery(data dtos.GetOrderDTO) (*Order, error) {
	order, err := GetOrderByIdQuery(data.OrderId)
	if err != nil {
		return nil, err
	}
	return &Order{
		ID:              order.ID,
		CustomerId:      order.CustomerId,
		Status:          order.Status,
		Total:           order.Total,
		Currency:        order.Currency,
		Carts:           order.Carts,
		ShippingAddress: order.ShippingAddress,
	}, nil
}
