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
		return renderElement(ctx, w, "nav", attrs, templ.GetChildren(ctx))
	})
}

func NavigationMenuList(props DOMProps) templ.Component            { return DropdownMenuGroup(props) }
func NavigationMenuItem(props DOMProps) templ.Component            { return DropdownMenuGroup(props) }
func NavigationMenuTrigger(props DOMProps) templ.Component         { return DialogTrigger(props) }
func NavigationMenuContent(props DOMProps) templ.Component         { return DropdownMenuContent(props) }
func NavigationMenuLink(props BreadcrumbLinkProps) templ.Component { return BreadcrumbLink(props) }
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
