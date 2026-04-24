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

func DrawerTrigger(props DOMProps) templ.Component     { return DialogTrigger(props) }
func DrawerPortal(props DOMProps) templ.Component      { return DialogPortal(props) }
func DrawerOverlay(props DOMProps) templ.Component     { return DialogOverlay(props) }
func DrawerHeader(props DOMProps) templ.Component      { return DialogHeader(props) }
func DrawerFooter(props DOMProps) templ.Component      { return DialogFooter(props) }
func DrawerTitle(props DOMProps) templ.Component       { return DialogTitle(props) }
func DrawerDescription(props DOMProps) templ.Component { return DialogDescription(props) }
func DrawerClose(props DOMProps) templ.Component       { return DialogClose(props) }

func DrawerContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "drawer-content", "relative z-50 h-full w-full bg-background p-6 shadow-lg"), templ.GetChildren(ctx))
	})
}
