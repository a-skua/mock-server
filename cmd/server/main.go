package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/a-skua/mock-server"
)

var (
	port = flag.Int("port", 8080, "Port to run the HTTP server on")
)

func init() {
	flag.Parse()
}

func main() {
	addr := ":" + strconv.Itoa(*port)
	mux := server.NewServeMux()
	log.Printf("Starting server on %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
