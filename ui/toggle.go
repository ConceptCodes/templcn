package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type ToggleProps struct {
	DOMProps
	Variant        string
	Size           string
	Pressed        bool
	DefaultPressed bool
	Disabled       bool
	Value          string
}

func Toggle(props ToggleProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "inline-flex items-center justify-center rounded-md text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 data-[state=on]:bg-accent data-[state=on]:text-accent-foreground"
		if props.Variant == "outline" {
			className = cn(className, "border border-input bg-transparent shadow-xs")
		}
		switch props.Size {
		case "sm":
			className = cn(className, "h-8 px-3")
		case "lg":
			className = cn(className, "h-10 px-4")
		default:
			className = cn(className, "h-9 px-3")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "toggle", className)
		pressed := props.Pressed || props.DefaultPressed
		attrs["type"] = "button"
		attrs["aria-pressed"] = "false"
		attrs["data-state"] = "off"
		if pressed {
			attrs["aria-pressed"] = "true"
			attrs["data-state"] = "on"
		}
		if props.DefaultPressed {
			attrs["data-default-pressed"] = "true"
		}
		if props.Disabled {
			attrs["disabled"] = true
		}
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
