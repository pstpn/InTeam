package main

import (
	"context"
	"fmt"
	"net"
	net_http "net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"

	"backend/config"
	"backend/internal/router/handlers"
	api "backend/internal/router/ogen"
	"backend/internal/service"
	mypostgres2 "backend/internal/storage/postgres"
	"backend/pkg/common"
	"backend/pkg/http"
	"backend/pkg/logger"
	"backend/pkg/storage/postgres"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()

	gr, ctx := errgroup.WithContext(ctx)

	cfg, err := config.NewConfig()
	if err != nil {
		logger.New(zerolog.ErrorLevel.String(), os.Stdout).Errorf("new config: %s", err.Error())
		return
	}
	l := logger.New(cfg.Logger.Level, os.Stdout)

	db, err := postgres.New(fmt.Sprintf("postgres://%s:%s@%s/%s",
		cfg.Database.Postgres.User,
		cfg.Database.Postgres.Password,
		net.JoinHostPort(cfg.Database.Postgres.Host, strconv.Itoa(cfg.Database.Postgres.Port)),
		cfg.Database.Postgres.Database,
	))
	if err != nil {
		l.Errorf("new postgres: %s", err.Error())
		return
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
		l.Errorf("new oas server: %s", err.Error())
		return
	}

	srv := http.NewServer(l)
	metricsServer := http.NewServer(l)

	gr.Go(func() error {
		metricsMux := net_http.NewServeMux()
		metricsMux.Handle("/metrics", promhttp.Handler())
		return metricsServer.Run(
			fmt.Sprintf("0.0.0.0:%d", cfg.HTTP.MetricsPort),
			metricsMux,
		)
	})

	gr.Go(func() error {
		<-ctx.Done()
		metricsServer.Shutdown()
		srv.Shutdown()
		return nil
	})

	gr.Go(func() error {
		return srv.Run(
			fmt.Sprintf("0.0.0.0:%d", cfg.HTTP.Port),
			http.Wrap(
				oas,
				http.MetricsMiddleware(common.NewMetrics()),
				http.CORSMiddleware(l),
				http.HeartbeatMiddleware("/healthcheck"),
			),
		)
	})
	if err := gr.Wait(); err != nil {
		l.Errorf("err group wait: %s", err.Error())
	}
}
