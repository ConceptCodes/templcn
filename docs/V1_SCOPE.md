# templcn/ui v1 scope

v1 targets a dependable Go/templ component distribution workflow, not byte-for-byte React parity.

Every advertised item must be source-owned, installable, documented, and covered by the appropriate rendered or browser verification. Differences from shadcn/ui are acceptable when they are explicit and tested.

## Adapted v1 items

These items are intentionally Go-native adaptations:

- `calendar`, `date-picker`: server-rendered date values with vanilla interaction.
- `carousel`, `drawer`: templ composition with a small browser runtime.
- `chart`: TanStack Charts enhancement with a server-rendered fallback, including Cartesian, pie, and radial charts.
- `combobox`, `command`: native HTML plus a client-side searchable-list enhancement.
- `data-table`: server-rendered rows with optional TanStack Table sorting enhancement.
- `sonner`, `toast`: DOM event/toast runtime rather than React state.

The v1 contract is defined by [V1_COMPONENT_MATRIX.md](V1_COMPONENT_MATRIX.md), generated from the docs index, UI source tree, and test inventory.

## Supported toolchain

- Go 1.23 or newer
- `github.com/a-h/templ` v0.3.1001 across UI, CLI output, and docs
- Tailwind CSS v4.2.x through the docs build
- Bun with the checked-in `docs-site/bun.lock`
- Chromium-based browsers; Playwright uses managed Chromium in CI and installed Chrome during local development when available
