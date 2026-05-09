const checkboxSelector = '[data-slot="checkbox"], [data-slot="switch"]';
const inputOTPRoot = '[data-slot="input-otp"]';
const carouselRoot = '[data-slot="carousel"]';
const resizableRoot = '[data-slot="resizable-panel-group"]';
const toastRoot = '[data-slot="toast"]';
const sidebarProvider = '[data-slot="sidebar-provider"]';
const datePickerRoot = '[data-slot="date-picker"]';

export function initLongTail() {
  document.querySelectorAll(checkboxSelector).forEach(syncCheckedInput);
  document.querySelectorAll(inputOTPRoot).forEach(syncOTP);
  document.querySelectorAll(carouselRoot).forEach(syncCarousel);
  document.querySelectorAll(resizableRoot).forEach(syncResizable);
  document.querySelectorAll(toastRoot).forEach(syncToast);
  document.querySelectorAll(sidebarProvider).forEach(syncSidebar);
  document.querySelectorAll(datePickerRoot).forEach(syncDatePicker);
  document.addEventListener("change", onChange);
  document.addEventListener("input", onInput);
  document.addEventListener("click", onClick);
  document.addEventListener("keydown", onKeydown);
  document.addEventListener("pointerdown", onPointerDown);
}

function onChange(event) {
  const input = event.target.closest(checkboxSelector);
  if (input) syncCheckedInput(input);
}

function onInput(event) {
  const slot = event.target.closest('[data-slot="input-otp-slot"]');
  const root = slot?.closest(inputOTPRoot);
  if (root) syncOTP(root);
}

function onClick(event) {
  const previous = event.target.closest('[data-slot="carousel-previous"]');
  const next = event.target.closest('[data-slot="carousel-next"]');
  const carousel = (previous || next)?.closest(carouselRoot);
  if (carousel) {
    event.preventDefault();
    moveCarousel(carousel, next ? 1 : -1);
    return;
  }

  const toastClose = event.target.closest('[data-slot="toast-close"]');
  const toast = toastClose?.closest(toastRoot);
  if (toast) {
    event.preventDefault();
    closeToast(toast);
    return;
  }

  const sidebarTrigger = event.target.closest('[data-slot="sidebar-trigger"]');
  const sidebar = sidebarTrigger?.closest(sidebarProvider);
  if (sidebar) {
    event.preventDefault();
    const open = sidebar.getAttribute("data-open") !== "true";
    setSidebar(sidebar, open);
    return;
  }

  const dateDay = event.target.closest('[data-slot="date-picker"] [data-slot="calendar-day-button"]');
  const datePicker = dateDay?.closest(datePickerRoot);
  if (datePicker && !dateDay.disabled) {
    const value = dateDay.getAttribute("data-date");
    datePicker.setAttribute("data-value", value);
    const input = datePicker.querySelector('input[type="hidden"]');
    if (input) input.value = value;
  }
}

function onKeydown(event) {
  const carousel = event.target.closest(carouselRoot);
  if (carousel && ["ArrowLeft", "ArrowRight"].includes(event.key)) {
    event.preventDefault();
    moveCarousel(carousel, event.key === "ArrowRight" ? 1 : -1);
  }
}

function onPointerDown(event) {
  const handle = event.target.closest('[data-slot="resizable-handle"]');
  const root = handle?.closest(resizableRoot);
  if (!root) return;
  const panels = Array.from(root.querySelectorAll('[data-slot="resizable-panel"]'));
  const index = Array.from(root.children).indexOf(handle);
  const before = panels.findLast?.((panel) => Array.from(root.children).indexOf(panel) < index) || panels[0];
  const after = panels.find((panel) => Array.from(root.children).indexOf(panel) > index) || panels[1];
  if (!before || !after) return;
  event.preventDefault();
  const startX = event.clientX;
  const startBefore = Number.parseFloat(before.getAttribute("data-size") || "50");
  const startAfter = Number.parseFloat(after.getAttribute("data-size") || "50");
  const width = root.getBoundingClientRect().width || 1;
  const move = (moveEvent) => {
    const delta = ((moveEvent.clientX - startX) / width) * 100;
    setPanelSize(before, clamp(startBefore + delta, before));
    setPanelSize(after, clamp(startAfter - delta, after));
  };
  const up = () => {
    window.removeEventListener("pointermove", move);
    window.removeEventListener("pointerup", up);
  };
  window.addEventListener("pointermove", move);
  window.addEventListener("pointerup", up, { once: true });
}

