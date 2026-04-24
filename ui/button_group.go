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
		className := "inline-flex"
		if orientation == ButtonGroupOrientationVertical {
			className = "inline-flex flex-col"
		}
		return renderElement(ctx, w, "div", attrsFromDOMProps(props.DOMProps, "button-group", className), templ.GetChildren(ctx))
	})
}

func ButtonGroupText(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "button-group-text", "px-3 py-2 text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func ButtonGroupSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "button-group-separator", "w-px self-stretch bg-border"), nil)
	})
}
