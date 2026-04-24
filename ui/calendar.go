package ui

import (
	"context"
	"io"

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
}

func Calendar(props CalendarProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "calendar", "grid gap-2 rounded-lg border p-3")
		if props.Mode != "" {
			attrs["data-mode"] = props.Mode
		}
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
		if props.ShowOutsideDays {
			attrs["data-show-outside-days"] = "true"
		}
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
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func CalendarDayButton(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "calendar-day-button", "inline-flex size-8 items-center justify-center rounded-md text-sm hover:bg-accent hover:text-accent-foreground"), templ.GetChildren(ctx))
	})
}
