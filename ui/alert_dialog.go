package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type AlertDialogProps struct {
	DialogProps
}

func AlertDialog(props AlertDialogProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "alert-dialog", "")
		attrs["data-state"] = openState(props.Open || props.DefaultOpen)
		if props.Open || props.DefaultOpen {
			attrs["data-open"] = "true"
		}
		if props.Open {
			attrs["open"] = true
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		attrs["data-modal"] = "true"
		return renderElement(ctx, w, "dialog", attrs, templ.GetChildren(ctx))
	})
}

func AlertDialogTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "alert-dialog-trigger", ""), templ.GetChildren(ctx))
	})
}

func AlertDialogPortal(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "alert-dialog-portal", ""), templ.GetChildren(ctx))
	})
}

func AlertDialogOverlay(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "alert-dialog-overlay", "fixed inset-0 z-50 bg-black/80 data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0")
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = "open"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func AlertDialogContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "alert-dialog-content", "fixed left-1/2 top-1/2 z-50 grid w-full max-w-lg -translate-x-1/2 -translate-y-1/2 gap-4 rounded-lg border bg-background p-6 shadow-lg duration-200 data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95")
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = "open"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func AlertDialogHeader(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "alert-dialog-header", "flex flex-col gap-2 text-center sm:text-left"), templ.GetChildren(ctx))
	})
}

func AlertDialogFooter(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "alert-dialog-footer", "flex flex-col-reverse gap-2 sm:flex-row sm:justify-end"), templ.GetChildren(ctx))
	})
}

func AlertDialogTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "h2", attrsFromDOMProps(props, "alert-dialog-title", "text-lg font-semibold"), templ.GetChildren(ctx))
	})
}

func AlertDialogDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "p", attrsFromDOMProps(props, "alert-dialog-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func AlertDialogAction(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "alert-dialog-action", buttonClasses(ButtonVariantDestructive, ButtonSizeDefault, "")), templ.GetChildren(ctx))
	})
}

func AlertDialogCancel(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "alert-dialog-cancel", buttonClasses(ButtonVariantOutline, ButtonSizeDefault, "")), templ.GetChildren(ctx))
	})
}
