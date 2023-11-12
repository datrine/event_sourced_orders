package utils

import (
	"log"

	"github.com/jaevor/go-nanoid"
)

func GenerateOrderID() (string, error) {
	fn, err := nanoid.CustomASCII("ABCDEFGHIJKLMNOPQR0123456789", 10)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	orderId := "ORDER_" + fn()
	return orderId, nil
}

func GenerateCartID() (string, error) {
	fn, err := nanoid.CustomASCII("ABCDEFGHIJKLMNOPQR0123456789", 10)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	orderId := "CART_" + fn()
	return orderId, nil
}

func GenerateItemID() (string, error) {
	fn, err := nanoid.CustomASCII("ABCDEFGHIJKLMNOPQR0123456789", 10)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	orderId := "ITEM_" + fn()
	return orderId, nil
}
