import { focusFirst, focusable, getState, setOpen } from "./core.js";

const triggerSelector = '[data-slot="dialog-trigger"], [data-slot="sheet-trigger"], [data-slot="drawer-trigger"], [data-slot="alert-dialog-trigger"]';
const closeSelector = '[data-slot$="-close"], [data-slot="alert-dialog-action"], [data-slot="alert-dialog-cancel"]';

export function initDialogs() {
  document.addEventListener("click", onClick);
  document.addEventListener("keydown", onKeydown);
  document.querySelectorAll("dialog[open]").forEach((dialog) => setDialogState(dialog, true));
}

function onClick(event) {
  const trigger = event.target.closest(triggerSelector);
  if (trigger) {
    const dialog = resolveDialog(trigger);
    if (dialog) {
      event.preventDefault();
      openDialog(dialog, trigger);
    }
    return;
  }

  const close = event.target.closest(closeSelector);
  if (close) {
    const dialog = close.closest("dialog");
    if (dialog) {
      event.preventDefault();
      closeDialog(dialog);
    }
    return;
  }

  const dialog = event.target.closest("dialog");
  if (dialog && event.target === dialog && dialog.open) {
    const content = dialog.querySelector('[data-slot$="-content"]');
    if (!content || !content.contains(document.elementFromPoint(event.clientX, event.clientY))) {
      closeDialog(dialog);
    }
  }
}

function onKeydown(event) {
  const dialog = document.querySelector("dialog[open]");
  if (!dialog) return;

  if (event.key === "Escape") {
    event.preventDefault();
    closeDialog(dialog);
    return;
  }

  if (event.key === "Tab" && dialog.getAttribute("data-modal") !== "false") {
    const nodes = focusable(dialog);
    if (!nodes.length) return;
    const first = nodes[0];
    const last = nodes[nodes.length - 1];
    if (event.shiftKey && document.activeElement === first) {
      event.preventDefault();
      last.focus();
    } else if (!event.shiftKey && document.activeElement === last) {
      event.preventDefault();
      first.focus();
    }
  }
}

function resolveDialog(trigger) {
  const targetId = trigger.getAttribute("aria-controls") || trigger.getAttribute("data-target");
  if (targetId) return document.getElementById(targetId);

  const root = trigger.closest("dialog");
  if (root) return root;

  let next = trigger.nextElementSibling;
  while (next) {
    if (next.tagName === "DIALOG") return next;
    next = next.nextElementSibling;
  }
  return document.querySelector("dialog");
}

function openDialog(dialog, trigger) {
  getState(dialog).lastTrigger = trigger;
  const modal = dialog.getAttribute("data-modal") !== "false";
  if (!dialog.open) {
    if (modal && typeof dialog.showModal === "function") dialog.showModal();
    else dialog.show();
  }
  setDialogState(dialog, true);
  focusFirst(dialog);
}

function closeDialog(dialog) {
  if (dialog.open) dialog.close();
  setDialogState(dialog, false);
  const trigger = getState(dialog).lastTrigger;
  if (trigger && document.contains(trigger)) trigger.focus();
}

function setDialogState(dialog, open) {
  setOpen(dialog, open, { keepMounted: true, restoreFocus: false });
  dialog.querySelectorAll('[data-slot$="-content"], [data-slot$="-overlay"]').forEach((el) => {
    el.setAttribute("data-state", open ? "open" : "closed");
  });
}

