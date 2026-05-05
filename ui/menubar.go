package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type MenubarProps struct{ DOMProps }

func Menubar(props MenubarProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "menubar", "flex h-9 items-center gap-1 rounded-md border bg-background p-1 shadow-xs")
		attrs["role"] = "menubar"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func MenubarMenu(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "menubar-menu", "relative"), templ.GetChildren(ctx))
	})
}
func MenubarTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "menubar-trigger", "inline-flex cursor-default select-none items-center rounded-sm px-3 py-1 text-sm font-medium outline-none hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground"), templ.GetChildren(ctx))
	})
}
func MenubarPortal(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "menubar-portal", ""), templ.GetChildren(ctx))
	})
}
func MenubarContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "menubar-content", "z-50 min-w-32 rounded-md border bg-popover p-1 text-popover-foreground shadow-md")
		attrs["hidden"] = true
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func MenubarGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "menubar-group", "grid gap-1")
		attrs["role"] = "group"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func MenubarItem(props DropdownMenuItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "relative flex cursor-pointer select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground"
		if props.Inset {
			className = cn(className, "pl-8")
		}
		if props.Variant == "destructive" {
			className = cn(className, "text-destructive")
		}
		return renderElement(ctx, w, "button", attrsFromDOMProps(props.DOMProps, "menubar-item", className), templ.GetChildren(ctx))
	})
}
func MenubarCheckboxItem(props DropdownMenuItemProps) templ.Component {
	return MenubarItem(DropdownMenuItemProps{DOMProps: DOMProps{ID: props.ID, Class: props.Class, Element: props.Element, Attrs: props.Attrs}, Inset: true, Variant: props.Variant})
}
func MenubarRadioGroup(props DropdownMenuRadioGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "menubar-radio-group", "grid gap-1")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func MenubarRadioItem(props DropdownMenuItemProps) templ.Component {
	return MenubarItem(DropdownMenuItemProps{DOMProps: DOMProps{ID: props.ID, Class: props.Class, Element: props.Element, Attrs: props.Attrs}, Inset: true, Variant: props.Variant})
}
func MenubarLabel(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "menubar-label", "px-2 py-1.5 text-sm font-semibold"), templ.GetChildren(ctx))
	})
}
func MenubarSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "menubar-separator", "-mx-1 my-1 h-px bg-border"), nil)
	})
}
func MenubarShortcut(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "menubar-shortcut", "ml-auto text-xs tracking-widest text-muted-foreground"), templ.GetChildren(ctx))
	})
}
func MenubarSub(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "menubar-sub", "relative"), templ.GetChildren(ctx))
	})
}
func MenubarSubTrigger(props DropdownMenuItemProps) templ.Component {
	return MenubarItem(props)
}
func MenubarSubContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "menubar-sub-content", "z-50 min-w-32 rounded-md border bg-popover p-1 text-popover-foreground shadow-md")
		attrs["hidden"] = true
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
