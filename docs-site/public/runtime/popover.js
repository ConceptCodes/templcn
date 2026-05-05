import { closeOpenFloating, focusFirst, setOpen } from "./core.js";
import { positionFloating } from "./positioning.js";

const roots = '[data-slot="popover"], [data-slot="hover-card"], [data-slot="tooltip"], [data-slot="select"], [data-slot="combobox"]';
const triggerSelector = '[data-slot="popover-trigger"], [data-slot="hover-card-trigger"], [data-slot="tooltip-trigger"], [data-slot="select-trigger"], [data-slot="combobox-trigger"]';

export function initPopovers() {
  document.querySelectorAll(roots).forEach((root) => {
    setOpen(root, root.getAttribute("data-open") === "true", { keepMounted: root.getAttribute("data-open") === "true", restoreFocus: false });
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
      if (event.detail === 0) focusFirst(root);
    }
    return;
  }

  document.querySelectorAll(`${roots}[data-open="true"]`).forEach((openRoot) => {
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
  if (event.key !== "Escape") return;
  const root = document.querySelector(`${roots}[data-open="true"]`);
  if (!root) return;
  event.preventDefault();
  setOpen(root, false);
}
