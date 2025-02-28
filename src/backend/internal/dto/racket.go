package dto

import "backend/pkg/storage/postgres"

type CreateRacketReq struct {
	Brand    string  `json:"brand"`
	Weight   float32 `json:"weight"`
	Balance  float32 `json:"balance"`
	HeadSize float32 `json:"head_size"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
	Image    []byte  `json:"image"`
}

type UpdateRacketReq struct {
	ID       int `json:"id"`
	Quantity int `json:"quantity"`
}

type ListRacketsReq struct {
	Pagination *postgres.Pagination
}
