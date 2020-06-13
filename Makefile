SHELL := /bin/bash

GOPATH = $(shell if [[ -n $$GOPATH ]]; then echo $$GOPATH; else echo "."; fi)

default: install lint test vet

install:
	@GO111MODULE=off go get -u golang.org/x/lint/golint
	@GO111MODULE=off go build -o $(GOPATH)/bin/golint golang.org/x/lint/golint
	@GO111MODULE=on go mod vendor

test:
	@GO111MODULE=on go test -v -race -cover ./...

lint:
	@GO111MODULE=on go list ./... | xargs -n 1 -I {} $(GOPATH)/bin/golint -set_exit_status {}

vet:
	@GO111MODULE=on go vet ./...

help: 
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) \#/\1/'

.PHONY: install lint test vet help #