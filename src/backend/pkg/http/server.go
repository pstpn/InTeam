package http

import (
	"context"
	"net/http"
	"sync"
	"time"

	"backend/pkg/logger"
)

const (
	readHeaderTimeout = 5 * time.Second
	idleTimeout       = 30 * time.Second
	shutdownTimeout   = 2 * time.Second
)

type Server struct {
	l          logger.Interface
	httpServer *http.Server
	lock       sync.Mutex
}

func NewServer(l logger.Interface) *Server {
	return &Server{
		l: l,
	}
}

func (s *Server) Run(listen string, handler http.Handler) error {
	s.l.Infof("server started on %s", listen)

	s.lock.Lock()
	s.httpServer = &http.Server{
		Addr:              listen,
		Handler:           handler,
		ReadHeaderTimeout: readHeaderTimeout,
		IdleTimeout:       idleTimeout,
	}
	s.lock.Unlock()

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	s.lock.Lock()
	if s.httpServer != nil {
		_ = s.httpServer.Shutdown(ctx)
	}
	s.lock.Unlock()
}
