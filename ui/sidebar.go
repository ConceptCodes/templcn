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
		attrs := attrsFromDOMProps(props.DOMProps, "sidebar-provider", "")
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
		attrs := attrsFromDOMProps(props.DOMProps, "sidebar", "flex h-full flex-col border-r bg-sidebar text-sidebar-foreground")
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
		attrs := attrsFromDOMProps(props, "sidebar-trigger", "")
		attrs["type"] = "button"
		attrs["aria-expanded"] = "false"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func SidebarRail(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "sidebar-rail", "w-1 bg-sidebar-border"), templ.GetChildren(ctx))
	})
}

func SidebarInset(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "main", attrsFromDOMProps(props, "sidebar-inset", "flex-1"), templ.GetChildren(ctx))
	})
}

func SidebarInput(props InputProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "sidebar-input", inputClasses)
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
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "sidebar-header", "p-2"), templ.GetChildren(ctx))
	})
}

func SidebarFooter(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "sidebar-footer", "p-2"), templ.GetChildren(ctx))
	})
}

func SidebarSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sidebar-separator", "shrink-0 bg-border h-px w-full")
		attrs["aria-hidden"] = "true"
		return renderVoidElement(ctx, w, "hr", attrs)
	})
}

func SidebarContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "sidebar-content", "flex-1 overflow-auto p-2"), templ.GetChildren(ctx))
	})
}

func SidebarGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "sidebar-group", "grid gap-2 p-2"), templ.GetChildren(ctx))
	})
}

func SidebarGroupLabel(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "sidebar-group-label", "px-2 text-xs font-medium uppercase text-sidebar-foreground/70"), templ.GetChildren(ctx))
	})
}

func SidebarGroupAction(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "sidebar-group-action", "ml-auto"), templ.GetChildren(ctx))
	})
}

func SidebarGroupContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "sidebar-group-content", "grid gap-1"), templ.GetChildren(ctx))
	})
}

func SidebarMenu(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "ul", attrsFromDOMProps(props, "sidebar-menu", "grid gap-1"), templ.GetChildren(ctx))
	})
}

func SidebarMenuItem(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "li", attrsFromDOMProps(props, "sidebar-menu-item", ""), templ.GetChildren(ctx))
	})
}

type SidebarMenuButtonProps struct {
	DOMProps
	IsActive bool
	Size     string
}

func SidebarMenuButton(props SidebarMenuButtonProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "flex w-full items-center gap-2 rounded-md px-2 py-1.5 text-sm transition-colors hover:bg-sidebar-accent hover:text-sidebar-accent-foreground"
		if props.IsActive {
			className = cn(className, "bg-sidebar-accent text-sidebar-accent-foreground")
		}
		if props.Size == "sm" {
			className = cn(className, "py-1")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "sidebar-menu-button", className)
		attrs["type"] = "button"
		if props.IsActive {
			attrs["aria-current"] = "page"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func SidebarMenuAction(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "sidebar-menu-action", "ml-auto"), templ.GetChildren(ctx))
	})
}

func SidebarMenuBadge(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "sidebar-menu-badge", "ml-auto inline-flex min-w-5 items-center justify-center rounded-md px-1 text-xs font-medium text-sidebar-foreground tabular-nums"), templ.GetChildren(ctx))
	})
}

func SidebarMenuSkeleton(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "sidebar-menu-skeleton", "animate-pulse rounded-md bg-accent"), templ.GetChildren(ctx))
	})
}

func SidebarMenuSub(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "ul", attrsFromDOMProps(props, "sidebar-menu-sub", "ml-4 grid gap-1 border-l border-sidebar-border pl-3"), templ.GetChildren(ctx))
	})
}

func SidebarMenuSubItem(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "li", attrsFromDOMProps(props, "sidebar-menu-sub-item", ""), templ.GetChildren(ctx))
	})
}

func SidebarMenuSubButton(props SidebarMenuButtonProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "flex w-full items-center gap-2 rounded-md px-2 py-1.5 text-sm transition-colors hover:bg-sidebar-accent hover:text-sidebar-accent-foreground"
		if props.IsActive {
			className = cn(className, "bg-sidebar-accent text-sidebar-accent-foreground")
		}
		if props.Size == "sm" {
			className = cn(className, "py-1")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "sidebar-menu-sub-button", className)
		attrs["type"] = "button"
		if props.IsActive {
			attrs["aria-current"] = "page"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
