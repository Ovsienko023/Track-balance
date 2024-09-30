.PHONY: all
all: build docs

.PHONY: docs
docs:
	swag init
	npx @redocly/cli@v1.19.0 build-docs docs/swagger.yaml -o web/apidoc/v1/index.html

.PHONY: build
build:
	GOARCH=arm64 GOOS=darwin go build -o api-tb-darwin
	CC=x86_64-unknown-linux-gnu-gcc CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -o api-tb-linux

.PHONY: install
install:
	brew tap SergioBenitez/osxct
	brew install x86_64-unknown-linux-gnu
