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
		attrs["data-delay-duration"] = props.DelayDuration
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
		open := props.Open || props.DefaultOpen
		attrs := attrsFromDOMProps(props.DOMProps, "tooltip", "relative inline-block")
		attrs["data-state"] = openState(open)
		if open {
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
		ctx = context.WithValue(ctx, floatingRenderStateKey{}, floatingRenderState{open: open})
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func TooltipTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "tooltip-trigger", "")
		if _, ok := attrs["type"]; !ok {
			attrs["type"] = "button"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func TooltipContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		ctx, ownChildren := childrenFromContext(ctx)
		attrs := attrsFromDOMProps(props, "tooltip-content", "absolute left-1/2 top-full z-50 mt-1.5 w-fit -translate-x-1/2 origin-(--radix-tooltip-content-transform-origin) animate-in rounded-md bg-foreground px-3 py-1.5 text-xs text-balance text-background fade-in-0 zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=closed]:zoom-out-95")
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = floatingStateFromContext(ctx)
		}
		if _, ok := attrs["data-side"]; !ok {
			attrs["data-side"] = "bottom"
		}
		if _, ok := attrs["role"]; !ok {
			attrs["role"] = "tooltip"
		}
		if !floatingOpenFromContext(ctx) {
			attrs["hidden"] = true
		}
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := renderChildren(ctx, w, ownChildren); err != nil {
				return err
			}
			return renderElement(ctx, w, "div", templ.Attributes{
				"class":       "z-50 size-2.5 translate-y-[calc(-50%_-_2px)] rotate-45 rounded-[2px] bg-foreground fill-foreground",
				"data-slot":   "tooltip-arrow",
				"aria-hidden": "true",
			}, nil)
		})
		return renderElement(ctx, w, "div", attrs, children)
	})
}
