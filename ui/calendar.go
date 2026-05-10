package ui

import (
	"context"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/a-h/templ"
)

type CalendarProps struct {
	DOMProps
	Mode            string
	Selected        string
	Range           string
	Month           string
	DefaultMonth    string
	ShowOutsideDays bool
	ShowWeekNumber  bool
	CaptionLayout   string
	ButtonVariant   string
	Locale          string
	DisabledDates   string
	FromYear        int
	ToYear          int
	NumberOfMonths  int
	Dir             string
	HideOutsideDays bool
	TimeZone        string
}

func Calendar(props CalendarProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "calendar", "group/calendar relative bg-background p-3 [--cell-size:--spacing(8)] [[data-slot=card-content]_&]:bg-transparent [[data-slot=popover-content]_&]:bg-transparent")
		mode := props.Mode
		if mode == "" {
			mode = "single"
		}
		month := calendarMonth(props)
		attrs["data-mode"] = mode
		if props.Selected != "" {
			attrs["data-selected"] = props.Selected
		}
		if props.Range != "" {
			attrs["data-range"] = props.Range
		}
		if props.Month != "" {
			attrs["data-month"] = props.Month
		}
		if props.DefaultMonth != "" {
			attrs["data-default-month"] = props.DefaultMonth
		}
		showOutsideDays := props.ShowOutsideDays || !props.HideOutsideDays
		attrs["data-show-outside-days"] = strconv.FormatBool(showOutsideDays)
		if props.ShowWeekNumber {
			attrs["data-show-week-number"] = "true"
		}
		if props.CaptionLayout != "" {
			attrs["data-caption-layout"] = props.CaptionLayout
		}
		if props.ButtonVariant != "" {
			attrs["data-button-variant"] = props.ButtonVariant
		}
		if props.Locale != "" {
			attrs["data-locale"] = props.Locale
		}
		if props.DisabledDates != "" {
			attrs["data-disabled-dates"] = props.DisabledDates
		}
		if props.FromYear != 0 {
			attrs["data-from-year"] = strconv.Itoa(props.FromYear)
		}
		if props.ToYear != 0 {
			attrs["data-to-year"] = strconv.Itoa(props.ToYear)
		}
		if props.NumberOfMonths > 1 {
			attrs["data-number-of-months"] = strconv.Itoa(props.NumberOfMonths)
		}
		if props.Dir != "" {
			attrs["dir"] = props.Dir
		}
		if props.TimeZone != "" {
			attrs["data-time-zone"] = props.TimeZone
		}
		attrs["data-current-month"] = month.Format("2006-01")
		ctx, customChildren := childrenFromContext(ctx)
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := calendarGrid(month, props).Render(ctx, w); err != nil {
				return err
			}
			return renderChildren(ctx, w, customChildren)
		})
		return renderElement(ctx, w, "div", attrs, children)
	})
}

