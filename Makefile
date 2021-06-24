.PHONY: build up down test

# ----------------------------------------------------------------------------
# VARIABLES
#  All images will be tagged with `latest` for ease of use in the local dev
#  system. Eventually when we get to CI/CD we should tag via GH

PROJECT=refactored-memory
TAG=latest

build:
	docker build -f Dockerfile -t $(PROJECT)/backend:$(TAG) .

up: build
	docker-compose -f docker-compose.yaml up --remove-orphans

down:
	docker-compose -f docker-compose.yaml down

# TODO: implement testing...
test:
	docker-compose -f docker-compose.test.yaml up --remove-orphans

