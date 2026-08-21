# templcn Architecture & Theming

This document explains the technical architecture, CSS variable token system, vanilla JS runtime event system, and Go-native composition model of `templcn`.

---

## 1. Source-Owned Distribution Model

Unlike traditional Go libraries packaged as monolithic modules, `templcn` uses a copy-and-own distribution model inspired by shadcn/ui:

```
[Registry / CLI] ──templcn add──> [User App: ./ui/*.go]
```

- Each component file (e.g., `ui/button.go`, `ui/dialog.go`) is written directly to the target project.
- Shared utilities (`cn.go`, `html.go`, `render.go`, `types.go`) are copied automatically when needed.
- Runtime assets (`assets/runtime.js`) are placed into the project's static asset directory.
- Developers have 100% control over component internals, modifications, and extensions.

---

## 2. Tailwind CSS v4 & OKLCH Theming

`templcn` is built on modern Tailwind CSS v4 using `@theme inline` and OKLCH color spaces for high-fidelity dark and light themes:

```css
/* styles/globals.css */
@import "tailwindcss";

@theme inline {
  --color-background: var(--background);
  --color-foreground: var(--foreground);
  --color-card: var(--card);
  --color-card-foreground: var(--card-foreground);
  --color-popover: var(--popover);
  --color-popover-foreground: var(--popover-foreground);
  --color-primary: var(--primary);
  --color-primary-foreground: var(--primary-foreground);
  --color-secondary: var(--secondary);
  --color-secondary-foreground: var(--secondary-foreground);
  --color-muted: var(--muted);
  --color-muted-foreground: var(--muted-foreground);
  --color-accent: var(--accent);
  --color-accent-foreground: var(--accent-foreground);
  --color-destructive: var(--destructive);
  --color-destructive-foreground: var(--destructive-foreground);
  --color-border: var(--border);
  --color-input: var(--input);
  --color-ring: var(--ring);
  --radius-sm: calc(var(--radius) - 4px);
  --radius-md: calc(var(--radius) - 2px);
  --radius-lg: var(--radius);
  --radius-xl: calc(var(--radius) + 4px);
}

:root {
  --background: oklch(1 0 0);
  --foreground: oklch(0.145 0 0);
  --card: oklch(1 0 0);
  --card-foreground: oklch(0.145 0 0);
  --primary: oklch(0.205 0 0);
  --primary-foreground: oklch(0.985 0 0);
  --secondary: oklch(0.97 0 0);
  --secondary-foreground: oklch(0.205 0 0);
  --muted: oklch(0.97 0 0);
  --muted-foreground: oklch(0.556 0 0);
  --accent: oklch(0.97 0 0);
  --accent-foreground: oklch(0.205 0 0);
  --destructive: oklch(0.577 0.245 27.325);
  --destructive-foreground: oklch(0.577 0.245 27.325);
  --border: oklch(0.922 0 0);
  --input: oklch(0.922 0 0);
  --ring: oklch(0.708 0 0);
  --radius: 0.625rem;
}

.dark {
  --background: oklch(0.145 0 0);
  --foreground: oklch(0.985 0 0);
  --card: oklch(0.145 0 0);
  --card-foreground: oklch(0.985 0 0);
  --primary: oklch(0.985 0 0);
  --primary-foreground: oklch(0.205 0 0);
  --secondary: oklch(0.269 0 0);
  --secondary-foreground: oklch(0.985 0 0);
  --muted: oklch(0.269 0 0);
  --muted-foreground: oklch(0.708 0 0);
  --accent: oklch(0.269 0 0);
  --accent-foreground: oklch(0.985 0 0);
  --destructive: oklch(0.396 0.141 25.723);
  --border: oklch(0.269 0 0);
  --input: oklch(0.269 0 0);
  --ring: oklch(0.439 0 0);
}
```

---

## 3. Vanilla JavaScript Runtime & Floating UI

Interactive components require focus traps, keyboard navigation (Escape, Arrow keys, Enter), overlay positioning, and click-outside dismissal without React.

`assets/runtime.js` provides:
- **Floating Positioning**: Bundles `@floating-ui/dom` to position dropdown menus, popovers, tooltips, comboboxes, and hover cards.
- **Data Attributes & State**: Components declare `data-state="open"` or `data-state="closed"`, which CSS transitions and animations hook into.
- **Roving Focus & Keyboard**: Handles arrow key navigation across items in select dropdowns, context menus, and tabs.
- **Native `<dialog>` integration**: Enhances HTML dialogs with backdrop animation and focus management.

---

## 4. `components.json` Configuration

The project root contains `components.json` to guide CLI sync operations:

```json
{
  "$schema": "https://templcn.com/schema.json",
  "style": "new-york",
  "tailwind": {
    "config": "",
    "css": "styles/globals.css",
    "baseColor": "zinc",
    "cssVariables": true
  },
  "aliases": {
    "components": "components",
    "utils": "ui",
    "ui": "ui",
    "lib": "lib",
    "hooks": "hooks"
  },
  "module": "example.com/my-app",
  "uiDir": "ui",
  "assetsDir": "assets",
  "stylesDir": "styles"
}
```
