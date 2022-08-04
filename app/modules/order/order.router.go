package order

type OrderRouter struct {
}

func NewOrderRouter(orderModule *OrderModule) *OrderRouter {
	orderModule.engine.GET("orders", orderModule.orderController.GetOrders)
	orderModule.engine.GET("orders/:id", orderModule.orderController.GetOrder)

	return new(OrderRouter)
}
