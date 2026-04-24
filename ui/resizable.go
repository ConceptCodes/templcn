package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type ResizablePanelGroupProps struct {
	DOMProps
	Direction string
}

func ResizablePanelGroup(props ResizablePanelGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "resizable-panel-group", "flex min-h-0 min-w-0")
		if props.Direction != "" {
			attrs["data-direction"] = props.Direction
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

type ResizablePanelProps struct {
	DOMProps
	DefaultSize float64
	MinSize     float64
	MaxSize     float64
	Collapsible bool
}

func ResizablePanel(props ResizablePanelProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "resizable-panel", "min-h-0 min-w-0")
		if props.DefaultSize > 0 {
			attrs["data-default-size"] = props.DefaultSize
		}
		if props.MinSize > 0 {
			attrs["data-min-size"] = props.MinSize
		}
		if props.MaxSize > 0 {
			attrs["data-max-size"] = props.MaxSize
		}
		if props.Collapsible {
			attrs["data-collapsible"] = "true"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

type ResizableHandleProps struct {
	DOMProps
	WithHandle bool
}

func ResizableHandle(props ResizableHandleProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "resizable-handle", "relative flex items-center justify-center bg-border")
		if props.WithHandle {
			attrs["data-with-handle"] = "true"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
