package eventresolutions

import (
	"encoding/json"
	"fmt"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"github.com/datrine/domains"
	"github.com/datrine/events"
	"github.com/datrine/projections"
)

func CartCreatedResolution(resolvedEvent *esdb.RecordedEvent) {
	eventDats := resolvedEvent.Data
	createdCartData := &events.CartCreatedEventData{}
	err := json.Unmarshal(eventDats, createdCartData)
	if err != nil {
		panic(err)
	}
	order := projections.OrdersMap[createdCartData.OrderId]
	if order == nil {
		fmt.Println("Nil order reference")
		return
	}
	carts := order.Carts
	cart := domains.Cart{
		ID:    createdCartData.Id,
		Items: createdCartData.Items,
	}
	for _, item := range createdCartData.Items {
		cart.Items = append(cart.Items, domains.Item{
			Name:      item.Name,
			UnitPrice: item.UnitPrice,
			Quantity:  item.Quantity,
		})
	}
	carts = (append(carts, cart))
	order.Carts = carts
}

func CartDeletedResolution(resolvedEvent *esdb.RecordedEvent) {
	eventDats := resolvedEvent.Data
	deletedCartData := &events.CartDeletedEventData{}
	err := json.Unmarshal(eventDats, deletedCartData)
	if err != nil {
		panic(err)
	}
	order := projections.OrdersMap[deletedCartData.OrderId]
	if order == nil {
		fmt.Println("Nil order reference")
		return
	}
	carts := order.Carts
	for index, cart := range carts {
		if deletedCartData.Id == cart.ID {
			if index > 0 {
				carts = append(carts[:index-1], carts[index+1:]...)
			} else {
				carts = carts[1:]
			}

			break
		}
	}
	order.Carts = carts
}

func ItemAddedToCartResolution(resolvedEvent *esdb.RecordedEvent) {
	eventDats := resolvedEvent.Data
	itemAddedToCartData := &events.ItemAddedToCartEventData{}
	err := json.Unmarshal(eventDats, itemAddedToCartData)
	if err != nil {
		panic(err)
	}
	order := projections.OrdersMap[itemAddedToCartData.OrderId]
	if order == nil {
		fmt.Println("Nil order reference")
		return
	}
	carts := order.Carts
	for index, cart := range carts {
		if itemAddedToCartData.CartId == cart.ID {
			cart := &carts[index]
			cart.Items = append(cart.Items, itemAddedToCartData.Item)
			break
		}
	}
	order.Carts = carts
}

func ItemRemovedFromCartResolution(resolvedEvent *esdb.RecordedEvent) {
	eventDats := resolvedEvent.Data
	itemRemovedFromCartData := &events.ItemRemovedFromCartEventData{}
	err := json.Unmarshal(eventDats, itemRemovedFromCartData)
	if err != nil {
		panic(err)
	}
	order := projections.OrdersMap[itemRemovedFromCartData.OrderId]
	if order == nil {
		fmt.Println("Nil order reference")
		return
	}
	carts := order.Carts
	for index, cart := range carts {
		if itemRemovedFromCartData.CartId == cart.ID {
			cart := &carts[index]
			if index == 0 {
				cart.Items = cart.Items[1:]
			} else {
				cart.Items = append(cart.Items[:index-1], cart.Items[index+1:]...)
			}
			break
		}
	}
	order.Carts = carts
}
