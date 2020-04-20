SHELL := /bin/bash

.PHONY: publish
publish:
	docker build -f deploy/Dockerfile -t registry.replicated.com/kots-bots/migrations:$${GITHUB_SHA} .
	docker push registry.replicated.com/kots-bots/migrations:$${GITHUB_SHA}


