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
		attrs := attrsFromDOMProps(props, "drawer-overlay", "fixed inset-0 z-50 bg-black/50 data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:animate-in data-[state=open]:fade-in-0")
		attrs["aria-hidden"] = "true"
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = dialogStateFromContext(ctx)
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func DrawerHeader(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "drawer-header", "flex flex-col gap-0.5 p-4 group-data-[vaul-drawer-direction=bottom]/drawer-content:text-center group-data-[vaul-drawer-direction=top]/drawer-content:text-center md:gap-1.5 md:text-left"), templ.GetChildren(ctx))
	})
}
func DrawerFooter(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "drawer-footer", "mt-auto flex flex-col gap-2 p-4"), templ.GetChildren(ctx))
	})
}
func DrawerTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "drawer-title", "font-semibold text-foreground")
		if state, ok := dialogAccessibilityFromContext(ctx); ok {
			if _, exists := attrs["id"]; !exists {
				attrs["id"] = state.titleID
			}
		}
		return renderElement(ctx, w, "h2", attrs, templ.GetChildren(ctx))
	})
}
func DrawerDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "drawer-description", "text-sm text-muted-foreground")
		if state, ok := dialogAccessibilityFromContext(ctx); ok {
			if _, exists := attrs["id"]; !exists {
				attrs["id"] = state.descriptionID
			}
		}
		return renderElement(ctx, w, "p", attrs, templ.GetChildren(ctx))
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
		ctx, ownChildren := childrenFromContext(ctx)
		attrs := attrsFromDOMProps(props, "drawer-content", "group/drawer-content fixed z-50 flex h-auto flex-col bg-background backdrop:bg-black/50 [&:not([open])]:hidden data-[vaul-drawer-direction=top]:inset-x-0 data-[vaul-drawer-direction=top]:top-0 data-[vaul-drawer-direction=top]:mb-24 data-[vaul-drawer-direction=top]:max-h-[80vh] data-[vaul-drawer-direction=top]:rounded-b-lg data-[vaul-drawer-direction=top]:border-b data-[vaul-drawer-direction=bottom]:inset-x-0 data-[vaul-drawer-direction=bottom]:bottom-0 data-[vaul-drawer-direction=bottom]:mt-24 data-[vaul-drawer-direction=bottom]:max-h-[80vh] data-[vaul-drawer-direction=bottom]:rounded-t-lg data-[vaul-drawer-direction=bottom]:border-t data-[vaul-drawer-direction=right]:inset-y-0 data-[vaul-drawer-direction=right]:right-0 data-[vaul-drawer-direction=right]:w-3/4 data-[vaul-drawer-direction=right]:border-l data-[vaul-drawer-direction=right]:sm:max-w-sm data-[vaul-drawer-direction=left]:inset-y-0 data-[vaul-drawer-direction=left]:left-0 data-[vaul-drawer-direction=left]:w-3/4 data-[vaul-drawer-direction=left]:border-r data-[vaul-drawer-direction=left]:sm:max-w-sm")
		ctx, attrs = prepareDialogAccessibility(ctx, attrs, "drawer")
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
		if _, ok := attrs["data-vaul-drawer-direction"]; !ok {
			attrs["data-vaul-drawer-direction"] = "bottom"
		}
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := renderElement(ctx, w, "div", templ.Attributes{"class": "mx-auto mt-4 hidden h-2 w-[100px] shrink-0 rounded-full bg-muted group-data-[vaul-drawer-direction=bottom]/drawer-content:block"}, nil); err != nil {
				return err
			}
			return renderChildren(ctx, w, ownChildren)
		})
		return renderElement(ctx, w, "dialog", attrs, children)
	})
}
