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
		if props.Value != "" {
			attrs["data-value"] = props.Value
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
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func CommandDialog(props CommandProps) templ.Component        { return Command(props) }
func CommandInput(props InputProps) templ.Component           { return Input(props) }
func CommandList(props DOMProps) templ.Component              { return SelectGroup(props) }
func CommandEmpty(props DOMProps) templ.Component             { return Empty(props) }
func CommandGroup(props DOMProps) templ.Component             { return SelectGroup(props) }
func CommandItem(props DropdownMenuItemProps) templ.Component { return DropdownMenuItem(props) }
func CommandShortcut(props DOMProps) templ.Component          { return DropdownMenuShortcut(props) }
func CommandSeparator(props DOMProps) templ.Component         { return DropdownMenuSeparator(props) }
