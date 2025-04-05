package dto

import (
	"time"

	"backend/internal/model"
	"backend/pkg/storage/postgres"
)

type UpdateOrderReq struct {
	OrderID int               `json:"order_id"`
	Status  model.OrderStatus `json:"status"`
}

type PlaceOrderReq struct {
	UserID        int       `json:"user_id"`
	DeliveryDate  time.Time `json:"delivery_date" format:"2006-01-02T15:07:00Z"`
	Address       string    `json:"address"`
	RecepientName string    `json:"recipient_name"`
}

type ListOrdersReq struct {
	UserID     int
	Pagination *postgres.Pagination
}
