package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func Empty(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "empty", "flex min-w-0 flex-1 flex-col items-center justify-center gap-6 rounded-lg border-dashed p-6 text-center text-balance md:p-12"), templ.GetChildren(ctx))
	})
}

func EmptyHeader(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "empty-header", "flex max-w-sm flex-col items-center gap-2 text-center"), templ.GetChildren(ctx))
	})
}
func EmptyMedia(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		variant, _ := props.Attrs["data-variant"].(string)
		className := "mb-2 flex shrink-0 items-center justify-center bg-transparent [&_svg]:pointer-events-none [&_svg]:shrink-0"
		if variant == "icon" {
			className = "mb-2 flex size-10 shrink-0 items-center justify-center rounded-lg bg-muted text-foreground [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-6"
		}
		attrs := attrsFromDOMProps(props, "empty-icon", className)
		if _, ok := attrs["data-variant"]; !ok {
			attrs["data-variant"] = "default"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func EmptyTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "empty-title", "text-lg font-medium tracking-tight"), templ.GetChildren(ctx))
	})
}
func EmptyDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "empty-description", "text-sm/relaxed text-muted-foreground [&>a]:underline [&>a]:underline-offset-4 [&>a:hover]:text-primary"), templ.GetChildren(ctx))
	})
}
func EmptyContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "empty-content", "flex w-full max-w-sm min-w-0 flex-col items-center gap-4 text-sm text-balance"), templ.GetChildren(ctx))
	})
}
