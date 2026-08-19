package ui

import (
	"context"
	"io"
	"strconv"
	"sync/atomic"

	"github.com/a-h/templ"
)

type DialogProps struct {
	DOMProps
	Open            bool
	DefaultOpen     bool
	Modal           bool
	ShowCloseButton bool
}

type dialogRenderState struct {
	open            bool
	slot            string
	modal           bool
	showCloseButton bool
}

type dialogRenderStateKey struct{}

type dialogAccessibilityState struct {
	contentID     string
	titleID       string
	descriptionID string
}

type dialogAccessibilityStateKey struct{}

var dialogAccessibilitySequence uint64

func prepareDialogAccessibility(ctx context.Context, attrs templ.Attributes, prefix string) (context.Context, templ.Attributes) {
	contentID, _ := attrs["id"].(string)
	if contentID == "" {
		contentID = prefix + "-content-" + strconv.FormatUint(atomic.AddUint64(&dialogAccessibilitySequence, 1), 10)
		attrs["id"] = contentID
	}
	state := dialogAccessibilityState{
		contentID:     contentID,
		titleID:       contentID + "-title",
		descriptionID: contentID + "-description",
	}
	if _, ok := attrs["aria-labelledby"]; !ok {
		attrs["aria-labelledby"] = state.titleID
	}
	if _, ok := attrs["aria-describedby"]; !ok {
		attrs["aria-describedby"] = state.descriptionID
	}
	return context.WithValue(ctx, dialogAccessibilityStateKey{}, state), attrs
}

func dialogAccessibilityFromContext(ctx context.Context) (dialogAccessibilityState, bool) {
	state, ok := ctx.Value(dialogAccessibilityStateKey{}).(dialogAccessibilityState)
	return state, ok
}

func Dialog(props DialogProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		open := props.Open || props.DefaultOpen
		attrs := attrsFromDOMProps(props.DOMProps, "dialog", "")
		attrs["data-state"] = openState(open)
		if open {
			attrs["data-open"] = "true"
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
		ctx = context.WithValue(ctx, dialogRenderStateKey{}, dialogRenderState{open: open, slot: "dialog", modal: props.Modal, showCloseButton: props.ShowCloseButton})
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func DialogTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "dialog-trigger", "")
		attrs["type"] = "button"
		attrs["aria-haspopup"] = "dialog"
		if _, ok := attrs["aria-expanded"]; !ok {
			attrs["aria-expanded"] = "false"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func DialogPortal(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dialog-portal", ""), templ.GetChildren(ctx))
	})
}

func DialogOverlay(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "dialog-overlay", "fixed inset-0 z-50 bg-black/50 data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:animate-in data-[state=open]:fade-in-0")
		attrs["aria-hidden"] = "true"
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = dialogStateFromContext(ctx)
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func DialogContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		ctx, ownChildren := childrenFromContext(ctx)
		attrs := attrsFromDOMProps(props, "dialog-content", "fixed inset-0 z-50 h-dvh max-h-none w-dvw max-w-none overflow-visible border-0 bg-transparent p-0 text-foreground outline-none backdrop:bg-black/50 [&:not([open])]:hidden")
		ctx, attrs = prepareDialogAccessibility(ctx, attrs, "dialog")
		state := dialogStateFromContextValue(ctx)
		attrs["role"] = "dialog"
		attrs["tabindex"] = "-1"
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = openState(state.open)
		}
		if state.open {
			attrs["open"] = true
			attrs["data-open"] = "true"
		}
		if state.modal {
			attrs["data-modal"] = "true"
		}
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			panelAttrs := templ.Attributes{
				"data-dialog-panel": true,
				"class":             "fixed left-[50%] top-[50%] z-50 grid translate-x-[-50%] translate-y-[-50%] gap-4 rounded-lg border bg-background p-6 shadow-lg duration-200 outline-none data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=closed]:zoom-out-95 data-[state=open]:animate-in data-[state=open]:fade-in-0 data-[state=open]:zoom-in-95",
				"data-state":        openState(state.open),
			}
			panel := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				if err := renderChildren(ctx, w, ownChildren); err != nil {
					return err
				}
				if dialogStateFromContextValue(ctx).showCloseButton {
					return renderDialogCloseIcon(ctx, w, "dialog-close")
				}
				return nil
			})
			return renderElement(ctx, w, "div", panelAttrs, panel)
		})
		return renderElement(ctx, w, "dialog", attrs, children)
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
		attrs := attrsFromDOMProps(props, "dialog-title", "text-lg leading-none font-semibold")
		if state, ok := dialogAccessibilityFromContext(ctx); ok {
			if _, exists := attrs["id"]; !exists {
				attrs["id"] = state.titleID
			}
		}
		return renderElement(ctx, w, "h2", attrs, templ.GetChildren(ctx))
	})
}

func DialogDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "dialog-description", "text-sm text-muted-foreground")
		if state, ok := dialogAccessibilityFromContext(ctx); ok {
			if _, exists := attrs["id"]; !exists {
				attrs["id"] = state.descriptionID
			}
		}
		return renderElement(ctx, w, "p", attrs, templ.GetChildren(ctx))
	})
}

func DialogClose(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "dialog-close", "")
		attrs["type"] = "button"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func openState(open bool) string {
	if open {
		return "open"
	}
	return "closed"
}

func dialogStateFromContext(ctx context.Context) string {
	return openState(dialogStateFromContextValue(ctx).open)
}

func dialogStateFromContextValue(ctx context.Context) dialogRenderState {
	state, ok := ctx.Value(dialogRenderStateKey{}).(dialogRenderState)
	if !ok {
		return dialogRenderState{}
	}
	return state
}

func renderDialogCloseIcon(ctx context.Context, w io.Writer, slot string) error {
	attrs := templ.Attributes{
		"type":        "button",
		"data-slot":   slot,
		"aria-label":  "Close",
		"class":       "absolute top-4 right-4 rounded-xs opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:ring-2 focus:ring-ring focus:ring-offset-2 focus:outline-hidden disabled:pointer-events-none data-[state=open]:bg-accent data-[state=open]:text-muted-foreground [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4",
		"data-action": "close",
	}
	icon := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, `<svg xmlns="http://www.w3.org/2000/svg" class="size-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M18 6 6 18"></path><path d="m6 6 12 12"></path></svg>`)
		return err
	})
	return renderElement(ctx, w, "button", attrs, icon)
}
