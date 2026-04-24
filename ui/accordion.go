package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type AccordionProps struct {
	DOMProps
	Type         string
	Value        []string
	DefaultValue []string
	Collapsible  bool
}

type AccordionItemProps struct {
	DOMProps
	Value    string
	Disabled bool
}

func Accordion(props AccordionProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "accordion", "grid gap-2")
		if props.Type != "" {
			attrs["data-type"] = props.Type
		}
		if len(props.Value) > 0 {
			attrs["data-value"] = props.Value
		}
		if len(props.DefaultValue) > 0 {
			attrs["data-default-value"] = props.DefaultValue
		}
		if props.Collapsible {
			attrs["data-collapsible"] = "true"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func AccordionItem(props AccordionItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "accordion-item", "rounded-md border border-border")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.Disabled {
			attrs["data-disabled"] = "true"
		}
		return renderElement(ctx, w, "details", attrs, templ.GetChildren(ctx))
	})
}

func AccordionTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "summary", attrsFromDOMProps(props, "accordion-trigger", "flex cursor-pointer items-center justify-between gap-4 px-4 py-3 text-sm font-medium outline-none"), templ.GetChildren(ctx))
	})
}

func AccordionContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "accordion-content", "px-4 pb-4 pt-0 text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}
