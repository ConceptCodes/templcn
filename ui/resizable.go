package ui

import (
	"context"
	"io"
	"strconv"

	"github.com/a-h/templ"
)

type ResizablePanelGroupProps struct {
	DOMProps
	Direction string
}

func ResizablePanelGroup(props ResizablePanelGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "resizable-panel-group", "flex h-full w-full aria-[orientation=vertical]:flex-col")
		attrs["data-direction"] = "horizontal"
		attrs["aria-orientation"] = "horizontal"
		if props.Direction != "" {
			attrs["data-direction"] = props.Direction
			attrs["aria-orientation"] = props.Direction
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
		size := props.DefaultSize
		if size <= 0 {
			size = 50
		}
		attrs["data-size"] = size
		attrs["style"] = "flex-basis: " + strconv.FormatFloat(size, 'f', -1, 64) + "%"
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
		attrs := attrsFromDOMProps(props.DOMProps, "resizable-handle", "relative flex w-px items-center justify-center bg-border after:absolute after:inset-y-0 after:left-1/2 after:w-1 after:-translate-x-1/2 focus-visible:ring-1 focus-visible:ring-ring focus-visible:ring-offset-1 focus-visible:outline-hidden aria-[orientation=horizontal]:h-px aria-[orientation=horizontal]:w-full aria-[orientation=horizontal]:after:left-0 aria-[orientation=horizontal]:after:h-1 aria-[orientation=horizontal]:after:w-full aria-[orientation=horizontal]:after:translate-x-0 aria-[orientation=horizontal]:after:-translate-y-1/2 [&[aria-orientation=horizontal]>div]:rotate-90")
		attrs["role"] = "separator"
		attrs["tabindex"] = "0"
		attrs["aria-orientation"] = "vertical"
		if props.WithHandle {
			attrs["data-with-handle"] = "true"
		}
		children := templ.GetChildren(ctx)
		if props.WithHandle && children == nil {
			children = templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				icon := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
					_, err := io.WriteString(w, `<svg xmlns="http://www.w3.org/2000/svg" class="size-2.5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><circle cx="9" cy="12" r="1"></circle><circle cx="9" cy="5" r="1"></circle><circle cx="9" cy="19" r="1"></circle><circle cx="15" cy="12" r="1"></circle><circle cx="15" cy="5" r="1"></circle><circle cx="15" cy="19" r="1"></circle></svg>`)
					return err
				})
				return renderElement(ctx, w, "div", templ.Attributes{"class": "z-10 flex h-4 w-3 items-center justify-center rounded-xs border bg-border"}, icon)
			})
		}
		return renderElement(ctx, w, "div", attrs, children)
	})
}
