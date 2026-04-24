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
		attrs := attrsFromDOMProps(props.DOMProps, "sheet", "fixed inset-0 z-50")
		if props.Side != "" {
			attrs["data-side"] = props.Side
		}
		if props.ShowCloseButton {
			attrs["data-show-close-button"] = "true"
		}
		if props.Open {
			attrs["open"] = true
		}
		return renderElement(ctx, w, "dialog", attrs, templ.GetChildren(ctx))
	})
}

func SheetTrigger(props DOMProps) templ.Component     { return DialogTrigger(props) }
func SheetPortal(props DOMProps) templ.Component      { return DialogPortal(props) }
func SheetOverlay(props DOMProps) templ.Component     { return DialogOverlay(props) }
func SheetHeader(props DOMProps) templ.Component      { return DialogHeader(props) }
func SheetFooter(props DOMProps) templ.Component      { return DialogFooter(props) }
func SheetTitle(props DOMProps) templ.Component       { return DialogTitle(props) }
func SheetDescription(props DOMProps) templ.Component { return DialogDescription(props) }
func SheetClose(props DOMProps) templ.Component       { return DialogClose(props) }

func SheetContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "sheet-content", "relative z-50 h-full w-full max-w-lg bg-background p-6 shadow-lg"), templ.GetChildren(ctx))
	})
}
