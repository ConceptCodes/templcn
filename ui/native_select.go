package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type NativeSelectProps struct {
	DOMProps
	Name     string
	Value    string
	Size     string
	Disabled bool
	Required bool
	Invalid  bool
}

func NativeSelect(props NativeSelectProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		sizeClass := "h-9"
		if props.Size == "sm" {
			sizeClass = "h-8"
		}
		attrs := attrsFromDOMProps(props.DOMProps, "native-select", cn("flex w-full appearance-none rounded-md border border-input bg-background px-3 py-2 text-sm shadow-xs outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:cursor-not-allowed disabled:opacity-50", sizeClass))
		if props.Name != "" {
			attrs["name"] = props.Name
		}
		if props.Value != "" {
			attrs["value"] = props.Value
		}
		if props.Disabled {
			attrs["disabled"] = true
		}
		if props.Required {
			attrs["required"] = true
		}
		if props.Invalid {
			attrs["aria-invalid"] = "true"
		}
		return renderElement(ctx, w, "select", attrs, templ.GetChildren(ctx))
	})
}

func NativeSelectOption(props DOMProps, value, text string, selected bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "native-select-option", "")
		attrs["value"] = value
		if selected {
			attrs["selected"] = true
		}
		return renderTextElement(ctx, w, "option", attrs, text)
	})
}

func NativeSelectOptGroup(props DOMProps, label string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "native-select-optgroup", "")
		if label != "" {
			attrs["label"] = label
		}
		return renderElement(ctx, w, "optgroup", attrs, templ.GetChildren(ctx))
	})
}
