# templcn/ui

Source-owned UI components for Go projects using templ, Tailwind CSS, and a small vanilla JavaScript runtime for interactive primitives.

## Status

This repository is preparing for a v1 release. The CLI installs source files into user projects instead of depending on a shared component package.

## Install From Source

```sh
go install github.com/conceptcodes/templcn/cli@latest
```

## Create A Project

```sh
templcn init --name my-app
cd my-app
go mod tidy
go build ./...
```

The generated project includes:

- `ui/` component source files
- `assets/runtime.js` for interactive primitives
- `styles/globals.css`
- `app.templ` and generated starter Go code

## Add Components

```sh
templcn add button
templcn add dialog select
templcn add --all
templcn add select --dry-run
templcn view button
templcn diff button
```

`add` copies only the requested component, declared dependencies, shared render helpers, and runtime file when needed.

## AI Agent Skill

Install the official `templcn` skill for your AI coding assistant (Cursor, Claude Code, Windsurf, Copilot, Antigravity, etc.) via [skills.sh](https://skills.sh):

```sh
npx skills add conceptcodes/templcn
```


## Development

Run module tests:

```sh
(cd ui && go test ./...)
(cd cli && go test ./...)
(cd docs-site && go test ./...)
```

The repository is a multi-module Go workspace, so the equivalent root command is:

```sh
make test
```

Run browser behavior tests:

```sh
(cd docs-site && TMPDIR=$PWD/.tmp bun run test:e2e)
```

Build docs:

```sh
(cd docs-site && bun install && bun run build:css && go run ./cmd/generate -output ./dist)
```

Prepare a self-contained CLI release:

```sh
(cd docs-site && bun run build:tanstack)
(cd cli && go generate ./... && go install .)
```

## Release Criteria

- Component-scoped CLI install behavior is covered by acceptance tests.
- Interactive primitives have browser tests for keyboard, focus, state, dismissal, and form value behavior.
- CLI metadata includes files, dependencies, runtime requirements, CSS, docs URLs, and examples.
- Docs pages use component-specific install commands and concrete previews.
- Known Go-native adaptations are documented in [`docs/V1_SCOPE.md`](docs/V1_SCOPE.md), with generated evidence in [`docs/V1_COMPONENT_MATRIX.md`](docs/V1_COMPONENT_MATRIX.md).
