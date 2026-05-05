package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type TooltipProviderProps struct {
	DOMProps
	DelayDuration int
}

func TooltipProvider(props TooltipProviderProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "tooltip-provider", "")
		if props.DelayDuration > 0 {
			attrs["data-delay-duration"] = props.DelayDuration
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

type TooltipProps struct {
	DOMProps
	Open        bool
	DefaultOpen bool
	Side        string
	Align       string
	SideOffset  string
	AlignOffset string
}

func Tooltip(props TooltipProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "tooltip", "relative inline-block")
		attrs["data-state"] = openState(props.Open || props.DefaultOpen)
		if props.Open || props.DefaultOpen {
			attrs["data-open"] = "true"
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
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
		if props.AlignOffset != "" {
			attrs["data-align-offset"] = props.AlignOffset
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func TooltipTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "tooltip-trigger", ""), templ.GetChildren(ctx))
	})
}

func TooltipContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "tooltip-content", "absolute left-1/2 top-full z-50 mt-1.5 w-max -translate-x-1/2 rounded-md bg-primary px-3 py-1.5 text-xs text-primary-foreground shadow-md animate-in fade-in-0 zoom-in-95 data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=closed]:zoom-out-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2")
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = "delayed-open"
		}
		if _, ok := attrs["data-side"]; !ok {
			attrs["data-side"] = "bottom"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
