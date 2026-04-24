package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type InputProps struct {
	DOMProps
	Type        string
	Name        string
	Value       string
	Placeholder string
	Disabled    bool
	Required    bool
	Invalid     bool
}

func Input(props InputProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "input", "flex h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-base shadow-xs transition-[color,box-shadow] outline-none file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:cursor-not-allowed disabled:opacity-50 md:text-sm")
		if props.Type == "" {
			props.Type = "text"
		}
		attrs["type"] = props.Type
		if props.Name != "" {
			attrs["name"] = props.Name
		}
		if props.Value != "" {
			attrs["value"] = props.Value
		}
		if props.Placeholder != "" {
			attrs["placeholder"] = props.Placeholder
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
		return renderVoidElement(ctx, w, "input", attrs)
	})
}
