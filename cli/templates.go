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
			<script src="/assets/runtime.js" defer></script>
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
 * shadcn-go minimal JS runtime
 * 
 * Hydrates headless UI behaviors purely based on data-* attributes.
 */

document.addEventListener('DOMContentLoaded', () => {
    initRuntime();
});

function initRuntime() {
    document.addEventListener('click', (e) => {
        handleToggles(e);
        handleOutsideClick(e);
    });
    document.addEventListener('keydown', (e) => {
        handleEscape(e);
        handleRovingFocus(e);
    });

    // Initialize immediate behaviors
    initCharts();
}

function initCharts() {
    // Look for Chart Containers and initialize them if a charting library is present
    const charts = document.querySelectorAll('[data-chart-container]');
    charts.forEach(chart => {
        const configStr = chart.getAttribute('data-config');
        if (configStr) {
            // Note: In a vanilla JS port, we expect users to bring their own charting library 
            // (like Chart.js or ApexCharts). This is a hook point for initializing them based 
            // on the container's data attributes.
            console.warn("ChartContainer found: inject your preferred JS charting library here to render.");
        }
    });
}

function handleRovingFocus(e) {
    if (e.key === 'ArrowDown' || e.key === 'ArrowUp') {
        const itemSelectors = '[data-dropdown-menu-item], [data-select-item], [data-combobox-item], [data-command-item], [data-alert-dialog-action], [data-alert-dialog-cancel]';
        
        // Find if we are currently focused on an item or a trigger
        const activeContainer = e.target.closest('[data-open="true"] [data-dropdown-menu-content], [data-open="true"] [data-select-content], [data-open="true"]:not(dialog), dialog[open]');
        
        if (!activeContainer) return;
        
        const items = Array.from(activeContainer.querySelectorAll(itemSelectors));
        if (items.length === 0) return;

        e.preventDefault();
        
        const currentIndex = items.indexOf(document.activeElement);
        let nextIndex;

        if (e.key === 'ArrowDown') {
            nextIndex = currentIndex >= items.length - 1 ? 0 : currentIndex + 1;
        } else {
            nextIndex = currentIndex <= 0 ? items.length - 1 : currentIndex - 1;
        }

        items[nextIndex].focus();
    }
}

