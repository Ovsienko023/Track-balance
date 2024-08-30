#.PHONY: all
#all: build_darwin2 build_linux2

.PHONY: build
build:
	gox -osarch="linux/amd64 darwin/arm64" \
	-output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}"

.PHONY: build_darwin
build_darwin:
	gox -osarch="darwin/arm64" \
	-output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}"

.PHONY: build_linux
build_linux:
	gox -osarch="linux/amd64" \
	-output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}"

.PHONY: build_linux2
build_linux2:
	GO111MODULE=on CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -tags "linux" -o bin/main_linux main.go

.PHONY: build_darwin2
build_darwin2:
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o bin/main_darwin main.go


.PHONY: docs
docs:
	swag init
	npx @redocly/cli@v1.19.0 build-docs docs/swagger.yaml -o web/apidoc/v1/index.html
