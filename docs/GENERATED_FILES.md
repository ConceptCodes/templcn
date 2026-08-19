# Generated files policy

Generated `*_templ.go` files are checked into this repository because the CLI, docs generator, and clean-checkout builds consume them directly. A source change that affects a templ file must include its regenerated Go output.

Release-generated assets are also checked in when they are required for a source-only install:

- `docs-site/public/tanstack-runtime.js`
- `cli/runtime_embed.go`
- `docs/V1_COMPONENT_MATRIX.md`

CI regenerates these artifacts and fails if the working tree changes. This keeps the CLI release and documentation site reproducible without requiring users to run a generator first.
