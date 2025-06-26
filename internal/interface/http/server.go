package web

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"majula/internal/core"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Service interface {
	GetPkg(name string) (core.GetPackumentRes, error)
	GetTarball(id string) (core.GetTarballRes, error)
	PublishPkg(name, version string, tags []string, manifest json.RawMessage, tar []byte) error
}

type Server struct {
	server  *http.Server
	service Service
	logger  *slog.Logger
}

func NewServer(s Service, l *slog.Logger) *Server {
	return &Server{
		service: s,
		logger:  l.With(slog.String("component", "http")),
	}
}

func (s *Server) Start(addr string) error {
	s.logger.Info(fmt.Sprintf("Starting majula server at %s", addr))

	r := chi.NewRouter()
	h := newHandler(s.service)

	r.Use(logger(s.logger))
	r.Use(recoverer(s.logger))

	r.Get("/-/whoami", h.handleGetWhoAmI)
	r.Get("/{package}", h.handleGetPkg)
	r.Put("/{package}", h.handlePutPkg)
	r.Get("/{package}/-/{tarball}", h.handleGetTarball)

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	s.server = srv

	return srv.ListenAndServe()
}

func (s *Server) Stop() {
	s.logger.Info("Gracefully shutting down majula server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error(fmt.Sprintf("Error shutting down majula server: %s", err))
	}

	if err := s.server.Close(); err != nil {
		s.logger.Error(fmt.Sprintf("Error forcibly shutting down majula server: %s", err))
	}
}
