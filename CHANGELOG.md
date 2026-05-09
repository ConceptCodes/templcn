# Changelog

## Unreleased

- Reworked CLI installs to copy only requested components, declared dependencies, shared helpers, and required runtime.
- Added explicit CLI metadata for dependencies, runtime, CSS, docs URLs, source files, and examples.
- Added generated-project acceptance coverage for `init`, `add button`, `add select`, `add --all`, dry-run, view, diff, and build readiness.
- Added browser behavior coverage for dialog, select, combobox, command, menus, floating content, tabs, accordion, radio group, toggle group, slider, calendar, and long-tail interactive components.
- Replaced placeholder component previews with concrete docs examples.
- Added release CI gates for Go modules and Playwright browser tests.
