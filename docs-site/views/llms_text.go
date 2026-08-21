package views

import (
	"strings"
)

func LLMSText() string {
	var b strings.Builder
	b.WriteString(`# templcn/ui

> Source-owned templcn/ui components for Go applications using templ, Tailwind CSS, and a small vanilla JavaScript runtime.

## Project Rules

- Components are copied into user projects as source files, not consumed as a shared package.
- Use the CLI for installation: ` + "`templcn add <component>`" + `, ` + "`templcn add --all`" + `, ` + "`templcn view <component>`" + `, ` + "`templcn diff <component>`" + `.
- AI Agent Skill: install official skill via skills.sh with ` + "`npx skills add conceptcodes/templcn`" + `.
- Component source lives in ` + "`ui/*.go`" + `.
- Runtime behavior is copied to ` + "`assets/runtime.js`" + ` when selected components need JavaScript.
- Theme tokens live in ` + "`styles/globals.css`" + ` and use Tailwind v4 ` + "`@theme inline`" + ` with OKLCH CSS variables.
- Docs and examples should use idiomatic templ syntax, not React or TSX syntax.

## Core Docs

- Introduction: /docs
- Installation: /docs/installation
- components.json: /docs/components-json
- Package Imports: /docs/package-imports
- Theming: /docs/theming
- Dark Mode: /docs/dark-mode
- CLI: /docs/cli
- Monorepo: /docs/monorepo
- JavaScript Runtime: /docs/javascript
- Forms: /docs/forms
- Components: /docs/components
- Blocks: /blocks
- Charts: /charts/area

## Important Files

- README.md
- CHANGELOG.md
- cli/main.go
- cli/project.go
- cli/registry_metadata.go
- cli/templates.go
- docs-site/views/site_data.go
- docs-site/views/component_examples.go
- docs-site/public/runtime.js
- docs-site/public/runtime/
- docs-site/styles/globals.css

## Component Docs

`)
	for _, doc := range ComponentIndex() {
		b.WriteString("- ")
		b.WriteString(doc.Title)
		b.WriteString(": /docs/components/")
		b.WriteString(doc.Slug)
		b.WriteString(" - install with `")
		b.WriteString(doc.Install)
		b.WriteString("`\n")
	}
	b.WriteString(`
## Blocks

`)
	for _, block := range Blocks() {
		b.WriteString("- ")
		b.WriteString(block.Title)
		b.WriteString(": /blocks/")
		b.WriteString(block.Slug)
		b.WriteString("\n")
	}
	return strings.TrimSpace(b.String()) + "\n"
}
