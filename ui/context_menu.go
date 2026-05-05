package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type ContextMenuProps struct{ DOMProps }

func ContextMenu(props ContextMenuProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props.DOMProps, "context-menu", "relative inline-block"), templ.GetChildren(ctx))
	})
}
func ContextMenuTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "context-menu-trigger", ""), templ.GetChildren(ctx))
	})
}
func ContextMenuPortal(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "context-menu-portal", ""), templ.GetChildren(ctx))
	})
}
func ContextMenuContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "context-menu-content", "z-50 min-w-32 rounded-md border bg-popover p-1 text-popover-foreground shadow-md")
		attrs["hidden"] = true
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func ContextMenuItem(props DropdownMenuItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "relative flex cursor-pointer select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground"
		if props.Inset {
			className = cn(className, "pl-8")
		}
		if props.Variant == "destructive" {
			className = cn(className, "text-destructive")
		}
		return renderElement(ctx, w, "button", attrsFromDOMProps(props.DOMProps, "context-menu-item", className), templ.GetChildren(ctx))
	})
}
func ContextMenuCheckboxItem(props DropdownMenuItemProps) templ.Component {
	return ContextMenuItem(DropdownMenuItemProps{DOMProps: DOMProps{ID: props.ID, Class: props.Class, Element: props.Element, Attrs: props.Attrs}, Inset: true, Variant: props.Variant})
}
func ContextMenuRadioItem(props DropdownMenuItemProps) templ.Component {
	return ContextMenuItem(DropdownMenuItemProps{DOMProps: DOMProps{ID: props.ID, Class: props.Class, Element: props.Element, Attrs: props.Attrs}, Inset: true, Variant: props.Variant})
}
func ContextMenuRadioGroup(props DropdownMenuRadioGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "context-menu-radio-group", "grid gap-1")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func ContextMenuLabel(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "context-menu-label", "px-2 py-1.5 text-sm font-semibold"), templ.GetChildren(ctx))
	})
}
func ContextMenuSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "context-menu-separator", "-mx-1 my-1 h-px bg-border"), nil)
	})
}
func ContextMenuShortcut(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "context-menu-shortcut", "ml-auto text-xs tracking-widest text-muted-foreground"), templ.GetChildren(ctx))
	})
}
func ContextMenuGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "context-menu-group", "grid gap-1")
		attrs["role"] = "group"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func ContextMenuSub(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "context-menu-sub", "relative"), templ.GetChildren(ctx))
	})
}
func ContextMenuSubContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "context-menu-sub-content", "z-50 min-w-32 rounded-md border bg-popover p-1 text-popover-foreground shadow-md")
		attrs["hidden"] = true
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func ContextMenuSubTrigger(props DropdownMenuItemProps) templ.Component {
	return ContextMenuItem(props)
}