func CalendarDayButton(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "calendar-day-button", calendarDayButtonClass())
		if _, ok := attrs["type"]; !ok {
			attrs["type"] = "button"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func calendarMonth(props CalendarProps) time.Time {
	for _, value := range []string{props.Month, props.DefaultMonth, props.Selected} {
		if value == "" {
			continue
		}
		if parsed, err := time.Parse("2006-01-02", value); err == nil {
			return time.Date(parsed.Year(), parsed.Month(), 1, 0, 0, 0, 0, time.UTC)
		}
		if parsed, err := time.Parse("2006-01", value); err == nil {
			return time.Date(parsed.Year(), parsed.Month(), 1, 0, 0, 0, 0, time.UTC)
		}
	}
	now := time.Now().UTC()
	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
}

func calendarGrid(month time.Time, props CalendarProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		if err := renderElement(ctx, w, "div", attrsFromDOMProps(DOMProps{}, "calendar-nav", "absolute inset-x-0 top-0 flex w-full items-center justify-between gap-1"), templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := renderTextElement(ctx, w, "button", templ.Attributes{"type": "button", "data-slot": "calendar-prev", "aria-label": "Previous month", "class": calendarNavButtonClass()}, "‹"); err != nil {
				return err
			}
			if props.CaptionLayout == "dropdown" {
				if err := calendarDropdownCaption(ctx, w, month, props); err != nil {
					return err
				}
			} else if err := renderTextElement(ctx, w, "div", templ.Attributes{"data-slot": "calendar-caption", "aria-live": "polite", "class": "text-sm font-medium"}, month.Format("January 2006")); err != nil {
				return err
			}
			return renderTextElement(ctx, w, "button", templ.Attributes{"type": "button", "data-slot": "calendar-next", "aria-label": "Next month", "class": calendarNavButtonClass()}, "›")
		})); err != nil {
			return err
		}

		return renderElement(ctx, w, "div", templ.Attributes{"data-slot": "calendar-months", "class": "relative flex flex-col gap-4 md:flex-row"}, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			months := props.NumberOfMonths
			if months < 1 {
				months = 1
			}
			for i := 0; i < months; i++ {
				if err := calendarMonthGrid(month.AddDate(0, i, 0), props).Render(ctx, w); err != nil {
					return err
				}
			}
			return nil
		}))
	})
}

func calendarMonthGrid(month time.Time, props CalendarProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		columns := "grid-cols-7"
		if props.ShowWeekNumber {
			columns = "grid-cols-8"
		}
		return renderElement(ctx, w, "div", templ.Attributes{"data-slot": "calendar-month", "class": "flex w-full flex-col gap-4", "data-month": month.Format("2006-01")}, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := renderTextElement(ctx, w, "div", templ.Attributes{"data-slot": "calendar-month-caption", "class": "sr-only"}, month.Format("January 2006")); err != nil {
				return err
			}
			if err := renderElement(ctx, w, "div", templ.Attributes{"data-slot": "calendar-grid", "role": "grid", "aria-label": month.Format("January 2006"), "class": "w-full border-collapse"}, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				return calendarRows(ctx, w, month, props, columns)
			})); err != nil {
				return err
			}
			return nil
		}))
	})
}

func calendarRows(ctx context.Context, w io.Writer, month time.Time, props CalendarProps, columns string) error {
	weekdays := []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"}
	if err := renderElement(ctx, w, "div", templ.Attributes{"role": "row", "class": cn("mt-2 flex w-full", columns)}, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		if props.ShowWeekNumber {
			if err := renderTextElement(ctx, w, "div", templ.Attributes{"role": "columnheader", "aria-label": "Week number", "class": "w-(--cell-size) select-none"}, ""); err != nil {
				return err
			}
		}
		for _, weekday := range weekdays {
			if err := renderTextElement(ctx, w, "div", templ.Attributes{"role": "columnheader", "class": "flex-1 rounded-md text-[0.8rem] font-normal text-muted-foreground select-none"}, weekday); err != nil {
				return err
			}
		}
		return nil
	})); err != nil {
		return err
	}
	start := month.AddDate(0, 0, -int(month.Weekday()))
	for week := 0; week < 6; week++ {
		if err := renderElement(ctx, w, "div", templ.Attributes{"role": "row", "class": cn("mt-2 flex w-full", columns)}, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if props.ShowWeekNumber {
				if err := renderTextElement(ctx, w, "div", templ.Attributes{"data-slot": "calendar-week-number", "class": "flex size-(--cell-size) items-center justify-center text-[0.8rem] text-muted-foreground select-none"}, isoWeekLabel(start.AddDate(0, 0, week*7))); err != nil {
					return err
				}
			}
			for day := 0; day < 7; day++ {
				date := start.AddDate(0, 0, week*7+day)
				if err := renderCalendarDay(ctx, w, date, month, props); err != nil {
					return err
				}
			}
			return nil
		})); err != nil {
			return err
		}
	}
	return nil
}

