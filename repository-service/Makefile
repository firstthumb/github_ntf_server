# Makefile
#

BINARY := repository-service
VERSION := 1.0.0
PWD := $(shell pwd)

.PHONY: windows
windows:
	mkdir -p release
	GOOS=windows GOARCH=amd64 go build -o release/$(BINARY)-v$(VERSION)-windows-amd64

.PHONY: linux
linux:
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/$(BINARY)-v$(VERSION)-linux-amd64

.PHONY: darwin
darwin:
	mkdir -p release
	GOOS=darwin GOARCH=amd64 go build -o release/$(BINARY)-v$(VERSION)-darwin-amd64

.PHONY: release
release: windows linux darwin

.PHONY: generate
generate:
	go run github.com/99designs/gqlgen -v
	cd $(PWD)/app/config/wire
	wire
	cd $(PWD)

start:
	go run cmd/server/main.go
