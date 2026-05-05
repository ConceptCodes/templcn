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
		attrs["data-state"] = "closed"
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

func HoverCardTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "hover-card-trigger", ""), templ.GetChildren(ctx))
	})
}
func HoverCardContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "hover-card-content", "absolute left-0 top-full z-50 mt-2 w-64 rounded-md border bg-popover p-4 text-popover-foreground shadow-md outline-none data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2")
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = "open"
		}
		if _, ok := attrs["data-side"]; !ok {
			attrs["data-side"] = "bottom"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
