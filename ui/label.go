package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type LabelProps struct {
	DOMProps
	For string
}

func Label(props LabelProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "label", "text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70")
		if props.For != "" {
			attrs["for"] = props.For
		}
		return renderElement(ctx, w, "label", attrs, templ.GetChildren(ctx))
	})
}
