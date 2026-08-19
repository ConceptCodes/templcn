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
		attrs := attrsFromDOMProps(props.DOMProps, "navigation-menu", "group/navigation-menu relative flex max-w-max flex-1 items-center justify-center")
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
		if _, ok := attrs["data-viewport"]; !ok {
			attrs["data-viewport"] = true
		}
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
		ctx, ownChildren := childrenFromContext(ctx)
		attrs := attrsFromDOMProps(props, "navigation-menu-trigger", "group inline-flex h-9 w-max items-center justify-center rounded-md bg-background px-4 py-2 text-sm font-medium transition-[color,box-shadow] outline-none hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground focus-visible:ring-[3px] focus-visible:ring-ring/50 focus-visible:outline-1 disabled:pointer-events-none disabled:opacity-50 data-[state=open]:bg-accent/50 data-[state=open]:text-accent-foreground data-[state=open]:hover:bg-accent data-[state=open]:focus:bg-accent")
		if _, ok := attrs["type"]; !ok {
			attrs["type"] = "button"
		}
		if _, ok := attrs["aria-expanded"]; !ok {
			attrs["aria-expanded"] = "false"
		}
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := renderChildren(ctx, w, ownChildren); err != nil {
				return err
			}
			_, err := io.WriteString(w, ` <svg xmlns="http://www.w3.org/2000/svg" class="relative top-[1px] ml-1 size-3 transition duration-300 group-data-[state=open]:rotate-180" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="m6 9 6 6 6-6"></path></svg>`)
			return err
		})
		return renderElement(ctx, w, "button", attrs, children)
	})
}
func NavigationMenuContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "navigation-menu-content", "top-0 left-0 w-full p-2 pr-2.5 data-[motion=from-end]:slide-in-from-right-52 data-[motion=from-start]:slide-in-from-left-52 data-[motion=to-end]:slide-out-to-right-52 data-[motion=to-start]:slide-out-to-left-52 data-[motion^=from-]:animate-in data-[motion^=from-]:fade-in data-[motion^=to-]:animate-out data-[motion^=to-]:fade-out md:absolute md:w-auto group-data-[viewport=false]/navigation-menu:top-full group-data-[viewport=false]/navigation-menu:mt-1.5 group-data-[viewport=false]/navigation-menu:overflow-hidden group-data-[viewport=false]/navigation-menu:rounded-md group-data-[viewport=false]/navigation-menu:border group-data-[viewport=false]/navigation-menu:bg-popover group-data-[viewport=false]/navigation-menu:text-popover-foreground group-data-[viewport=false]/navigation-menu:shadow group-data-[viewport=false]/navigation-menu:duration-200 **:data-[slot=navigation-menu-link]:focus:ring-0 **:data-[slot=navigation-menu-link]:focus:outline-none group-data-[viewport=false]/navigation-menu:data-[state=closed]:animate-out group-data-[viewport=false]/navigation-menu:data-[state=closed]:fade-out-0 group-data-[viewport=false]/navigation-menu:data-[state=closed]:zoom-out-95 group-data-[viewport=false]/navigation-menu:data-[state=open]:animate-in group-data-[viewport=false]/navigation-menu:data-[state=open]:fade-in-0 group-data-[viewport=false]/navigation-menu:data-[state=open]:zoom-in-95")
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
		attrs := attrsFromDOMProps(props.DOMProps, "navigation-menu-link", "flex flex-col gap-1 rounded-sm p-2 text-sm transition-all outline-none hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground focus-visible:ring-[3px] focus-visible:ring-ring/50 focus-visible:outline-1 data-[active=true]:bg-accent/50 data-[active=true]:text-accent-foreground data-[active=true]:hover:bg-accent data-[active=true]:focus:bg-accent [&_svg:not([class*='size-'])]:size-4 [&_svg:not([class*='text-'])]:text-muted-foreground")
		if props.Href != "" {
			attrs["href"] = props.Href
		}
		return renderElement(ctx, w, "a", attrs, templ.GetChildren(ctx))
	})
}
func NavigationMenuIndicator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		inner := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return renderElement(ctx, w, "div", templ.Attributes{"class": "relative top-[60%] h-2 w-2 rotate-45 rounded-tl-sm bg-border shadow-md"}, nil)
		})
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "navigation-menu-indicator", "top-full z-[1] flex h-1.5 items-end justify-center overflow-hidden data-[state=hidden]:animate-out data-[state=hidden]:fade-out data-[state=visible]:animate-in data-[state=visible]:fade-in"), inner)
	})
}
func NavigationMenuViewport(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		viewport := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return renderElement(ctx, w, "div", attrsFromDOMProps(props, "navigation-menu-viewport", "origin-top-center relative mt-1.5 h-[var(--radix-navigation-menu-viewport-height)] w-full overflow-hidden rounded-md border bg-popover text-popover-foreground shadow data-[state=closed]:animate-out data-[state=closed]:zoom-out-95 data-[state=open]:animate-in data-[state=open]:zoom-in-90 md:w-[var(--radix-navigation-menu-viewport-width)]"), templ.GetChildren(ctx))
		})
		return renderElement(ctx, w, "div", templ.Attributes{"class": "absolute top-full left-0 isolate z-50 flex justify-center"}, viewport)
	})
}
