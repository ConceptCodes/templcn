package main

import "strings"

func starterMainGo() string {
	return `package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	mux.HandleFunc("/", homeHandler)

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if err := Home().Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
`
}

func starterAppTempl(modulePath string) string {
	uiImport := modulePath + "/ui"
	return strings.TrimSpace(`package main

import "` + uiImport + `"

templ Home() {
	<!DOCTYPE html>
	<html lang="en" class="h-full bg-background text-foreground antialiased">
		<head>
			<meta charset="UTF-8"/>
			<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
			<title>Shadcn for Go</title>
			<script type="module" src="/assets/runtime.js"></script>
		</head>
		<body class="min-h-svh bg-background text-foreground">
			<main class="mx-auto flex min-h-svh max-w-4xl flex-col justify-center gap-8 px-6 py-16">
				@ui.Card(ui.DOMProps{}) {
					@ui.CardHeader(ui.DOMProps{}) {
						@ui.CardTitle(ui.DOMProps{}) { <span>Shadcn for Go</span> }
						@ui.CardDescription(ui.DOMProps{}) { <span>Templ-native components copied into your project.</span> }
					}
					@ui.CardContent(ui.DOMProps{}) {
						<div class="flex flex-wrap gap-3">
							@ui.Button(ui.ButtonProps{Label: "Get started"})
							@ui.Button(ui.ButtonProps{Label: "Browse docs", Variant: ui.ButtonVariantOutline, Href: "/docs"})
						</div>
					}
				}
			</main>
		</body>
	</html>
}
`)
}

func starterCSS() string {
	return `@import "tailwindcss";

@source "./**/*.go";
@source "./**/*.templ";

@custom-variant dark (&:is(.dark *));

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
  --radius-sm: calc(var(--radius) * 0.6);
  --radius-md: calc(var(--radius) * 0.8);
  --radius-lg: var(--radius);
  --radius-xl: calc(var(--radius) * 1.4);
}

:root {
  --radius: 0.625rem;
  --background: oklch(1 0 0);
  --foreground: oklch(0.145 0 0);
  --card: oklch(1 0 0);
  --card-foreground: oklch(0.145 0 0);
  --popover: oklch(1 0 0);
  --popover-foreground: oklch(0.145 0 0);
  --primary: oklch(0.205 0 0);
  --primary-foreground: oklch(0.985 0 0);
  --secondary: oklch(0.97 0 0);
  --secondary-foreground: oklch(0.205 0 0);
  --muted: oklch(0.97 0 0);
  --muted-foreground: oklch(0.556 0 0);
  --accent: oklch(0.97 0 0);
  --accent-foreground: oklch(0.205 0 0);
  --destructive: oklch(0.577 0.245 27.325);
  --destructive-foreground: oklch(0.985 0 0);
  --border: oklch(0.922 0 0);
  --input: oklch(0.922 0 0);
  --ring: oklch(0.708 0 0);
}

.dark {
  --background: oklch(0.145 0 0);
  --foreground: oklch(0.985 0 0);
  --card: oklch(0.205 0 0);
  --card-foreground: oklch(0.985 0 0);
  --popover: oklch(0.205 0 0);
  --popover-foreground: oklch(0.985 0 0);
  --primary: oklch(0.922 0 0);
  --primary-foreground: oklch(0.205 0 0);
  --secondary: oklch(0.269 0 0);
  --secondary-foreground: oklch(0.985 0 0);
  --muted: oklch(0.269 0 0);
  --muted-foreground: oklch(0.708 0 0);
  --accent: oklch(0.269 0 0);
  --accent-foreground: oklch(0.985 0 0);
  --destructive: oklch(0.704 0.191 22.216);
  --destructive-foreground: oklch(0.985 0 0);
  --border: oklch(1 0 0 / 10%);
  --input: oklch(1 0 0 / 15%);
  --ring: oklch(0.556 0 0);
}

@layer base {
  * {
    @apply border-border outline-ring/50;
  }

  body {
    @apply bg-background text-foreground;
  }
}
`
}

