/**
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
        }
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
                let hiddenInput = root.querySelector(`input[name="${name}"]`);
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
            const triggers = document.querySelectorAll(`[aria-controls="${triggerId}"], [data-target="${triggerId}"]`);
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
