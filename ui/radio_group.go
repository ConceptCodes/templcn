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

type radioGroupRenderState struct {
	name     string
	value    string
	disabled bool
}

type radioGroupRenderStateKey struct{}

func RadioGroup(props RadioGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		value := props.Value
		if value == "" {
			value = props.DefaultValue
		}
		attrs := attrsFromDOMProps(props.DOMProps, "radio-group", "grid gap-3")
		attrs["role"] = "radiogroup"
		if props.Name != "" {
			attrs["data-name"] = props.Name
		}
		if value != "" {
			attrs["data-value"] = value
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
		ctx = context.WithValue(ctx, radioGroupRenderStateKey{}, radioGroupRenderState{name: props.Name, value: value, disabled: props.Disabled})
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
		state, _ := ctx.Value(radioGroupRenderStateKey{}).(radioGroupRenderState)
		name := props.Name
		if name == "" {
			name = state.name
		}
		disabled := props.Disabled || state.disabled
		checked := props.Value != "" && state.value == props.Value
		attrs["data-state"] = map[bool]string{true: "checked", false: "unchecked"}[checked]
		attrs["aria-checked"] = map[bool]string{true: "true", false: "false"}[checked]
		if disabled {
			attrs["data-disabled"] = "true"
		}
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			inputAttrs := templ.Attributes{
				"type":  "radio",
				"class": "aspect-square size-4 shrink-0 rounded-full border border-input text-primary shadow-xs transition-[color,box-shadow] outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:cursor-not-allowed disabled:opacity-50 aria-invalid:border-destructive aria-invalid:ring-destructive/20 dark:bg-input/30 dark:aria-invalid:ring-destructive/40",
				"value": props.Value,
			}
			if name != "" {
				inputAttrs["name"] = name
			}
			if checked {
				inputAttrs["checked"] = true
			}
			if disabled {
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
