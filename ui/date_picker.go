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
		open := props.Open || props.DefaultOpen
		attrs := attrsFromDOMProps(props.DOMProps, "date-picker", "grid gap-2")
		attrs["data-state"] = openState(open)
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
		if open {
			attrs["data-open"] = "true"
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if props.Name != "" {
				inputAttrs := templ.Attributes{"type": "hidden", "name": props.Name, "value": props.Value}
				if err := renderVoidElement(ctx, w, "input", inputAttrs); err != nil {
					return err
				}
			}
			return renderChildren(ctx, w, templ.GetChildren(ctx))
		})
		return renderElement(ctx, w, "div", attrs, children)
	})
}

type DateRangePickerProps struct {
	DatePickerProps
}

func DateRangePicker(props DateRangePickerProps) templ.Component {
	return DatePicker(props.DatePickerProps)
}
