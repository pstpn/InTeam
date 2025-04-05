package http

import (
	"net/http"
	"strings"
	"time"

	"github.com/rs/cors"

	"backend/pkg/common"
	"backend/pkg/logger"
)

type Middleware = func(http.Handler) http.Handler

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func HeartbeatMiddleware(endpoint string) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if (r.Method == http.MethodGet || r.Method == http.MethodHead) && strings.EqualFold(r.URL.Path, endpoint) {
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte("."))
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}

func CORSMiddleware(l logger.Interface) Middleware {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Authorization"},
		AllowCredentials: true,
		Logger:           l,
	}).Handler
}

func MetricsMiddleware(metrics *common.Metrics) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rw := &responseWriter{w, http.StatusOK}
			start := time.Now()
			h.ServeHTTP(rw, r)
			elapsed := time.Since(start)
			metrics.RequestCounter.WithLabelValues(r.URL.Path, r.Method, http.StatusText(rw.statusCode)).Inc()
			metrics.RequestDuration.WithLabelValues(r.URL.Path, r.Method).Observe(elapsed.Seconds())
		})
	}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func Wrap(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return h
}
