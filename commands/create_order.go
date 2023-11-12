package commands

import (
	"github.com/datrine/events"
	eventdtos "github.com/datrine/events/event_dtos"
	"github.com/datrine/http/apis/dtos"
	"github.com/datrine/utils"
)

func CreateOrderCommand(data dtos.CreateOrderDTO) (string, error) {
	orderId, _ := utils.GenerateOrderID()
	eventInput := eventdtos.OrderCreatedDTO{
		ID:              orderId,
		CustomerId:      data.CustomerId,
		Status:          "PENDING",
		Total:           data.Total,
		Currency:        data.Currency,
		Cart:            data.Carts,
		ShippingAddress: data.ShippingAddress,
	}
	_, err := events.SaveOrderCreatedEventHandler(eventInput)
	return orderId, err
}
