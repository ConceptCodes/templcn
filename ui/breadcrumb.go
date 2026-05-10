package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func Breadcrumb(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "breadcrumb", "")
		if _, ok := attrs["aria-label"]; !ok {
			attrs["aria-label"] = "breadcrumb"
		}
		return renderElement(ctx, w, "nav", attrs, templ.GetChildren(ctx))
	})
}

func BreadcrumbList(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "ol", attrsFromDOMProps(props, "breadcrumb-list", "flex flex-wrap items-center gap-1.5 text-sm break-words text-muted-foreground sm:gap-2.5"), templ.GetChildren(ctx))
	})
}

func BreadcrumbItem(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "li", attrsFromDOMProps(props, "breadcrumb-item", "inline-flex items-center gap-1.5"), templ.GetChildren(ctx))
	})
}

type BreadcrumbLinkProps struct {
	DOMProps
	Href string
}

func BreadcrumbLink(props BreadcrumbLinkProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "breadcrumb-link", "transition-colors hover:text-foreground")
		if props.Href != "" {
			attrs["href"] = props.Href
		}
		return renderElement(ctx, w, "a", attrs, templ.GetChildren(ctx))
	})
}

func BreadcrumbPage(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "breadcrumb-page", "font-normal text-foreground")
		attrs["role"] = "link"
		attrs["aria-disabled"] = "true"
		attrs["aria-current"] = "page"
		return renderElement(ctx, w, "span", attrs, templ.GetChildren(ctx))
	})
}

func BreadcrumbSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "breadcrumb-separator", "[&>svg]:size-3.5")
		attrs["role"] = "presentation"
		attrs["aria-hidden"] = "true"
		children := templ.GetChildren(ctx)
		if children == nil {
			children = templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				_, err := io.WriteString(w, `<svg xmlns="http://www.w3.org/2000/svg" class="size-3.5" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="m9 18 6-6-6-6"></path></svg>`)
				return err
			})
		}
		return renderElement(ctx, w, "li", attrs, children)
	})
}

func BreadcrumbEllipsis(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "breadcrumb-ellipsis", "flex size-9 items-center justify-center")
		attrs["role"] = "presentation"
		attrs["aria-hidden"] = "true"
		children := templ.GetChildren(ctx)
		if children == nil {
			children = templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				_, err := io.WriteString(w, `<svg xmlns="http://www.w3.org/2000/svg" class="size-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><circle cx="12" cy="12" r="1"></circle><circle cx="19" cy="12" r="1"></circle><circle cx="5" cy="12" r="1"></circle></svg><span class="sr-only">More</span>`)
				return err
			})
		}
		return renderElement(ctx, w, "span", attrs, children)
	})
}