func starterRuntimeJS() string {
	return `/**
 * shadcn-go JS runtime
 *
 * Starter projects get a bundled runtime. The docs site keeps the same behavior
 * split into component modules under /public/runtime/.
 */

const floatingRoots = '[data-slot="popover"], [data-slot="hover-card"], [data-slot="tooltip"], [data-slot="select"], [data-slot="combobox"], [data-slot="dropdown-menu"], [data-slot="context-menu"], [data-slot="menubar"], [data-slot="navigation-menu"], [data-dropdown-menu-root]';
const floatingContent = '[data-slot$="-content"], [data-dropdown-menu-content]';

function isOpen(root) {
  return root?.getAttribute('data-open') === 'true' || root?.hasAttribute('open');
}

function contentFor(root) {
  return root?.querySelector(floatingContent);
}

function triggerFor(root) {
  return root?.querySelector('[data-slot$="-trigger"], [data-dropdown-menu-trigger], [aria-haspopup]');
}

function setOpen(root, open, trigger = triggerFor(root)) {
  const content = contentFor(root);
  root.toggleAttribute('open', open && root.tagName === 'DETAILS');
  root.setAttribute('data-state', open ? 'open' : 'closed');
  if (open) root.setAttribute('data-open', 'true');
  else root.removeAttribute('data-open');
  trigger?.setAttribute('aria-expanded', open ? 'true' : 'false');
  trigger?.setAttribute('data-state', open ? 'open' : 'closed');
  content?.setAttribute('data-state', open ? 'open' : 'closed');
  if (content) content.toggleAttribute('hidden', !open);
  if (open) positionFloating(root);
}

function positionFloating(root) {
  const trigger = triggerFor(root);
  const content = contentFor(root);
  if (!trigger || !content) return;
  content.style.position = 'fixed';
  content.style.margin = '0';
  content.style.zIndex ||= '50';
  const side = root.getAttribute('data-side') || content.getAttribute('data-side') || 'bottom';
  const align = root.getAttribute('data-align') || content.getAttribute('data-align') || 'center';
  const offset = Number.parseFloat(root.getAttribute('data-side-offset') || '4') || 4;
  const anchor = trigger.getBoundingClientRect();
  const rect = content.getBoundingClientRect();
  let left = align === 'start' ? anchor.left : align === 'end' ? anchor.right - rect.width : anchor.left + (anchor.width - rect.width) / 2;
  let top = side === 'top' ? anchor.top - rect.height - offset : anchor.bottom + offset;
  left = Math.max(8, Math.min(left, window.innerWidth - rect.width - 8));
  top = Math.max(8, Math.min(top, window.innerHeight - rect.height - 8));
  content.style.left = Math.round(left) + 'px';
  content.style.top = Math.round(top) + 'px';
  content.setAttribute('data-side', side);
  content.setAttribute('data-align', align);
  content.style.setProperty('--anchor-width', anchor.width + 'px');
  content.style.setProperty('--anchor-height', anchor.height + 'px');
  content.style.setProperty('--radix-popover-trigger-width', anchor.width + 'px');
  content.style.setProperty('--radix-dropdown-menu-trigger-width', anchor.width + 'px');
}

function focusable(root) {
  return Array.from(root.querySelectorAll('a[href], button:not([disabled]), input:not([disabled]), [tabindex]:not([tabindex="-1"])')).filter((el) => !el.hidden);
}

function initRuntime() {
  document.querySelectorAll(floatingRoots).forEach((root) => setOpen(root, isOpen(root)));
  document.addEventListener('click', (event) => {
    const trigger = event.target.closest('[data-slot$="-trigger"], [data-dropdown-menu-trigger], [aria-haspopup]');
    const root = trigger?.closest(floatingRoots);
    if (root) {
      event.preventDefault();
      document.querySelectorAll(floatingRoots + '[data-open="true"]').forEach((openRoot) => {
        if (openRoot !== root) setOpen(openRoot, false);
      });
      setOpen(root, !isOpen(root), trigger);
      return;
    }

    const close = event.target.closest('[data-slot$="-close"], [data-slot="alert-dialog-action"], [data-slot="alert-dialog-cancel"]');
    if (close) {
      const dialog = close.closest('dialog');
      if (dialog) {
        dialog.close();
        setOpen(dialog, false);
      }
    }

    const item = event.target.closest('[data-slot$="-item"], [data-dropdown-menu-item]');
    const selectRoot = item?.closest('[data-slot="select"], [data-slot="combobox"]');
    if (item && selectRoot) {
      const value = item.getAttribute('data-value') || item.textContent.trim();
      selectRoot.setAttribute('data-value', value);
      const display = selectRoot.querySelector('[data-slot$="-value"]');
      if (display) display.textContent = item.textContent.trim();
      setOpen(selectRoot, false);
    }

    document.querySelectorAll(floatingRoots + '[data-open="true"]').forEach((openRoot) => {
      if (!openRoot.contains(event.target)) setOpen(openRoot, false);
    });
  });

  document.addEventListener('keydown', (event) => {
    const root = event.target.closest(floatingRoots + '[data-open="true"]');
    if (event.key === 'Escape') {
      const openRoot = root || document.querySelector(floatingRoots + '[data-open="true"]');
      if (openRoot) {
        event.preventDefault();
        setOpen(openRoot, false);
      }
    }
    if (root && (event.key === 'ArrowDown' || event.key === 'ArrowUp')) {
      const items = focusable(root).filter((el) => el.matches('[data-slot$="-item"], [data-dropdown-menu-item]'));
      if (!items.length) return;
      event.preventDefault();
      const current = items.indexOf(document.activeElement);
      const next = event.key === 'ArrowDown' ? (current + 1 + items.length) % items.length : (current - 1 + items.length) % items.length;
      items[next].focus();
    }
  });

  window.addEventListener('resize', () => document.querySelectorAll(floatingRoots + '[data-open="true"]').forEach(positionFloating));
  window.addEventListener('scroll', () => document.querySelectorAll(floatingRoots + '[data-open="true"]').forEach(positionFloating), true);
}

if (document.readyState === 'loading') document.addEventListener('DOMContentLoaded', initRuntime, { once: true });
else initRuntime();
`
}
