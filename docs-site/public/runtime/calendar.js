const calendarRoot = '[data-slot="calendar"]';
const dayButton = '[data-slot="calendar-day-button"]';
const monthNames = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
const shortMonthNames = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];
const weekdays = ["Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"];

export function initCalendars() {
  document.querySelectorAll(calendarRoot).forEach(renderCalendar);
  document.addEventListener("click", onClick);
  document.addEventListener("change", onChange);
  document.addEventListener("keydown", onKeydown);
}

function onClick(event) {
  const target = event.target instanceof Element ? event.target : event.target?.parentElement;
  const root = target?.closest(calendarRoot);
  if (!root) return;
  if (target.closest('[data-slot="calendar-prev"]')) {
    event.preventDefault();
    shiftMonth(root, -1);
    return;
  }
  if (target.closest('[data-slot="calendar-next"]')) {
    event.preventDefault();
    shiftMonth(root, 1);
    return;
  }
  const day = target.closest(dayButton);
  if (day && !day.disabled) {
    event.preventDefault();
    selectDay(root, day.getAttribute("data-date"));
  }
}

function onChange(event) {
  const select = event.target.closest('[data-slot="calendar-month-select"], [data-slot="calendar-year-select"]');
  const root = select?.closest(calendarRoot);
  if (!root) return;
  const month = Number.parseInt(root.querySelector('[data-slot="calendar-month-select"]')?.value || "1", 10);
  const year = Number.parseInt(root.querySelector('[data-slot="calendar-year-select"]')?.value || String(new Date().getUTCFullYear()), 10);
  root.setAttribute("data-current-month", `${year}-${pad(month)}`);
  renderCalendar(root);
  root.dispatchEvent(new CustomEvent("calendar-month-change", { bubbles: true, detail: { month: root.getAttribute("data-current-month") } }));
}

function onKeydown(event) {
  const day = event.target.closest(dayButton);
  const root = day?.closest(calendarRoot);
  if (!root) return;
  const offsets = { ArrowRight: 1, ArrowLeft: -1, ArrowDown: 7, ArrowUp: -7 };
  if (event.key in offsets) {
    event.preventDefault();
    moveDayFocus(root, day, offsets[event.key]);
    return;
  }
  if (event.key === "Home" || event.key === "End") {
    event.preventDefault();
    const days = visibleDays(root);
    days[event.key === "Home" ? 0 : days.length - 1]?.focus();
    return;
  }
  if (event.key === "PageUp" || event.key === "PageDown") {
    event.preventDefault();
    shiftMonth(root, event.key === "PageUp" ? -1 : 1);
  }
}

function selectDay(root, value) {
  const mode = root.getAttribute("data-mode") || "single";
  if (mode === "range") {
    const [start, end] = parseRange(root.getAttribute("data-range") || root.getAttribute("data-selected") || "");
    if (!start || end || value < start) root.setAttribute("data-range", `${value}/`);
    else root.setAttribute("data-range", `${start}/${value}`);
  } else if (mode === "multiple") {
    const values = splitValues(root.getAttribute("data-selected"));
    const index = values.indexOf(value);
    if (index === -1) values.push(value);
    else values.splice(index, 1);
    root.setAttribute("data-selected", values.join(","));
  } else {
    root.setAttribute("data-selected", value);
  }
  syncFormValue(root);
  renderCalendar(root);
  root.dispatchEvent(new CustomEvent("calendar-select", { bubbles: true, detail: calendarValue(root) }));
}

function shiftMonth(root, delta) {
  const current = parseMonth(root.getAttribute("data-current-month"));
  current.setUTCMonth(current.getUTCMonth() + delta);
  root.setAttribute("data-current-month", formatMonth(current));
  renderCalendar(root);
  root.dispatchEvent(new CustomEvent("calendar-month-change", { bubbles: true, detail: { month: formatMonth(current) } }));
}

