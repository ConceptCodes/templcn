package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type HoverCardProps struct {
	DOMProps
	Open         bool
	DefaultOpen  bool
	OpenDelayMs  int
	CloseDelayMs int
	Side         string
	Align        string
	SideOffset   string
}

func HoverCard(props HoverCardProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		open := props.Open || props.DefaultOpen
		attrs := attrsFromDOMProps(props.DOMProps, "hover-card", "relative inline-block")
		attrs["data-state"] = openState(open)
		if open {
			attrs["data-open"] = "true"
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
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
		ctx = context.WithValue(ctx, floatingRenderStateKey{}, floatingRenderState{open: open})
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func HoverCardTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "hover-card-trigger", "")
		if _, ok := attrs["type"]; !ok {
			attrs["type"] = "button"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
func HoverCardContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "hover-card-content", "absolute left-0 top-full z-50 mt-1 w-64 origin-(--radix-hover-card-content-transform-origin) rounded-md border bg-popover p-4 text-popover-foreground shadow-md outline-hidden data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=closed]:zoom-out-95 data-[state=open]:animate-in data-[state=open]:fade-in-0 data-[state=open]:zoom-in-95")
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = floatingStateFromContext(ctx)
		}
		if _, ok := attrs["data-side"]; !ok {
			attrs["data-side"] = "bottom"
		}
		if _, ok := attrs["tabindex"]; !ok {
			attrs["tabindex"] = "-1"
		}
		if !floatingOpenFromContext(ctx) {
			attrs["hidden"] = true
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
