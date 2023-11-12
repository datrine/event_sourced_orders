package conn

import (
	"log"

	"github.com/EventStore/EventStore-Client-Go/v3/esdb"
)

var DB *esdb.Client

func StartConn() {
	log.Println("Connecting to stream...")
	settings, err := esdb.ParseConnectionString("esdb://localhost:2113?keepAliveTimeout=10000&keepAliveInterval=10000&tls=false")
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	db, err := esdb.NewClient(settings)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	log.Println("Connected to stream...")
	DB = db
}
