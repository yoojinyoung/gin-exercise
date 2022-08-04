package order

import (
	"time"
)

type OrderService struct {
}

func NewOrderService(module *OrderModule) *OrderService {
	orderService := new(OrderService)

	return orderService
}

func (orderService *OrderService) GetOrders() []*Order {
	order1 := &Order{
		Id:        1,
		Price:     20000,
		CreatedAt: time.Now(),
	}
	order2 := &Order{
		Id:        2,
		Price:     10000,
		CreatedAt: time.Now(),
	}

	orders := make([]*Order, 2)

	orders[0] = order1
	orders[1] = order2

	return orders
}

func (orderService *OrderService) GetOrder(id int64) *Order {
	order := &Order{
		Id:        id,
		Price:     10000,
		CreatedAt: time.Now(),
	}

	return order
}
