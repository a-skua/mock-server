package server

import (
	"net/http"

	"github.com/a-skua/mock-server/handler"
)

func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Hello)
	mux.HandleFunc("/sleep", handler.Sleep)
	mux.HandleFunc("/notfound", handler.NotFound)
	return mux
}
