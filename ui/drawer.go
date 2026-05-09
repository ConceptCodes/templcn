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
		open := props.Open || props.DefaultOpen
		attrs := attrsFromDOMProps(props.DOMProps, "drawer", "")
		attrs["data-state"] = openState(open)
		if open {
			attrs["data-open"] = "true"
		}
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
		ctx = context.WithValue(ctx, dialogRenderStateKey{}, dialogRenderState{open: open, slot: "drawer", modal: props.Modal})
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func DrawerTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "drawer-trigger", "")
		attrs["type"] = "button"
		attrs["aria-haspopup"] = "dialog"
		if _, ok := attrs["aria-expanded"]; !ok {
			attrs["aria-expanded"] = "false"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
func DrawerPortal(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "drawer-portal", ""), templ.GetChildren(ctx))
	})
}
func DrawerOverlay(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "drawer-overlay", "fixed inset-0 bg-black/50")
		attrs["aria-hidden"] = "true"
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = dialogStateFromContext(ctx)
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
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
		attrs := attrsFromDOMProps(props, "drawer-close", "")
		attrs["type"] = "button"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func DrawerContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "drawer-content", "fixed inset-x-0 bottom-0 top-auto z-50 m-0 h-fit max-h-[85vh] w-full overflow-auto rounded-t-lg bg-background p-6 shadow-lg backdrop:bg-black/50 [&:not([open])]:hidden")
		state := dialogStateFromContextValue(ctx)
		attrs["role"] = "dialog"
		attrs["tabindex"] = "-1"
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = openState(state.open)
		}
		if state.open {
			attrs["open"] = true
			attrs["data-open"] = "true"
		}
		if state.modal {
			attrs["data-modal"] = "true"
		}
		return renderElement(ctx, w, "dialog", attrs, templ.GetChildren(ctx))
	})
}
