package handler

import (
	"io"
	"net/http"
	"time"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func Sleep(w http.ResponseWriter, req *http.Request) {
	time.Sleep(3 * time.Second)
	io.WriteString(w, "Awake now!\n")
}

func NotFound(w http.ResponseWriter, req *http.Request) {
	time.Sleep(1 * time.Second)
	http.NotFound(w, req)
}
