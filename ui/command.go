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
		attrs := attrsFromDOMProps(props.DOMProps, "command", "flex h-full w-full flex-col overflow-hidden rounded-md bg-popover text-popover-foreground")
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
		attrs := attrsFromDOMProps(props.DOMProps, "command-input", "flex h-10 w-full rounded-md bg-transparent py-3 text-sm outline-hidden placeholder:text-muted-foreground disabled:cursor-not-allowed disabled:opacity-50")
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
		input := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if _, err := io.WriteString(w, `<svg xmlns="http://www.w3.org/2000/svg" class="size-4 shrink-0 opacity-50" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="m21 21-4.34-4.34"></path><circle cx="11" cy="11" r="8"></circle></svg>`); err != nil {
				return err
			}
			return renderVoidElement(ctx, w, "input", attrs)
		})
		return renderElement(ctx, w, "div", templ.Attributes{"data-slot": "command-input-wrapper", "class": "flex h-9 items-center gap-2 border-b px-3"}, input)
	})
}
func CommandList(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "command-list", "max-h-[300px] scroll-py-1 overflow-x-hidden overflow-y-auto")
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
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "command-group", "overflow-hidden p-1 text-foreground [&_[cmdk-group-heading]]:px-2 [&_[cmdk-group-heading]]:py-1.5 [&_[cmdk-group-heading]]:text-xs [&_[cmdk-group-heading]]:font-medium [&_[cmdk-group-heading]]:text-muted-foreground"), templ.GetChildren(ctx))
	})
}
func CommandItem(props DropdownMenuItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "relative flex cursor-default items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-hidden select-none data-[disabled=true]:pointer-events-none data-[disabled=true]:opacity-50 data-[selected=true]:bg-accent data-[selected=true]:text-accent-foreground [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4 [&_svg:not([class*='text-'])]:text-muted-foreground"
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
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "command-separator", "-mx-1 h-px bg-border"), nil)
	})
}
