package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type DialogProps struct {
	DOMProps
	Open            bool
	DefaultOpen     bool
	Modal           bool
	ShowCloseButton bool
}

func Dialog(props DialogProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "dialog", "fixed inset-0 z-50 grid place-items-center")
		if props.Open {
			attrs["open"] = true
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		if props.Modal {
			attrs["data-modal"] = "true"
		}
		if props.ShowCloseButton {
			attrs["data-show-close-button"] = "true"
		}
		return renderElement(ctx, w, "dialog", attrs, templ.GetChildren(ctx))
	})
}

func DialogTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "dialog-trigger", ""), templ.GetChildren(ctx))
	})
}

func DialogPortal(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dialog-portal", ""), templ.GetChildren(ctx))
	})
}

func DialogOverlay(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dialog-overlay", "fixed inset-0 bg-black/50"), templ.GetChildren(ctx))
	})
}

func DialogContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dialog-content", "relative z-50 w-full max-w-lg rounded-xl border bg-background p-6 shadow-lg"), templ.GetChildren(ctx))
	})
}

func DialogHeader(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dialog-header", "flex flex-col gap-2 text-center sm:text-left"), templ.GetChildren(ctx))
	})
}

func DialogFooter(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dialog-footer", "flex flex-col-reverse gap-2 sm:flex-row sm:justify-end"), templ.GetChildren(ctx))
	})
}

func DialogTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "h2", attrsFromDOMProps(props, "dialog-title", "text-lg font-semibold tracking-tight"), templ.GetChildren(ctx))
	})
}

func DialogDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "p", attrsFromDOMProps(props, "dialog-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func DialogClose(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "dialog-close", ""), templ.GetChildren(ctx))
	})
}
