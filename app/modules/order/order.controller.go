package order

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	ordersService *OrderService
}

func NewOrderController(module *OrderModule) *OrderController {
	orderController := new(OrderController)
	orderController.ordersService = module.orderService

	return orderController
}

func (orderController *OrderController) GetOrders(ctx *gin.Context) {
	ctx.JSON(200, orderController.ordersService.GetOrders())
}

func (orderController *OrderController) GetOrder(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 0, 0)
	ctx.JSON(200, orderController.ordersService.GetOrder(id))
}
