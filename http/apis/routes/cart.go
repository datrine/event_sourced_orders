package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/datrine/commands"
	"github.com/datrine/http/apis/dtos"
	"github.com/datrine/queries"
	"github.com/datrine/utils"
	"github.com/labstack/echo"
)

func CreateCart(c echo.Context) error {
	var input dtos.CreateCartDTO
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	orderId := c.Param("id")
	input.OrderId = orderId
	err = Validate.Struct(&input)
	fmt.Println(input)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	cartId, err := commands.CreateCartCommand(input)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	log.Println(cartId)
	return c.JSON(http.StatusOK, &utils.CreateCartOkResponse{
		Message: "Cart created",
		Data: utils.CreateCartOkData{
			CartId: cartId,
		},
	})
}

func DeleteCart(c echo.Context) error {
	var input dtos.DeleteCartDTO
	orderId := c.Param("id")
	cart_id := c.Param("cartId")
	input.OrderId = orderId
	input.CartId = cart_id
	err := Validate.Struct(&input)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	cartId, err := commands.DeleteCartCommand(input)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	log.Println(cartId)
	return c.JSON(http.StatusOK, &utils.DeleteCartOkResponse{
		Message: "Cart deleted",
		Data: utils.DeleteCartOkData{
			CartId: cartId,
		},
	})
}

func AddItemToCart(c echo.Context) error {
	var input dtos.AddItemToCartDTO
	var itemDTO dtos.ItemDTO
	err := c.Bind(&itemDTO)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})
	}
	orderId := c.Param("id")
	cart_id := c.Param("cartId")
	input.OrderId = orderId
	input.CartId = cart_id
	input.Item = itemDTO
	err = Validate.Struct(&input)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})
	}

	itemId, err := commands.AddItemToCartCommand(input)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})
	}

	log.Println(itemId)
	return c.JSON(http.StatusOK, &utils.AddItemToCartOkResponse{
		Message: "Item Added To cart successfully",
		Data: utils.AddItemToCartOkData{
			ItemId: itemId,
		},
	})
}

func EditItemToCart(c echo.Context) error {
	var editItem dtos.EditItemInCartDTO
	orderId := c.Param("id")
	cart_id := c.Param("cartId")
	item_id := c.Param("itemId")
	editItem.OrderId = orderId
	editItem.CartId = cart_id
	editItem.ItemId = item_id
	err := c.Bind(&editItem)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	err = Validate.Struct(&editItem)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	itemId, err := commands.EditItemInCartCommand(editItem)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	log.Println(itemId)
	return c.JSON(http.StatusOK, &utils.AddItemToCartOkResponse{
		Message: "Item Added To cart successfully",
		Data: utils.AddItemToCartOkData{
			ItemId: itemId,
		},
	})
}

func RemoveItemFromCart(c echo.Context) error {
	var editItem dtos.RemoveItemFromCartDTO
	orderId := c.Param("id")
	cart_id := c.Param("cartId")
	item_id := c.Param("itemId")
	editItem.OrderId = orderId
	editItem.CartId = cart_id
	editItem.ItemId = item_id
	err := c.Bind(&editItem)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	err = Validate.Struct(&editItem)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	itemId, err := commands.RemoveItemFromCartCommand(editItem)
	if err != nil {
		return c.JSON(400, &utils.ErrResponse{
			Message: err.Error(),
		})

	}
	log.Println(itemId)
	return c.JSON(http.StatusOK, &utils.RemoveItemFromCartOkResponse{
		Message: "Item removed from cart successfully",
		Data: utils.RemoveItemFromCartOkData{
			ItemId: itemId,
		},
	})
}

func GetCartByCartId(c echo.Context) error {
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