function syncCheckedInput(input) {
  input.setAttribute("data-state", input.checked ? "checked" : "unchecked");
  input.setAttribute("aria-checked", input.checked ? "true" : "false");
}

function syncOTP(root) {
  const slots = Array.from(root.querySelectorAll('[data-slot="input-otp-slot"]'));
  const value = slots.map((slot) => slot.textContent.trim()).join("");
  root.setAttribute("data-value", value);
  const input = root.querySelector('input[type="hidden"]');
  if (input) input.value = value;
}

function syncCarousel(root) {
  const slides = Array.from(root.querySelectorAll('[data-slot="carousel-item"]'));
  const index = Math.max(0, Math.min(Number.parseInt(root.getAttribute("data-index") || root.getAttribute("data-start-index") || "0", 10), slides.length - 1));
  root.setAttribute("data-index", String(index));
  slides.forEach((slide, slideIndex) => {
    slide.toggleAttribute("hidden", slideIndex !== index);
    slide.setAttribute("aria-hidden", slideIndex === index ? "false" : "true");
  });
}

function moveCarousel(root, direction) {
  const slides = root.querySelectorAll('[data-slot="carousel-item"]');
  const loop = root.getAttribute("data-loop") === "true";
  let index = Number.parseInt(root.getAttribute("data-index") || "0", 10) + direction;
  if (loop) index = (index + slides.length) % slides.length;
  else index = Math.max(0, Math.min(index, slides.length - 1));
  root.setAttribute("data-index", String(index));
  syncCarousel(root);
}

function syncResizable(root) {
  root.querySelectorAll('[data-slot="resizable-panel"]').forEach((panel) => {
    const size = Number.parseFloat(panel.getAttribute("data-size") || panel.getAttribute("data-default-size") || "50");
    setPanelSize(panel, size);
  });
}

function setPanelSize(panel, size) {
  panel.setAttribute("data-size", String(Math.round(size * 100) / 100));
  panel.style.flexBasis = `${size}%`;
}

function clamp(value, panel) {
  const min = Number.parseFloat(panel.getAttribute("data-min-size") || "5");
  const max = Number.parseFloat(panel.getAttribute("data-max-size") || "95");
  return Math.max(min, Math.min(max, value));
}

function syncToast(toast) {
  const open = toast.getAttribute("data-open") === "true" || toast.getAttribute("data-default-open") === "true";
  toast.setAttribute("data-state", open ? "open" : "closed");
  toast.toggleAttribute("hidden", !open);
  const duration = Number.parseInt(toast.getAttribute("data-duration") || "0", 10);
  if (open && duration > 0) window.setTimeout(() => closeToast(toast), duration);
}

function closeToast(toast) {
  toast.removeAttribute("data-open");
  toast.setAttribute("data-state", "closed");
  toast.setAttribute("hidden", "");
}

function syncSidebar(root) {
  setSidebar(root, root.getAttribute("data-open") === "true" || root.getAttribute("data-default-open") === "true");
}

function setSidebar(root, open) {
  if (open) root.setAttribute("data-open", "true");
  else root.removeAttribute("data-open");
  root.setAttribute("data-state", open ? "open" : "closed");
  root.querySelector('[data-slot="sidebar-trigger"]')?.setAttribute("aria-expanded", open ? "true" : "false");
}

function syncDatePicker(root) {
  const input = root.querySelector('input[type="hidden"]');
  if (input && root.getAttribute("data-value")) input.value = root.getAttribute("data-value");
}
