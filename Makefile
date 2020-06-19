APP_NAME = scheduler
EPOCH_TIME = $(shell date +%s)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
SHELL = /bin/bash
VERSION=$(shell date +%s)
GO_LDFLAGS = '-X "github.com/clintjedwards/${APP_NAME}/cmd.appVersion=$(VERSION)" \
			   -X "github.com/clintjedwards/${APP_NAME}/api.appVersion=$(VERSION)"'

SEMVER = 0.0.1
SHELL = /bin/bash
VERSION = ${SEMVER}_${EPOCH_TIME}_${GIT_COMMIT}

## build-protos: build required protobuf files
build-protos:
	protoc --proto_path=proto --go_out=plugins=grpc:proto proto/*.proto
	protoc --js_out=import_style=commonjs,binary:./frontend/src/ --grpc-web_out=import_style=typescript,mode=grpcwebtext:./frontend/src/ -I ./proto/ proto/*.proto
## build: run tests and compile application
build: check-path-included
	protoc --proto_path=proto --go_out=plugins=grpc:proto proto/*.proto
	protoc --js_out=import_style=commonjs,binary:./frontend/src/ --grpc-web_out=import_style=typescript,mode=grpcwebtext:./frontend/src/ -I ./proto/ proto/*.proto
	go mod tidy
	npm run --prefix ./frontend build:production
	go test ./utils
	go generate
	go build -ldflags $(GO_LDFLAGS) -o $(path)

## run: build application and run server
run: export DEBUG=true
run:
	protoc --proto_path=proto --go_out=plugins=grpc:proto proto/*.proto
	protoc --js_out=import_style=commonjs,binary:./frontend/src/ --grpc-web_out=import_style=typescript,mode=grpcwebtext:./frontend/src/ -I ./proto/ proto/*.proto
	go mod tidy
	npm run --prefix ./frontend build:development
	go generate
	go build -ldflags $(GO_LDFLAGS) -o /tmp/${APP_NAME} && /tmp/${APP_NAME} server

## run: build application and run server
run-backend: export DEBUG=true
run-backend:
	protoc --proto_path=proto --go_out=plugins=grpc:proto proto/*.proto
	protoc --js_out=import_style=commonjs,binary:./frontend/src/ --grpc-web_out=import_style=typescript,mode=grpcwebtext:./frontend/src/ -I ./proto/ proto/*.proto
	go mod tidy
	go generate
	go build -ldflags $(GO_LDFLAGS) -o /tmp/${APP_NAME} && /tmp/${APP_NAME} server

## install: build application and install on system
install:
	protoc --proto_path=proto --go_out=plugins=grpc:proto proto/*.proto
	protoc --js_out=import_style=commonjs,binary:./frontend/src/ --grpc-web_out=import_style=typescript,mode=grpcwebtext:./frontend/src/ -I ./proto/ proto/*.proto
	go mod tidy
	npm run --prefix ./frontend build:production
	go generate
	go build -ldflags $(GO_LDFLAGS) -o /tmp/${APP_NAME}
	sudo mv /tmp/${APP_NAME} /usr/local/bin/
	chmod +x /usr/local/bin/${APP_NAME}

## help: prints this help message
help:
	@echo "Usage: "
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

check-path-included:
ifndef path
	$(error path is undefined; ex. path=/tmp/${APP_NAME})
endif
