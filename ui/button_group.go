package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type ButtonGroupOrientation string

const (
	ButtonGroupOrientationHorizontal ButtonGroupOrientation = "horizontal"
	ButtonGroupOrientationVertical   ButtonGroupOrientation = "vertical"
)

type ButtonGroupProps struct {
	DOMProps
	Orientation ButtonGroupOrientation
}

func ButtonGroup(props ButtonGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		orientation := props.Orientation
		if orientation == "" {
			orientation = ButtonGroupOrientationHorizontal
		}
		className := "flex w-fit items-stretch has-[>[data-slot=button-group]]:gap-2 [&>*]:focus-visible:relative [&>*]:focus-visible:z-10 has-[select[aria-hidden=true]:last-child]:[&>[data-slot=select-trigger]:last-of-type]:rounded-r-md [&>[data-slot=select-trigger]:not([class*='w-'])]:w-fit [&>input]:flex-1 [&>*:not(:first-child)]:rounded-l-none [&>*:not(:first-child)]:border-l-0 [&>*:not(:last-child)]:rounded-r-none"
		if orientation == ButtonGroupOrientationVertical {
			className = "flex w-fit items-stretch flex-col has-[>[data-slot=button-group]]:gap-2 [&>*]:focus-visible:relative [&>*]:focus-visible:z-10 has-[select[aria-hidden=true]:last-child]:[&>[data-slot=select-trigger]:last-of-type]:rounded-r-md [&>[data-slot=select-trigger]:not([class*='w-'])]:w-fit [&>input]:flex-1 [&>*:not(:first-child)]:rounded-t-none [&>*:not(:first-child)]:border-t-0 [&>*:not(:last-child)]:rounded-b-none"
		}
		attrs := attrsFromDOMProps(props.DOMProps, "button-group", className)
		attrs["role"] = "group"
		attrs["data-orientation"] = string(orientation)
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func ButtonGroupText(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "button-group-text", "flex items-center gap-2 rounded-md border bg-muted px-4 text-sm font-medium shadow-xs [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4"), templ.GetChildren(ctx))
	})
}

func ButtonGroupSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "button-group-separator", "relative m-0! self-stretch bg-input shrink-0 data-[orientation=horizontal]:h-px data-[orientation=horizontal]:w-full data-[orientation=vertical]:h-auto data-[orientation=vertical]:w-px")
		attrs["data-orientation"] = SeparatorOrientationVertical
		attrs["aria-hidden"] = "true"
		return renderVoidElement(ctx, w, "hr", attrs)
	})
}
