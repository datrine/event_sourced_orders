package eventresolutions

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
	"github.com/datrine/conn"
	"github.com/datrine/events"
)

func ResolveOrders() {
	db := conn.DB
	options := esdb.ReadAllOptions{
		From:      esdb.Start{},
		Direction: esdb.Forwards,
	}
	stream, err := db.ReadAll(context.Background(), options, math.MaxUint64)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	defer stream.Close()

	for {
		resolvedEvent, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println(err.Error())
			break
		}

		if err != nil {
			panic(err)
		}
		event := resolvedEvent.Event
		eventType := event.EventType
		if eventType == events.OrderCreated.String() {
			CartCreatedResolution(event)
		}

		if eventType == events.ShippingAddressUpdated.String() {
			ShippingAddressUpdatedResolution(event)
		}

		if eventType == events.CartCreated.String() {
			CartCreatedResolution(event)
		}

		if eventType == events.CartDeleted.String() {
			CartDeletedResolution(event)
		}
		if eventType == events.ItemAddedToCart.String() {
			fmt.Println("ItemAddedToCart")
			ItemAddedToCartResolution(event)
		}
	}
}
