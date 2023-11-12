package eventresolutions

import (
	"encoding/json"
	"fmt"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"github.com/datrine/events"
	"github.com/datrine/projections"
)

func ShippingAddressUpdatedResolution(resolvedEvent *esdb.RecordedEvent) {
	Orders := projections.OrdersMap
	eventDats := resolvedEvent.Data
	shippingAddressUpdatedData := &events.ShippingAddressUpdatedEventData{}
	err := json.Unmarshal(eventDats, shippingAddressUpdatedData)
	if err != nil {
		panic(err)
	}
	order := Orders[shippingAddressUpdatedData.OrderId]
	if order == nil {
		fmt.Println("Nil reference")
		return
	}
	shippingAddress := (&(*order).ShippingAddress)
	if shippingAddressUpdatedData.City != "" {
		shippingAddress.City = shippingAddressUpdatedData.City
	}
	if shippingAddressUpdatedData.State != "" {
		shippingAddress.State = shippingAddressUpdatedData.State
	}
	if shippingAddressUpdatedData.Street != "" {
		shippingAddress.Street = shippingAddressUpdatedData.Street
	}
	if shippingAddressUpdatedData.ZipCode != "" {
		shippingAddress.ZipCode = shippingAddressUpdatedData.ZipCode
	}
	//order.ShippingAddress = shippingAddress

}
