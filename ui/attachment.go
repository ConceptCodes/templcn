package ui

import (
	"context"
	"io"
)

import "github.com/a-h/templ"

type AttachmentProps struct {
	DOMProps
	State       string
	Size        string
	Orientation string
}

func Attachment(props AttachmentProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		state, size, orientation := props.State, props.Size, props.Orientation
		if state == "" {
			state = "done"
		}
		if size == "" {
			size = "default"
		}
		if orientation == "" {
			orientation = "horizontal"
		}
		className := "group/attachment relative flex w-fit max-w-full min-w-0 shrink-0 flex-wrap rounded-xl border bg-card text-card-foreground transition-colors focus-within:ring-1 focus-within:ring-ring/50 has-[>a,>button]:hover:bg-muted/50 data-[state=error]:border-destructive/30 data-[state=idle]:border-dashed"
		if orientation == "horizontal" {
			className = cn(className, "min-w-40 items-center")
		} else {
			className = cn(className, "w-24 flex-col")
		}
		if size == "sm" {
			className = cn(className, "gap-2.5 text-xs")
		} else if size == "xs" {
			className = cn(className, "gap-1.5 rounded-lg text-xs")
		} else {
			className = cn(className, "gap-2 text-sm")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "attachment", className)
		attrs["data-state"], attrs["data-size"], attrs["data-orientation"] = state, size, orientation
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func AttachmentMedia(props DOMProps) templ.Component {
	return attachmentBlock("attachment-media", "relative flex aspect-square w-10 shrink-0 items-center justify-center overflow-hidden rounded-lg bg-muted text-foreground", props)
}
func AttachmentContent(props DOMProps) templ.Component {
	return attachmentBlock("attachment-content", "max-w-full min-w-0 flex-1 leading-tight", props)
}
func AttachmentTitle(props DOMProps) templ.Component {
	return attachmentBlock("attachment-title", "block max-w-full min-w-0 truncate font-medium", props)
}
func AttachmentDescription(props DOMProps) templ.Component {
	return attachmentBlock("attachment-description", "mt-0.5 block min-w-0 max-w-full truncate text-xs text-muted-foreground", props)
}
func AttachmentActions(props DOMProps) templ.Component {
	return attachmentBlock("attachment-actions", "relative z-20 flex shrink-0 items-center", props)
}
func AttachmentGroup(props DOMProps) templ.Component {
	return attachmentBlock("attachment-group", "flex min-w-0 snap-x snap-mandatory gap-3 overflow-x-auto overscroll-x-contain py-1", props)
}

func AttachmentAction(props ButtonProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		variant, size := props.Variant, props.Size
		if variant == "" {
			variant = ButtonVariantGhost
		}
		if size == "" {
			size = ButtonSizeIconXS
		}
		attrs := attrsFromDOMProps(props.DOMProps, "attachment-action", buttonClasses(variant, size, ""))
		attrs["type"] = "button"
		if props.Disabled {
			attrs["disabled"] = true
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func AttachmentTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		tag := props.Element
		if tag == "" {
			tag = "button"
		}
		attrs := attrsFromDOMProps(props, "attachment-trigger", "absolute inset-0 z-10 outline-none")
		if tag == "button" {
			attrs["type"] = "button"
		}
		return renderElement(ctx, w, tag, attrs, templ.GetChildren(ctx))
	})
}

func attachmentBlock(slot, className string, props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, slot, className), templ.GetChildren(ctx))
	})
}
