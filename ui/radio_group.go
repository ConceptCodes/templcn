package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type RadioGroupProps struct {
	DOMProps
	Name         string
	Value        string
	DefaultValue string
	Disabled     bool
	Required     bool
}

func RadioGroup(props RadioGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "radio-group", "grid gap-2")
		attrs["role"] = "radiogroup"
		if props.Name != "" {
			attrs["data-name"] = props.Name
		}
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.DefaultValue != "" {
			attrs["data-default-value"] = props.DefaultValue
		}
		if props.Disabled {
			attrs["data-disabled"] = "true"
		}
		if props.Required {
			attrs["aria-required"] = "true"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

type RadioGroupItemProps struct {
	DOMProps
	Name     string
	Value    string
	Disabled bool
	Label    string
}

func RadioGroupItem(props RadioGroupItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "flex items-center gap-2 text-sm"
		attrs := attrsFromDOMProps(props.DOMProps, "radio-group-item", className)
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			inputAttrs := templ.Attributes{
				"type":  "radio",
				"class": "h-4 w-4 border border-input text-primary focus-visible:ring-[3px] focus-visible:ring-ring/50",
				"value": props.Value,
			}
			if props.Name != "" {
				inputAttrs["name"] = props.Name
			}
			if props.Disabled {
				inputAttrs["disabled"] = true
			}
			if err := renderVoidElement(ctx, w, "input", inputAttrs); err != nil {
				return err
			}
			if props.Label != "" {
				_, err := io.WriteString(w, templ.EscapeString(props.Label))
				return err
			}
			return renderChildren(ctx, w, templ.GetChildren(ctx))
		})
		return renderElement(ctx, w, "label", attrs, children)
	})
}
