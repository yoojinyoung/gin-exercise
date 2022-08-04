package order

import "github.com/gin-gonic/gin"

type OrderModule struct {
	engine          *gin.Engine
	orderController *OrderController
	orderRouter     *OrderRouter
	orderService    *OrderService
}

func NewOrderModule(engine *gin.Engine) *OrderModule {
	orderModule := new(OrderModule)

	orderModule.engine = engine
	orderModule.orderController = NewOrderController(orderModule)
	orderModule.orderRouter = NewOrderRouter(orderModule)
	orderModule.orderService = NewOrderService(orderModule)

	return orderModule
}
