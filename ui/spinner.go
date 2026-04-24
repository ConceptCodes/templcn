package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type SpinnerSize string

const (
	SpinnerSizeSM SpinnerSize = "sm"
	SpinnerSizeMD SpinnerSize = "md"
	SpinnerSizeLG SpinnerSize = "lg"
)

type SpinnerProps struct {
	DOMProps
	Size  SpinnerSize
	Label string
}

func Spinner(props SpinnerProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		size := props.Size
		if size == "" {
			size = SpinnerSizeMD
		}

		dimensions := map[SpinnerSize]string{
			SpinnerSizeSM: "size-4",
			SpinnerSizeMD: "size-5",
			SpinnerSizeLG: "size-6",
		}[size]

		attrs := attrsFromDOMProps(props.DOMProps, "spinner", cn("inline-flex items-center justify-center animate-spin text-muted-foreground", dimensions))
		if props.Label != "" {
			attrs["role"] = "status"
			attrs["aria-label"] = props.Label
		} else {
			attrs["aria-hidden"] = "true"
		}

		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return renderVoidElement(ctx, w, "svg", templ.Attributes{
				"viewBox":     "0 0 24 24",
				"fill":        "none",
				"aria-hidden": "true",
			})
		})

		return renderElement(ctx, w, "div", attrs, children)
	})
}
