package handler

import (
	"fmt"
	"io"
	"log/slog"
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

func Formdata(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseMultipartForm(1024 * 1024); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}
	keyValues := make([]any, 0, len(req.PostForm)*2)
	io.WriteString(w, "Received form data:\n")
	for key, values := range req.PostForm {
		for _, value := range values {
			keyValues = append(keyValues, key, value)
			io.WriteString(w, fmt.Sprintf("%s: %s\n", key, value))
		}
	}
	slog.Info("Parsed form data", keyValues...)
}