function renderCalendar(root) {
  const current = parseMonth(root.getAttribute("data-current-month"));
  root.setAttribute("data-current-month", formatMonth(current));
  renderCaption(root, current);
  let months = root.querySelector('[data-slot="calendar-months"]');
  if (!months) {
    months = document.createElement("div");
    months.setAttribute("data-slot", "calendar-months");
    months.className = "flex flex-col gap-4 sm:flex-row";
    root.appendChild(months);
  }
  months.innerHTML = "";
  const count = Math.max(1, Number.parseInt(root.getAttribute("data-number-of-months") || "1", 10));
  for (let index = 0; index < count; index += 1) {
    const month = new Date(Date.UTC(current.getUTCFullYear(), current.getUTCMonth() + index, 1));
    months.appendChild(renderMonth(root, month));
  }
  syncExistingDayButtons(root);
  syncFormValue(root);
}

function renderCaption(root, month) {
  const caption = root.querySelector('[data-slot="calendar-caption"]');
  if (!caption) return;
  if (root.getAttribute("data-caption-layout") !== "dropdown") {
    caption.textContent = `${monthNames[month.getUTCMonth()]} ${month.getUTCFullYear()}`;
    return;
  }
  caption.className = "flex items-center gap-1";
  caption.textContent = "";
  const monthSelect = document.createElement("select");
  monthSelect.setAttribute("data-slot", "calendar-month-select");
  monthSelect.setAttribute("aria-label", "Month");
  monthSelect.className = "h-8 rounded-md border bg-background px-2 text-sm";
  shortMonthNames.forEach((label, index) => monthSelect.appendChild(new Option(label, String(index + 1), false, index === month.getUTCMonth())));
  const yearSelect = document.createElement("select");
  yearSelect.setAttribute("data-slot", "calendar-year-select");
  yearSelect.setAttribute("aria-label", "Year");
  yearSelect.className = "h-8 rounded-md border bg-background px-2 text-sm";
  const from = Number.parseInt(root.getAttribute("data-from-year") || String(month.getUTCFullYear() - 50), 10);
  const to = Number.parseInt(root.getAttribute("data-to-year") || String(month.getUTCFullYear() + 50), 10);
  for (let year = Math.min(from, to); year <= Math.max(from, to); year += 1) {
    yearSelect.appendChild(new Option(String(year), String(year), false, year === month.getUTCFullYear()));
  }
  caption.append(monthSelect, yearSelect);
}

function renderMonth(root, month) {
  const wrapper = document.createElement("div");
  wrapper.setAttribute("data-slot", "calendar-month");
  wrapper.setAttribute("data-month", formatMonth(month));
  wrapper.className = "space-y-2";
  const grid = document.createElement("div");
  grid.setAttribute("data-slot", "calendar-grid");
  grid.setAttribute("role", "grid");
  grid.setAttribute("aria-label", `${monthNames[month.getUTCMonth()]} ${month.getUTCFullYear()}`);
  grid.className = "space-y-1";
  grid.appendChild(renderHeaderRow(root));
  const start = new Date(Date.UTC(month.getUTCFullYear(), month.getUTCMonth(), 1 - month.getUTCDay()));
  for (let week = 0; week < 6; week += 1) {
    grid.appendChild(renderWeek(root, month, start, week));
  }
  wrapper.appendChild(grid);
  return wrapper;
}

function renderHeaderRow(root) {
  const row = document.createElement("div");
  row.setAttribute("role", "row");
  row.className = rowClass(root);
  if (root.getAttribute("data-show-week-number") === "true") {
    row.appendChild(cellText("", "columnheader", "Week number"));
  }
  weekdays.forEach((weekday) => row.appendChild(cellText(weekday, "columnheader")));
  return row;
}

