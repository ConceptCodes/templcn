# templcn CLI Reference

The `templcn` CLI manages source-owned component distribution, scaffolding, presets, diffs, and local/remote registry integration.

## Installation

```bash
go install github.com/conceptcodes/templcn/cli@latest
```

---

## Commands

### `templcn init` / `templcn create`

Scaffolds a new Go + `templ` project or sets up `components.json` in an existing repository.

```bash
templcn init [components...] [flags]
```

#### Flags
- `-n, --name <string>`: Project name (creates a subfolder with this name).
- `-t, --template <string>`: Template to scaffold (default: `next`).
- `-b, --base <string>`: Component base (default: `radix`).
- `-p, --preset <string>`: Theme preset to apply (default: `nova`, options: `default`, `new-york`, `zinc`, `slate`, `stone`, `gray`, `neutral`, `red`, `rose`, `orange`, `green`, `blue`, `yellow`, `violet`).
- `-c, --cwd <path>`: Working directory (default: `.`).
- `-y, --yes`: Skip confirmation prompts (default: `true`).
- `-d, --defaults`: Use default configuration.
- `-f, --force`: Overwrite existing files.
- `-s, --silent`: Mute output.
- `--monorepo`: Scaffold a monorepo structure.
- `--rtl`: Enable RTL layout support.
- `--reinstall`: Reinstall existing components.

#### Examples
```bash
# Create a fresh project named 'blog'
templcn init --name blog

# Scaffold with pre-selected components
templcn init button card dialog select --name dashboard
```

---

### `templcn add`

Syncs component source code, declared dependencies, shared render helpers, and required runtime assets into the user's project.

```bash
templcn add [components...] [flags]
```

#### Flags
- `-a, --all`: Sync every available component in the registry.
- `-o, --overwrite`: Overwrite existing component files.
- `-p, --path <string>`: Target UI package path (overrides `components.json`).
- `-c, --cwd <path>`: Working directory (default: `.`).
- `-y, --yes`: Skip confirmation prompts.
- `-s, --silent`: Mute output.
- `--dry-run`: Preview file modifications without writing to disk.
- `--diff <file>`: Show diff for a specific file.
- `--view <file>`: Show file contents.

#### Examples
```bash
# Add single or multiple components
templcn add button
templcn add dialog select tabs dropdown-menu

# Install all components
templcn add --all

# Add a pre-built UI block (e.g. login form, dashboard)
templcn add login-01 dashboard-01

# Add from a custom or local registry item
templcn add ./custom-card.json
```

---

### `templcn apply`

Applies a theme preset or font configuration to an existing project.

```bash
templcn apply [preset] [flags]
```

#### Flags
- `-p, --preset <string>`: Preset name to apply.
- `--only <strings>`: Apply only parts of a preset (e.g. `--only theme` or `--only font`).
- `-c, --cwd <path>`: Working directory (default: `.`).
- `-y, --yes`: Skip confirmation prompts.
- `-s, --silent`: Mute output.

#### Examples
```bash
# Apply the 'new-york' preset
templcn apply new-york

# Update only CSS theme variables
templcn apply zinc --only theme
```

---

### `templcn view`

Displays the source code of any component or registry item in the terminal.

```bash
templcn view <items...> [flags]
```

#### Examples
```bash
templcn view button
templcn view select
```

---

### `templcn diff`

Compares local component files against registry source to review local customizations or pending updates.

```bash
templcn diff <items...> [flags]
```

#### Examples
```bash
templcn diff button
templcn diff dialog
```

---

### `templcn search` / `templcn list`

Searches local and remote registries for available components and blocks.

```bash
templcn search [registries...] [flags]
```

#### Flags
- `-q, --query <string>`: Search query term.
- `-l, --limit <int>`: Maximum results to return (default: 100).
- `-o, --offset <int>`: Number of results to skip (default: 0).
- `-c, --cwd <path>`: Working directory (default: `.`).

#### Examples
```bash
templcn search
templcn search -q dialog
```

---

### `templcn info`

Inspects and outputs the current project configuration, module path, installed components, and registered presets.

```bash
templcn info [--json]
```

---

### `templcn build`

Generates a registry distribution manifest and bundle for publishing custom component registries.

```bash
templcn build [registry] [-o, --output <dir>]
```

---

### `templcn docs`

Outputs component documentation, required CSS variables, dependencies, and code examples directly in the terminal or as JSON.

```bash
templcn docs button
templcn docs dialog --json
```

---

### `templcn parity`

Compares local component catalog and slots with upstream shadcn/ui.

```bash
templcn parity [--source <url|file>] [--json]
```
