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
		attrs := attrsFromDOMProps(props.DOMProps, "label", "flex items-center gap-2 text-sm leading-none font-medium select-none group-data-[disabled=true]:pointer-events-none group-data-[disabled=true]:opacity-50 peer-disabled:cursor-not-allowed peer-disabled:opacity-50")
		if props.For != "" {
			attrs["for"] = props.For
		}
		return renderElement(ctx, w, "label", attrs, templ.GetChildren(ctx))
	})
}
