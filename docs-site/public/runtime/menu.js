import { closeOpenFloating, contentFor, focusFirst, moveFocus, setOpen } from "./core.js";
import { positionFloating } from "./positioning.js";

const menuRoots = '[data-slot="dropdown-menu"], [data-dropdown-menu-root], [data-slot="context-menu"], [data-slot="menubar"], [data-slot="navigation-menu"]';
const menuTrigger = '[data-slot="dropdown-menu-trigger"], [data-dropdown-menu-trigger], [data-slot="menubar-trigger"], [data-slot="context-menu-trigger"], [data-slot="navigation-menu-trigger"]';

export function initMenus() {
  document.querySelectorAll(menuRoots).forEach((root) => {
    setOpen(root, root.getAttribute("data-open") === "true", { keepMounted: root.getAttribute("data-open") === "true", restoreFocus: false });
  });
  document.addEventListener("click", onClick);
  document.addEventListener("contextmenu", onContextMenu);
  document.addEventListener("keydown", onKeydown);
}

function onClick(event) {
  const trigger = event.target.closest(menuTrigger);
  const root = trigger?.closest(menuRoots);
  if (root) {
    event.preventDefault();
    const open = root.getAttribute("data-open") !== "true";
    if (open) closeOpenFloating(root);
    setOpen(root, open, { trigger });
    if (open) positionFloating(root);
    return;
  }

  const item = event.target.closest('[data-slot="dropdown-menu-item"], [data-slot="select-item"], [data-slot="combobox-item"], [data-dropdown-menu-item]');
  if (item) {
    handleSelection(item);
    return;
  }

  document.querySelectorAll(`${menuRoots}[data-open="true"]`).forEach((root) => {
    if (!root.contains(event.target)) setOpen(root, false, { restoreFocus: false });
  });
}

function onContextMenu(event) {
  const root = event.target.closest('[data-slot="context-menu"]');
  if (!root) return;
  event.preventDefault();
  root.style.setProperty("--context-menu-x", `${event.clientX}px`);
  root.style.setProperty("--context-menu-y", `${event.clientY}px`);
  setOpen(root, true);
  const content = contentFor(root);
  if (content) {
    content.style.position = "fixed";
    content.style.left = `${event.clientX}px`;
    content.style.top = `${event.clientY}px`;
    content.removeAttribute("hidden");
  }
  focusFirst(root);
}

function onKeydown(event) {
  const root = event.target.closest(`${menuRoots}[data-open="true"]`);

  if (!root) {
    const trigger = event.target.closest(menuTrigger);
    const triggerRoot = trigger?.closest(menuRoots);
    if (triggerRoot && ["ArrowDown", "Enter", " "].includes(event.key)) {
      event.preventDefault();
      setOpen(triggerRoot, true, { trigger });
      positionFloating(triggerRoot);
      focusFirst(triggerRoot);
    }
    return;
  }

  if (event.key === "Escape") {
    event.preventDefault();
    setOpen(root, false);
  } else if (event.key === "ArrowDown") {
    event.preventDefault();
    moveFocus(root, 1);
  } else if (event.key === "ArrowUp") {
    event.preventDefault();
    moveFocus(root, -1);
  } else if (event.key === "Home") {
    event.preventDefault();
    focusFirst(root);
  }
}

function handleSelection(item) {
  const selectRoot = item.closest('[data-slot="select"], [data-slot="combobox"]');
  if (selectRoot) {
    const value = item.getAttribute("data-value") || item.textContent.trim();
    selectRoot.setAttribute("data-value", value);
    const display = selectRoot.querySelector('[data-slot="select-value"], [data-slot="combobox-value"]');
    if (display) display.textContent = item.textContent.trim();
    const name = selectRoot.getAttribute("data-name");
    if (name) {
      let input = selectRoot.querySelector(`input[name="${CSS.escape(name)}"]`);
      if (!input) {
        input = document.createElement("input");
        input.type = "hidden";
        input.name = name;
        selectRoot.appendChild(input);
      }
      input.value = value;
      input.dispatchEvent(new Event("input", { bubbles: true }));
      input.dispatchEvent(new Event("change", { bubbles: true }));
    }
    setOpen(selectRoot, false);
    return;
  }

  const menuRoot = item.closest(menuRoots);
  if (menuRoot) setOpen(menuRoot, false);
}
