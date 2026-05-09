package ui

import (
	"strings"
	"testing"
)

func TestCalendarRendersMonthGridAndSelectedDate(t *testing.T) {
	html := mustRender(t, Calendar(CalendarProps{DefaultMonth: "2026-05", Selected: "2026-05-07"}))
	if attrValue(html, "data-current-month") != "2026-05" {
		t.Fatalf("calendar current month = %q, want 2026-05", attrValue(html, "data-current-month"))
	}
	if !strings.Contains(html, `data-slot="calendar-grid"`) || !strings.Contains(html, `role="grid"`) {
		t.Fatalf("calendar should render a grid, got %s", html)
	}
	if !strings.Contains(html, `data-date="2026-05-07"`) || !strings.Contains(html, `aria-selected="true"`) {
		t.Fatalf("calendar selected day attrs missing, got %s", html)
	}
	if !strings.Contains(html, `data-show-outside-days="true"`) {
		t.Fatalf("calendar should show outside days by default, got %s", html)
	}
}

func TestCalendarSupportsDropdownCaptionWeekNumbersRangeAndDisabledDays(t *testing.T) {
	html := mustRender(t, Calendar(CalendarProps{
		DefaultMonth:   "2026-05",
		Mode:           "range",
		Range:          "2026-05-07/2026-05-10",
		CaptionLayout:  "dropdown",
		ShowWeekNumber: true,
		DisabledDates:  "2026-05-09",
		FromYear:       2025,
		ToYear:         2027,
	}))
	for _, want := range []string{
		`data-slot="calendar-month-select"`,
		`data-slot="calendar-year-select"`,
		`data-slot="calendar-week-number"`,
		`data-range-start="true"`,
		`data-range-middle="true"`,
		`data-range-end="true"`,
		`data-disabled="true"`,
		`aria-disabled="true"`,
	} {
		if !strings.Contains(html, want) {
			t.Fatalf("calendar missing %s, got %s", want, html)
		}
	}
}

func TestCalendarCanHideOutsideDays(t *testing.T) {
	html := mustRender(t, Calendar(CalendarProps{DefaultMonth: "2026-05", HideOutsideDays: true}))
	if !strings.Contains(html, `data-show-outside-days="false"`) {
		t.Fatalf("calendar should mark outside days hidden mode, got %s", html)
	}
	if !strings.Contains(html, `hidden`) {
		t.Fatalf("calendar should hide outside day buttons, got %s", html)
	}
}

func TestCalendarDayButtonDefaultsToButtonType(t *testing.T) {
	html := mustRender(t, CalendarDayButton(DOMProps{}))
	if attrValue(html, "type") != "button" {
		t.Fatalf("CalendarDayButton type = %q, want button", attrValue(html, "type"))
	}
}
