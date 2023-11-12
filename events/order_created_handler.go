package events

import (
	"encoding/json"
	"log"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"github.com/datrine/conn"
	"github.com/datrine/domains"
	eventdtos "github.com/datrine/events/event_dtos"
)

func SaveOrderCreatedEventHandler(inputData eventdtos.OrderCreatedDTO) (*OrderCreatedResponse, error) {

	testEvent := OrderCreatedEventData{
		ID:              inputData.ID,
		CustomerId:      inputData.CustomerId,
		ShippingAddress: inputData.ShippingAddress,
		Carts:           inputData.Cart,
		Total:           inputData.Total,
		Currency:        inputData.Currency,
		Status:          "Creating",
	}

	data, err := json.Marshal(testEvent)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	eventData := esdb.EventData{
		ContentType: esdb.ContentTypeJson,
		EventType:   OrderCreated.String(),
		Data:        data,
	}
	_, err = conn.AppendOrderEvent(inputData.ID, &eventData)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &OrderCreatedResponse{
		OrderId:   inputData.ID,
		EventData: &eventData,
	}, nil
}

type OrderCreatedEventData struct {
	ID              string
	CustomerId      string
	Status          string
	Total           string
	Currency        string
	Carts           []domains.Cart
	ShippingAddress domains.Address
}

type OrderCreatedResponse struct {
	OrderId   string
	CartId    string
	EventData *esdb.EventData
}
