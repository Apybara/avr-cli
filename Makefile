GOPATH:=$(shell go env GOPATH)

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o avr-cli

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t avr-cli:latest .