package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type AlertVariant string

const (
	AlertVariantDefault     AlertVariant = "default"
	AlertVariantDestructive AlertVariant = "destructive"
)

type AlertProps struct {
	DOMProps
	Variant AlertVariant
}

func alertClasses(variant AlertVariant, className string) string {
	if variant == "" {
		variant = AlertVariantDefault
	}

	base := "relative grid w-full grid-cols-[0_1fr] items-start gap-y-0.5 rounded-lg border px-4 py-3 text-sm has-[>svg]:grid-cols-[calc(var(--spacing)*4)_1fr] has-[>svg]:gap-x-3 [&>svg]:size-4 [&>svg]:translate-y-0.5 [&>svg]:text-current"
	switch variant {
	case AlertVariantDestructive:
		return cn(base, "bg-card text-destructive *:data-[slot=alert-description]:text-destructive/90 [&>svg]:text-current", className)
	default:
		return cn(base, "bg-card text-card-foreground", className)
	}
}

func Alert(props AlertProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "alert", alertClasses(props.Variant, ""))
		attrs["role"] = "alert"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func AlertTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "alert-title", "col-start-2 line-clamp-1 min-h-4 font-medium tracking-tight"), templ.GetChildren(ctx))
	})
}

func AlertDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "alert-description", "col-start-2 grid justify-items-start gap-1 text-sm text-muted-foreground [&_p]:leading-relaxed"), templ.GetChildren(ctx))
	})
}
