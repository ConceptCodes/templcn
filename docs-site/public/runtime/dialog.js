import { focusFirst, focusable, getState, setOpen } from "./core.js";

const triggerSelector = '[data-slot="dialog-trigger"], [data-slot="sheet-trigger"], [data-slot="drawer-trigger"], [data-slot="alert-dialog-trigger"]';
const closeSelector = '[data-slot$="-close"], [data-slot="alert-dialog-action"], [data-slot="alert-dialog-cancel"]';

export function initDialogs() {
  document.addEventListener("click", onClick);
  document.addEventListener("keydown", onKeydown);
  document.querySelectorAll("dialog").forEach((dialog) => {
    setDialogState(dialog, dialog.open || dialog.getAttribute("data-open") === "true");
    dialog.addEventListener("cancel", (event) => {
      event.preventDefault();
      closeDialog(dialog);
    });
    dialog.addEventListener("close", () => setDialogState(dialog, false));
  });
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
    const content = dialog.querySelector("[data-dialog-panel]") || (dialog.matches('[data-slot$="-content"]') ? dialog : dialog.querySelector('[data-slot$="-content"]'));
    if (!content || !content.contains(document.elementFromPoint(event.clientX, event.clientY))) {
      closeDialog(dialog);
    }
  }
}

function onKeydown(event) {
  const dialogs = Array.from(document.querySelectorAll("dialog[open]"));
  const dialog = dialogs[dialogs.length - 1];
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

  const owner = trigger.closest('[data-slot="dialog"], [data-slot="alert-dialog"], [data-slot="sheet"], [data-slot="drawer"]');
  const ownedDialog = owner?.querySelector("dialog");
  if (ownedDialog) return ownedDialog;

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
  setDialogState(dialog, true, trigger);
  focusFirst(dialog);
}

function closeDialog(dialog) {
  if (dialog.open) dialog.close();
  setDialogState(dialog, false);
  const trigger = getState(dialog).lastTrigger;
  if (trigger && document.contains(trigger)) trigger.focus();
}

function triggerForDialog(dialog) {
  const stored = getState(dialog).lastTrigger;
  if (stored && document.contains(stored)) return stored;
  if (!dialog.id) return null;
  return document.querySelector(`${triggerSelector}[aria-controls="${CSS.escape(dialog.id)}"]`);
}

function setDialogState(dialog, open, trigger = triggerForDialog(dialog)) {
  setOpen(dialog, open, { keepMounted: true, restoreFocus: false });
  const owner = dialog.closest('[data-slot="dialog"], [data-slot="alert-dialog"], [data-slot="sheet"], [data-slot="drawer"]');
  if (owner && owner !== dialog) {
    owner.setAttribute("data-state", open ? "open" : "closed");
    if (open) owner.setAttribute("data-open", "true");
    else owner.removeAttribute("data-open");
  }
  trigger?.setAttribute("aria-expanded", open ? "true" : "false");
  trigger?.setAttribute("data-state", open ? "open" : "closed");
  document.querySelectorAll(triggerSelector).forEach((candidate) => {
    if (candidate === trigger || candidate.getAttribute("aria-controls") === dialog.id || candidate.getAttribute("data-target") === dialog.id) {
      candidate.setAttribute("aria-expanded", open ? "true" : "false");
      candidate.setAttribute("data-state", open ? "open" : "closed");
    }
  });
  dialog.querySelectorAll('[data-dialog-panel], [data-slot$="-content"], [data-slot$="-overlay"]').forEach((el) => {
    el.setAttribute("data-state", open ? "open" : "closed");
  });
}
