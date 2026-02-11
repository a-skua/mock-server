PORT ?= 8080

.PHONY: serve
serve:
	go run ./cmd/server -port $(PORT)

.PHONY: fmt
fmt:
	go fmt ./...
