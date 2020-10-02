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

## build: run tests and compile application
build: check-path-included
	go test ./utils
	npm run --prefix ./frontend build:production
	go generate
	go mod tidy
	go build -ldflags $(GO_LDFLAGS) -o $(path)

## run: build application and run server
run: export DEBUG=true
run:
	npm run --prefix ./frontend build:development
	go generate
	go build -ldflags $(GO_LDFLAGS) -o /tmp/${APP_NAME} && /tmp/${APP_NAME} server

## help: prints this help message
help:
	@echo "Usage: "
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

check-path-included:
ifndef path
	$(error path is undefined; ex. path=/tmp/${APP_NAME})
endif