function renderWeek(root, month, start, week) {
  const row = document.createElement("div");
  row.setAttribute("role", "row");
  row.className = rowClass(root);
  if (root.getAttribute("data-show-week-number") === "true") {
    const weekDate = addDays(start, week * 7);
    const weekNumber = cellText(String(isoWeek(weekDate)).padStart(2, "0"));
    weekNumber.setAttribute("data-slot", "calendar-week-number");
    row.appendChild(weekNumber);
  }
  for (let day = 0; day < 7; day += 1) {
    row.appendChild(renderDay(root, month, addDays(start, week * 7 + day)));
  }
  return row;
}

function renderDay(root, month, date) {
  const value = formatDate(date);
  const state = dayState(root, value);
  const outside = date.getUTCMonth() !== month.getUTCMonth();
  const cell = document.createElement("div");
  cell.setAttribute("role", "gridcell");
  cell.setAttribute("data-day", value);
  cell.className = "relative p-0 text-center text-sm focus-within:relative focus-within:z-20";
  setBool(cell, "data-outside", outside);
  setBool(cell, "data-selected", state.selected);
  setBool(cell, "data-range-start", state.rangeStart);
  setBool(cell, "data-range-end", state.rangeEnd);
  setBool(cell, "data-range-middle", state.rangeMiddle);
  const button = document.createElement("button");
  button.type = "button";
  button.setAttribute("data-slot", "calendar-day-button");
  button.setAttribute("data-date", value);
  button.setAttribute("data-day", date.toLocaleDateString(root.getAttribute("data-locale") || undefined, { timeZone: "UTC" }));
  button.setAttribute("aria-label", date.toLocaleDateString(root.getAttribute("data-locale") || undefined, { dateStyle: "long", timeZone: "UTC" }));
  button.setAttribute("aria-selected", state.selected ? "true" : "false");
  button.className = "inline-flex size-(--cell-size) items-center justify-center rounded-md text-sm font-normal tabular-nums hover:bg-accent hover:text-accent-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-40 aria-selected:bg-primary aria-selected:text-primary-foreground data-[today=true]:bg-accent data-[today=true]:text-accent-foreground data-[outside=true]:text-muted-foreground data-[outside=true]:opacity-50 data-[range-middle=true]:bg-accent data-[range-middle=true]:text-accent-foreground data-[range-start=true]:rounded-s-(--cell-radius) data-[range-end=true]:rounded-e-(--cell-radius)";
  button.textContent = String(date.getUTCDate());
  setBool(button, "data-selected", state.selected);
  setBool(button, "data-range-start", state.rangeStart);
  setBool(button, "data-range-end", state.rangeEnd);
  setBool(button, "data-range-middle", state.rangeMiddle);
  setBool(button, "data-outside", outside);
  setBool(button, "data-today", value === formatDate(new Date()));
  if (isDisabled(root, value)) {
    button.disabled = true;
    button.setAttribute("aria-disabled", "true");
    button.setAttribute("data-disabled", "true");
  }
  if (outside && root.getAttribute("data-show-outside-days") === "false") button.hidden = true;
  cell.appendChild(button);
  return cell;
}

function moveDayFocus(root, current, offset) {
  const date = parseDate(current.getAttribute("data-date"));
  const nextDate = addDays(date, offset);
  const month = root.getAttribute("data-current-month");
  const nextMonth = formatMonth(nextDate);
  if (nextMonth !== month) {
    root.setAttribute("data-current-month", nextMonth);
    renderCalendar(root);
  }
  root.querySelector(`${dayButton}[data-date="${formatDate(nextDate)}"]`)?.focus();
}

function visibleDays(root) {
  return Array.from(root.querySelectorAll(dayButton)).filter((button) => !button.hidden && !button.disabled);
}

function rowClass(root) {
  return `grid gap-1 ${root.getAttribute("data-show-week-number") === "true" ? "grid-cols-8" : "grid-cols-7"}`;
}

