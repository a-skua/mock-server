package server

import (
	"log/slog"
	"net/http"

	"github.com/a-skua/mock-server/handler"
)

type RequestLog struct {
	next http.Handler
}

func NewRequestLog(next http.Handler) *RequestLog {
	return &RequestLog{next: next}
}

func (l *RequestLog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slog.Info("Received request", "method", r.Method, "path", r.URL.Path)
	l.next.ServeHTTP(w, r)
}

func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", NewRequestLog(http.HandlerFunc(handler.Hello)))
	mux.Handle("/sleep", NewRequestLog(http.HandlerFunc(handler.Sleep)))
	mux.Handle("/notfound", NewRequestLog(http.HandlerFunc(handler.NotFound)))
	mux.Handle("/formdata", NewRequestLog(http.HandlerFunc(handler.Formdata)))
	return mux
}
