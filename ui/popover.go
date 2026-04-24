package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type PopoverProps struct {
	DOMProps
	Open        bool
	DefaultOpen bool
	Side        string
	Align       string
	SideOffset  string
	AlignOffset string
	Modal       bool
}

func Popover(props PopoverProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "popover", "relative inline-block")
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
		if props.Modal {
			attrs["data-modal"] = "true"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func PopoverTrigger(props DOMProps) templ.Component { return DialogTrigger(props) }
func PopoverAnchor(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "popover-anchor", ""), templ.GetChildren(ctx))
	})
}
func PopoverContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "popover-content", "z-50 rounded-md border bg-popover p-4 text-popover-foreground shadow-md outline-none"), templ.GetChildren(ctx))
	})
}
func PopoverHeader(props DOMProps) templ.Component      { return DialogHeader(props) }
func PopoverTitle(props DOMProps) templ.Component       { return DialogTitle(props) }
func PopoverDescription(props DOMProps) templ.Component { return DialogDescription(props) }
