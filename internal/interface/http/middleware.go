package web

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)

func logger(l *slog.Logger) func(http.Handler) http.Handler {
	return func(n http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			n.ServeHTTP(w, r)

			l.Info("incoming HTTP request",
				slog.String("url", r.URL.Path),
				slog.String("method", r.Method),
				slog.String("remote_address", r.RemoteAddr),
			)
		})
	}
}

func recoverer(l *slog.Logger) func(http.Handler) http.Handler {
	return func(n http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					l.Error("panic caught",
						slog.String("url", r.URL.Path),
						slog.String("method", r.Method),
						slog.String("remote_address", r.RemoteAddr),
						slog.Any("error", rec),
						slog.String("stack", "\n"+string(debug.Stack())),
					)

					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()

			n.ServeHTTP(w, r)
		})
	}
}
