SHELL=/bin/bash

tidy:
	@echo "Tidying up..."
	go mod tidy

test:
	@echo "Running tests..."
	@go test ./...

install:
	go mod tidy
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/kisielk/errcheck@latest
	go install github.com/axw/gocov/gocov@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install github.com/client9/misspell/cmd/misspell@latest

lint:
	staticcheck ./...
	# go vet ./...
	golint -set_exit_status  ./...
	errcheck -ignore 'os:.*,' ./... 
	gosec ./...
