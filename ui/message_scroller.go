package ui

import (
	"context"
	"github.com/a-h/templ"
	"io"
)

type MessageScrollerProps struct{ DOMProps }
type MessageScrollerItemProps struct {
	DOMProps
	ScrollAnchor bool
}
type MessageScrollerButtonProps struct {
	DOMProps
	Direction string
	Variant   ButtonVariant
	Size      ButtonSize
}

func MessageScrollerProvider(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "message-scroller-provider", "contents"), templ.GetChildren(ctx))
	})
}
func MessageScroller(props MessageScrollerProps) templ.Component {
	return messageScrollerBlock("message-scroller", "group/message-scroller relative flex size-full min-h-0 flex-col overflow-hidden", props.DOMProps)
}
func MessageScrollerViewport(props DOMProps) templ.Component {
	return messageScrollerBlock("message-scroller-viewport", "size-full min-h-0 min-w-0 overflow-y-auto overscroll-contain", props)
}
func MessageScrollerContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "message-scroller-content", "flex h-max min-h-full flex-col gap-8")
		if _, ok := attrs["role"]; !ok {
			attrs["role"] = "log"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func MessageScrollerItem(props MessageScrollerItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "message-scroller-item", "min-w-0 shrink-0")
		if props.ScrollAnchor {
			attrs["data-scroll-anchor"] = true
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func MessageScrollerButton(props MessageScrollerButtonProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		direction, variant, size := props.Direction, props.Variant, props.Size
		if direction == "" {
			direction = "end"
		}
		if variant == "" {
			variant = ButtonVariantSecondary
		}
		if size == "" {
			size = ButtonSizeIconSM
		}
		attrs := attrsFromDOMProps(props.DOMProps, "message-scroller-button", "absolute left-1/2 -translate-x-1/2 bottom-4 z-20 inline-flex size-8 items-center justify-center rounded-full border border-border bg-background text-foreground shadow-md transition-all hover:bg-muted cursor-pointer")
		attrs["data-direction"], attrs["data-variant"], attrs["data-size"] = direction, variant, size
		attrs["type"] = "button"
		children := templ.GetChildren(ctx)
		if children == nil {
			children = templ.ComponentFunc(func(_ context.Context, out io.Writer) error {
				arrow := "↓"
				if direction == "start" {
					arrow = "↑"
				}
				_, err := io.WriteString(out, `<span aria-hidden="true">`+arrow+`</span><span class="sr-only">Scroll to `+map[bool]string{true: "end", false: "start"}[direction == "end"]+`</span>`)
				return err
			})
		}
		return renderElement(ctx, w, "button", attrs, children)
	})
}
func messageScrollerBlock(slot, className string, props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, slot, className), templ.GetChildren(ctx))
	})
}
