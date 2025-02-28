package main

import (
	"backend/internal/router/handlers"
	api "backend/internal/router/ogen"
	mypostgres2 "backend/internal/storage/postgres"
	"backend/pkg/http"
	"backend/pkg/logger"
	"fmt"
	"log"
	"os"

	"backend/config"

	"backend/internal/service"
	"backend/pkg/storage/postgres"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	loggerFile, err := os.OpenFile(
		cfg.Logger.File,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		log.Fatal(err)
	}
	l := logger.New(cfg.Logger.Level, loggerFile)

	db, err := postgres.New(fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.Database.Postgres.User,
		cfg.Database.Postgres.Password,
		cfg.Database.Postgres.Host,
		cfg.Database.Postgres.Port,
		cfg.Database.Postgres.Database,
	))
	if err != nil {
		log.Fatal(err)
	}

	userRepo := mypostgres2.NewUserStorage(db)

	authService := service.NewAuthService(l, userRepo, cfg.Auth.SigningKey, cfg.Auth.AccessTokenTTL)

	h := handlers.NewHandler(l, authService)
	oas, err := api.NewServer(h)
	if err != nil {
		log.Fatal(err)
	}
	srv := http.NewServer(l)

	err = srv.Run(
		fmt.Sprintf("0.0.0.0:%d", cfg.HTTP.Port),
		http.Wrap(oas, http.HeartbeatMiddleware("/healthcheck")),
	)
	if err != nil {
		log.Fatal(err)
	}
}
