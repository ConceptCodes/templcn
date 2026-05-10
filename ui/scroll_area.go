package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func ScrollArea(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		viewport := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := renderElement(ctx, w, "div", templ.Attributes{"data-slot": "scroll-area-viewport", "class": "size-full rounded-[inherit] transition-[color,box-shadow] outline-none focus-visible:ring-[3px] focus-visible:ring-ring/50 focus-visible:outline-1"}, templ.GetChildren(ctx)); err != nil {
				return err
			}
			return ScrollBar(ScrollBarProps{}).Render(ctx, w)
		})
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "scroll-area", "relative"), viewport)
	})
}

type ScrollBarProps struct {
	DOMProps
	Orientation string
}

func ScrollBar(props ScrollBarProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "flex touch-none p-px transition-colors select-none"
		if props.Orientation == "horizontal" {
			className = cn(className, "h-2.5 flex-col border-t border-t-transparent")
		} else {
			className = cn(className, "h-full w-2.5 border-l border-l-transparent")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "scroll-area-scrollbar", className)
		if props.Orientation == "" {
			attrs["data-orientation"] = "vertical"
		} else {
			attrs["data-orientation"] = props.Orientation
		}
		children := templ.GetChildren(ctx)
		if children == nil {
			children = templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				return renderElement(ctx, w, "div", templ.Attributes{"data-slot": "scroll-area-thumb", "class": "relative flex-1 rounded-full bg-border"}, nil)
			})
		}
		return renderElement(ctx, w, "div", attrs, children)
	})
}