function cellText(value, role, label) {
  const el = document.createElement("div");
  if (role) el.setAttribute("role", role);
  if (label) el.setAttribute("aria-label", label);
  el.className = "flex size-(--cell-size) items-center justify-center text-xs text-muted-foreground";
  el.textContent = value;
  return el;
}

function dayState(root, value) {
  const selected = splitValues(root.getAttribute("data-selected"));
  const [start, end] = parseRange(root.getAttribute("data-range") || (root.getAttribute("data-mode") === "range" ? root.getAttribute("data-selected") : ""));
  const rangeStart = value === start;
  const rangeEnd = value === end;
  const rangeMiddle = Boolean(start && end && value > start && value < end);
  return { selected: selected.includes(value) || rangeStart || rangeEnd, rangeStart, rangeEnd, rangeMiddle };
}

function syncExistingDayButtons(root) {
  root.querySelectorAll(dayButton).forEach((button) => {
    const value = button.getAttribute("data-date");
    if (!value) return;
    const state = dayState(root, value);
    button.setAttribute("aria-selected", state.selected ? "true" : "false");
    setBool(button, "data-selected", state.selected);
    setBool(button, "data-range-start", state.rangeStart);
    setBool(button, "data-range-end", state.rangeEnd);
    setBool(button, "data-range-middle", state.rangeMiddle);
    const cell = button.closest('[role="gridcell"]');
    if (cell) {
      setBool(cell, "data-selected", state.selected);
      setBool(cell, "data-range-start", state.rangeStart);
      setBool(cell, "data-range-end", state.rangeEnd);
      setBool(cell, "data-range-middle", state.rangeMiddle);
    }
  });
}

function syncFormValue(root) {
  const value = calendarValue(root);
  const picker = root.closest('[data-slot="date-picker"]');
  if (picker) {
    picker.setAttribute("data-value", value.value || value.range || "");
    picker.querySelector('input[type="hidden"]')?.setAttribute("value", value.value || value.range || "");
  }
}

function calendarValue(root) {
  if ((root.getAttribute("data-mode") || "single") === "range") return { range: root.getAttribute("data-range") || "" };
  return { value: root.getAttribute("data-selected") || "" };
}

function isDisabled(root, value) {
  return splitValues(root.getAttribute("data-disabled-dates")).includes(value);
}

function splitValues(value) {
  return (value || "").split(/[,\s]+/).map((part) => part.trim()).filter(Boolean);
}

function parseRange(value) {
  const parts = (value || "").split(/[/:,]/).map((part) => part.trim());
  return [parts[0] || "", parts[1] || ""];
}

function parseMonth(value) {
  const [year, month] = (value || new Date().toISOString().slice(0, 7)).split("-").map(Number);
  return new Date(Date.UTC(year, month - 1, 1));
}

function parseDate(value) {
  const [year, month, day] = value.split("-").map(Number);
  return new Date(Date.UTC(year, month - 1, day));
}

function addDays(date, days) {
  const next = new Date(date);
  next.setUTCDate(next.getUTCDate() + days);
  return next;
}

function formatMonth(date) {
  return `${date.getUTCFullYear()}-${pad(date.getUTCMonth() + 1)}`;
}

function formatDate(date) {
  return `${date.getUTCFullYear()}-${pad(date.getUTCMonth() + 1)}-${pad(date.getUTCDate())}`;
}

function pad(value) {
  return String(value).padStart(2, "0");
}

function setBool(el, attr, value) {
  if (value) el.setAttribute(attr, "true");
  else el.removeAttribute(attr);
}

function isoWeek(date) {
  const d = new Date(Date.UTC(date.getUTCFullYear(), date.getUTCMonth(), date.getUTCDate()));
  d.setUTCDate(d.getUTCDate() + 4 - (d.getUTCDay() || 7));
  const yearStart = new Date(Date.UTC(d.getUTCFullYear(), 0, 1));
  return Math.ceil(((d - yearStart) / 86400000 + 1) / 7);
}
