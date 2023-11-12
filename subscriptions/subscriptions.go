package subscriptions

import (
	"context"
	"fmt"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"github.com/datrine/conn"
	eventresolutions "github.com/datrine/event_resolutions"
	"github.com/datrine/events"
)

func SubscribeToAllStreams(c chan bool) {
	subscription, err := conn.DB.SubscribeToAll(context.Background(),
		esdb.SubscribeToAllOptions{Filter: esdb.ExcludeSystemEventsFilter(), From: esdb.Start{}})
	if err != nil {
		panic(err)
	}

	//defer subscription.Close()

	for {
		subEvent := subscription.Recv()

		if subEvent.SubscriptionDropped != nil {
			fmt.Println(subEvent.SubscriptionDropped.Error.Error())
			break
		}

		if subEvent.EventAppeared == nil {
			fmt.Println("No event appeared in subscription")
			continue
		}

		event := subEvent.EventAppeared.Event
		eventType := event.EventType
		if eventType == events.OrderCreated.String() {
			eventresolutions.OrderCreatedResolution(event)
		}

		if eventType == events.ShippingAddressUpdated.String() {
			eventresolutions.ShippingAddressUpdatedResolution(event)
		}
		if eventType == events.CartCreated.String() {
			eventresolutions.CartCreatedResolution(event)
		}
		if eventType == events.CartDeleted.String() {
			eventresolutions.CartDeletedResolution(event)
		}
		if eventType == events.ItemAddedToCart.String() {
			eventresolutions.ItemAddedToCartResolution(event)
		}
	}
	c <- true
}
