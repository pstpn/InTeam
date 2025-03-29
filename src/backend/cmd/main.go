package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"backend/config"
	"backend/internal/router/handlers"
	api "backend/internal/router/ogen"
	"backend/internal/service"
	mypostgres2 "backend/internal/storage/postgres"
	"backend/pkg/http"
	"backend/pkg/logger"
	"backend/pkg/storage/postgres"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	l := logger.New(cfg.Logger.Level, os.Stdout)

	db, err := postgres.New(fmt.Sprintf("postgres://%s:%s@%s/%s",
		cfg.Database.Postgres.User,
		cfg.Database.Postgres.Password,
		net.JoinHostPort(cfg.Database.Postgres.Host, strconv.Itoa(cfg.Database.Postgres.Port)),
		cfg.Database.Postgres.Database,
	))
	if err != nil {
		log.Fatal(err)
	}

	userStorage := mypostgres2.NewUserStorage(db)
	cartStorage := mypostgres2.NewCartStorage(db)
	feedbackStorage := mypostgres2.NewFeedbackStorage(db)
	orderStorage := mypostgres2.NewOrderStorage(db)
	racketStorage := mypostgres2.NewRacketStorage(db)

	authService := service.NewAuthService(l, userStorage, cfg.Auth.SigningKey, cfg.Auth.AccessTokenTTL)
	userService := service.NewUserService(l, userStorage)
	cartService := service.NewCartService(l, cartStorage, racketStorage)
	feedbackService := service.NewFeedbackService(l, feedbackStorage)
	orderService := service.NewOrderService(l, orderStorage, cartStorage, racketStorage)
	racketService := service.NewRacketService(l, racketStorage)

	h := handlers.NewHandler(
		l,
		authService,
		cartService,
		userService,
		feedbackService,
		orderService,
		racketService,
	)
	securityHandler := handlers.NewSecurityHandler(userService, []byte(cfg.Auth.SigningKey))
	oas, err := api.NewServer(h, securityHandler)
	if err != nil {
		log.Fatal(err)
	}
	srv := http.NewServer(l)

	err = srv.Run(
		fmt.Sprintf("0.0.0.0:%d", cfg.HTTP.Port),
		http.Wrap(
			oas,
			http.CORSMiddleware(l),
			http.HeartbeatMiddleware("/healthcheck"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
