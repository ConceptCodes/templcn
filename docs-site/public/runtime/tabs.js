import { setOpen } from "./core.js";

export function initTabsAndDisclosures() {
  document.addEventListener("click", onClick);
  document.addEventListener("keydown", onKeydown);
  document.querySelectorAll('[data-slot="tabs"], [data-tabs-root]').forEach(syncTabsState);
  document.querySelectorAll('[data-slot="accordion"]').forEach(syncAccordionDefaults);
  document.querySelectorAll('[data-slot="accordion-item"], [data-slot="collapsible"]').forEach(syncDetailsState);
  document.addEventListener("toggle", (event) => {
    if (event.target.matches('[data-slot="accordion-item"], [data-slot="collapsible"]')) {
      syncDetailsState(event.target);
    }
  }, true);
}

function onKeydown(event) {
  const accordionTrigger = event.target.closest('[data-slot="accordion-trigger"]');
  if (accordionTrigger) {
    const root = accordionTrigger.closest('[data-slot="accordion"]');
    const triggers = root ? Array.from(root.querySelectorAll('[data-slot="accordion-trigger"]')).filter((trigger) => trigger.getAttribute("aria-disabled") !== "true") : [];
    const current = triggers.indexOf(accordionTrigger);
    let next = -1;
    if (event.key === "ArrowDown") next = (current + 1) % triggers.length;
    else if (event.key === "ArrowUp") next = (current - 1 + triggers.length) % triggers.length;
    else if (event.key === "Home") next = 0;
    else if (event.key === "End") next = triggers.length - 1;
    if (current >= 0 && next >= 0) {
      event.preventDefault();
      triggers[next].focus();
      return;
    }
  }

  const trigger = event.target.closest('[data-slot="tabs-trigger"]');
  if (!trigger) return;
  const root = trigger.closest('[data-slot="tabs"], [data-tabs-root]');
  if (!root) return;

  const orientation = root.getAttribute("data-orientation") || "horizontal";
  const nextKeys = orientation === "vertical" ? ["ArrowDown"] : ["ArrowRight", "ArrowDown"];
  const previousKeys = orientation === "vertical" ? ["ArrowUp"] : ["ArrowLeft", "ArrowUp"];
  const triggers = enabledTabs(root);
  const current = triggers.indexOf(trigger);
  if (current === -1) return;

  let next = -1;
  if (nextKeys.includes(event.key)) next = (current + 1) % triggers.length;
  else if (previousKeys.includes(event.key)) next = (current - 1 + triggers.length) % triggers.length;
  else if (event.key === "Home") next = 0;
  else if (event.key === "End") next = triggers.length - 1;
  else if (event.key === "Enter" || event.key === " ") next = current;

  if (next >= 0) {
    event.preventDefault();
    triggers[next].focus();
    activateTab(triggers[next]);
  }
}

function syncTabsState(root) {
  const value = root.getAttribute("data-value") || root.getAttribute("data-default-value") || root.querySelector('[data-slot="tabs-trigger"]')?.getAttribute("data-value");
  if (!value) return;
  const trigger = root.querySelector(`[data-slot="tabs-trigger"][data-value="${CSS.escape(value)}"]`);
  if (trigger) activateTab(trigger);
}

function onClick(event) {
  const tab = event.target.closest('[data-slot="tabs-trigger"]');
  if (tab) {
    event.preventDefault();
    activateTab(tab);
    return;
  }

  const accordionTrigger = event.target.closest('[data-slot="accordion-trigger"]');
  if (accordionTrigger) syncAccordion(accordionTrigger);

  const carouselButton = event.target.closest('[data-slot="carousel-previous"], [data-slot="carousel-next"]');
  if (carouselButton) scrollCarousel(carouselButton);
}

function enabledTabs(root) {
  return Array.from(root.querySelectorAll('[data-slot="tabs-trigger"]')).filter((trigger) => !trigger.disabled && trigger.getAttribute("aria-disabled") !== "true");
}

function activateTab(trigger) {
  const root = trigger.closest('[data-slot="tabs"], [data-tabs-root]');
  const value = trigger.getAttribute("data-value");
  if (!root || !value) return;

  root.querySelectorAll('[data-slot="tabs-trigger"]').forEach((candidate) => {
    const active = candidate.getAttribute("data-value") === value;
    candidate.setAttribute("data-state", active ? "active" : "inactive");
    candidate.setAttribute("aria-selected", active ? "true" : "false");
    candidate.setAttribute("tabindex", active ? "0" : "-1");
  });

  root.querySelectorAll('[data-slot="tabs-content"]').forEach((panel) => {
    const active = panel.getAttribute("data-value") === value;
    panel.setAttribute("data-state", active ? "active" : "inactive");
    panel.toggleAttribute("hidden", !active);
  });
}

function syncAccordion(trigger) {
  const item = trigger.closest('[data-slot="accordion-item"]');
  const root = item?.closest('[data-slot="accordion"]');
  if (!item || !root) return;
  if (item.getAttribute("data-disabled") === "true") {
    item.removeAttribute("open");
    syncDetailsState(item);
    return;
  }

  if (root.getAttribute("data-type") === "single" && !item.open) {
    root.querySelectorAll('[data-slot="accordion-item"][open]').forEach((other) => {
      if (other !== item) other.removeAttribute("open");
      syncDetailsState(other);
    });
  }
}

function syncAccordionDefaults(root) {
  const raw = root.getAttribute("data-value") || root.getAttribute("data-default-value") || "";
  if (!raw) return;
  const values = raw.split(/[\s,]+/).filter(Boolean);
  root.querySelectorAll('[data-slot="accordion-item"]').forEach((item) => {
    item.toggleAttribute("open", values.includes(item.getAttribute("data-value")));
    syncDetailsState(item);
  });
}

function syncDetailsState(details) {
  setOpen(details, details.open, { keepMounted: true, restoreFocus: false });
  const content = details.querySelector('[data-slot$="-content"]');
  if (content) content.setAttribute("data-state", details.open ? "open" : "closed");
}

function scrollCarousel(button) {
  const carousel = button.closest('[data-slot="carousel"]');
  const content = carousel?.querySelector('[data-slot="carousel-content"]');
  if (!content) return;

  const item = content.querySelector('[data-slot="carousel-item"]');
  const amount = item ? item.getBoundingClientRect().width : content.clientWidth;
  content.scrollBy({
    left: button.matches('[data-slot="carousel-next"]') ? amount : -amount,
    behavior: "smooth",
  });
}
