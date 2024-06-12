package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type Middleware struct {
	log *slog.Logger
}

func NewMiddleware(log *slog.Logger) *Middleware {
	return &Middleware{log: log}
}

func (m *Middleware) Logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		m.log.Info("started handling request", "method", r.Method, "url", r.URL.String())
		f(w, r)
		m.log.Info("sent response to request", "method", r.Method, "url", r.URL.String(), "duration", time.Since(start))
	}
}
