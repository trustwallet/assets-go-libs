.PHONY: default clean build

PWD := $(shell pwd)

APPS    := sanitychecker
BILDDIR ?= bin

default: clean build

build: $(APPS)

$(BILDDIR)/%:
	go build -o $@ ./cmd/$*

$(APPS): %: $(BILDDIR)/%

clean:
	@rm -f ${BILDDIR}/*

## test: Run unit tests
test:
	@echo "  >  Running unit tests"
	GOBIN=$(GOBIN) go test -cover -v ./...

## lint: Install and run linter
lint: go-lint-install go-lint

go-lint-install:
ifeq (,$(wildcard test -f bin/golangci-lint))
	@echo "  >  Installing golint"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s
endif

go-lint:
	@echo "  >  Running golint"
	bin/golangci-lint run --timeout=2m
