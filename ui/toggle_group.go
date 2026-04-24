package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type ToggleGroupProps struct {
	DOMProps
	Type         string
	Value        []string
	DefaultValue []string
	Variant      string
	Size         string
	Spacing      string
}

type ToggleGroupItemProps struct {
	DOMProps
	Value    string
	Disabled bool
}

func ToggleGroup(props ToggleGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "toggle-group", "inline-flex gap-1")
		if props.Type != "" {
			attrs["data-type"] = props.Type
		}
		if len(props.Value) > 0 {
			attrs["data-value"] = props.Value
		}
		if len(props.DefaultValue) > 0 {
			attrs["data-default-value"] = props.DefaultValue
		}
		if props.Variant != "" {
			attrs["data-variant"] = props.Variant
		}
		if props.Size != "" {
			attrs["data-size"] = props.Size
		}
		if props.Spacing != "" {
			attrs["data-spacing"] = props.Spacing
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func ToggleGroupItem(props ToggleGroupItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "toggle-group-item", "rounded-md")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.Disabled {
			attrs["disabled"] = true
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
