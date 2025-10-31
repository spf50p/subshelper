.PHONY: all build build-linux deps

all: deps build

build:
	@go build -o ./.bin/subshelper ./main.go

build-linux-amd64:
	@GOOS=linux GOARCH=amd64 go build -o ./.bin/subshelper-linux-amd64 ./main.go

build-darwin-arm64:
	@GOOS=darwin GOARCH=arm64 go build -o ./.bin/subshelper-darwin-arm64 ./main.go

deps:
	@go mod tidy
	@go mod download
