package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func Breadcrumb(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "nav", attrsFromDOMProps(props, "breadcrumb", ""), templ.GetChildren(ctx))
	})
}

func BreadcrumbList(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "ol", attrsFromDOMProps(props, "breadcrumb-list", "flex flex-wrap items-center gap-1.5 text-sm text-muted-foreground"), templ.GetChildren(ctx))
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
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "breadcrumb-page", "font-normal text-foreground"), templ.GetChildren(ctx))
	})
}

func BreadcrumbSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "li", attrsFromDOMProps(props, "breadcrumb-separator", "mx-1 text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func BreadcrumbEllipsis(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "breadcrumb-ellipsis", "flex h-9 w-9 items-center justify-center"), templ.GetChildren(ctx))
	})
}
