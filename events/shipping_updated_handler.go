package events

import (
	"encoding/json"
	"log"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"github.com/datrine/conn"
	"github.com/datrine/domains"
	eventdtos "github.com/datrine/events/event_dtos"
	"github.com/gofrs/uuid"
)

func SaveShippingAddressUpdatedEventHandler(inputData eventdtos.ShippingAddressUpdatedDTO) (*ShippingAddressUpdatedResponse, error) {
	orderId := uuid.Must(uuid.NewV4()).String()
	updateId := uuid.Must(uuid.NewV4()).String()
	testEvent := ShippingAddressUpdatedEventData{
		OrderId:  inputData.OrderId,
		UpdateId: updateId,
		Address: domains.Address{
			Street:  inputData.Street,
			State:   inputData.State,
			ZipCode: inputData.ZipCode,
			City:    inputData.City,
		},
	}

	data, err := json.Marshal(testEvent)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	eventData := esdb.EventData{
		ContentType: esdb.ContentTypeJson,
		EventType:   ShippingAddressUpdated.String(),
		Data:        data,
	}
	_, err = conn.AppendOrderEvent(orderId, &eventData)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &ShippingAddressUpdatedResponse{
		OrderId:   orderId,
		UpdateId:  updateId,
		EventData: &eventData,
	}, nil
}

type ShippingAddressUpdatedEventData struct {
	OrderId  string
	UpdateId string
	domains.Address
}

type ShippingAddressUpdatedResponse struct {
	OrderId   string
	UpdateId  string
	EventData *esdb.EventData
}
