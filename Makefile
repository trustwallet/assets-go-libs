#! /usr/bin/make -f

# Go related variables.
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin

# Go files.
GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

all: fmt lint test

test:
	@echo "  >  Running unit tests"
	GOBIN=$(GOBIN) go test -cover -race -coverprofile=coverage.txt -covermode=atomic -v ./...

fmt:
	@echo "  >  Format all go files"
	GOBIN=$(GOBIN) gofmt -w ${GOFMT_FILES}

lint: go-lint-install go-lint

go-lint-install:
ifeq (,$(wildcard test -f bin/golangci-lint))
	@echo "  >  Installing golint"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/v1.45.2/install.sh | sh -s
endif

go-lint:
	@echo "  >  Running golint"
	bin/golangci-lint run --timeout=2m