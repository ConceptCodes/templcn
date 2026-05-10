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
		open := props.Open || props.DefaultOpen
		attrs := attrsFromDOMProps(props.DOMProps, "alert-dialog", "")
		attrs["data-state"] = openState(open)
		if open {
			attrs["data-open"] = "true"
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		attrs["data-modal"] = "true"
		ctx = context.WithValue(ctx, dialogRenderStateKey{}, dialogRenderState{open: open, slot: "alert-dialog", modal: true})
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func AlertDialogTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "alert-dialog-trigger", "")
		attrs["type"] = "button"
		attrs["aria-haspopup"] = "dialog"
		if _, ok := attrs["aria-expanded"]; !ok {
			attrs["aria-expanded"] = "false"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func AlertDialogPortal(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "alert-dialog-portal", ""), templ.GetChildren(ctx))
	})
}

func AlertDialogOverlay(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "alert-dialog-overlay", "fixed inset-0 z-50 bg-black/50 data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:animate-in data-[state=open]:fade-in-0")
		attrs["aria-hidden"] = "true"
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = dialogStateFromContext(ctx)
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func AlertDialogContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "alert-dialog-content", "group/alert-dialog-content fixed top-[50%] left-[50%] z-50 grid w-full max-w-[calc(100%-2rem)] translate-x-[-50%] translate-y-[-50%] gap-4 rounded-lg border bg-background p-6 shadow-lg duration-200 backdrop:bg-black/50 [&:not([open])]:hidden data-[size=sm]:max-w-xs data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=closed]:zoom-out-95 data-[state=open]:animate-in data-[state=open]:fade-in-0 data-[state=open]:zoom-in-95 data-[size=default]:sm:max-w-lg")
		state := dialogStateFromContextValue(ctx)
		attrs["role"] = "alertdialog"
		attrs["aria-modal"] = "true"
		attrs["tabindex"] = "-1"
		if _, ok := attrs["data-size"]; !ok {
			attrs["data-size"] = "default"
		}
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = openState(state.open)
		}
		if state.open {
			attrs["open"] = true
			attrs["data-open"] = "true"
		}
		attrs["data-modal"] = "true"
		return renderElement(ctx, w, "dialog", attrs, templ.GetChildren(ctx))
	})
}

func AlertDialogHeader(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "alert-dialog-header", "grid grid-rows-[auto_1fr] place-items-center gap-1.5 text-center has-data-[slot=alert-dialog-media]:grid-rows-[auto_auto_1fr] has-data-[slot=alert-dialog-media]:gap-x-6 sm:group-data-[size=default]/alert-dialog-content:place-items-start sm:group-data-[size=default]/alert-dialog-content:text-left sm:group-data-[size=default]/alert-dialog-content:has-data-[slot=alert-dialog-media]:grid-rows-[auto_1fr]"), templ.GetChildren(ctx))
	})
}

func AlertDialogFooter(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "alert-dialog-footer", "flex flex-col-reverse gap-2 group-data-[size=sm]/alert-dialog-content:grid group-data-[size=sm]/alert-dialog-content:grid-cols-2 sm:flex-row sm:justify-end"), templ.GetChildren(ctx))
	})
}

func AlertDialogTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "h2", attrsFromDOMProps(props, "alert-dialog-title", "text-lg font-semibold sm:group-data-[size=default]/alert-dialog-content:group-has-data-[slot=alert-dialog-media]/alert-dialog-content:col-start-2"), templ.GetChildren(ctx))
	})
}

func AlertDialogDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "p", attrsFromDOMProps(props, "alert-dialog-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func AlertDialogMedia(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "alert-dialog-media", "mb-2 inline-flex size-16 items-center justify-center rounded-md bg-muted sm:group-data-[size=default]/alert-dialog-content:row-span-2 *:[svg:not([class*='size-'])]:size-8"), templ.GetChildren(ctx))
	})
}

func AlertDialogAction(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "alert-dialog-action", buttonClasses(ButtonVariantDestructive, ButtonSizeDefault, ""))
		attrs["type"] = "button"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func AlertDialogCancel(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "alert-dialog-cancel", buttonClasses(ButtonVariantOutline, ButtonSizeDefault, ""))
		attrs["type"] = "button"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
