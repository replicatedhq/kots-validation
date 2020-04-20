SHELL := /bin/bash

export GO111MODULE=on

.PHONY: modules
modules:
	go mod download

.PHONY: test
test:
	go test -v `go list ./... | grep -v /vendor/ | grep -v /pact$$ | grep -v /e2e/`

.PHONY: build
build: fmt vet build-api

.PHONY: build-api
build-api:
	mkdir -p bin
	go build \
		-i \
		-o ./bin/kotsbots-api \
		./cmd


.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: run
run:
	./bin/kotsbots-api run

.PHONY: publish
publish:
	docker build -f deploy/Dockerfile -t registry.replicated.com/kots-bots/api:$${GITHUB_SHA} .
	docker push registry.replicated.com/kots-bots/api:$${GITHUB_SHA}


