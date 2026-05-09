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

type collapsibleRenderState struct {
	open bool
}

type collapsibleRenderStateKey struct{}

func Collapsible(props CollapsibleProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		open := props.Open || props.DefaultOpen
		attrs := attrsFromDOMProps(props.DOMProps, "collapsible", "grid gap-2")
		attrs["data-state"] = openState(open)
		if open {
			attrs["data-open"] = "true"
			attrs["open"] = true
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		if props.Disabled {
			attrs["data-disabled"] = "true"
		}
		ctx = context.WithValue(ctx, collapsibleRenderStateKey{}, collapsibleRenderState{open: open})
		return renderElement(ctx, w, "details", attrs, templ.GetChildren(ctx))
	})
}

func CollapsibleTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "collapsible-trigger", "cursor-pointer list-none")
		if state, ok := ctx.Value(collapsibleRenderStateKey{}).(collapsibleRenderState); ok {
			attrs["aria-expanded"] = map[bool]string{true: "true", false: "false"}[state.open]
		}
		return renderElement(ctx, w, "summary", attrs, templ.GetChildren(ctx))
	})
}

func CollapsibleContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "collapsible-content", "overflow-hidden pt-2 data-[state=closed]:animate-collapsible-up data-[state=open]:animate-collapsible-down")
		open := false
		if state, ok := ctx.Value(collapsibleRenderStateKey{}).(collapsibleRenderState); ok {
			open = state.open
		}
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = openState(open)
		}
		if !open {
			attrs["hidden"] = true
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
