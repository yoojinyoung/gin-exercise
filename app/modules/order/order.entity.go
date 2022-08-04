package order

import "time"

type Order struct {
	Id        int64     `json:"id"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}
