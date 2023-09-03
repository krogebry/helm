REPO=docker.com/arene/build-infra/github-actions-runner
VERSION := $(shell cat VERSION)

build:
	docker build . -t ${REPO}:${VERSION}

push:
	docker push ${REPO}:${VERSION}
