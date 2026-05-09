package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type CommandProps struct {
	DOMProps
	Value           string
	DefaultValue    string
	Open            bool
	Placeholder     string
	Title           string
	Description     string
	ShowCloseButton bool
}

func Command(props CommandProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "command", "grid gap-2 rounded-lg border bg-popover p-2 text-popover-foreground shadow-md")
		value := props.Value
		if value == "" {
			value = props.DefaultValue
		}
		if value != "" {
			attrs["data-value"] = value
		}
		if props.DefaultValue != "" {
			attrs["data-default-value"] = props.DefaultValue
		}
		if props.Open {
			attrs["data-open"] = "true"
		}
		if props.Placeholder != "" {
			attrs["data-placeholder"] = props.Placeholder
		}
		if props.Title != "" {
			attrs["data-title"] = props.Title
		}
		if props.Description != "" {
			attrs["data-description"] = props.Description
		}
		if props.ShowCloseButton {
			attrs["data-show-close-button"] = "true"
		}
		attrs["role"] = "combobox"
		attrs["aria-expanded"] = "true"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func CommandDialog(props CommandProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "command-dialog", "fixed inset-0 z-50 grid place-items-center")
		if props.Open {
			attrs["data-open"] = "true"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func CommandInput(props InputProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "command-input", inputClasses)
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
		if _, ok := attrs["role"]; !ok {
			attrs["role"] = "searchbox"
		}
		if _, ok := attrs["aria-autocomplete"]; !ok {
			attrs["aria-autocomplete"] = "list"
		}
		return renderVoidElement(ctx, w, "input", attrs)
	})
}
func CommandList(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "command-list", "grid gap-1")
		if _, ok := attrs["role"]; !ok {
			attrs["role"] = "listbox"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func CommandEmpty(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "command-empty", "py-6 text-center text-sm"), templ.GetChildren(ctx))
	})
}
func CommandGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "command-group", "grid gap-1"), templ.GetChildren(ctx))
	})
}
func CommandItem(props DropdownMenuItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "relative flex cursor-pointer select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground"
		if props.Inset {
			className = cn(className, "pl-8")
		}
		if props.Variant == "destructive" {
			className = cn(className, "text-destructive")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "command-item", className)
		if _, ok := attrs["type"]; !ok {
			attrs["type"] = "button"
		}
		if _, ok := attrs["role"]; !ok {
			attrs["role"] = "option"
		}
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.Disabled {
			attrs["disabled"] = true
			attrs["aria-disabled"] = "true"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
func CommandShortcut(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "command-shortcut", "ml-auto text-xs tracking-widest text-muted-foreground"), templ.GetChildren(ctx))
	})
}
func CommandSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "command-separator", "-mx-1 my-1 h-px bg-border"), nil)
	})
}
