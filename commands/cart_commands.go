package commands

import (
	"github.com/datrine/domains"
	"github.com/datrine/events"
	eventdtos "github.com/datrine/events/event_dtos"
	"github.com/datrine/http/apis/dtos"
	"github.com/datrine/utils"
)

func CreateCartCommand(data dtos.CreateCartDTO) (string, error) {
	var itemsAdded []domains.Item
	cartId, _ := utils.GenerateCartID()
	for _, item := range data.Items {
		itemId, _ := utils.GenerateItemID()
		itemsAdded = append(itemsAdded, domains.Item{
			ID:        itemId,
			Name:      item.Name,
			UnitPrice: item.UnitPrice,
			Quantity:  item.Quantity,
		})
	}
	eventInput := eventdtos.CartCreatedDTO{
		ID:      cartId,
		OrderId: data.OrderId,
		Items:   itemsAdded,
	}
	_, err := events.CartCreatedEventHandler(eventInput)
	return cartId, err
}

func DeleteCartCommand(data dtos.DeleteCartDTO) (string, error) {
	eventInput := eventdtos.CartDeletedDTO{
		OrderId: data.OrderId,
		CartId:  data.CartId,
	}
	ooo, err := events.CartDeletedEventHandler(eventInput)
	cartId := ooo.CartId
	return cartId, err
}

func AddItemToCartCommand(data dtos.AddItemToCartDTO) (string, error) {
	itemId, _ := utils.GenerateItemID()
	item := domains.Item{
		ID:        itemId,
		Name:      data.Item.Name,
		UnitPrice: data.Item.UnitPrice,
		Quantity:  data.Item.Quantity,
	}
	eventInput := eventdtos.ItemAddedToCartDTO{
		OrderId: data.OrderId,
		CartId:  data.CartId,
		Item:    item,
	}
	_, err := events.ItemAddedToCartEventHandler(eventInput)
	return itemId, err
}

func EditItemInCartCommand(data dtos.EditItemInCartDTO) (string, error) {
	eventInput := eventdtos.ItemEditedInCartDTO{
		OrderId:  data.OrderId,
		CartId:   data.CartId,
		ItemId:   data.ItemId,
		Quantity: data.Quantity,
	}
	_, err := events.ItemEditedInCartEventHandler(eventInput)
	return data.ItemId, err
}

func RemoveItemFromCartCommand(data dtos.RemoveItemFromCartDTO) (string, error) {
	eventInput := eventdtos.ItemEditedInCartDTO{
		OrderId: data.OrderId,
		CartId:  data.CartId,
		ItemId:  data.ItemId,
	}
	_, err := events.ItemRemovedFromCartEventHandler(eventInput)
	return data.ItemId, err
}
