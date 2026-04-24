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
		attrs := attrsFromDOMProps(props.DOMProps, "tooltip", "inline-block")
		if props.Open {
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

func TooltipTrigger(props DOMProps) templ.Component { return DialogTrigger(props) }

func TooltipContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "tooltip-content", "z-50 rounded-md bg-primary px-3 py-1.5 text-xs text-primary-foreground shadow-md"), templ.GetChildren(ctx))
	})
}
