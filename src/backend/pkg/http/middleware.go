package http

import (
	"net/http"
	"strings"

	"github.com/rs/cors"
)

type Middleware = func(http.Handler) http.Handler

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

func CORSMiddleware() Middleware {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-XSRF-Token",
			"Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Origin"},
	}).Handler
}

func Wrap(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}

	return h
}
