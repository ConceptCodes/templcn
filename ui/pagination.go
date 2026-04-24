package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func Pagination(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "nav", attrsFromDOMProps(props, "pagination", "mx-auto flex w-full justify-center"), templ.GetChildren(ctx))
	})
}

func PaginationContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "ul", attrsFromDOMProps(props, "pagination-content", "flex flex-row items-center gap-1"), templ.GetChildren(ctx))
	})
}

func PaginationItem(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "li", attrsFromDOMProps(props, "pagination-item", ""), templ.GetChildren(ctx))
	})
}

type PaginationLinkProps struct {
	DOMProps
	Href     string
	IsActive bool
	Size     string
}

func PaginationLink(props PaginationLinkProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "inline-flex h-9 min-w-9 items-center justify-center rounded-md border border-input bg-background px-3 text-sm shadow-xs transition-colors hover:bg-accent hover:text-accent-foreground"
		if props.IsActive {
			className = cn(className, "bg-primary text-primary-foreground hover:bg-primary/90")
		}
		if props.Size == "sm" {
			className = cn(className, "h-8 min-w-8 px-2")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "pagination-link", className)
		if props.Href != "" {
			attrs["href"] = props.Href
		}
		return renderElement(ctx, w, "a", attrs, templ.GetChildren(ctx))
	})
}

func PaginationPrevious(props PaginationLinkProps) templ.Component { return PaginationLink(props) }
func PaginationNext(props PaginationLinkProps) templ.Component     { return PaginationLink(props) }

func PaginationEllipsis(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "pagination-ellipsis", "inline-flex h-9 w-9 items-center justify-center"), templ.GetChildren(ctx))
	})
}
