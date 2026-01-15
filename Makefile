APP_NAME := app
CMD_DIR := ./backend/cmd/api

.PHONY: run build test fmt tidy lint check

run:
	go run ./backend/cmd/api

build:
	go build ./backend/cmd/api

test:
	go test ./...

fmt:
	goimports -w .

tidy:
	go mod tidy

lint:
	golangci-lint run

check: fmt tidy lint