function handleToggles(e) {
    // 1. Dialogs and Sheets (Native <dialog> elements)
    const dialogTrigger = e.target.closest('[data-dialog-trigger], [data-sheet-trigger]');
    if (dialogTrigger) {
        let targetId = dialogTrigger.getAttribute('aria-controls') || dialogTrigger.getAttribute('data-target');
        let target = targetId ? document.getElementById(targetId) : null;
        
        // Fallback: look for next sibling that is a dialog
        if (!target) {
            let next = dialogTrigger.nextElementSibling;
            while(next) {
                if(next.tagName === 'DIALOG') {
                    target = next;
                    break;
                }
                next = next.nextElementSibling;
            }
        }
        
        if (target && target.tagName === 'DIALOG') {
            const isModal = target.getAttribute('data-modal') !== 'false';
            if (!target.open && target.getAttribute('open') === null) {
                if (isModal) {
                    target.showModal();
                } else {
                    target.show();
                }
                target.setAttribute('data-open', 'true');
            }
        }
    }

    // 2. Dialog Closes
    const dialogClose = e.target.closest('[data-dialog-close], [data-sheet-close]');
    if (dialogClose) {
        const dialog = dialogClose.closest('dialog');
        if (dialog) {
            dialog.close();
            dialog.removeAttribute('data-open');
        }
    }

    // 3. Popovers, Dropdowns, Comboboxes
    const popoverTrigger = e.target.closest('[data-popover-trigger], [data-dropdown-menu-trigger], [data-select-trigger], [data-combobox-trigger], [data-collapsible-trigger]');
    if (popoverTrigger) {
        // Find the closest root container
        const rootSelectors = ['[data-popover]', '[data-dropdown-menu-root]', '[data-select]', '[data-combobox]', '[data-collapsible]'];
        let root = null;
        for (const selector of rootSelectors) {
            root = popoverTrigger.closest(selector);
            if (root) break;
        }
        
        if (root) {
            const isOpen = root.getAttribute('data-open') === 'true' || root.getAttribute('open') !== null;
            if (isOpen) {
                root.removeAttribute('data-open');
                root.removeAttribute('open');
                popoverTrigger.setAttribute('aria-expanded', 'false');
            } else {
                root.setAttribute('data-open', 'true');
                if(root.tagName === 'DETAILS') root.setAttribute('open', '');
                popoverTrigger.setAttribute('aria-expanded', 'true');
            }


    // 4. Selections (Select, Combobox)
    const selectItem = e.target.closest('[data-select-item], [data-combobox-item], [data-dropdown-menu-item]');
    if (selectItem) {
        const value = selectItem.getAttribute('data-value') || selectItem.textContent.trim();
        const root = selectItem.closest('[data-select], [data-combobox]');
        
        if (root) {
            // Update root's data-value
            root.setAttribute('data-value', value);
            
            // Find hidden input if exists and update it
            const name = root.getAttribute('data-name');
            if (name) {
                let hiddenInput = root.querySelector(` + "`input[name=\"${name}\"]`" + `);
                if (!hiddenInput) {
                    hiddenInput = document.createElement('input');
                    hiddenInput.type = 'hidden';
                    hiddenInput.name = name;
                    root.appendChild(hiddenInput);
                }
                hiddenInput.value = value;
                // Dispatch input/change event
                hiddenInput.dispatchEvent(new Event('input', { bubbles: true }));
                hiddenInput.dispatchEvent(new Event('change', { bubbles: true }));
            }

            // Update text node in SelectValue if present
            const valueDisplay = root.querySelector('[data-select-value], [data-combobox-value]');
            if (valueDisplay) {
                if (valueDisplay.tagName === 'INPUT') {
                    valueDisplay.value = selectItem.textContent.trim();
                } else {
                    valueDisplay.textContent = selectItem.textContent.trim();
                }
            }

            // Close the popover
            root.removeAttribute('data-open');
            root.removeAttribute('open');
            const trigger = root.querySelector('[aria-expanded="true"]');
            if (trigger) trigger.setAttribute('aria-expanded', 'false');
        } else {
            // Just close for Dropdown Menus if no Select/Combobox root
            const menuRoot = selectItem.closest('[data-dropdown-menu-root], [data-popover]');
            if (menuRoot) {
                menuRoot.removeAttribute('data-open');
                menuRoot.removeAttribute('open');
            }
        }
    }

    // 5. Tabs
    const tabTrigger = e.target.closest('[data-tabs-trigger]');
    if (tabTrigger) {
        const value = tabTrigger.getAttribute('data-value');
        const root = tabTrigger.closest('[data-tabs]');
        if (root && value) {
            // Update triggers
            const triggers = root.querySelectorAll('[data-tabs-trigger]');
            triggers.forEach(t => {
                if (t.getAttribute('data-value') === value) {
                    t.setAttribute('data-state', 'active');
                } else {
                    t.setAttribute('data-state', 'inactive');
                }
            });

            // Update content panels
            const contents = root.querySelectorAll('[data-tabs-content]');
            contents.forEach(c => {
                if (c.getAttribute('data-value') === value) {
                    c.setAttribute('data-state', 'active');
                    c.removeAttribute('hidden');
                } else {
                    c.setAttribute('data-state', 'inactive');
                    c.setAttribute('hidden', 'true');
                }
            });
        }
    }

    // 6. Accordion (Single & Collapsible support for native <details>)
    const accordionTrigger = e.target.closest('[data-accordion-trigger]');
    if (accordionTrigger) {
        const item = accordionTrigger.closest('details');
        const root = item ? item.closest('[data-accordion]') : null;
        if (root && root.getAttribute('data-type') === 'single') {
            const isCollapsible = root.getAttribute('data-collapsible') === 'true';
            
            // If already open and not collapsible, prevent closing
            if (item.hasAttribute('open') && !isCollapsible) {
                e.preventDefault();
            } else if (!item.hasAttribute('open')) {
                // If opening, close all other open details in this accordion
                const others = root.querySelectorAll('details[open]');
                others.forEach(other => {
                    if (other !== item) {
                        other.removeAttribute('open');
                    }
                });
            }
        }
    }
    // 7. Carousel
    const carouselBtn = e.target.closest('[data-carousel-previous], [data-carousel-next]');
    if (carouselBtn) {
        const carousel = carouselBtn.closest('[data-carousel]');
        if (carousel) {
            const content = carousel.querySelector('[data-carousel-content]');
            if (content) {
                // Determine item width (fallback to container width)
                const firstItem = content.querySelector('[data-carousel-item]');
                const scrollAmount = firstItem ? firstItem.getBoundingClientRect().width : content.clientWidth;
                
                // Allow native smooth scrolling if the container overflow is managed by CSS,
                // or forcefully adjust scrollLeft for Embla-like headless setups.
                const isNext = carouselBtn.hasAttribute('data-carousel-next');
                
                // If CSS isn't natively snappy, we do manual scroll mapping
                content.scrollBy({
                    left: isNext ? scrollAmount : -scrollAmount,
                    behavior: 'smooth'
                });
            }
        }
    }
}

function handleOutsideClick(e) {
    // Close active popovers/dropdowns if click is outside their tree
    const activePopovers = document.querySelectorAll('[data-open="true"]:not(dialog)');
    activePopovers.forEach(popover => {
        // If the popover contains the click, ignore
        if (popover.contains(e.target)) return;
        
        // If the click is on a trigger for this popover, ignore (handled by handleToggles)
        const triggerId = popover.id;
        if (triggerId) {
            const triggers = document.querySelectorAll(` + "`[aria-controls=\"${triggerId}\"], [data-target=\"${triggerId}\"]`" + `);
            for (let t of triggers) {
                if (t.contains(e.target)) return;
            }
        }
        
        // Close it
        popover.removeAttribute('data-open');
        popover.removeAttribute('open');
        const trigger = popover.querySelector('[aria-expanded="true"]');
        if(trigger) trigger.setAttribute('aria-expanded', 'false');
    });

    // Handle native Dialog backdrop clicks (clicking outside the dialog content)
    if (e.target.tagName === 'DIALOG' && e.target.open) {
        const rect = e.target.getBoundingClientRect();
        const isInDialog = (rect.top <= e.clientY && e.clientY <= rect.top + rect.height &&
                            rect.left <= e.clientX && e.clientX <= rect.left + rect.width);
        if (!isInDialog) {
            e.target.close();
            e.target.removeAttribute('data-open');
        }
    }
}

function handleEscape(e) {
    if (e.key === 'Escape') {
        const activePopovers = document.querySelectorAll('[data-open="true"]:not(dialog)');
        // Close the most recently opened popover
        if (activePopovers.length > 0) {
            const lastPopover = activePopovers[activePopovers.length - 1];
            lastPopover.removeAttribute('data-open');
            lastPopover.removeAttribute('open');
            const trigger = lastPopover.querySelector('[aria-expanded="true"]');
            if(trigger) trigger.setAttribute('aria-expanded', 'false');
        }
    }
}
`
}
