package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type SeparatorOrientation string

const (
	SeparatorOrientationHorizontal SeparatorOrientation = "horizontal"
	SeparatorOrientationVertical   SeparatorOrientation = "vertical"
)

type SeparatorProps struct {
	DOMProps
	Orientation SeparatorOrientation
	Decorative  bool
}

func Separator(props SeparatorProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		orientation := props.Orientation
		if orientation == "" {
			orientation = SeparatorOrientationHorizontal
		}

		className := "shrink-0 bg-border"
		if orientation == SeparatorOrientationHorizontal {
			className = cn(className, "h-px w-full")
		} else {
			className = cn(className, "h-full w-px")
		}

		attrs := attrsFromDOMProps(props.DOMProps, "separator", className)
		if !props.Decorative {
			attrs["role"] = "separator"
		} else {
			attrs["aria-hidden"] = "true"
		}
		return renderVoidElement(ctx, w, "hr", attrs)
	})
}
