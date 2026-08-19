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
		ctx = context.WithValue(ctx, dialogRenderStateKey{}, dialogRenderState{open: open, slot: "sheet", modal: props.Modal, showCloseButton: props.ShowCloseButton})
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
		attrs := attrsFromDOMProps(props, "sheet-overlay", "fixed inset-0 z-50 bg-black/50 data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:animate-in data-[state=open]:fade-in-0")
		attrs["aria-hidden"] = "true"
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = dialogStateFromContext(ctx)
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func SheetHeader(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "sheet-header", "flex flex-col gap-1.5 p-4"), templ.GetChildren(ctx))
	})
}
func SheetFooter(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "sheet-footer", "mt-auto flex flex-col gap-2 p-4"), templ.GetChildren(ctx))
	})
}
func SheetTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sheet-title", "font-semibold text-foreground")
		if state, ok := dialogAccessibilityFromContext(ctx); ok {
			if _, exists := attrs["id"]; !exists {
				attrs["id"] = state.titleID
			}
		}
		return renderElement(ctx, w, "h2", attrs, templ.GetChildren(ctx))
	})
}
func SheetDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "sheet-description", "text-sm text-muted-foreground")
		if state, ok := dialogAccessibilityFromContext(ctx); ok {
			if _, exists := attrs["id"]; !exists {
				attrs["id"] = state.descriptionID
			}
		}
		return renderElement(ctx, w, "p", attrs, templ.GetChildren(ctx))
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
		ctx, ownChildren := childrenFromContext(ctx)
		attrs := attrsFromDOMProps(props, "sheet-content", "fixed z-50 flex flex-col gap-4 bg-background shadow-lg transition ease-in-out backdrop:bg-black/50 [&:not([open])]:hidden data-[state=closed]:animate-out data-[state=closed]:duration-300 data-[state=open]:animate-in data-[state=open]:duration-500 data-[side=bottom]:inset-x-0 data-[side=bottom]:bottom-0 data-[side=bottom]:h-auto data-[side=bottom]:border-t data-[side=bottom]:data-[state=closed]:slide-out-to-bottom data-[side=bottom]:data-[state=open]:slide-in-from-bottom data-[side=left]:inset-y-0 data-[side=left]:left-0 data-[side=left]:h-full data-[side=left]:w-3/4 data-[side=left]:border-r data-[side=left]:data-[state=closed]:slide-out-to-left data-[side=left]:data-[state=open]:slide-in-from-left data-[side=right]:inset-y-0 data-[side=right]:right-0 data-[side=right]:h-full data-[side=right]:w-3/4 data-[side=right]:border-l data-[side=right]:data-[state=closed]:slide-out-to-right data-[side=right]:data-[state=open]:slide-in-from-right data-[side=top]:inset-x-0 data-[side=top]:top-0 data-[side=top]:h-auto data-[side=top]:border-b data-[side=top]:data-[state=closed]:slide-out-to-top data-[side=top]:data-[state=open]:slide-in-from-top sm:max-w-sm")
		ctx, attrs = prepareDialogAccessibility(ctx, attrs, "sheet")
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
			if err := renderChildren(ctx, w, ownChildren); err != nil {
				return err
			}
			if dialogStateFromContextValue(ctx).showCloseButton {
				return renderSheetCloseIcon(ctx, w)
			}
			return nil
		})
		return renderElement(ctx, w, "dialog", attrs, children)
	})
}

func renderSheetCloseIcon(ctx context.Context, w io.Writer) error {
	attrs := templ.Attributes{
		"type":        "button",
		"data-slot":   "sheet-close",
		"aria-label":  "Close",
		"class":       "absolute top-4 right-4 rounded-xs opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:ring-2 focus:ring-ring focus:ring-offset-2 focus:outline-hidden disabled:pointer-events-none data-[state=open]:bg-secondary",
		"data-action": "close",
	}
	icon := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, `<svg xmlns="http://www.w3.org/2000/svg" class="size-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M18 6 6 18"></path><path d="m6 6 12 12"></path></svg>`)
		return err
	})
	return renderElement(ctx, w, "button", attrs, icon)
}

func sheetSide(side string) string {
	switch side {
	case "top", "bottom", "left", "right":
		return side
	default:
		return "right"
	}
}
