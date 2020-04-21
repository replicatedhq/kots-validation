 SHELL := /bin/bash
PROJECT_NAME ?= web

.PHONY: deps
deps:
	npm install

.PHONY: serve
serve:
	NODE_ENV=dev node --max-old-space-size=6144 `npm bin`/webpack-dev-server --config webpack.config.js --progress -w --color --debug --env development

.PHONY: build-kots
build-kots: deps
	# NODE_ENV=production `npm bin`/webpack -p --config webpack.config.js --env kots

.PHONY: publish
publish:
	docker build -f deploy/Dockerfile -t registry.replicated.com/kots-validation/web:$${GITHUB_SHA} .
	docker push registry.replicated.com/kots-validation/web:$${GITHUB_SHA}
