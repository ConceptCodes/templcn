import { setOpen } from "./core.js";

export function initTabsAndDisclosures() {
  document.addEventListener("click", onClick);
  document.querySelectorAll('[data-slot="accordion-item"], [data-slot="collapsible"]').forEach(syncDetailsState);
  document.addEventListener("toggle", (event) => {
    if (event.target.matches('[data-slot="accordion-item"], [data-slot="collapsible"]')) {
      syncDetailsState(event.target);
    }
  }, true);
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

  if (root.getAttribute("data-type") === "single" && !item.open) {
    root.querySelectorAll('[data-slot="accordion-item"][open]').forEach((other) => {
      if (other !== item) other.removeAttribute("open");
      syncDetailsState(other);
    });
  }
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

