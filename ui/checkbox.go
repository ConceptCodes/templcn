package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type CheckboxProps struct {
	DOMProps
	Name           string
	Value          string
	Checked        bool
	DefaultChecked bool
	Disabled       bool
	Required       bool
	Invalid        bool
}

func Checkbox(props CheckboxProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "checkbox", "peer size-4 shrink-0 rounded-[4px] border border-input shadow-xs transition-shadow outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:cursor-not-allowed disabled:opacity-50 aria-invalid:border-destructive aria-invalid:ring-destructive/20 checked:border-primary checked:bg-primary checked:text-primary-foreground data-[state=checked]:border-primary data-[state=checked]:bg-primary data-[state=checked]:text-primary-foreground dark:bg-input/30 dark:aria-invalid:ring-destructive/40 dark:data-[state=checked]:bg-primary")
		attrs["type"] = "checkbox"
		checked := props.Checked || props.DefaultChecked
		attrs["data-state"] = map[bool]string{true: "checked", false: "unchecked"}[checked]
		attrs["aria-checked"] = map[bool]string{true: "true", false: "false"}[checked]
		if props.Name != "" {
			attrs["name"] = props.Name
		}
		if props.Value != "" {
			attrs["value"] = props.Value
		}
		if props.Checked {
			attrs["checked"] = true
		}
		if props.DefaultChecked {
			attrs["checked"] = true
			attrs["data-default-checked"] = "true"
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
