package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type NavigationMenuProps struct {
	DOMProps
	Value         string
	DefaultValue  string
	Orientation   string
	DelayDuration int
}

func NavigationMenu(props NavigationMenuProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "navigation-menu", "relative")
		value := props.Value
		if value == "" {
			value = props.DefaultValue
		}
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.DefaultValue != "" {
			attrs["data-default-value"] = props.DefaultValue
		}
		if props.Orientation != "" {
			attrs["data-orientation"] = props.Orientation
		}
		if props.DelayDuration > 0 {
			attrs["data-delay-duration"] = props.DelayDuration
		}
		attrs["data-state"] = openState(value != "")
		ctx = context.WithValue(ctx, menuRenderStateKey{}, menuRenderState{open: value != ""})
		return renderElement(ctx, w, "nav", attrs, templ.GetChildren(ctx))
	})
}

func NavigationMenuList(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "ul", attrsFromDOMProps(props, "navigation-menu-list", "group flex flex-1 list-none items-center justify-center gap-1"), templ.GetChildren(ctx))
	})
}
func NavigationMenuItem(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "li", attrsFromDOMProps(props, "navigation-menu-item", "relative"), templ.GetChildren(ctx))
	})
}
func NavigationMenuTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "navigation-menu-trigger", "group inline-flex h-9 w-max items-center justify-center rounded-md bg-background px-4 py-2 text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground focus:outline-none disabled:pointer-events-none disabled:opacity-50")
		if _, ok := attrs["type"]; !ok {
			attrs["type"] = "button"
		}
		if _, ok := attrs["aria-expanded"]; !ok {
			attrs["aria-expanded"] = "false"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
func NavigationMenuContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "navigation-menu-content", "left-0 top-0 w-full md:absolute md:w-auto")
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = menuStateFromContext(ctx)
		}
		if !menuOpenFromContext(ctx) {
			attrs["hidden"] = true
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func NavigationMenuLink(props BreadcrumbLinkProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "navigation-menu-link", "block select-none rounded-md p-3 leading-none no-underline outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground")
		if props.Href != "" {
			attrs["href"] = props.Href
		}
		return renderElement(ctx, w, "a", attrs, templ.GetChildren(ctx))
	})
}
func NavigationMenuIndicator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "navigation-menu-indicator", ""), templ.GetChildren(ctx))
	})
}
func NavigationMenuViewport(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "navigation-menu-viewport", ""), templ.GetChildren(ctx))
	})
}
