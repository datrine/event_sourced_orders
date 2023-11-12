package conn

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
)

func ReadStream(streamName string) {
	stream, err := DB.ReadStream(context.Background(), streamName, esdb.ReadStreamOptions{}, uint64(10))
	if err != nil {
		panic(err)
	}

	defer stream.Close()

	for {
		event, err := stream.Recv()

		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			panic(err)
		}

		// Doing something productive with the event
		fmt.Println(event)
	}
}
