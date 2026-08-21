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
		open := props.Open || props.DefaultOpen
		attrs["role"] = "status"
		attrs["aria-live"] = "polite"
		attrs["data-state"] = openState(open)
		if open {
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
		attrs := attrsFromDOMProps(props, "toast-action", "")
		attrs["type"] = "button"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func ToastClose(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "toast-close", "absolute right-2 top-2 rounded-md p-1 text-foreground/50 opacity-70 transition-opacity hover:opacity-100 focus:opacity-100 focus:outline-none cursor-pointer")
		attrs["type"] = "button"
		attrs["aria-label"] = "Close toast"
		children := templ.GetChildren(ctx)
		if children == nil {
			children = templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				_, err := io.WriteString(w, `<svg class="size-4 pointer-events-none" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path></svg>`)
				return err
			})
		}
		return renderElement(ctx, w, "button", attrs, children)
	})
}
