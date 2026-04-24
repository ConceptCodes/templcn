package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func ScrollArea(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "scroll-area", "relative overflow-hidden"), templ.GetChildren(ctx))
	})
}

type ScrollBarProps struct {
	DOMProps
	Orientation string
}

func ScrollBar(props ScrollBarProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "flex touch-none select-none transition-colors"
		if props.Orientation == "horizontal" {
			className = cn(className, "h-2 w-full flex-col border-t border-t-transparent p-px")
		} else {
			className = cn(className, "h-full w-2 border-l border-l-transparent p-px")
		}
		return renderElement(ctx, w, "div", attrsFromDOMProps(props.DOMProps, "scrollbar", className), templ.GetChildren(ctx))
	})
}
