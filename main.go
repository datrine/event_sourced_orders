package main

import (
	"fmt"

	"github.com/datrine/conn"
	"github.com/datrine/http/apis/routes"
	"github.com/datrine/subscriptions"
)

func init() {
	conn.StartConn()
}
func main() {
	//eventresolutions.ResolveOrders()
	fmt.Println("running")
	ch1 := make(chan bool)
	go subscriptions.SubscribeToAllStreams(ch1)
	go routes.SetupRoutes()

	subDone := <-ch1
	fmt.Println(subDone)
}
