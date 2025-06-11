package main

import (
	"majula/internal/core"
	"majula/internal/infrastructure"
	iface_http "majula/internal/interface/http"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("MAJULA_REGISTRY_PORT")

	if port == "" {
		port = "8080"
	}

	addr := ":" + port

	st := infrastructure.NewMemoryStorage()
	s := core.NewService(st)
	r := iface_http.NewRouter(s)

	http.ListenAndServe(addr, r)
}
