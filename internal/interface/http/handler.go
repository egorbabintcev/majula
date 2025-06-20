package web

import (
	"encoding/json"
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

func (h *handler) handleGetWhoAmI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, _ := json.Marshal(GetWhoAmIRes{
		Username: "system",
	})

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func NewRouter(s *core.Service, l *slog.Logger) http.Handler {
	r := chi.NewRouter()
	h := newHandler(s)

	r.Use(logger(l))
	r.Use(recoverer(l))

	r.Get("/-/whoami", h.handleGetWhoAmI)

	return r
}
