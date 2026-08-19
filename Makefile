.PHONY: test test-go build-css build-docs test-browser generate-matrix release-cli

export TMPDIR := $(CURDIR)/.tmp
export GOWORK := off

test: test-go

test-go:
	mkdir -p .tmp
	(cd ui && go test ./...)
	(cd cli && go test ./...)
	(cd docs-site && go test ./...)

build-css:
	(cd docs-site && bun run build:css)

build-docs:
	(cd docs-site && bun run build:css && bun run build:tanstack && go run ./cmd/generate -output ./dist)

test-browser:
	(cd docs-site && TMPDIR=$$PWD/.tmp bun run test:e2e)

generate-matrix:
	(cd docs-site && go run ./cmd/matrix)

release-cli:
	(cd docs-site && bun run build:tanstack)
	(cd cli && go generate ./... && go build -o templcn .)
