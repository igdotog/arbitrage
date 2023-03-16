export GO11MODULE=true

ifneq (,$(wildcard .env))
	include .env
	export
endif

.PHONY: update
update:
	go mod tidy

.PHONY: test
test:
	gotestsum --format=testname -- ./... -tags=units,integrations -cover

.PHONY: run-local
run-local:
	go build -o dist/app main.go && dist/app

.PHONY: run-local
build:
	docker-compose up -d --build