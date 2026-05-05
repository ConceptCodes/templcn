package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type FieldOrientation string

const (
	FieldOrientationVertical   FieldOrientation = "vertical"
	FieldOrientationHorizontal FieldOrientation = "horizontal"
	FieldOrientationResponsive FieldOrientation = "responsive"
)

type FieldSetProps struct {
	DOMProps
	Orientation FieldOrientation
	Invalid     bool
	Disabled    bool
}

func FieldSet(props FieldSetProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		orientation := props.Orientation
		if orientation == "" {
			orientation = FieldOrientationVertical
		}
		attrs := attrsFromDOMProps(props.DOMProps, "fieldset", cn("grid gap-2 rounded-lg border border-border p-4", map[FieldOrientation]string{
			FieldOrientationVertical:   "grid",
			FieldOrientationHorizontal: "grid gap-4 md:grid-cols-[180px_1fr]",
			FieldOrientationResponsive: "grid gap-4 md:grid-cols-[180px_1fr]",
		}[orientation]))
		if props.Invalid {
			attrs["data-invalid"] = "true"
		}
		if props.Disabled {
			attrs["disabled"] = true
		}
		return renderElement(ctx, w, "fieldset", attrs, templ.GetChildren(ctx))
	})
}

func FieldLegend(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "legend", attrsFromDOMProps(props, "field-legend", "px-1 text-sm font-medium"), templ.GetChildren(ctx))
	})
}

func FieldGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "field-group", "grid gap-4"), templ.GetChildren(ctx))
	})
}

func Field(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "field", "grid gap-2")
		attrs["role"] = "group"
		attrs["data-orientation"] = "vertical"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func FieldContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "field-content", "flex flex-1 flex-col gap-1.5 leading-snug"), templ.GetChildren(ctx))
	})
}

func FieldLabel(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "label", attrsFromDOMProps(props, "field-label", "text-sm font-medium"), templ.GetChildren(ctx))
	})
}

func FieldTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "field-title", "text-sm font-medium"), templ.GetChildren(ctx))
	})
}

func FieldDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "field-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func FieldSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderVoidElement(ctx, w, "hr", attrsFromDOMProps(props, "field-separator", "my-2 border-border"))
	})
}

type FieldErrorProps struct {
	DOMProps
	Errors []string
}

func FieldError(props FieldErrorProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		text := ""
		if len(props.Errors) > 0 {
			text = props.Errors[0]
		}
		attrs := attrsFromDOMProps(props.DOMProps, "field-error", "text-sm text-destructive")
		attrs["role"] = "alert"
		return renderTextElement(ctx, w, "p", attrs, text)
	})
}
