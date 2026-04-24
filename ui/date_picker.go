package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type DatePickerProps struct {
	DOMProps
	Mode        string
	Name        string
	Value       string
	Range       string
	Placeholder string
	Open        bool
	DefaultOpen bool
}

func DatePicker(props DatePickerProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "date-picker", "grid gap-2")
		if props.Mode != "" {
			attrs["data-mode"] = props.Mode
		}
		if props.Name != "" {
			attrs["data-name"] = props.Name
		}
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.Range != "" {
			attrs["data-range"] = props.Range
		}
		if props.Placeholder != "" {
			attrs["data-placeholder"] = props.Placeholder
		}
		if props.Open {
			attrs["data-open"] = "true"
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

type DateRangePickerProps struct {
	DatePickerProps
}

func DateRangePicker(props DateRangePickerProps) templ.Component {
	return DatePicker(props.DatePickerProps)
}
