---
name: templcn
description: Comprehensive guide and component patterns for templcn/ui — source-owned UI components for Go applications using templ, Tailwind CSS v4, and vanilla JavaScript. Use when setting up templcn, installing components via the CLI (templcn init, templcn add), configuring components.json and Tailwind CSS OKLCH tokens, embedding the vanilla JS runtime for interactive primitives (dialog, dropdown, select, tabs, sheet, etc.), or composing type-safe Go templ components.
allowed-tools: Read, Write, Bash, Edit, Glob
---

# templcn Component Patterns for Go + templ

## Overview

Expert guide for building accessible, customizable, production-ready web user interfaces in Go using `templ`, Tailwind CSS v4, and a lightweight vanilla JavaScript runtime.

`templcn` brings the `shadcn/ui` developer experience and New York design system to Go and `templ`. Components are copied directly into user projects as pure Go and `templ` source code rather than imported as a monolithic external dependency.

---

## Table of Contents

- [When to Use](#when-to-use)
- [Instructions for AI Agents](#instructions-for-ai-agents)
- [Constraints and Warnings](#constraints-and-warnings)
- [Quick Start](#quick-start)
- [What is templcn?](#what-is-templcn)
- [Project Configuration](#project-configuration)
  - [components.json](#componentsjson)
  - [Tailwind CSS v4 & Theming](#tailwind-css-v4--theming)
  - [HTML Base Layout Setup](#html-base-layout-setup)
- [Core Components](#core-components)
  - [Button & Button Group](#button--button-group)
  - [Input, Textarea & Form Controls](#input-textarea--form-controls)
  - [Forms & HTMX Validation](#forms--htmx-validation)
  - [Card](#card)
  - [Dialog (Modal) & Alert Dialog](#dialog-modal--alert-dialog)
  - [Sheet (Slide-Over) & Drawer](#sheet-slide-over--drawer)
  - [Select (Dropdown)](#select-dropdown)
  - [Dropdown Menu, Context Menu & Menubar](#dropdown-menu-context-menu--menubar)
  - [Tabs, Accordion & Collapsible](#tabs-accordion--collapsible)
  - [Table & Data Table](#table--data-table)
  - [Popover, Tooltip & Hover Card](#popover-tooltip--hover-card)
  - [Sonner & Toast Notifications](#sonner--toast-notifications)
  - [Charts](#charts)
  - [Sidebar & Navigation Menu](#sidebar--navigation-menu)
- [Advanced Patterns](#advanced-patterns)
  - [Complete Form with Server-Side Go Validation & HTMX](#complete-form-with-server-side-go-validation--htmx)
  - [Modal Confirmation Dialog with HTMX Action](#modal-confirmation-dialog-with-htmx-action)
  - [Dashboard Layout with Sidebar](#dashboard-layout-with-sidebar)
- [Best Practices](#best-practices)
- [Reference Links](#reference-links)

---

## When to Use

- Setting up a new Go web project with `templ` and Tailwind CSS.
- Installing or configuring individual UI components via `templcn add`.
- Building forms with Go server validation and HTMX updates.
- Creating accessible modals, dialogs, slide-overs, dropdowns, tooltips, and tabs without React.
- Customizing component styling and design tokens with Tailwind CSS v4 OKLCH variables.
- Implementing dark mode and theme switching in Go web applications.
- Composing complex dashboard layouts, data tables, and interactive charts in `templ`.

---

## Instructions for AI Agents

1. **Initialize or Inspect Project**:
   - For new projects: Run `templcn init --name <app-name>`
   - For existing projects: Verify `components.json` exists, or run `templcn init` to configure it.
2. **Install Components**:
   - Run `templcn add <components...>` (e.g. `templcn add button input form card dialog select`).
   - Use `templcn add --all` when building full applications.
3. **Import Local UI Package**:
   - **CRITICAL**: Always import the local application's UI package (e.g. `import "my-app/ui"` or `import "{{module}}/ui"`).
   - **NEVER** import `github.com/conceptcodes/templcn/ui` in user code.
4. **Include Static Assets in Root Layout**:
   - Link `styles/globals.css` in the `<head>`.
   - Add `<script src="/assets/runtime.js" defer></script>` for interactive primitives (dialogs, dropdowns, selects, tabs, etc.).
5. **Compose Components Idiomatically**:
   - Use typed Go structs for props (e.g., `ui.ButtonProps{Variant: ui.ButtonVariantOutline}`).
   - Use `Href` or `Element` instead of React's `asChild`.
   - Pass custom classes or HTMX attributes via `ui.DOMProps{Class: "...", Attrs: templ.Attributes{"hx-post": "/api"}}`.
6. **Compile Templates**:
   - Always run `templ generate` after creating or editing `.templ` files before building or running tests (`go build ./...`).

---

## Constraints and Warnings

- **Source-Owned Code**: Components are copied into the user's project (`ui/`). You own the code and can edit it directly.
- **No React / Node in Production**: There is no React runtime, Radix package, or Node.js server in production. The app compiles to a single static Go binary.
- **Go-Native Composition**:
  - React's `asChild` is replaced with `Href string` and `Element string`.
  - React callbacks (`onOpenChange`, `onChange`) are replaced with standard HTML forms, HTMX triggers, hidden inputs, and custom DOM events.
- **Tailwind CSS v4 Requirement**: Components require Tailwind v4 with `@theme inline` and semantic OKLCH tokens.
- **Client Runtime Asset**: Interactive components rely on `assets/runtime.js` (and `@floating-ui/dom` bundled within it) for focus traps, floating positioning, keyboard navigation, and transitions.
- **Always Run `templ generate`**: Any change to `.templ` files requires regeneration to produce `_templ.go` files before Go compilation.

---

## Quick Start

### New Project Setup

```bash
# 1. Initialize project with templcn CLI
templcn init --name my-app
cd my-app

# 2. Install essential components
templcn add button input form card dialog select tabs

# 3. Tidy dependencies, generate templates, and run
go mod tidy
templ generate
go run main.go
```

### Adding to an Existing Go Project

```bash
# 1. Install the templcn CLI
go install github.com/conceptcodes/templcn/cli@latest

# 2. Initialize components.json configuration
templcn init

# 3. Add desired components
templcn add button card dialog select dropdown-menu
```

---

## What is templcn?

`templcn` is **not** an imported shared package. Instead:
- It is a **curated collection of source-owned UI components** copied directly into your repository.
- Built specifically for **Go + `templ`** with first-class type safety.
- Styled with **Tailwind CSS v4** utilizing `@theme inline` and OKLCH color variables.
- Uses a **lightweight vanilla JS runtime** (`assets/runtime.js`) for interactive primitives (dialogs, dropdowns, tooltips, popovers, tabs, sheets) without React.

---

## Project Configuration

### components.json

`components.json` at the root of the project specifies project layout and configuration:

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

### Tailwind CSS v4 & Theming

Theme variables are configured in `styles/globals.css` using Tailwind v4 `@theme inline` and OKLCH tokens:

```css
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
  --destructive-foreground: oklch(0.637 0.237 25.331);
  --border: oklch(0.269 0 0);
  --input: oklch(0.269 0 0);
  --ring: oklch(0.439 0 0);
}
```

### HTML Base Layout Setup

```templ
package views

templ Layout(title string) {
    <!DOCTYPE html>
    <html lang="en">
        <head>
            <meta charset="UTF-8"/>
            <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
            <title>{ title }</title>
            <link rel="stylesheet" href="/styles/globals.css"/>
            <!-- Include runtime.js for interactive primitives -->
            <script src="/assets/runtime.js" defer></script>
            <!-- Optional HTMX support -->
            <script src="https://unpkg.com/htmx.org@2.0.0" defer></script>
        </head>
        <body class="bg-background text-foreground antialiased min-h-screen">
            { children... }
        </body>
    </html>
}
```

---

## Core Components

All components embed `ui.DOMProps` for `ID`, `Class`, `Element`, and `Attrs`.

### Button & Button Group

```templ
package views

import "my-app/ui"

templ ButtonCatalog() {
    <div class="flex flex-wrap items-center gap-3">
        <!-- Default button -->
        @ui.Button(ui.ButtonProps{Label: "Default Button"})

        <!-- Variants -->
        @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantSecondary, Label: "Secondary"})
        @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline, Label: "Outline"})
        @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantGhost, Label: "Ghost"})
        @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantDestructive, Label: "Destructive"})

        <!-- Button as Link (Href replaces React's asChild) -->
        @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantLink, Href: "/dashboard", Label: "Dashboard Link"})

        <!-- Button Sizes -->
        @ui.Button(ui.ButtonProps{Size: ui.ButtonSizeXS, Label: "XS"})
        @ui.Button(ui.ButtonProps{Size: ui.ButtonSizeSM, Label: "Small"})
        @ui.Button(ui.ButtonProps{Size: ui.ButtonSizeLG, Label: "Large"})
        @ui.Button(ui.ButtonProps{Size: ui.ButtonSizeIcon}) {
            <svg class="size-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
        }

        <!-- Button Group -->
        @ui.ButtonGroup(ui.ButtonGroupProps{}) {
            @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline, Label: "Years"})
            @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline, Label: "Months"})
            @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline, Label: "Days"})
        }
    </div>
}
```

---

### Input, Textarea & Form Controls

```templ
package views

import "my-app/ui"

templ FormControlsCatalog() {
    <div class="max-w-md space-y-4">
        <!-- Input with Label -->
        <div class="space-y-2">
            @ui.Label(ui.DOMProps{Attrs: templ.Attributes{"for": "email"}}) { Email Address }
            @ui.Input(ui.InputProps{
                ID:          "email",
                Type:        "email",
                Placeholder: "name@company.com",
            })
        </div>

        <!-- Input Group with Addon / Icon -->
        <div class="space-y-2">
            @ui.Label(ui.DOMProps{Attrs: templ.Attributes{"for": "search"}}) { Search }
            @ui.InputGroup(ui.DOMProps{}) {
                @ui.InputGroupAddon(ui.DOMProps{}) {
                    <svg class="size-4 text-muted-foreground" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                    </svg>
                }
                @ui.Input(ui.InputProps{ID: "search", Placeholder: "Search components..."})
            }
        </div>

        <!-- Textarea -->
        <div class="space-y-2">
            @ui.Label(ui.DOMProps{Attrs: templ.Attributes{"for": "bio"}}) { Bio }
            @ui.Textarea(ui.TextareaProps{
                ID:          "bio",
                Placeholder: "Tell us a little bit about yourself",
                Rows:        4,
            })
        </div>

        <!-- Checkbox & Switch -->
        <div class="flex items-center space-x-2">
            @ui.Checkbox(ui.CheckboxProps{ID: "remember", Name: "remember"})
            @ui.Label(ui.DOMProps{Attrs: templ.Attributes{"for": "remember"}}) { Remember me }
        </div>

        <div class="flex items-center space-x-2">
            @ui.Switch(ui.SwitchProps{ID: "airplane-mode", Name: "airplane_mode"})
            @ui.Label(ui.DOMProps{Attrs: templ.Attributes{"for": "airplane-mode"}}) { Airplane Mode }
        </div>
    </div>
}
```

---

### Card

```templ
package views

import "my-app/ui"

templ CardCatalog() {
    @ui.Card(ui.DOMProps{Class: "w-full max-w-md"}) {
        @ui.CardHeader(ui.DOMProps{}) {
            @ui.CardTitle(ui.DOMProps{}) { Create an Account }
            @ui.CardDescription(ui.DOMProps{}) { Enter your details below to create your account. }
        }
        @ui.CardContent(ui.DOMProps{}) {
            <div class="space-y-4">
                <div class="space-y-2">
                    @ui.Label(ui.DOMProps{Attrs: templ.Attributes{"for": "username"}}) { Username }
                    @ui.Input(ui.InputProps{ID: "username", Placeholder: "johndoe"})
                </div>
                <div class="space-y-2">
                    @ui.Label(ui.DOMProps{Attrs: templ.Attributes{"for": "password"}}) { Password }
                    @ui.Input(ui.InputProps{ID: "password", Type: "password"})
                </div>
            </div>
        }
        @ui.CardFooter(ui.DOMProps{Class: "flex justify-between"}) {
            @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline, Label: "Cancel"})
            @ui.Button(ui.ButtonProps{Label: "Create Account"})
        }
    }
}
```

---

### Dialog (Modal) & Alert Dialog

```templ
package views

import "my-app/ui"

templ DialogCatalog() {
    <div class="flex gap-4">
        <!-- Standard Dialog -->
        @ui.Dialog(ui.DialogProps{}) {
            @ui.DialogTrigger(ui.DOMProps{}) {
                @ui.Button(ui.ButtonProps{Label: "Edit Profile"})
            }
            @ui.DialogContent(ui.DOMProps{}) {
                @ui.DialogHeader(ui.DOMProps{}) {
                    @ui.DialogTitle(ui.DOMProps{}) { Edit Profile }
                    @ui.DialogDescription(ui.DOMProps{}) {
                        Make changes to your profile here. Click save when you are done.
                    }
                }
                <div class="grid gap-4 py-4">
                    <div class="grid grid-cols-4 items-center gap-4">
                        @ui.Label(ui.DOMProps{Class: "text-right", Attrs: templ.Attributes{"for": "name"}}) { Name }
                        @ui.Input(ui.InputProps{
                            DOMProps: ui.DOMProps{Class: "col-span-3"},
                            ID:       "name",
                            Value:    "Alex Johnson",
                        })
                    </div>
                </div>
                @ui.DialogFooter(ui.DOMProps{}) {
                    @ui.Button(ui.ButtonProps{Type: "submit", Label: "Save Changes"})
                }
            }
        }

        <!-- Alert Dialog (Destructive Action Confirmation) -->
        @ui.AlertDialog(ui.AlertDialogProps{}) {
            @ui.AlertDialogTrigger(ui.DOMProps{}) {
                @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantDestructive, Label: "Delete Project"})
            }
            @ui.AlertDialogContent(ui.DOMProps{}) {
                @ui.AlertDialogHeader(ui.DOMProps{}) {
                    @ui.AlertDialogTitle(ui.DOMProps{}) { Are you absolutely sure? }
                    @ui.AlertDialogDescription(ui.DOMProps{}) {
                        This action cannot be undone. This will permanently delete your project and remove all data from our servers.
                    }
                }
                @ui.AlertDialogFooter(ui.DOMProps{}) {
                    @ui.AlertDialogCancel(ui.DOMProps{}) { Cancel }
                    @ui.AlertDialogAction(ui.DOMProps{Class: "bg-destructive text-destructive-foreground"}) {
                        Continue
                    }
                }
            }
        }
    </div>
}
```

---

### Sheet (Slide-Over) & Drawer

```templ
package views

import "my-app/ui"

templ SheetCatalog() {
    <!-- Side Sheet (options: right, left, top, bottom) -->
    @ui.Sheet(ui.SheetProps{}) {
        @ui.SheetTrigger(ui.DOMProps{}) {
            @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline, Label: "Open Navigation"})
        }
        @ui.SheetContent(ui.SheetContentProps{Side: "right"}) {
            @ui.SheetHeader(ui.DOMProps{}) {
                @ui.SheetTitle(ui.DOMProps{}) { Navigation Menu }
                @ui.SheetDescription(ui.DOMProps{}) { Browse application sections. }
            }
            <nav class="flex flex-col gap-3 py-6">
                <a href="/dashboard" class="text-sm font-medium hover:underline">Dashboard</a>
                <a href="/projects" class="text-sm font-medium hover:underline">Projects</a>
                <a href="/settings" class="text-sm font-medium hover:underline">Settings</a>
            </nav>
            @ui.SheetFooter(ui.DOMProps{}) {
                @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline, Label: "Close"})
            }
        }
    }
}
```

---

### Select (Dropdown)

```templ
package views

import "my-app/ui"

templ SelectCatalog() {
    @ui.Select(ui.SelectProps{
        Name:         "theme",
        Placeholder:  "Select a theme",
        DefaultValue: "system",
    }) {
        @ui.SelectTrigger(ui.DOMProps{Class: "w-[180px]"}) {
            @ui.SelectValue(ui.DOMProps{})
        }
        @ui.SelectContent(ui.DOMProps{}) {
            @ui.SelectGroup(ui.DOMProps{}) {
                @ui.SelectLabel(ui.DOMProps{}) { Appearance }
                @ui.SelectItem(ui.SelectItemProps{Value: "light"}) { Light }
                @ui.SelectItem(ui.SelectItemProps{Value: "dark"}) { Dark }
                @ui.SelectItem(ui.SelectItemProps{Value: "system"}) { System }
            }
        }
    }
}
```

---

### Dropdown Menu, Context Menu & Menubar

```templ
package views

import "my-app/ui"

templ MenuCatalog() {
    @ui.DropdownMenu(ui.DropdownMenuProps{}) {
        @ui.DropdownMenuTrigger(ui.DOMProps{}) {
            @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline, Label: "My Account"})
        }
        @ui.DropdownMenuContent(ui.DOMProps{Class: "w-56"}) {
            @ui.DropdownMenuLabel(ui.DOMProps{}) { alex@example.com }
            @ui.DropdownMenuSeparator(ui.DOMProps{})
            @ui.DropdownMenuGroup(ui.DOMProps{}) {
                @ui.DropdownMenuItem(ui.DOMProps{}) {
                    <span>Profile</span>
                    @ui.DropdownMenuShortcut(ui.DOMProps{}) { ⇧⌘P }
                }
                @ui.DropdownMenuItem(ui.DOMProps{}) {
                    <span>Billing</span>
                    @ui.DropdownMenuShortcut(ui.DOMProps{}) { ⌘B }
                }
                @ui.DropdownMenuItem(ui.DOMProps{}) {
                    <span>Settings</span>
                    @ui.DropdownMenuShortcut(ui.DOMProps{}) { ⌘S }
                }
            }
            @ui.DropdownMenuSeparator(ui.DOMProps{})
            @ui.DropdownMenuItem(ui.DOMProps{Class: "text-destructive"}) {
                <span>Log out</span>
                @ui.DropdownMenuShortcut(ui.DOMProps{}) { ⇧⌘Q }
            }
        }
    }
}
```

---

### Tabs, Accordion & Collapsible

```templ
package views

import "my-app/ui"

templ TabsAndAccordionCatalog() {
    <div class="space-y-8 max-w-md">
        <!-- Tabs -->
        @ui.Tabs(ui.TabsProps{DefaultValue: "overview"}) {
            @ui.TabsList(ui.DOMProps{Class: "grid w-full grid-cols-2"}) {
                @ui.TabsTrigger(ui.TabsTriggerProps{Value: "overview"}) { Overview }
                @ui.TabsTrigger(ui.TabsTriggerProps{Value: "analytics"}) { Analytics }
            }
            @ui.TabsContent(ui.TabsContentProps{Value: "overview"}) {
                <div class="p-4 border rounded-md mt-2 text-sm">
                    Overview content and summary cards.
                </div>
            }
            @ui.TabsContent(ui.TabsContentProps{Value: "analytics"}) {
                <div class="p-4 border rounded-md mt-2 text-sm">
                    Detailed analytics and performance metrics.
                </div>
            }
        }

        <!-- Accordion -->
        @ui.Accordion(ui.AccordionProps{Type: "single", Collapsible: true}) {
            @ui.AccordionItem(ui.AccordionItemProps{Value: "item-1"}) {
                @ui.AccordionTrigger(ui.DOMProps{}) { What is templcn? }
                @ui.AccordionContent(ui.DOMProps{}) {
                    Source-owned components for Go applications using templ and Tailwind CSS.
                }
            }
            @ui.AccordionItem(ui.AccordionItemProps{Value: "item-2"}) {
                @ui.AccordionTrigger(ui.DOMProps{}) { How do I customize styles? }
                @ui.AccordionContent(ui.DOMProps{}) {
                    Edit globals.css tokens or modify the copied component code directly.
                }
            }
        }
    </div>
}
```

---

### Table & Data Table

```templ
package views

import "my-app/ui"

type Invoice struct {
    ID     string
    Status string
    Method string
    Amount string
}

templ InvoiceTable(invoices []Invoice) {
    @ui.Table(ui.DOMProps{}) {
        @ui.TableCaption(ui.DOMProps{}) { A list of your recent invoices. }
        @ui.TableHeader(ui.DOMProps{}) {
            @ui.TableRow(ui.DOMProps{}) {
                @ui.TableHead(ui.DOMProps{Class: "w-[100px]"}) { Invoice }
                @ui.TableHead(ui.DOMProps{}) { Status }
                @ui.TableHead(ui.DOMProps{}) { Method }
                @ui.TableHead(ui.DOMProps{Class: "text-right"}) { Amount }
            }
        }
        @ui.TableBody(ui.DOMProps{}) {
            for _, inv := range invoices {
                @ui.TableRow(ui.DOMProps{}) {
                    @ui.TableCell(ui.DOMProps{Class: "font-medium"}) { { inv.ID } }
                    @ui.TableCell(ui.DOMProps{}) { { inv.Status } }
                    @ui.TableCell(ui.DOMProps{}) { { inv.Method } }
                    @ui.TableCell(ui.DOMProps{Class: "text-right"}) { { inv.Amount } }
                }
            }
        }
    }
}
```

---

### Popover, Tooltip & Hover Card

```templ
package views

import "my-app/ui"

templ PopoverAndTooltipCatalog() {
    <div class="flex items-center gap-4">
        <!-- Tooltip -->
        @ui.TooltipProvider(ui.TooltipProviderProps{}) {
            @ui.Tooltip(ui.TooltipProps{}) {
                @ui.TooltipTrigger(ui.DOMProps{}) {
                    @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline, Label: "Hover Me"})
                }
                @ui.TooltipContent(ui.DOMProps{}) {
                    <p>Add to library</p>
                }
            }
        }

        <!-- Popover -->
        @ui.Popover(ui.PopoverProps{}) {
            @ui.PopoverTrigger(ui.DOMProps{}) {
                @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline, Label: "Dimensions"})
            }
            @ui.PopoverContent(ui.DOMProps{Class: "w-80 p-4"}) {
                <div class="grid gap-4">
                    <div class="space-y-1">
                        <h4 class="font-medium text-sm">Dimensions</h4>
                        <p class="text-xs text-muted-foreground">Set the dimensions for the layer.</p>
                    </div>
                    <div class="grid gap-2">
                        <div class="grid grid-cols-3 items-center gap-4">
                            @ui.Label(ui.DOMProps{Attrs: templ.Attributes{"for": "width"}}) { Width }
                            @ui.Input(ui.InputProps{DOMProps: ui.DOMProps{Class: "col-span-2"}, ID: "width", Value: "100%"})
                        </div>
                    </div>
                </div>
            }
        }
    </div>
}
```

---

### Sonner & Toast Notifications

```templ
package views

import "my-app/ui"

templ ToastExample() {
    <!-- Place Toaster or Sonner in root layout -->
    @ui.Sonner(ui.SonnerProps{})

    <!-- Trigger toast via custom event in runtime.js -->
    <button
        onclick="window.dispatchEvent(new CustomEvent('templcn:toast', { detail: { title: 'Event Created', description: 'Monday, January 1st at 6:00pm' } }))"
        class="rounded-md bg-primary px-4 py-2 text-sm text-primary-foreground"
    >
        Show Notification
    </button>
}
```

---

## Advanced Patterns

### Complete Form with Server-Side Go Validation & HTMX

```templ
package views

import "my-app/ui"

type SignupErrors struct {
    Email    string
    Password string
}

templ SignupForm(errors SignupErrors, successMessage string) {
    <div class="w-full max-w-sm mx-auto" id="signup-container">
        if successMessage != "" {
            @ui.Alert(ui.AlertProps{}) {
                @ui.AlertTitle(ui.DOMProps{}) { Success! }
                @ui.AlertDescription(ui.DOMProps{}) { { successMessage } }
            }
        } else {
            @ui.Card(ui.DOMProps{}) {
                @ui.CardHeader(ui.DOMProps{}) {
                    @ui.CardTitle(ui.DOMProps{}) { Register }
                    @ui.CardDescription(ui.DOMProps{}) { Create your account to get started. }
                }
                @ui.CardContent(ui.DOMProps{}) {
                    <form
                        hx-post="/api/signup"
                        hx-target="#signup-container"
                        hx-swap="outerHTML"
                        class="space-y-4"
                    >
                        @ui.Field(ui.FieldProps{Invalid: errors.Email != ""}) {
                            @ui.FieldLabel(ui.DOMProps{Attrs: templ.Attributes{"for": "email"}}) { Email }
                            @ui.Input(ui.InputProps{
                                ID:          "email",
                                Name:        "email",
                                Type:        "email",
                                Placeholder: "you@example.com",
                                Required:    true,
                            })
                            if errors.Email != "" {
                                @ui.FieldError(ui.DOMProps{}) { { errors.Email } }
                            }
                        }

                        @ui.Field(ui.FieldProps{Invalid: errors.Password != ""}) {
                            @ui.FieldLabel(ui.DOMProps{Attrs: templ.Attributes{"for": "password"}}) { Password }
                            @ui.Input(ui.InputProps{
                                ID:       "password",
                                Name:     "password",
                                Type:     "password",
                                Required: true,
                            })
                            if errors.Password != "" {
                                @ui.FieldError(ui.DOMProps{}) { { errors.Password } }
                            }
                        }

                        @ui.Button(ui.ButtonProps{
                            Type:  "submit",
                            Label: "Create Account",
                            DOMProps: ui.DOMProps{Class: "w-full"},
                        })
                    </form>
                }
            }
        }
    </div>
}
```

---

### Modal Confirmation Dialog with HTMX Action

```templ
package views

import "my-app/ui"

templ DeleteProjectModal(projectID string) {
    @ui.Dialog(ui.DialogProps{}) {
        @ui.DialogTrigger(ui.DOMProps{}) {
            @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantDestructive, Label: "Delete Project"})
        }
        @ui.DialogContent(ui.DOMProps{}) {
            @ui.DialogHeader(ui.DOMProps{}) {
                @ui.DialogTitle(ui.DOMProps{}) { Confirm Deletion }
                @ui.DialogDescription(ui.DOMProps{}) {
                    Are you sure you want to delete this project? All associated resources will be permanently removed.
                }
            }
            @ui.DialogFooter(ui.DOMProps{}) {
                @ui.Button(ui.ButtonProps{
                    Variant: ui.ButtonVariantDestructive,
                    Label:   "Confirm Delete",
                    DOMProps: ui.DOMProps{
                        Attrs: templ.Attributes{
                            "hx-delete": "/projects/" + projectID,
                            "hx-target": "#project-row-" + projectID,
                            "hx-swap":   "outerHTML swap:1s",
                        },
                    },
                })
            }
        }
    }
}
```

---

## Best Practices

1. **Keep Imports Local**:
   Always import your local project package (`my-app/ui`). Never import `conceptcodes/templcn/ui`.
2. **Use Semantic OKLCH Colors**:
   Utilize Tailwind v4 design tokens (`text-primary`, `bg-card`, `text-muted-foreground`, `border-border`, `ring-ring`) so that light and dark modes work out of the box.
3. **Handle Interactivity with Clean Primitives**:
   - Use standard HTML forms and HTMX for server requests.
   - Use `assets/runtime.js` for keyboard focus traps, dismiss-on-click-outside, and floating positioning.
4. **Compile Frequently**:
   Run `templ generate` followed by `go test ./...` or `go build ./...` to ensure templates compile smoothly.

---

## Reference Links

- [CLI Reference](references/cli.md) — Comprehensive guide to `init`, `add`, `apply`, `view`, `diff`, `search`, etc.
- [Component API Matrix](references/components.md) — Exhaustive list of all 40+ components, props structs, and slots.
- [Architecture & Theming](references/architecture.md) — Deep dive into Tailwind CSS v4 OKLCH tokens and runtime JS mechanics.
