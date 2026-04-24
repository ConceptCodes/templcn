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

	base := "relative w-full rounded-xl border px-4 py-3 text-sm grid gap-2"
	switch variant {
	case AlertVariantDestructive:
		return cn(base, "border-destructive/50 text-destructive dark:border-destructive [&_svg]:text-destructive", className)
	default:
		return cn(base, "bg-background text-foreground", className)
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
		return renderElement(ctx, w, "h5", attrsFromDOMProps(props, "alert-title", "mb-1 font-medium leading-none tracking-tight"), templ.GetChildren(ctx))
	})
}

func AlertDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "alert-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}
