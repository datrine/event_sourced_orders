package routes

import (
	"errors"
	"log"
	"net/http"

	"github.com/datrine/commands"
	"github.com/datrine/http/apis/dtos"
	"github.com/datrine/queries"
	"github.com/datrine/utils"
	"github.com/labstack/echo"
)

func CreateOrder(c echo.Context) error {
	var input dtos.CreateOrderDTO
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	err = Validate.Struct(&input)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	orderId, err := commands.CreateOrderCommand(input)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	log.Println(orderId)
	return c.JSON(http.StatusOK, &utils.CreateOrderOkResponse{
		Message: "Order created",
		Data: utils.CreateOrderOkData{
			OrderId: orderId,
		},
	})
}

func GetOrders(c echo.Context) error {
	var orderId = c.Param("id")
	order, err := queries.GetOrders(0)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	log.Println(orderId)
	return c.JSON(http.StatusOK, &utils.GetOrdersOkResponse{
		Message: "Order list",
		Data:    order,
	})
}

func GetOrderByOrderId(c echo.Context) error {
	var orderId = c.Param("id")
	order, err := queries.GetOrderByIdQuery(orderId)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	log.Println(orderId)
	return c.JSON(http.StatusOK, &utils.GetOrderOkResponse{
		Message: "Order created",
		Data: utils.GetOrderOkData{
			Order: order,
		},
	})
}

func UpdateShippingAddress(c echo.Context) error {
	orderId := c.Param("id")
	if orderId == "" {
		return c.JSON(400, &utils.ErrResponse{
			Message: errors.New("no order id").Error(),
		})

	}
	var input dtos.UpdateShippingAddressDTO
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	input.OrderId = orderId
	err = Validate.Struct(&input)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	updateId, err := commands.UpdateShippingAddressCommand(input)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	log.Println(updateId)
	return c.JSON(http.StatusOK, &utils.UpdateShippingAddressOkResponse{
		Message: "Shipping Address updated",
		Data: utils.UpdateShippingAddressOkData{
			UpdateId: updateId,
		},
	})
}
