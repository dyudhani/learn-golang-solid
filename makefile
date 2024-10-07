IS_IN_PROGRESS="is in progress..."

.PHONY: all
all: env install mod

# help: prints this help message
.PHONY: help
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## env: will setup env
.PHONY: env
env:
	@echo "make env ${IS_IN_PROGRESS}"
	@go env -w GO111MODULE=on
	@go env -w GOBIN=`go env GOPATH`/bin
	@go env -w GOPROXY=https://proxy.golang.org,direct

## mod: will pull all dependency
.PHONY: mod
mod:
	@echo "make mod ${IS_IN_PROGRESS}"
	@rm -rf ./vendor ./go.sum
	@go mod tidy
	@go mod vendor

.PHONY: run
run:
	@go run cmd/api/main.go