package events

import (
	"encoding/json"
	"log"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"github.com/datrine/conn"
	"github.com/datrine/domains"
	eventdtos "github.com/datrine/events/event_dtos"
)

func CartCreatedEventHandler(inputData eventdtos.CartCreatedDTO) (*CartCreatedResponse, error) {
	testEvent := CartCreatedEventData{
		Id:      inputData.ID,
		OrderId: inputData.OrderId,
		Items:   inputData.Items,
	}

	data, err := json.Marshal(testEvent)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	eventData := esdb.EventData{
		ContentType: esdb.ContentTypeJson,
		EventType:   CartCreated.String(),
		Data:        data,
	}
	_, err = conn.AppendOrderEvent(inputData.OrderId, &eventData)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &CartCreatedResponse{
		OrderId:   inputData.OrderId,
		CartId:    inputData.ID,
		EventData: &eventData,
	}, nil
}

func CartDeletedEventHandler(inputData eventdtos.CartDeletedDTO) (*CartDeletedResponse, error) {
	testEvent := CartDeletedEventData{
		Id:      inputData.CartId,
		OrderId: inputData.OrderId,
	}
	data, err := json.Marshal(testEvent)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	eventData := esdb.EventData{
		ContentType: esdb.ContentTypeJson,
		EventType:   CartDeleted.String(),
		Data:        data,
	}
	_, err = conn.AppendOrderEvent(inputData.OrderId, &eventData)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &CartDeletedResponse{
		OrderId:   inputData.OrderId,
		CartId:    inputData.CartId,
		EventData: &eventData,
	}, nil
}

func ItemAddedToCartEventHandler(inputData eventdtos.ItemAddedToCartDTO) (*ItemAddedToCartResponse, error) {
	testEvent := ItemAddedToCartEventData{
		OrderId: inputData.OrderId,
		CartId:  inputData.CartId,
		Item:    inputData.Item,
	}

	data, err := json.Marshal(testEvent)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	eventData := esdb.EventData{
		ContentType: esdb.ContentTypeJson,
		EventType:   ItemAddedToCart.String(),
		Data:        data,
	}
	_, err = conn.AppendOrderEvent(inputData.OrderId, &eventData)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &ItemAddedToCartResponse{
		OrderId:   inputData.OrderId,
		CartId:    inputData.CartId,
		ItemId:    inputData.Item.ID,
		EventData: &eventData,
	}, nil
}

func ItemEditedInCartEventHandler(inputData eventdtos.ItemEditedInCartDTO) (*ItemAddedToCartResponse, error) {
	testEvent := ItemEditedInCartEventData{
		OrderId:  inputData.OrderId,
		CartId:   inputData.CartId,
		ItemId:   inputData.ItemId,
		Quantity: inputData.Quantity,
	}

	data, err := json.Marshal(testEvent)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	eventData := esdb.EventData{
		ContentType: esdb.ContentTypeJson,
		EventType:   ItemEditedInCart.String(),
		Data:        data,
	}
	_, err = conn.AppendOrderEvent(inputData.OrderId, &eventData)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &ItemAddedToCartResponse{
		OrderId:   inputData.OrderId,
		CartId:    inputData.CartId,
		ItemId:    inputData.ItemId,
		EventData: &eventData,
	}, nil
}

func ItemRemovedFromCartEventHandler(inputData eventdtos.ItemEditedInCartDTO) (*ItemRemovedFromCartResponse, error) {
	testEvent := ItemRemovedFromCartEventData{
		OrderId: inputData.OrderId,
		CartId:  inputData.CartId,
		ItemId:  inputData.ItemId,
	}

	data, err := json.Marshal(testEvent)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	eventData := esdb.EventData{
		ContentType: esdb.ContentTypeJson,
		EventType:   ItemRemovedFromCart.String(),
		Data:        data,
	}
	_, err = conn.AppendOrderEvent(inputData.OrderId, &eventData)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &ItemRemovedFromCartResponse{
		OrderId:   inputData.OrderId,
		CartId:    inputData.CartId,
		ItemId:    inputData.ItemId,
		EventData: &eventData,
	}, nil
}

type CartCreatedEventData struct {
	Id      string
	OrderId string
	Items   []domains.Item
}

type CartCreatedResponse struct {
	OrderId   string
	CartId    string
	EventData *esdb.EventData
}

type CartDeletedEventData struct {
	Id      string
	OrderId string
}

type CartDeletedResponse struct {
	OrderId   string
	CartId    string
	EventData *esdb.EventData
}

type ItemAddedToCartEventData struct {
	OrderId string
	CartId  string
	domains.Item
}

type ItemAddedToCartResponse struct {
	OrderId   string
	CartId    string
	ItemId    string
	EventData *esdb.EventData
}

type ItemEditedInCartEventData struct {
	OrderId  string
	CartId   string
	ItemId   string
	Quantity uint
}

type ItemEditedInCartResponse struct {
	OrderId   string
	CartId    string
	ItemId    string
	EventData *esdb.EventData
}

type ItemRemovedFromCartEventData struct {
	OrderId string
	CartId  string
	ItemId  string
}

type ItemRemovedFromCartResponse struct {
	OrderId   string
	CartId    string
	ItemId    string
	EventData *esdb.EventData
}