func renderCalendarDay(ctx context.Context, w io.Writer, date time.Time, month time.Time, props CalendarProps) error {
	dateValue := date.Format("2006-01-02")
	outside := date.Month() != month.Month()
	selected, rangeStart, rangeEnd, rangeMiddle := calendarSelectionState(dateValue, props)
	disabled := calendarDateDisabled(dateValue, props.DisabledDates)
	today := dateValue == time.Now().UTC().Format("2006-01-02")
	attrs := templ.Attributes{"role": "gridcell", "data-day": dateValue, "class": "group/day relative aspect-square h-full w-full p-0 text-center select-none [&:last-child[data-selected=true]_button]:rounded-r-md [&:first-child[data-selected=true]_button]:rounded-l-md"}
	if outside {
		attrs["data-outside"] = "true"
	}
	if selected {
		attrs["data-selected"] = "true"
	}
	if rangeStart {
		attrs["data-range-start"] = "true"
	}
	if rangeEnd {
		attrs["data-range-end"] = "true"
	}
	if rangeMiddle {
		attrs["data-range-middle"] = "true"
	}
	if err := renderElement(ctx, w, "div", attrs, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		buttonAttrs := attrsFromDOMProps(DOMProps{}, "calendar-day-button", calendarDayButtonClass())
		buttonAttrs["type"] = "button"
		buttonAttrs["data-date"] = dateValue
		buttonAttrs["data-day"] = date.Format("1/2/2006")
		buttonAttrs["aria-label"] = date.Format("January 2, 2006")
		buttonAttrs["aria-selected"] = map[bool]string{true: "true", false: "false"}[selected]
		if selected {
			buttonAttrs["data-selected"] = "true"
		}
		if selected && !rangeStart && !rangeEnd && !rangeMiddle {
			buttonAttrs["data-selected-single"] = "true"
		}
		if today {
			buttonAttrs["data-today"] = "true"
		}
		if outside {
			buttonAttrs["data-outside"] = "true"
		}
		if rangeStart {
			buttonAttrs["data-range-start"] = "true"
		}
		if rangeEnd {
			buttonAttrs["data-range-end"] = "true"
		}
		if rangeMiddle {
			buttonAttrs["data-range-middle"] = "true"
		}
		if disabled {
			buttonAttrs["disabled"] = true
			buttonAttrs["aria-disabled"] = "true"
			buttonAttrs["data-disabled"] = "true"
		}
		if outside && props.HideOutsideDays {
			buttonAttrs["hidden"] = true
		}
		return renderTextElement(ctx, w, "button", buttonAttrs, strconv.Itoa(date.Day()))
	})); err != nil {
		return err
	}
	return nil
}

func calendarDropdownCaption(ctx context.Context, w io.Writer, month time.Time, props CalendarProps) error {
	return renderElement(ctx, w, "div", templ.Attributes{"data-slot": "calendar-caption", "class": "flex items-center gap-1", "aria-live": "polite"}, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		if err := renderElement(ctx, w, "select", templ.Attributes{"data-slot": "calendar-month-select", "aria-label": "Month", "class": "h-8 rounded-md border bg-background px-2 text-sm"}, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			for i := time.January; i <= time.December; i++ {
				attrs := templ.Attributes{"value": strconv.Itoa(int(i))}
				if i == month.Month() {
					attrs["selected"] = true
				}
				if err := renderTextElement(ctx, w, "option", attrs, time.Date(2000, i, 1, 0, 0, 0, 0, time.UTC).Format("Jan")); err != nil {
					return err
				}
			}
			return nil
		})); err != nil {
			return err
		}
		from, to := calendarYearBounds(month, props)
		return renderElement(ctx, w, "select", templ.Attributes{"data-slot": "calendar-year-select", "aria-label": "Year", "class": "h-8 rounded-md border bg-background px-2 text-sm"}, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			for year := from; year <= to; year++ {
				attrs := templ.Attributes{"value": strconv.Itoa(year)}
				if year == month.Year() {
					attrs["selected"] = true
				}
				if err := renderTextElement(ctx, w, "option", attrs, strconv.Itoa(year)); err != nil {
					return err
				}
			}
			return nil
		}))
	}))
}

