SHELL=/bin/bash

tidy:
	@echo "Tidying up..."
	go mod tidy

test:
	@echo "Running tests..."
	@go test ./...

lint:
	@echo "Linting ..."
	@golangci-lint run --timeout 5m
