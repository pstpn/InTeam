package model

import "time"

type OrderStatus string

const (
	OrderStatusInProgress = "InProgress"
	OrderStatusDone       = "Done"
)

type OrderLine struct {
	RacketID int `json:"racket_id"`
	Quantity int `json:"quantity"`
}

type Order struct {
	ID            int          `json:"id"`
	UserID        int          `json:"user_id"`
	CreationDate  time.Time    `json:"creation_date" format:"2006-01-02T15:07:00Z"`
	DeliveryDate  time.Time    `json:"delivery_date" format:"2006-01-02T15:07:00Z"`
	Address       string       `json:"address"`
	RecipientName string       `json:"recipient_name"`
	Status        OrderStatus  `json:"status"`
	Lines         []*OrderLine `json:"lines"`
	TotalPrice    float32      `json:"total_price"`
}
