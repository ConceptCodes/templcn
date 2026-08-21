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
  if (!open) {
    toast.style.display = "none";
  } else {
    toast.style.display = "";
  }
  const duration = Number.parseInt(toast.getAttribute("data-duration") || "0", 10);
  if (open && duration > 0) window.setTimeout(() => closeToast(toast), duration);
}

function closeToast(toast) {
  toast.removeAttribute("data-open");
  toast.setAttribute("data-state", "closed");
  toast.classList.remove("slide-in-from-bottom-5", "fade-in");
  toast.classList.add("fade-out", "slide-out-to-right-full", "duration-200");
  toast.style.opacity = "0";
  toast.style.transform = "translateX(100%)";
  setTimeout(() => {
    toast.style.display = "none";
    toast.remove();
  }, 200);
}

export function toast(title, options = {}) {
  let toaster = document.querySelector('[data-slot="toaster"]') || document.querySelector('[data-toaster]');
  if (!toaster) {
    toaster = document.createElement('div');
    toaster.setAttribute('data-slot', 'toaster');
    toaster.className = 'fixed bottom-0 right-0 z-50 flex flex-col gap-2 p-4 max-w-md w-full pointer-events-none';
    document.body.appendChild(toaster);
  }

  const toastEl = document.createElement('div');
  toastEl.setAttribute('data-slot', 'toast');
  toastEl.setAttribute('role', 'status');
  toastEl.setAttribute('aria-live', 'polite');
  toastEl.className = 'pointer-events-auto relative flex w-full items-center justify-between space-x-4 overflow-hidden rounded-md border p-4 pr-6 shadow-lg transition-all bg-background text-foreground animate-in fade-in slide-in-from-bottom-5 duration-300';
  
  if (options.variant === 'destructive' || options.type === 'error') {
    toastEl.className += ' border-destructive/50 text-destructive bg-destructive/10';
  } else if (options.type === 'success') {
    toastEl.className += ' border-emerald-500/50 text-emerald-600 dark:text-emerald-400 bg-emerald-500/10';
  } else if (options.type === 'warning') {
    toastEl.className += ' border-amber-500/50 text-amber-600 dark:text-amber-400 bg-amber-500/10';
  } else if (options.type === 'info') {
    toastEl.className += ' border-blue-500/50 text-blue-600 dark:text-blue-400 bg-blue-500/10';
  }

  const content = document.createElement('div');
  content.className = 'grid gap-1';

  const titleEl = document.createElement('div');
  titleEl.className = 'text-sm font-semibold';
  titleEl.textContent = title;
  content.appendChild(titleEl);

  if (options.description) {
    const descEl = document.createElement('div');
    descEl.className = 'text-xs text-muted-foreground';
    descEl.textContent = options.description;
    content.appendChild(descEl);
  }

  toastEl.appendChild(content);

  if (options.action) {
    const actionBtn = document.createElement('button');
    actionBtn.type = 'button';
    actionBtn.className = 'inline-flex h-8 shrink-0 items-center justify-center rounded-md border bg-transparent px-3 text-xs font-medium transition-colors hover:bg-secondary focus:outline-none focus:ring-1 focus:ring-ring cursor-pointer';
    actionBtn.textContent = options.action.label || 'Action';
    actionBtn.addEventListener('click', (e) => {
      e.stopPropagation();
      options.action.onClick?.();
      dismiss();
    });
    toastEl.appendChild(actionBtn);
  }

  const closeBtn = document.createElement('button');
  closeBtn.type = 'button';
  closeBtn.setAttribute('data-slot', 'toast-close');
  closeBtn.setAttribute('aria-label', 'Close toast');
  closeBtn.className = 'absolute right-2 top-2 rounded-md p-1 text-foreground/50 opacity-70 transition-opacity hover:opacity-100 focus:opacity-100 focus:outline-none cursor-pointer';
  closeBtn.innerHTML = '<svg class="size-4 pointer-events-none" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/></svg>';
  closeBtn.onclick = (e) => {
    e.stopPropagation();
    e.preventDefault();
    dismiss();
  };
  toastEl.appendChild(closeBtn);

  function dismiss() {
    toastEl.classList.remove('slide-in-from-bottom-5', 'fade-in');
    toastEl.classList.add('fade-out', 'slide-out-to-right-full', 'duration-200');
    toastEl.style.opacity = '0';
    toastEl.style.transform = 'translateX(100%)';
    setTimeout(() => {
      toastEl.style.display = 'none';
      toastEl.remove();
    }, 200);
  }

  toaster.appendChild(toastEl);

  const duration = options.duration || 4000;
  if (duration > 0) {
    setTimeout(dismiss, duration);
  }

  return { dismiss };
}

toast.success = (title, options = {}) => toast(title, { ...options, type: 'success' });
toast.error = (title, options = {}) => toast(title, { ...options, type: 'error' });
toast.warning = (title, options = {}) => toast(title, { ...options, type: 'warning' });
toast.info = (title, options = {}) => toast(title, { ...options, type: 'info' });
toast.message = (title, options = {}) => toast(title, options);

if (typeof window !== 'undefined') {
  window.toast = toast;
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
