package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type CollapsibleProps struct {
	DOMProps
	Open        bool
	DefaultOpen bool
	Disabled    bool
}

func Collapsible(props CollapsibleProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "collapsible", "grid gap-2")
		if props.Open {
			attrs["open"] = true
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		if props.Disabled {
			attrs["data-disabled"] = "true"
		}
		return renderElement(ctx, w, "details", attrs, templ.GetChildren(ctx))
	})
}

func CollapsibleTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "summary", attrsFromDOMProps(props, "collapsible-trigger", "cursor-pointer list-none"), templ.GetChildren(ctx))
	})
}

func CollapsibleContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "collapsible-content", "pt-2"), templ.GetChildren(ctx))
	})
}
