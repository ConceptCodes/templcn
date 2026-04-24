package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type HoverCardProps struct {
	DOMProps
	OpenDelayMs  int
	CloseDelayMs int
	Side         string
	Align        string
	SideOffset   string
}

func HoverCard(props HoverCardProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "hover-card", "relative inline-block")
		if props.OpenDelayMs > 0 {
			attrs["data-open-delay"] = props.OpenDelayMs
		}
		if props.CloseDelayMs > 0 {
			attrs["data-close-delay"] = props.CloseDelayMs
		}
		if props.Side != "" {
			attrs["data-side"] = props.Side
		}
		if props.Align != "" {
			attrs["data-align"] = props.Align
		}
		if props.SideOffset != "" {
			attrs["data-side-offset"] = props.SideOffset
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func HoverCardTrigger(props DOMProps) templ.Component { return DialogTrigger(props) }
func HoverCardContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "hover-card-content", "z-50 w-80 rounded-md border bg-popover p-4 text-popover-foreground shadow-md"), templ.GetChildren(ctx))
	})
}
