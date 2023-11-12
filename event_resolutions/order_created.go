package eventresolutions

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"github.com/datrine/domains"
	"github.com/datrine/events"
	"github.com/datrine/projections"
)

var Orders map[string]domains.Order

func OrderCreatedResolution(resolvedEvent *esdb.RecordedEvent) {
	eventDats := resolvedEvent.Data
	createdOrderData := &events.OrderCreatedEventData{}
	err := json.Unmarshal(eventDats, createdOrderData)
	if err != nil {
		panic(err)
	}
	Total, err := strconv.ParseFloat(createdOrderData.Total, 64)

	if err != nil {
		fmt.Println("No total set")
		createdOrderData.Total = "0"
	}
	projections.OrdersMap[createdOrderData.ID] = &domains.Order{
		ID:              createdOrderData.ID,
		CustomerId:      createdOrderData.CustomerId,
		Status:          createdOrderData.Status,
		Total:           float64(Total),
		Currency:        createdOrderData.Currency,
		Carts:           createdOrderData.Carts,
		ShippingAddress: createdOrderData.ShippingAddress,
	}
}
