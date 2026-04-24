package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type SelectProps struct {
	DOMProps
	Name         string
	Value        string
	DefaultValue string
	Open         bool
	DefaultOpen  bool
	Disabled     bool
	Required     bool
}

type SelectTriggerProps struct {
	DOMProps
	Size string
}

type SelectContentProps struct {
	DOMProps
	Position string
	Align    string
}

func Select(props SelectProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "select", "relative")
		if props.Name != "" {
			attrs["data-name"] = props.Name
		}
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.DefaultValue != "" {
			attrs["data-default-value"] = props.DefaultValue
		}
		if props.Open {
			attrs["data-open"] = "true"
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		if props.Disabled {
			attrs["data-disabled"] = "true"
		}
		if props.Required {
			attrs["data-required"] = "true"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func SelectGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "select-group", "grid gap-1"), templ.GetChildren(ctx))
	})
}
func SelectValue(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "select-value", ""), templ.GetChildren(ctx))
	})
}
func SelectTrigger(props SelectTriggerProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "flex h-9 w-full items-center justify-between rounded-md border border-input bg-background px-3 py-2 text-sm shadow-xs outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:cursor-not-allowed disabled:opacity-50"
		if props.Size == "sm" {
			className = cn(className, "h-8")
		}
		return renderElement(ctx, w, "button", attrsFromDOMProps(props.DOMProps, "select-trigger", className), templ.GetChildren(ctx))
	})
}

func SelectContent(props SelectContentProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "select-content", "z-50 min-w-32 rounded-md border bg-popover p-1 text-popover-foreground shadow-md")
		if props.Position != "" {
			attrs["data-position"] = props.Position
		}
		if props.Align != "" {
			attrs["data-align"] = props.Align
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func SelectLabel(props DOMProps) templ.Component             { return DropdownMenuLabel(props) }
func SelectItem(props DropdownMenuItemProps) templ.Component { return DropdownMenuItem(props) }
func SelectSeparator(props DOMProps) templ.Component         { return DropdownMenuSeparator(props) }
func SelectScrollUpButton(props DOMProps) templ.Component    { return DialogTrigger(props) }
func SelectScrollDownButton(props DOMProps) templ.Component  { return DialogTrigger(props) }
