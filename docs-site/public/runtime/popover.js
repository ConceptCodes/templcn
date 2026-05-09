import { closeOpenFloating, focusFirst, moveFocus, setOpen } from "./core.js";
import { positionFloating } from "./positioning.js";

const roots = '[data-slot="popover"], [data-slot="hover-card"], [data-slot="tooltip"], [data-slot="select"], [data-slot="combobox"]';
const triggerSelector = '[data-slot="popover-trigger"], [data-slot="hover-card-trigger"], [data-slot="tooltip-trigger"], [data-slot="select-trigger"], [data-slot="combobox-trigger"]';
const selectableRoots = '[data-slot="select"], [data-slot="combobox"]';
const openSelectableRoots = '[data-slot="select"][data-open="true"], [data-slot="combobox"][data-open="true"]';
const openRoots = '[data-slot="popover"][data-open="true"], [data-slot="hover-card"][data-open="true"], [data-slot="tooltip"][data-open="true"], [data-slot="select"][data-open="true"], [data-slot="combobox"][data-open="true"]';
const selectableItems = '[data-slot="select-item"], [data-slot="combobox-item"], [role="option"]';
const typeahead = new WeakMap();

export function initPopovers() {
  document.querySelectorAll(roots).forEach((root) => {
    setOpen(root, root.getAttribute("data-open") === "true", { keepMounted: root.getAttribute("data-open") === "true", restoreFocus: false });
    if (root.matches(selectableRoots)) syncSelectValue(root);
  });
  document.addEventListener("click", onClick);
  document.addEventListener("mouseover", onHover);
  document.addEventListener("mouseout", onHoverOut);
  document.addEventListener("focusin", onFocusIn);
  document.addEventListener("keydown", onKeydown);
}

function onClick(event) {
  const trigger = event.target.closest(triggerSelector);
  const root = trigger?.closest(roots);
  if (root && !root.matches('[data-slot="tooltip"], [data-slot="hover-card"]')) {
    event.preventDefault();
    const open = root.getAttribute("data-open") !== "true";
    if (open) closeOpenFloating(root);
    setOpen(root, open, { trigger });
    if (open) {
      positionFloating(root);
      if (event.detail === 0) {
        if (root.matches(selectableRoots)) moveFocus(root, 1);
        else focusFirst(root);
      }
    }
    return;
  }

  document.querySelectorAll(openRoots).forEach((openRoot) => {
    if (!openRoot.contains(event.target)) setOpen(openRoot, false);
  });
}

function onHover(event) {
  const trigger = event.target.closest('[data-slot="tooltip-trigger"], [data-slot="hover-card-trigger"]');
  const root = trigger?.closest(roots);
  if (!root) return;
  setOpen(root, true, { trigger });
  positionFloating(root);
}

function onHoverOut(event) {
  const root = event.target.closest('[data-slot="tooltip"], [data-slot="hover-card"]');
  if (root && !root.contains(event.relatedTarget)) setOpen(root, false, { restoreFocus: false });
}

function onFocusIn(event) {
  const trigger = event.target.closest('[data-slot="tooltip-trigger"], [data-slot="hover-card-trigger"]');
  const root = trigger?.closest(roots);
  if (!root) return;
  setOpen(root, true, { trigger });
  positionFloating(root);
}

function onKeydown(event) {
  const activeRoot = event.target.closest(openSelectableRoots);
  if (activeRoot) {
    if (event.key === "Escape") {
      event.preventDefault();
      setOpen(activeRoot, false);
      return;
    }
    if (event.key === "ArrowDown") {
      event.preventDefault();
      moveFocus(activeRoot, 1);
      return;
    }
    if (event.key === "ArrowUp") {
      event.preventDefault();
      moveFocus(activeRoot, -1);
      return;
    }
    if (event.key === "Home") {
      event.preventDefault();
      focusFirst(activeRoot);
      return;
    }
    if (event.key === "End") {
      event.preventDefault();
      moveFocus(activeRoot, -1);
      return;
    }
    if (event.key === "Enter" || event.key === " ") {
      const item = event.target.closest(selectableItems);
      if (item) {
        event.preventDefault();
        item.click();
      }
      return;
    }
    if (event.key.length === 1 && !event.metaKey && !event.ctrlKey && !event.altKey) {
      focusByTypeahead(activeRoot, event.key);
    }
    return;
  }

  const trigger = event.target.closest('[data-slot="select-trigger"], [data-slot="combobox-trigger"]');
  const triggerRoot = trigger?.closest(selectableRoots);
  if (triggerRoot && ["ArrowDown", "ArrowUp", "Enter", " "].includes(event.key)) {
    event.preventDefault();
    closeOpenFloating(triggerRoot);
    setOpen(triggerRoot, true, { trigger });
    positionFloating(triggerRoot);
    if (event.key === "ArrowUp") moveFocus(triggerRoot, -1);
    else moveFocus(triggerRoot, 1);
    return;
  }

  if (event.key !== "Escape") return;
  const root = document.querySelector(openRoots);
  if (!root) return;
  event.preventDefault();
  setOpen(root, false);
}

function syncSelectValue(root) {
  const value = root.getAttribute("data-value") || root.getAttribute("data-default-value") || "";
  if (!value) return;
  root.setAttribute("data-value", value);
  const item = Array.from(root.querySelectorAll(selectableItems)).find((option) => option.getAttribute("data-value") === value);
  const display = root.querySelector('[data-slot="select-value"], [data-slot="combobox-value"]');
  if (item) {
    item.setAttribute("aria-selected", "true");
    item.setAttribute("data-state", "checked");
    if (display && !display.textContent.trim()) display.textContent = item.textContent.trim();
  }
  const name = root.getAttribute("data-name");
  if (name) {
    let input = root.querySelector(`input[type="hidden"][name="${CSS.escape(name)}"]`);
    if (!input) {
      input = document.createElement("input");
      input.type = "hidden";
      input.name = name;
      root.appendChild(input);
    }
    input.value = value;
  }
}

function focusByTypeahead(root, key) {
  const previous = typeahead.get(root) || { value: "", expires: 0 };
  const now = Date.now();
  const nextValue = (now > previous.expires ? "" : previous.value) + key.toLocaleLowerCase();
  typeahead.set(root, { value: nextValue, expires: now + 700 });
  const items = Array.from(root.querySelectorAll(selectableItems)).filter((item) => !item.disabled && item.getAttribute("aria-disabled") !== "true");
  const match = items.find((item) => item.textContent.trim().toLocaleLowerCase().startsWith(nextValue));
  if (match) {
    items.forEach((item) => item.removeAttribute("data-highlighted"));
    match.setAttribute("data-highlighted", "");
    match.focus();
  }
}
