package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type FormProps struct {
	DOMProps
	Name        string
	IDValue     string
	DescribedBy string
	Invalid     bool
	Message     string
}

func Form(props FormProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "form", "grid gap-4")
		if props.IDValue != "" {
			attrs["id"] = props.IDValue
		}
		if props.Name != "" {
			attrs["name"] = props.Name
		}
		if props.DescribedBy != "" {
			attrs["aria-describedby"] = props.DescribedBy
		}
		if props.Invalid {
			attrs["aria-invalid"] = "true"
		}
		if props.Message != "" {
			attrs["data-message"] = props.Message
		}
		tag := "form"
		if props.Element != "" {
			tag = props.Element
		}
		return renderElement(ctx, w, tag, attrs, templ.GetChildren(ctx))
	})
}

func FormItem(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "form-item", "grid gap-2"), templ.GetChildren(ctx))
	})
}

func FormLabel(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "label", attrsFromDOMProps(props, "form-label", "text-sm font-medium leading-none"), templ.GetChildren(ctx))
	})
}

func FormControl(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "form-control", "grid gap-1"), templ.GetChildren(ctx))
	})
}

func FormDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "p", attrsFromDOMProps(props, "form-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func FormMessage(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "form-message", "text-sm text-destructive")
		return renderElement(ctx, w, "p", attrs, templ.GetChildren(ctx))
	})
}
