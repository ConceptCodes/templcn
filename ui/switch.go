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
		size := "default"
		if props.Size == "sm" {
			size = "sm"
		}
		attrs := attrsFromDOMProps(props.DOMProps, "switch", "peer group/switch inline-flex appearance-none shrink-0 items-center rounded-full border border-transparent shadow-xs transition-all outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:cursor-not-allowed disabled:opacity-50 data-[size=default]:h-[1.15rem] data-[size=default]:w-8 data-[size=sm]:h-3.5 data-[size=sm]:w-6 data-[state=checked]:bg-primary data-[state=unchecked]:bg-input checked:bg-primary dark:data-[state=unchecked]:bg-input/80 after:pointer-events-none after:block after:rounded-full after:bg-background after:ring-0 after:transition-transform after:content-[''] data-[size=default]:after:size-4 data-[size=sm]:after:size-3 data-[state=checked]:after:translate-x-[calc(100%-2px)] data-[state=unchecked]:after:translate-x-0 dark:data-[state=checked]:after:bg-primary-foreground dark:data-[state=unchecked]:after:bg-foreground")
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
