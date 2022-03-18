# Makefile to build go-sdk library

VDIR=v3

all: build test lint tidy

build:
	cd ${VDIR} && go build ./...

test: unittest integrationtest

unittest:
	cd ${VDIR} && go test `go list ./... | grep -v examples`

integrationtest:
	cd ${VDIR} && go test `go list ./... | grep -v examples` -tags=integration

lint:
	cd ${VDIR} && golangci-lint run --build-tags=integration

tidy:
	cd ${VDIR} && go mod tidy
