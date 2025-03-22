package main

import (
	api "backend/internal/router/ogen"
	"bytes"
	"context"
	ht "github.com/ogen-go/ogen/http"
	"log"
	"os"
)

func main() {
	client, err := api.NewClient(
		"http://localhost:8081",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	b, err := os.ReadFile("file.png")
	if err != nil {
		log.Fatal(err)
	}

	racket, err := client.RacketsCreateRacket(
		context.Background(),
		&api.RacketsCreateRacketReq{
			Balance:  0,
			Brand:    "example",
			HeadSize: 0,
			Price:    0,
			Quantity: 0,
			Weight:   0,
			Image: ht.MultipartFile{
				Name: "file.png",
				File: bytes.NewReader(b),
				Size: int64(len(b)),
			},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("res_example.png", racket.Racket.Image, 0777)
	if err != nil {
		log.Fatal(err)
	}
}
