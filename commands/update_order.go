package commands

import (
	"github.com/datrine/domains"
	"github.com/datrine/events"
	eventdtos "github.com/datrine/events/event_dtos"
	"github.com/datrine/http/apis/dtos"
)

func UpdateOrderCommand(data dtos.UpdateShippingAddressDTO) (string, error) {
	eventInput := eventdtos.ShippingAddressUpdatedDTO{
		OrderId: data.OrderId,
		Address: domains.Address{
			City:    data.City,
			ZipCode: data.ZipCode,
			Country: data.Country,
			State:   data.State,
			Street:  data.Street,
		},
	}
	ooo, err := events.SaveShippingAddressUpdatedEventHandler(eventInput)
	updateId := ooo.UpdateId
	return updateId, err
}
