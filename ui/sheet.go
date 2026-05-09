package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type SheetProps struct {
	DialogProps
	Side            string
	ShowCloseButton bool
}

func Sheet(props SheetProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		open := props.Open || props.DefaultOpen
		attrs := attrsFromDOMProps(props.DOMProps, "sheet", "")
		attrs["data-side"] = sheetSide(props.Side)
		attrs["data-state"] = openState(open)
		if open {
			attrs["data-open"] = "true"
		}
		if props.ShowCloseButton {
			attrs["data-show-close-button"] = "true"
		}
		ctx = context.WithValue(ctx, dialogRenderStateKey{}, dialogRenderState{open: open, slot: "sheet", modal: props.Modal})
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func SheetTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sheet-trigger", "")
		attrs["type"] = "button"
		attrs["aria-haspopup"] = "dialog"
		if _, ok := attrs["aria-expanded"]; !ok {
			attrs["aria-expanded"] = "false"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
func SheetPortal(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "sheet-portal", ""), templ.GetChildren(ctx))
	})
}
func SheetOverlay(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sheet-overlay", "fixed inset-0 z-50 bg-black/80 data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0")
		attrs["aria-hidden"] = "true"
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = dialogStateFromContext(ctx)
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func SheetHeader(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "sheet-header", "flex flex-col gap-2 text-center sm:text-left"), templ.GetChildren(ctx))
	})
}
func SheetFooter(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "sheet-footer", "mt-auto flex flex-col gap-2 sm:flex-row sm:justify-end"), templ.GetChildren(ctx))
	})
}
func SheetTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "h2", attrsFromDOMProps(props, "sheet-title", "text-lg font-semibold text-foreground"), templ.GetChildren(ctx))
	})
}
func SheetDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "p", attrsFromDOMProps(props, "sheet-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}
func SheetClose(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sheet-close", "")
		attrs["type"] = "button"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func SheetContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sheet-content", "fixed z-50 m-0 flex flex-col gap-4 bg-background p-6 shadow-lg transition ease-in-out backdrop:bg-black/80 [&:not([open])]:hidden data-[state=closed]:duration-300 data-[state=open]:duration-500 data-[state=open]:animate-in data-[state=closed]:animate-out data-[side=bottom]:inset-x-0 data-[side=bottom]:bottom-0 data-[side=bottom]:top-auto data-[side=bottom]:border-t data-[side=bottom]:slide-in-from-bottom data-[side=left]:inset-y-0 data-[side=left]:left-0 data-[side=left]:right-auto data-[side=left]:h-full data-[side=left]:w-3/4 data-[side=left]:border-r data-[side=left]:slide-in-from-left data-[side=right]:inset-y-0 data-[side=right]:left-auto data-[side=right]:right-0 data-[side=right]:h-full data-[side=right]:w-3/4 data-[side=right]:border-l data-[side=right]:slide-in-from-right data-[side=top]:inset-x-0 data-[side=top]:bottom-auto data-[side=top]:top-0 data-[side=top]:border-b data-[side=top]:slide-in-from-top sm:max-w-sm")
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
		if _, ok := attrs["data-side"]; !ok {
			attrs["data-side"] = "right"
		}
		if state.modal {
			attrs["data-modal"] = "true"
		}
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := renderChildren(ctx, w, templ.GetChildren(ctx)); err != nil {
				return err
			}
			return renderDialogCloseIcon(ctx, w, "sheet-close")
		})
		return renderElement(ctx, w, "dialog", attrs, children)
	})
}

func sheetSide(side string) string {
	switch side {
	case "top", "bottom", "left", "right":
		return side
	default:
		return "right"
	}
}
