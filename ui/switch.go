package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type SwitchProps struct {
	DOMProps
	Name           string
	Value          string
	Checked        bool
	DefaultChecked bool
	Disabled       bool
	Required       bool
	Size           string
}

func Switch(props SwitchProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		sizeClass := "h-5 w-9"
		size := "default"
		if props.Size == "sm" {
			sizeClass = "h-4 w-7"
			size = "sm"
		}
		attrs := attrsFromDOMProps(props.DOMProps, "switch", cn("peer appearance-none shrink-0 rounded-full border border-transparent bg-input shadow-xs outline-none transition-all focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:cursor-not-allowed disabled:opacity-50 checked:bg-primary", sizeClass, "after:pointer-events-none after:block after:rounded-full after:bg-background after:shadow-sm after:transition-transform after:content-[''] checked:after:translate-x-4 rtl:checked:after:-translate-x-4"))
		attrs["type"] = "checkbox"
		attrs["role"] = "switch"
		attrs["data-size"] = size
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
		return renderVoidElement(ctx, w, "input", attrs)
	})
}
