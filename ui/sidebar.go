package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type SidebarProviderProps struct {
	DOMProps
	DefaultOpen bool
	Open        bool
}

func SidebarProvider(props SidebarProviderProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "sidebar-provider", "group/sidebar-wrapper flex min-h-svh w-full has-data-[variant=inset]:bg-sidebar")
		open := props.Open || props.DefaultOpen
		attrs["data-state"] = openState(open)
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		if open {
			attrs["data-open"] = "true"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

type SidebarProps struct {
	DOMProps
	Side        string
	Variant     string
	Collapsible string
}

func Sidebar(props SidebarProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "sidebar", "group peer hidden text-sidebar-foreground md:block")
		if props.Side != "" {
			attrs["data-side"] = props.Side
		}
		if props.Variant != "" {
			attrs["data-variant"] = props.Variant
		}
		if props.Collapsible != "" {
			attrs["data-collapsible"] = props.Collapsible
		}
		return renderElement(ctx, w, "aside", attrs, templ.GetChildren(ctx))
	})
}

func SidebarTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sidebar-trigger", buttonClasses(ButtonVariantGhost, ButtonSizeIcon, "size-7"))
		attrs["data-sidebar"] = "trigger"
		attrs["type"] = "button"
		attrs["aria-expanded"] = "false"
		children := templ.GetChildren(ctx)
		if children == nil {
			children = templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				_, err := io.WriteString(w, `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><rect width="18" height="18" x="3" y="3" rx="2"></rect><path d="M9 3v18"></path></svg><span class="sr-only">Toggle Sidebar</span>`)
				return err
			})
		}
		return renderElement(ctx, w, "button", attrs, children)
	})
}

func SidebarRail(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sidebar-rail", "absolute inset-y-0 z-20 hidden w-4 -translate-x-1/2 transition-all ease-linear group-data-[side=left]:-right-4 group-data-[side=right]:left-0 after:absolute after:inset-y-0 after:left-1/2 after:w-[2px] hover:after:bg-sidebar-border sm:flex in-data-[side=left]:cursor-w-resize in-data-[side=right]:cursor-e-resize [[data-side=left][data-state=collapsed]_&]:cursor-e-resize [[data-side=right][data-state=collapsed]_&]:cursor-w-resize group-data-[collapsible=offcanvas]:translate-x-0 group-data-[collapsible=offcanvas]:after:left-full hover:group-data-[collapsible=offcanvas]:bg-sidebar [[data-side=left][data-collapsible=offcanvas]_&]:-right-2 [[data-side=right][data-collapsible=offcanvas]_&]:-left-2")
		attrs["data-sidebar"] = "rail"
		attrs["aria-label"] = "Toggle Sidebar"
		attrs["tabindex"] = "-1"
		attrs["title"] = "Toggle Sidebar"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func SidebarInset(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "main", attrsFromDOMProps(props, "sidebar-inset", "relative flex w-full flex-1 flex-col bg-background md:peer-data-[variant=inset]:m-2 md:peer-data-[variant=inset]:ml-0 md:peer-data-[variant=inset]:rounded-xl md:peer-data-[variant=inset]:shadow-sm md:peer-data-[variant=inset]:peer-data-[state=collapsed]:ml-2"), templ.GetChildren(ctx))
	})
}

func SidebarInput(props InputProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "sidebar-input", cn(inputClasses, "h-8 w-full bg-background shadow-none"))
		attrs["data-sidebar"] = "input"
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
		if props.Disabled {
			attrs["disabled"] = true
		}
		if props.Required {
			attrs["required"] = true
		}
		if props.Invalid {
			attrs["aria-invalid"] = "true"
		}
		return renderVoidElement(ctx, w, "input", attrs)
	})
}

func SidebarHeader(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sidebar-header", "flex flex-col gap-2 p-2")
		attrs["data-sidebar"] = "header"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func SidebarFooter(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sidebar-footer", "flex flex-col gap-2 p-2")
		attrs["data-sidebar"] = "footer"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func SidebarSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sidebar-separator", "mx-2 w-auto bg-sidebar-border shrink-0 h-px")
		attrs["data-sidebar"] = "separator"
		attrs["aria-hidden"] = "true"
		return renderVoidElement(ctx, w, "hr", attrs)
	})
}

func SidebarContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sidebar-content", "flex min-h-0 flex-1 flex-col gap-2 overflow-auto group-data-[collapsible=icon]:overflow-hidden")
		attrs["data-sidebar"] = "content"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func SidebarGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sidebar-group", "relative flex w-full min-w-0 flex-col p-2")
		attrs["data-sidebar"] = "group"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func SidebarGroupLabel(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sidebar-group-label", "flex h-8 shrink-0 items-center rounded-md px-2 text-xs font-medium text-sidebar-foreground/70 ring-sidebar-ring outline-hidden transition-[margin,opacity] duration-200 ease-linear focus-visible:ring-2 [&>svg]:size-4 [&>svg]:shrink-0 group-data-[collapsible=icon]:-mt-8 group-data-[collapsible=icon]:opacity-0")
		attrs["data-sidebar"] = "group-label"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func SidebarGroupAction(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sidebar-group-action", "absolute top-3.5 right-3 flex aspect-square w-5 items-center justify-center rounded-md p-0 text-sidebar-foreground ring-sidebar-ring outline-hidden transition-transform hover:bg-sidebar-accent hover:text-sidebar-accent-foreground focus-visible:ring-2 [&>svg]:size-4 [&>svg]:shrink-0 after:absolute after:-inset-2 md:after:hidden group-data-[collapsible=icon]:hidden")
		attrs["data-sidebar"] = "group-action"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func SidebarGroupContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sidebar-group-content", "w-full text-sm")
		attrs["data-sidebar"] = "group-content"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func SidebarMenu(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sidebar-menu", "flex w-full min-w-0 flex-col gap-1")
		attrs["data-sidebar"] = "menu"
		return renderElement(ctx, w, "ul", attrs, templ.GetChildren(ctx))
	})
}

