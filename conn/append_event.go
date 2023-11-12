package conn

import (
	"context"
	"strings"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
)

func AppendOrderEvent(orderId string, event *esdb.EventData) (*esdb.WriteResult, error) {
	streamId := strings.Join([]string{"orders", orderId}, ":")
	result, err := DB.AppendToStream(context.Background(), streamId, esdb.AppendToStreamOptions{}, *event)
	return result, err
}
