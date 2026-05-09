import { focusFirst, moveFocus } from "./core.js";

const commandRoot = '[data-slot="command"]';
const commandInput = '[data-slot="command-input"]';
const commandItem = '[data-slot="command-item"]';
const commandEmpty = '[data-slot="command-empty"]';

export function initCommands() {
  document.querySelectorAll(commandRoot).forEach(syncCommand);
  document.addEventListener("input", onInput);
  document.addEventListener("keydown", onKeydown);
}

function onInput(event) {
  const input = event.target.closest(commandInput);
  const root = input?.closest(commandRoot);
  if (!root) return;
  filterCommand(root, input.value);
}

function onKeydown(event) {
  const root = event.target.closest(commandRoot);
  if (!root) return;
  if (event.key === "ArrowDown") {
    event.preventDefault();
    moveFocus(root, 1);
  } else if (event.key === "ArrowUp") {
    event.preventDefault();
    moveFocus(root, -1);
  } else if (event.key === "Home") {
    event.preventDefault();
    focusFirst(root);
  } else if (event.key === "Enter") {
    const item = event.target.closest(commandItem) || root.querySelector(`${commandItem}[data-highlighted]`);
    if (item && !item.hidden && !item.disabled && item.getAttribute("aria-disabled") !== "true") {
      event.preventDefault();
      item.click();
    }
  }
}

function syncCommand(root) {
  const input = root.querySelector(commandInput);
  filterCommand(root, input?.value || "");
}

function filterCommand(root, query) {
  const normalized = query.trim().toLocaleLowerCase();
  let visible = 0;
  root.querySelectorAll(commandItem).forEach((item) => {
    const disabled = item.disabled || item.getAttribute("aria-disabled") === "true";
    const matches = !normalized || item.textContent.trim().toLocaleLowerCase().includes(normalized);
    item.toggleAttribute("hidden", !matches);
    item.setAttribute("aria-selected", "false");
    if (!disabled && matches) visible += 1;
  });
  root.querySelectorAll(commandEmpty).forEach((empty) => empty.toggleAttribute("hidden", visible !== 0));
}