func SidebarMenuItem(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sidebar-menu-item", "group/menu-item relative")
		attrs["data-sidebar"] = "menu-item"
		return renderElement(ctx, w, "li", attrs, templ.GetChildren(ctx))
	})
}

type SidebarMenuButtonProps struct {
	DOMProps
	IsActive bool
	Size     string
}

func SidebarMenuButton(props SidebarMenuButtonProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "peer/menu-button flex w-full items-center gap-2 overflow-hidden rounded-md p-2 text-left text-sm ring-sidebar-ring outline-hidden transition-[width,height,padding] group-has-data-[sidebar=menu-action]/menu-item:pr-8 group-data-[collapsible=icon]:size-8! group-data-[collapsible=icon]:p-2! hover:bg-sidebar-accent hover:text-sidebar-accent-foreground focus-visible:ring-2 active:bg-sidebar-accent active:text-sidebar-accent-foreground disabled:pointer-events-none disabled:opacity-50 aria-disabled:pointer-events-none aria-disabled:opacity-50 data-[active=true]:bg-sidebar-accent data-[active=true]:font-medium data-[active=true]:text-sidebar-accent-foreground data-[state=open]:hover:bg-sidebar-accent data-[state=open]:hover:text-sidebar-accent-foreground [&>span:last-child]:truncate [&>svg]:size-4 [&>svg]:shrink-0 h-8"
		if props.IsActive {
			className = cn(className, "bg-sidebar-accent font-medium text-sidebar-accent-foreground")
		}
		if props.Size == "sm" {
			className = cn(className, "h-7 text-xs")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "sidebar-menu-button", className)
		attrs["data-sidebar"] = "menu-button"
		if props.Size == "" {
			attrs["data-size"] = "default"
		} else {
			attrs["data-size"] = props.Size
		}
		attrs["data-active"] = props.IsActive
		attrs["type"] = "button"
		if props.IsActive {
			attrs["aria-current"] = "page"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func SidebarMenuAction(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sidebar-menu-action", "absolute top-1.5 right-1 flex aspect-square w-5 items-center justify-center rounded-md p-0 text-sidebar-foreground ring-sidebar-ring outline-hidden transition-transform peer-hover/menu-button:text-sidebar-accent-foreground hover:bg-sidebar-accent hover:text-sidebar-accent-foreground focus-visible:ring-2 [&>svg]:size-4 [&>svg]:shrink-0 after:absolute after:-inset-2 md:after:hidden peer-data-[size=sm]/menu-button:top-1 peer-data-[size=default]/menu-button:top-1.5 peer-data-[size=lg]/menu-button:top-2.5 group-data-[collapsible=icon]:hidden")
		attrs["data-sidebar"] = "menu-action"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func SidebarMenuBadge(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sidebar-menu-badge", "pointer-events-none absolute right-1 flex h-5 min-w-5 items-center justify-center rounded-md px-1 text-xs font-medium text-sidebar-foreground tabular-nums select-none peer-hover/menu-button:text-sidebar-accent-foreground peer-data-[active=true]/menu-button:text-sidebar-accent-foreground peer-data-[size=sm]/menu-button:top-1 peer-data-[size=default]/menu-button:top-1.5 peer-data-[size=lg]/menu-button:top-2.5 group-data-[collapsible=icon]:hidden")
		attrs["data-sidebar"] = "menu-badge"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func SidebarMenuSkeleton(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sidebar-menu-skeleton", "flex h-8 items-center gap-2 rounded-md px-2")
		attrs["data-sidebar"] = "menu-skeleton"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func SidebarMenuSub(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sidebar-menu-sub", "mx-3.5 flex min-w-0 translate-x-px flex-col gap-1 border-l border-sidebar-border px-2.5 py-0.5 group-data-[collapsible=icon]:hidden")
		attrs["data-sidebar"] = "menu-sub"
		return renderElement(ctx, w, "ul", attrs, templ.GetChildren(ctx))
	})
}

func SidebarMenuSubItem(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sidebar-menu-sub-item", "group/menu-sub-item relative")
		attrs["data-sidebar"] = "menu-sub-item"
		return renderElement(ctx, w, "li", attrs, templ.GetChildren(ctx))
	})
}

func SidebarMenuSubButton(props SidebarMenuButtonProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "flex h-7 min-w-0 -translate-x-px items-center gap-2 overflow-hidden rounded-md px-2 text-sidebar-foreground ring-sidebar-ring outline-hidden hover:bg-sidebar-accent hover:text-sidebar-accent-foreground focus-visible:ring-2 active:bg-sidebar-accent active:text-sidebar-accent-foreground disabled:pointer-events-none disabled:opacity-50 aria-disabled:pointer-events-none aria-disabled:opacity-50 [&>span:last-child]:truncate [&>svg]:size-4 [&>svg]:shrink-0 [&>svg]:text-sidebar-accent-foreground group-data-[collapsible=icon]:hidden text-sm"
		if props.IsActive {
			className = cn(className, "bg-sidebar-accent text-sidebar-accent-foreground")
		}
		if props.Size == "sm" {
			className = cn(className, "text-xs")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "sidebar-menu-sub-button", className)
		attrs["data-sidebar"] = "menu-sub-button"
		if props.Size == "" {
			attrs["data-size"] = "md"
		} else {
			attrs["data-size"] = props.Size
		}
		attrs["data-active"] = props.IsActive
		attrs["type"] = "button"
		if props.IsActive {
			attrs["aria-current"] = "page"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
