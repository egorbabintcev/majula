package web

import (
	"log/slog"
	"majula/internal/core"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	service *core.Service
}

func newHandler(s *core.Service) *handler {
	return &handler{
		service: s,
	}
}

func NewRouter(s *core.Service, l *slog.Logger) http.Handler {
	r := chi.NewRouter()
	h := newHandler(s)

	r.Use(logger(l))
	r.Use(recoverer(l))

	_ = h

	return r
}
