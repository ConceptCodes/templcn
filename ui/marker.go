package ui

import (
	"context"
	"github.com/a-h/templ"
	"io"
)

type MarkerProps struct {
	DOMProps
	Variant string
	Element string
}

func Marker(props MarkerProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		variant, tag := props.Variant, props.Element
		if variant == "" {
			variant = "default"
		}
		if tag == "" {
			tag = "div"
		}
		className := "group/marker relative flex min-h-4 w-full items-center gap-2 text-left text-sm text-muted-foreground"
		if variant == "separator" {
			className = cn(className, "before:mr-1 before:h-px before:min-w-0 before:flex-1 before:bg-border after:ml-1 after:h-px after:min-w-0 after:flex-1 after:bg-border")
		} else if variant == "border" {
			className = cn(className, "border-b border-border pb-2")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "marker", className)
		attrs["data-variant"] = variant
		return renderElement(ctx, w, tag, attrs, templ.GetChildren(ctx))
	})
}
func MarkerIcon(props DOMProps) templ.Component {
	attrs := props
	attrs.Attrs = cloneAttributes(props.Attrs)
	attrs.Attrs["aria-hidden"] = true
	return markerBlock("marker-icon", "size-4 shrink-0", attrs)
}
func MarkerContent(props DOMProps) templ.Component {
	return markerBlock("marker-content", "min-w-0 wrap-break-word", props)
}
func markerBlock(slot, className string, props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, slot, className), templ.GetChildren(ctx))
	})
}
