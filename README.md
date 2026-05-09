# shadcn for Go/templ

Templ-native shadcn/ui components copied into Go projects with a component-scoped CLI and a small vanilla JavaScript runtime for interactive primitives.

## Status

This repository is preparing for a v1 release. The CLI installs source files into user projects instead of depending on a shared component package.

## Install From Source

```sh
go install ./cli
```

## Create A Project

```sh
shadcn init --name my-app
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
shadcn add button
shadcn add dialog select
shadcn add --all
shadcn add select --dry-run
shadcn view button
shadcn diff button
```

`add` copies only the requested component, declared dependencies, shared render helpers, and runtime file when needed.

## Development

Run module tests:

```sh
(cd ui && go test ./...)
(cd cli && go test ./...)
(cd docs-site && go test ./...)
```

Run browser behavior tests:

```sh
(cd docs-site && TMPDIR=$PWD/.tmp bun run test:e2e)
```

Build docs:

```sh
(cd docs-site && bun install && bun run build:css && go run ./cmd/generate -output ./dist)
```

## Release Criteria

- Component-scoped CLI install behavior is covered by acceptance tests.
- Interactive primitives have browser tests for keyboard, focus, state, dismissal, and form value behavior.
- CLI metadata includes files, dependencies, runtime requirements, CSS, docs URLs, and examples.
- Docs pages use component-specific install commands and concrete previews.
