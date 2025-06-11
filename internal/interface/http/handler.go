package iface_http

import (
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

func NewRouter(s *core.Service) http.Handler {
	r := chi.NewRouter()
	h := newHandler(s)

	_ = h

	return r
}
