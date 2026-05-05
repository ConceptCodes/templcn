export const SELECTORS = {
  floatingRoot:
    '[data-slot="popover"], [data-slot="hover-card"], [data-slot="tooltip"], [data-slot="select"], [data-slot="combobox"], [data-slot="dropdown-menu"], [data-slot="context-menu"], [data-slot="menubar"], [data-slot="navigation-menu"], [data-dropdown-menu-root]',
  floatingContent:
    '[data-slot="popover-content"], [data-slot="hover-card-content"], [data-slot="tooltip-content"], [data-slot="select-content"], [data-slot="combobox-content"], [data-slot="dropdown-menu-content"], [data-dropdown-menu-content]',
  item:
    '[data-slot$="-item"], [data-dropdown-menu-item], [data-select-item], [data-combobox-item], [role="menuitem"], [role="option"]',
  trigger:
    '[data-slot$="-trigger"], [data-dropdown-menu-trigger], [aria-haspopup]',
};

const state = new WeakMap();

export function getState(root) {
  if (!state.has(root)) state.set(root, {});
  return state.get(root);
}

export function isOpen(root) {
  return root?.getAttribute("data-open") === "true" || root?.getAttribute("data-state") === "open" || root?.hasAttribute("open");
}

export function contentFor(root, selector = SELECTORS.floatingContent) {
  return root?.querySelector(selector) || null;
}

export function triggerFor(root) {
  const explicit = root.querySelector(
    '[data-slot="popover-trigger"], [data-slot="hover-card-trigger"], [data-slot="tooltip-trigger"], [data-slot="dropdown-menu-trigger"], [data-dropdown-menu-trigger], [data-slot="select-trigger"], [data-slot="combobox-trigger"], [data-slot="context-menu-trigger"], [data-slot="menubar-trigger"], [data-slot="navigation-menu-trigger"]',
  );
  return explicit || null;
}

export function setOpen(root, open, options = {}) {
  if (!root) return;

  const trigger = options.trigger || triggerFor(root);
  const content = options.content || contentFor(root);
  const overlay = root.querySelector('[data-slot$="-overlay"]');

  if (open) {
    getState(root).lastTrigger = trigger || document.activeElement;
    root.setAttribute("data-open", "true");
    root.setAttribute("data-state", "open");
    if (root.tagName === "DETAILS") root.setAttribute("open", "");
    trigger?.setAttribute("aria-expanded", "true");
    trigger?.setAttribute("data-state", "open");
    content?.removeAttribute("hidden");
    content?.setAttribute("data-state", "open");
    overlay?.setAttribute("data-state", "open");
  } else {
    root.removeAttribute("data-open");
    root.setAttribute("data-state", "closed");
    if (root.tagName === "DETAILS") root.removeAttribute("open");
    trigger?.setAttribute("aria-expanded", "false");
    trigger?.setAttribute("data-state", "closed");
    content?.setAttribute("data-state", "closed");
    overlay?.setAttribute("data-state", "closed");
    if (content && !options.keepMounted) content.setAttribute("hidden", "");
    if (options.restoreFocus !== false) {
      const lastTrigger = getState(root).lastTrigger;
      if (lastTrigger && document.contains(lastTrigger)) lastTrigger.focus();
    }
  }
}

export function closeOpenFloating(except) {
  document.querySelectorAll(`${SELECTORS.floatingRoot}[data-open="true"]`).forEach((root) => {
    if (root !== except && !root.contains(except)) setOpen(root, false, { restoreFocus: false });
  });
}

export function focusable(container) {
  return Array.from(
    container.querySelectorAll(
      'a[href], button:not([disabled]), input:not([disabled]), select:not([disabled]), textarea:not([disabled]), [tabindex]:not([tabindex="-1"])',
    ),
  ).filter((el) => !el.hasAttribute("hidden") && el.offsetParent !== null);
}

export function focusFirst(container) {
  const first = focusable(container)[0];
  if (first) first.focus();
}

export function moveFocus(container, direction) {
  const items = focusable(container).filter((el) => el.matches(SELECTORS.item) || el.getAttribute("role") === "menuitem" || el.getAttribute("role") === "option");
  if (!items.length) return false;
  const current = items.indexOf(document.activeElement);
  const next = current === -1 ? 0 : (current + direction + items.length) % items.length;
  items.forEach((item) => item.removeAttribute("data-highlighted"));
  items[next].setAttribute("data-highlighted", "");
  items[next].focus();
  return true;
}

export function px(value, fallback = 0) {
  const parsed = Number.parseFloat(value);
  return Number.isFinite(parsed) ? parsed : fallback;
}
