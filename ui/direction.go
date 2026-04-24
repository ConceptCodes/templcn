package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type DirectionProviderProps struct {
	DOMProps
	Direction string
}

func DirectionProvider(props DirectionProviderProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "direction-provider", "")
		if props.Direction != "" {
			attrs["dir"] = props.Direction
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
