package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func Empty(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "empty", "grid place-items-center gap-4 rounded-xl border border-dashed p-8 text-center"), templ.GetChildren(ctx))
	})
}

func EmptyHeader(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "empty-header", "grid gap-1"), templ.GetChildren(ctx))
	})
}
func EmptyMedia(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "empty-media", "flex size-12 items-center justify-center rounded-full bg-muted"), templ.GetChildren(ctx))
	})
}
func EmptyTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "h3", attrsFromDOMProps(props, "empty-title", "text-base font-semibold"), templ.GetChildren(ctx))
	})
}
func EmptyDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "p", attrsFromDOMProps(props, "empty-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}
func EmptyContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "empty-content", ""), templ.GetChildren(ctx))
	})
}
