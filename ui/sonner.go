package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type ToasterProps struct {
	DOMProps
	Theme      string
	Position   string
	RichColors bool
	Duration   int
}

func Toaster(props ToasterProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "toaster", "fixed inset-0 z-50")
		if props.Theme != "" {
			attrs["data-theme"] = props.Theme
		}
		if props.Position != "" {
			attrs["data-position"] = props.Position
		}
		if props.RichColors {
			attrs["data-rich-colors"] = "true"
		}
		if props.Duration > 0 {
			attrs["data-duration"] = props.Duration
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