func calendarSelectionState(dateValue string, props CalendarProps) (selected bool, rangeStart bool, rangeEnd bool, rangeMiddle bool) {
	selectedDates := splitCalendarValues(props.Selected)
	for _, value := range selectedDates {
		if value == dateValue {
			selected = true
			break
		}
	}
	rangeValue := props.Range
	if rangeValue == "" && props.Mode == "range" {
		rangeValue = props.Selected
	}
	start, end, ok := parseCalendarRange(rangeValue)
	if !ok {
		return selected, false, false, false
	}
	rangeStart = dateValue == start
	rangeEnd = dateValue == end
	rangeMiddle = dateValue > start && dateValue < end
	selected = selected || rangeStart || rangeEnd
	return selected, rangeStart, rangeEnd, rangeMiddle
}

func splitCalendarValues(value string) []string {
	if value == "" {
		return nil
	}
	parts := strings.FieldsFunc(value, func(r rune) bool {
		return r == ',' || r == ' '
	})
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		if trimmed := strings.TrimSpace(part); trimmed != "" {
			out = append(out, trimmed)
		}
	}
	return out
}

func parseCalendarRange(value string) (string, string, bool) {
	if value == "" {
		return "", "", false
	}
	parts := strings.FieldsFunc(value, func(r rune) bool {
		return r == '/' || r == ':' || r == ','
	})
	if len(parts) < 2 {
		return "", "", false
	}
	start := strings.TrimSpace(parts[0])
	end := strings.TrimSpace(parts[1])
	if start == "" || end == "" {
		return "", "", false
	}
	if start > end {
		start, end = end, start
	}
	return start, end, true
}

func calendarDateDisabled(dateValue string, disabledDates string) bool {
	for _, value := range splitCalendarValues(disabledDates) {
		if value == dateValue {
			return true
		}
	}
	return false
}

func calendarYearBounds(month time.Time, props CalendarProps) (int, int) {
	from := props.FromYear
	to := props.ToYear
	if from == 0 {
		from = month.Year() - 50
	}
	if to == 0 {
		to = month.Year() + 50
	}
	if from > to {
		from, to = to, from
	}
	return from, to
}

func isoWeekLabel(date time.Time) string {
	_, week := date.ISOWeek()
	return strconv.FormatInt(int64(week), 10)
}

func calendarNavButtonClass() string {
	return buttonClasses(ButtonVariantGhost, ButtonSizeDefault, "size-(--cell-size) p-0 select-none aria-disabled:opacity-50")
}

func calendarDayButtonClass() string {
	return buttonClasses(ButtonVariantGhost, ButtonSizeIcon, "flex aspect-square size-auto w-full min-w-(--cell-size) flex-col gap-1 leading-none font-normal group-data-[focused=true]/day:relative group-data-[focused=true]/day:z-10 group-data-[focused=true]/day:border-ring group-data-[focused=true]/day:ring-[3px] group-data-[focused=true]/day:ring-ring/50 data-[range-end=true]:rounded-md data-[range-end=true]:rounded-r-md data-[range-end=true]:bg-primary data-[range-end=true]:text-primary-foreground data-[range-middle=true]:rounded-none data-[range-middle=true]:bg-accent data-[range-middle=true]:text-accent-foreground data-[range-start=true]:rounded-md data-[range-start=true]:rounded-l-md data-[range-start=true]:bg-primary data-[range-start=true]:text-primary-foreground data-[selected-single=true]:bg-primary data-[selected-single=true]:text-primary-foreground dark:hover:text-accent-foreground [&>span]:text-xs [&>span]:opacity-70")
}
