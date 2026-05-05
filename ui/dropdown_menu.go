package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func menuContainer(attrs templ.Attributes, children templ.Component, ctx context.Context, w io.Writer) error {
	return renderElement(ctx, w, "div", attrs, children)
}

type DropdownMenuProps struct {
	DOMProps
	Open        bool
	DefaultOpen bool
	Modal       bool
	Side        string
	Align       string
	SideOffset  string
	AlignOffset string
}

func DropdownMenu(props DropdownMenuProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "dropdown-menu", "relative inline-block")
		attrs["data-dropdown-menu-root"] = "true"
		attrs["data-state"] = openState(props.Open || props.DefaultOpen)
		if props.Open || props.DefaultOpen {
			attrs["data-open"] = "true"
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		if props.Modal {
			attrs["data-modal"] = "true"
		}
		if props.Side != "" {
			attrs["data-side"] = props.Side
		}
		if props.Align != "" {
			attrs["data-align"] = props.Align
		}
		if props.SideOffset != "" {
			attrs["data-side-offset"] = props.SideOffset
		}
		if props.AlignOffset != "" {
			attrs["data-align-offset"] = props.AlignOffset
		}
		return menuContainer(attrs, templ.GetChildren(ctx), ctx, w)
	})
}

func DropdownMenuTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "dropdown-menu-trigger", "")
		attrs["data-dropdown-menu-trigger"] = "true"
		attrs["aria-haspopup"] = "menu"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
func DropdownMenuPortal(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dropdown-menu-portal", ""), templ.GetChildren(ctx))
	})
}

func DropdownMenuContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "dropdown-menu-content", "absolute left-0 top-full z-50 mt-1 min-w-32 overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-md data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2")
		attrs["data-dropdown-menu-content"] = "true"
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = "open"
		}
		if _, ok := attrs["data-side"]; !ok {
			attrs["data-side"] = "bottom"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func DropdownMenuGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dropdown-menu-group", "grid gap-1"), templ.GetChildren(ctx))
	})
}

type DropdownMenuItemProps struct {
	DOMProps
	Inset   bool
	Variant string
}

func DropdownMenuItem(props DropdownMenuItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "relative flex cursor-pointer select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground"
		if props.Inset {
			className = cn(className, "pl-8")
		}
		if props.Variant == "destructive" {
			className = cn(className, "text-destructive")
		}
		return renderElement(ctx, w, "button", attrsFromDOMProps(props.DOMProps, "dropdown-menu-item", className), templ.GetChildren(ctx))
	})
}

func DropdownMenuCheckboxItem(props DropdownMenuItemProps) templ.Component {
	return DropdownMenuItem(props)
}

type DropdownMenuRadioGroupProps struct {
	DOMProps
	Value string
}

func DropdownMenuRadioGroup(props DropdownMenuRadioGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "dropdown-menu-radio-group", "grid gap-1")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func DropdownMenuRadioItem(props DropdownMenuItemProps) templ.Component {
	return DropdownMenuItem(props)
}
func DropdownMenuLabel(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dropdown-menu-label", "px-2 py-1.5 text-sm font-semibold"), templ.GetChildren(ctx))
	})
}
func DropdownMenuSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dropdown-menu-separator", "-mx-1 my-1 h-px bg-border"), nil)
	})
}
func DropdownMenuShortcut(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "dropdown-menu-shortcut", "ml-auto text-xs tracking-widest text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func DropdownMenuSub(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dropdown-menu-sub", "relative"), templ.GetChildren(ctx))
	})
}
func DropdownMenuSubTrigger(props DropdownMenuItemProps) templ.Component {
	return DropdownMenuItem(props)
}
func DropdownMenuSubContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "dropdown-menu-sub-content", "absolute left-full top-0 z-50 ml-1 min-w-32 overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-md data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95")
		attrs["data-dropdown-menu-content"] = "true"
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = "open"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
