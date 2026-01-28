.PHONY: all bootstrap lint test format docs

all: format lint test docs validate

bootstrap:
	./scripts/bootstrap

lint:
	./scripts/lint

test:
	./scripts/test

format:
	./scripts/format

docs:
	./scripts/docs

validate:
	go tool tfplugindocs validate --provider-name openprovider
