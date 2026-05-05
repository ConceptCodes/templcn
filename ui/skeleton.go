package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func Skeleton(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "skeleton", "animate-pulse rounded-md bg-accent"), templ.GetChildren(ctx))
	})
}
