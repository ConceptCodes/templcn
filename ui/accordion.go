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
		attrs := attrsFromDOMProps(props.DOMProps, "accordion-item", "border-b")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		attrs["data-state"] = "closed"
		if props.Disabled {
			attrs["data-disabled"] = "true"
		}
		return renderElement(ctx, w, "details", attrs, templ.GetChildren(ctx))
	})
}

func AccordionTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := renderChildren(ctx, w, templ.GetChildren(ctx)); err != nil {
				return err
			}
			_, err := io.WriteString(w, `<svg xmlns="http://www.w3.org/2000/svg" class="size-4 shrink-0 text-muted-foreground transition-transform duration-200 group-data-[state=open]:rotate-180" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="m6 9 6 6 6-6"></path></svg>`)
			return err
		})
		return renderElement(ctx, w, "summary", attrsFromDOMProps(props, "accordion-trigger", "group flex flex-1 cursor-pointer list-none items-start justify-between gap-4 rounded-md py-4 text-left text-sm font-medium transition-all outline-none hover:underline focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:pointer-events-none disabled:opacity-50"), children)
	})
}

func AccordionContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "accordion-content", "overflow-hidden text-sm data-[state=closed]:animate-accordion-up data-[state=open]:animate-accordion-down")
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = "open"
		}
		inner := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return renderElement(ctx, w, "div", templ.Attributes{"class": "pb-4 pt-0"}, templ.GetChildren(ctx))
		})
		return renderElement(ctx, w, "div", attrs, inner)
	})
}
