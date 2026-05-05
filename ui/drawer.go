package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type DrawerProps struct {
	SheetProps
	Dismissible      bool
	SnapPoints       []string
	DefaultSnapPoint string
}

func Drawer(props DrawerProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "drawer", "fixed inset-0 z-50")
		if props.Side != "" {
			attrs["data-side"] = props.Side
		}
		if props.Dismissible {
			attrs["data-dismissible"] = "true"
		}
		if len(props.SnapPoints) > 0 {
			attrs["data-snap-points"] = props.SnapPoints
		}
		if props.DefaultSnapPoint != "" {
			attrs["data-default-snap-point"] = props.DefaultSnapPoint
		}
		return renderElement(ctx, w, "dialog", attrs, templ.GetChildren(ctx))
	})
}

func DrawerTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "drawer-trigger", ""), templ.GetChildren(ctx))
	})
}
func DrawerPortal(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "drawer-portal", ""), templ.GetChildren(ctx))
	})
}
func DrawerOverlay(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "drawer-overlay", "fixed inset-0 bg-black/50"), templ.GetChildren(ctx))
	})
}
func DrawerHeader(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "drawer-header", "flex flex-col gap-2 text-center sm:text-left"), templ.GetChildren(ctx))
	})
}
func DrawerFooter(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "drawer-footer", "flex flex-col-reverse gap-2 sm:flex-row sm:justify-end"), templ.GetChildren(ctx))
	})
}
func DrawerTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "h2", attrsFromDOMProps(props, "drawer-title", "text-lg font-semibold tracking-tight"), templ.GetChildren(ctx))
	})
}
func DrawerDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "p", attrsFromDOMProps(props, "drawer-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}
func DrawerClose(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "drawer-close", ""), templ.GetChildren(ctx))
	})
}

func DrawerContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "drawer-content", "relative z-50 h-full w-full bg-background p-6 shadow-lg"), templ.GetChildren(ctx))
	})
}
