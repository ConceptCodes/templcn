const radioGroups = '[data-slot="radio-group"]';
const toggleGroups = '[data-slot="toggle-group"]';
const sliderRoots = '[data-slot="slider"]';

export function initFormControls() {
  document.querySelectorAll(radioGroups).forEach(syncRadioGroup);
  document.querySelectorAll(toggleGroups).forEach(syncToggleGroup);
  document.querySelectorAll(sliderRoots).forEach(syncSlider);
  document.addEventListener("change", onChange);
  document.addEventListener("click", onClick);
  document.addEventListener("keydown", onKeydown);
  document.addEventListener("input", onInput);
}

function onChange(event) {
  const radio = event.target.closest('[data-slot="radio-group-item"] input[type="radio"]');
  const radioRoot = radio?.closest(radioGroups);
  if (radioRoot) syncRadioGroup(radioRoot);
}

function onClick(event) {
  const item = event.target.closest('[data-slot="toggle-group-item"]');
  const root = item?.closest(toggleGroups);
  if (!root || item.disabled || item.getAttribute("aria-disabled") === "true") return;
  event.preventDefault();
  const type = root.getAttribute("data-type") || "multiple";
  const currentlyOn = item.getAttribute("data-state") === "on";
  if (type === "single") {
    root.querySelectorAll('[data-slot="toggle-group-item"]').forEach((candidate) => setToggleItem(candidate, false));
    setToggleItem(item, !currentlyOn);
  } else {
    setToggleItem(item, !currentlyOn);
  }
  syncToggleGroup(root);
}

function onKeydown(event) {
  const toggle = event.target.closest('[data-slot="toggle-group-item"]');
  const toggleRoot = toggle?.closest(toggleGroups);
  if (toggleRoot && ["ArrowRight", "ArrowDown", "ArrowLeft", "ArrowUp", "Home", "End"].includes(event.key)) {
    event.preventDefault();
    moveToggleFocus(toggleRoot, toggle, event.key);
    return;
  }

  const radioItem = event.target.closest('[data-slot="radio-group-item"]');
  const radioRoot = radioItem?.closest(radioGroups);
  if (radioRoot && ["ArrowRight", "ArrowDown", "ArrowLeft", "ArrowUp", "Home", "End"].includes(event.key)) {
    event.preventDefault();
    moveRadioFocus(radioRoot, radioItem, event.key);
  }
}

function onInput(event) {
  const input = event.target.closest('[data-slot="slider"] input[type="range"]');
  const root = input?.closest(sliderRoots);
  if (root) syncSlider(root);
}

function syncRadioGroup(root) {
  const checked = root.querySelector('[data-slot="radio-group-item"] input[type="radio"]:checked');
  if (checked) root.setAttribute("data-value", checked.value);
  root.querySelectorAll('[data-slot="radio-group-item"]').forEach((item) => {
    const input = item.querySelector('input[type="radio"]');
    const selected = input?.checked === true;
    item.setAttribute("data-state", selected ? "checked" : "unchecked");
    item.setAttribute("aria-checked", selected ? "true" : "false");
  });
}

function syncToggleGroup(root) {
  const values = Array.from(root.querySelectorAll('[data-slot="toggle-group-item"][data-state="on"]')).map((item) => item.getAttribute("data-value")).filter(Boolean);
  root.setAttribute("data-value", values.join(" "));
}

function setToggleItem(item, on) {
  item.setAttribute("data-state", on ? "on" : "off");
  item.setAttribute("aria-pressed", on ? "true" : "false");
}

function syncSlider(root) {
  const inputs = Array.from(root.querySelectorAll('input[type="range"]'));
  if (!inputs.length) return;
  const values = inputs.map((input) => input.value);
  root.setAttribute("data-value", values.join(" "));
  inputs.forEach((input) => input.setAttribute("aria-valuenow", input.value));
  root.querySelectorAll('[data-slot="slider-thumb"]').forEach((thumb, index) => {
    if (values[index] !== undefined) thumb.setAttribute("data-value", values[index]);
  });
  const min = Number.parseFloat(inputs[0].min || "0");
  const max = Number.parseFloat(inputs[0].max || "100");
  const percent = (value) => {
    if (!Number.isFinite(min) || !Number.isFinite(max) || max <= min) return 0;
    return Math.min(100, Math.max(0, ((Number.parseFloat(value) - min) / (max - min)) * 100));
  };
  const percentages = values.map(percent);
  const start = percentages.length > 1 ? Math.min(...percentages) : 0;
  const end = Math.max(...percentages);
  const range = root.querySelector('[data-slot="slider-range"]');
  if (range) {
    if (root.getAttribute("data-orientation") === "vertical") {
      range.style.bottom = `${start}%`;
      range.style.height = `${end - start}%`;
    } else {
      range.style.left = `${start}%`;
      range.style.width = `${end - start}%`;
    }
  }
}

function moveToggleFocus(root, current, key) {
  const items = enabledItems(root, '[data-slot="toggle-group-item"]');
  moveFocusInList(items, current, key)?.focus();
}

function moveRadioFocus(root, current, key) {
  const items = enabledItems(root, '[data-slot="radio-group-item"]');
  const next = moveFocusInList(items, current, key);
  const input = next?.querySelector('input[type="radio"]');
  if (input) {
    input.checked = true;
    input.focus();
    input.dispatchEvent(new Event("change", { bubbles: true }));
  }
}

function enabledItems(root, selector) {
  return Array.from(root.querySelectorAll(selector)).filter((item) => !item.hasAttribute("disabled") && item.getAttribute("aria-disabled") !== "true" && !item.querySelector(":disabled"));
}

function moveFocusInList(items, current, key) {
  const index = items.indexOf(current);
  if (index === -1 || !items.length) return null;
  if (key === "Home") return items[0];
  if (key === "End") return items[items.length - 1];
  const direction = key === "ArrowRight" || key === "ArrowDown" ? 1 : -1;
  return items[(index + direction + items.length) % items.length];
}
