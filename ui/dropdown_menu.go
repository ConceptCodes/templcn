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

type menuRenderState struct {
	open bool
}

type menuRenderStateKey struct{}

func DropdownMenu(props DropdownMenuProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		open := props.Open || props.DefaultOpen
		attrs := attrsFromDOMProps(props.DOMProps, "dropdown-menu", "relative inline-block")
		attrs["data-dropdown-menu-root"] = "true"
		attrs["data-state"] = openState(open)
		if open {
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
		ctx = context.WithValue(ctx, menuRenderStateKey{}, menuRenderState{open: open})
		return menuContainer(attrs, templ.GetChildren(ctx), ctx, w)
	})
}

func DropdownMenuTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "dropdown-menu-trigger", "")
		attrs["type"] = "button"
		attrs["data-dropdown-menu-trigger"] = "true"
		attrs["aria-haspopup"] = "menu"
		if _, ok := attrs["aria-expanded"]; !ok {
			attrs["aria-expanded"] = "false"
		}
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
		attrs := attrsFromDOMProps(props, "dropdown-menu-content", "absolute left-0 top-full z-50 mt-1 max-h-(--radix-dropdown-menu-content-available-height) min-w-[8rem] origin-(--radix-dropdown-menu-content-transform-origin) overflow-x-hidden overflow-y-auto rounded-md border bg-popover p-1 text-popover-foreground shadow-md data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=closed]:zoom-out-95 data-[state=open]:animate-in data-[state=open]:fade-in-0 data-[state=open]:zoom-in-95")
		attrs["data-dropdown-menu-content"] = "true"
		attrs["role"] = "menu"
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = menuStateFromContext(ctx)
		}
		if menuStateFromContext(ctx) == "closed" {
			attrs["hidden"] = true
		}
		if _, ok := attrs["data-side"]; !ok {
			attrs["data-side"] = "bottom"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func DropdownMenuGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dropdown-menu-group", ""), templ.GetChildren(ctx))
	})
}

type DropdownMenuItemProps struct {
	DOMProps
	Inset    bool
	Variant  string
	Value    string
	Disabled bool
}

func DropdownMenuItem(props DropdownMenuItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "relative flex cursor-default items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-hidden select-none focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50 data-[inset]:pl-8 data-[variant=destructive]:text-destructive data-[variant=destructive]:focus:bg-destructive/10 data-[variant=destructive]:focus:text-destructive dark:data-[variant=destructive]:focus:bg-destructive/20 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4 [&_svg:not([class*='text-'])]:text-muted-foreground data-[variant=destructive]:*:[svg]:text-destructive!"
		attrs := attrsFromDOMProps(props.DOMProps, "dropdown-menu-item", className)
		attrs["type"] = "button"
		attrs["role"] = "menuitem"
		if props.Inset {
			attrs["data-inset"] = true
		}
		if props.Variant != "" {
			attrs["data-variant"] = props.Variant
		} else {
			attrs["data-variant"] = "default"
		}
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.Disabled {
			attrs["disabled"] = true
			attrs["aria-disabled"] = "true"
			attrs["data-disabled"] = "true"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func DropdownMenuCheckboxItem(props DropdownMenuItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "dropdown-menu-checkbox-item", "relative flex cursor-default items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4")
		attrs["type"] = "button"
		attrs["role"] = "menuitemcheckbox"
		if props.Disabled {
			attrs["disabled"] = true
			attrs["aria-disabled"] = "true"
			attrs["data-disabled"] = "true"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

type DropdownMenuRadioGroupProps struct {
	DOMProps
	Value string
}

func DropdownMenuRadioGroup(props DropdownMenuRadioGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "dropdown-menu-radio-group", "")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func DropdownMenuRadioItem(props DropdownMenuItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "dropdown-menu-radio-item", "relative flex cursor-default items-center gap-2 rounded-sm py-1.5 pr-2 pl-8 text-sm outline-hidden select-none focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4")
		attrs["type"] = "button"
		attrs["role"] = "menuitemradio"
		if props.Disabled {
			attrs["disabled"] = true
			attrs["aria-disabled"] = "true"
			attrs["data-disabled"] = "true"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
func DropdownMenuLabel(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dropdown-menu-label", "px-2 py-1.5 text-sm font-medium data-[inset]:pl-8"), templ.GetChildren(ctx))
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
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		ctx, ownChildren := childrenFromContext(ctx)
		attrs := attrsFromDOMProps(props.DOMProps, "dropdown-menu-sub-trigger", "flex cursor-default items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-hidden select-none focus:bg-accent focus:text-accent-foreground data-[inset]:pl-8 data-[state=open]:bg-accent data-[state=open]:text-accent-foreground [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4 [&_svg:not([class*='text-'])]:text-muted-foreground")
		attrs["type"] = "button"
		if props.Inset {
			attrs["data-inset"] = true
		}
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := renderChildren(ctx, w, ownChildren); err != nil {
				return err
			}
			_, err := io.WriteString(w, `<svg xmlns="http://www.w3.org/2000/svg" class="ml-auto size-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="m9 18 6-6-6-6"></path></svg>`)
			return err
		})
		return renderElement(ctx, w, "button", attrs, children)
	})
}
func DropdownMenuSubContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "dropdown-menu-sub-content", "absolute left-full top-0 z-50 ml-1 min-w-[8rem] origin-(--radix-dropdown-menu-content-transform-origin) overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-lg data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=closed]:zoom-out-95 data-[state=open]:animate-in data-[state=open]:fade-in-0 data-[state=open]:zoom-in-95")
		attrs["data-dropdown-menu-content"] = "true"
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = "open"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func menuStateFromContext(ctx context.Context) string {
	state, ok := ctx.Value(menuRenderStateKey{}).(menuRenderState)
	if !ok {
		return "closed"
	}
	return openState(state.open)
}

func menuOpenFromContext(ctx context.Context) bool {
	state, ok := ctx.Value(menuRenderStateKey{}).(menuRenderState)
	return ok && state.open
}
