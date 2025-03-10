package model

type Racket struct {
	ID        int     `json:"id"`
	Brand     string  `json:"brand"`
	Weight    float32 `json:"weight"`
	Balance   float32 `json:"balance"`
	HeadSize  float32 `json:"head_size"`
	Available bool    `json:"available"`
	Quantity  int     `json:"quantity"`
	Price     int     `json:"price"`
	Image     []byte  `json:"image"`
}
