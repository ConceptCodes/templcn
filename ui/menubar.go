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
		attrs["data-state"] = "closed"
		ctx = context.WithValue(ctx, menuRenderStateKey{}, menuRenderState{open: false})
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
		attrs := attrsFromDOMProps(props, "menubar-trigger", "flex items-center rounded-sm px-2 py-1 text-sm font-medium outline-hidden select-none focus:bg-accent focus:text-accent-foreground data-[state=open]:bg-accent data-[state=open]:text-accent-foreground")
		if _, ok := attrs["type"]; !ok {
			attrs["type"] = "button"
		}
		if _, ok := attrs["role"]; !ok {
			attrs["role"] = "menuitem"
		}
		if _, ok := attrs["aria-haspopup"]; !ok {
			attrs["aria-haspopup"] = "menu"
		}
		if _, ok := attrs["aria-expanded"]; !ok {
			attrs["aria-expanded"] = "false"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
func MenubarPortal(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "menubar-portal", ""), templ.GetChildren(ctx))
	})
}
func MenubarContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "menubar-content", "z-50 min-w-[12rem] origin-(--radix-menubar-content-transform-origin) overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-md data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 data-[state=closed]:fade-out-0 data-[state=closed]:zoom-out-95 data-[state=open]:animate-in data-[state=open]:fade-in-0 data-[state=open]:zoom-in-95")
		if _, ok := attrs["role"]; !ok {
			attrs["role"] = "menu"
		}
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = menuStateFromContext(ctx)
		}
		if !menuOpenFromContext(ctx) {
			attrs["hidden"] = true
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func MenubarGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "menubar-group", "")
		attrs["role"] = "group"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func MenubarItem(props DropdownMenuItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "relative flex cursor-default items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-hidden select-none focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50 data-[inset]:pl-8 data-[variant=destructive]:text-destructive data-[variant=destructive]:focus:bg-destructive/10 data-[variant=destructive]:focus:text-destructive dark:data-[variant=destructive]:focus:bg-destructive/20 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4 [&_svg:not([class*='text-'])]:text-muted-foreground data-[variant=destructive]:*:[svg]:text-destructive!"
		attrs := attrsFromDOMProps(props.DOMProps, "menubar-item", className)
		if props.Inset {
			attrs["data-inset"] = true
		}
		if props.Variant != "" {
			attrs["data-variant"] = props.Variant
		} else {
			attrs["data-variant"] = "default"
		}
		if _, ok := attrs["type"]; !ok {
			attrs["type"] = "button"
		}
		if _, ok := attrs["role"]; !ok {
			attrs["role"] = "menuitem"
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
func MenubarCheckboxItem(props DropdownMenuItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "menubar-checkbox-item", "relative flex cursor-default items-center gap-2 rounded-xs py-1.5 pr-2 pl-8 text-sm outline-hidden select-none focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4")
		attrs["type"] = "button"
		attrs["role"] = "menuitemcheckbox"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
func MenubarRadioGroup(props DropdownMenuRadioGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "menubar-radio-group", "")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func MenubarRadioItem(props DropdownMenuItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "menubar-radio-item", "relative flex cursor-default items-center gap-2 rounded-xs py-1.5 pr-2 pl-8 text-sm outline-hidden select-none focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4")
		attrs["type"] = "button"
		attrs["role"] = "menuitemradio"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
func MenubarLabel(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "menubar-label", "px-2 py-1.5 text-sm font-medium data-[inset]:pl-8"), templ.GetChildren(ctx))
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
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		ctx, ownChildren := childrenFromContext(ctx)
		attrs := attrsFromDOMProps(props.DOMProps, "menubar-sub-trigger", "flex cursor-default items-center rounded-sm px-2 py-1.5 text-sm outline-none select-none focus:bg-accent focus:text-accent-foreground data-[inset]:pl-8 data-[state=open]:bg-accent data-[state=open]:text-accent-foreground")
		attrs["type"] = "button"
		if props.Inset {
			attrs["data-inset"] = true
		}
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := renderChildren(ctx, w, ownChildren); err != nil {
				return err
			}
			_, err := io.WriteString(w, `<svg xmlns="http://www.w3.org/2000/svg" class="ml-auto h-4 w-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="m9 18 6-6-6-6"></path></svg>`)
			return err
		})
		return renderElement(ctx, w, "button", attrs, children)
	})
}
func MenubarSubContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "menubar-sub-content", "z-50 min-w-[8rem] origin-(--radix-menubar-content-transform-origin) overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-lg data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=closed]:zoom-out-95 data-[state=open]:animate-in data-[state=open]:fade-in-0 data-[state=open]:zoom-in-95")
		attrs["role"] = "menu"
		attrs["data-state"] = "closed"
		attrs["hidden"] = true
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
