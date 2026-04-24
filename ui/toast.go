package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type ToastProviderProps struct {
	DOMProps
}

func ToastProvider(props ToastProviderProps) templ.Component {
	return Toaster(ToasterProps{DOMProps: props.DOMProps})
}

func ToastViewport(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "toast-viewport", "fixed bottom-0 right-0 z-50 flex flex-col gap-2 p-4"), templ.GetChildren(ctx))
	})
}

type ToastProps struct {
	DOMProps
	Open        bool
	DefaultOpen bool
	Variant     string
	Duration    int
}

func Toast(props ToastProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "toast", "grid gap-2 rounded-lg border bg-background p-4 shadow-lg")
		if props.Open {
			attrs["data-open"] = "true"
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		if props.Variant != "" {
			attrs["data-variant"] = props.Variant
		}
		if props.Duration > 0 {
			attrs["data-duration"] = props.Duration
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func ToastTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "toast-title", "font-semibold"), templ.GetChildren(ctx))
	})
}

func ToastDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "toast-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func ToastAction(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "toast-action", ""), templ.GetChildren(ctx))
	})
}

func ToastClose(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "toast-close", ""), templ.GetChildren(ctx))
	})
}
