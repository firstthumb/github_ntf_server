# Makefile
#

PWD := $(shell pwd)

.PHONY: generate
generate:
	cd $(PWD)/rpc/repository
	protoc --twirp_out=. --go_out=. ./repository_service.proto
	cd $(PWD)

start:
	go run cmd/server/main.go
