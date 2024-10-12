.PHONY: all
all: gen build docs

.PHONY: gen
gen:
	buf generate

.PHONY: docs
docs:
	npx @redocly/cli@v1.19.0 build-docs docs/generated/api/api.openapi.yaml -o web/doc/index.html

.PHONY: build
build:
	GOARCH=arm64 GOOS=darwin go build -o api-tb-darwin
	CC=x86_64-unknown-linux-gnu-gcc CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -o api-tb-linux

.PHONY: install
install:
	brew tap SergioBenitez/osxct
	brew install x86_64-unknown-linux-gnu
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
