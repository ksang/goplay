# Define parameters
BINARY=goplay
SHELL := /bin/bash
GOPACKAGES = $(shell go list ./... | grep -v vendor)
ROOTDIR = $(pwd)

.PHONY: build install test linux

default: build

build: main.go
	go build -v -o ./build/${BINARY} main.go

install:
	go install  ./...

test:
	go test -v ${GOPACKAGES}

clean:
	rm -rf build

linux: main.go
	GOOS=linux GOARCH=amd64 go build -o ./build/linux/${BINARY} main.go
