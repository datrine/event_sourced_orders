package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

var Validate = validator.New()

func SetupRoutes() {
	e := echo.New()
	ordersRouters := e.Group("orders")
	ordersRouters.GET("", GetOrders)
	ordersRouters.POST("/", CreateOrder)
	ordersRouters.GET("/:id", GetOrderByOrderId)
	ordersRouters.PUT("/:id/shipping", UpdateShippingAddress)
	orderCartsRouters := ordersRouters.Group("/:id/carts")
	orderCartsRouters.DELETE("/:cartId", DeleteCart)
	orderCartsRouters.POST("", CreateCart)
	cartProdRouters := orderCartsRouters.Group("/:cartId/items")
	cartProdRouters.POST("", AddItemToCart)
	cartProdRouters.PUT("/:itemId", AddItemToCart)

	//AddProductToCart

	e.Logger.Fatal(e.Start(":1323"))
}
