package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func Pagination(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "pagination", "mx-auto flex w-full justify-center")
		attrs["role"] = "navigation"
		if _, ok := attrs["aria-label"]; !ok {
			attrs["aria-label"] = "pagination"
		}
		return renderElement(ctx, w, "nav", attrs, templ.GetChildren(ctx))
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
		size := ButtonSizeIcon
		if props.Size != "" {
			size = ButtonSize(props.Size)
		}
		variant := ButtonVariantGhost
		if props.IsActive {
			variant = ButtonVariantOutline
		}
		className := buttonClasses(variant, size, "")
		attrs := attrsFromDOMProps(props.DOMProps, "pagination-link", className)
		attrs["data-active"] = props.IsActive
		if props.IsActive {
			attrs["aria-current"] = "page"
		}
		if props.Href != "" {
			attrs["href"] = props.Href
		}
		return renderElement(ctx, w, "a", attrs, templ.GetChildren(ctx))
	})
}

func PaginationPrevious(props PaginationLinkProps) templ.Component {
	props.Size = string(ButtonSizeDefault)
	props.Class = cn("gap-1 px-2.5 sm:pl-2.5", props.Class)
	if props.Attrs == nil {
		props.Attrs = templ.Attributes{}
	}
	if _, ok := props.Attrs["aria-label"]; !ok {
		props.Attrs["aria-label"] = "Go to previous page"
	}
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		ctx, callerChildren := childrenFromContext(ctx)
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := renderElement(ctx, w, "svg", templ.Attributes{
				"xmlns":           "http://www.w3.org/2000/svg",
				"viewBox":         "0 0 24 24",
				"fill":            "none",
				"stroke":          "currentColor",
				"stroke-width":    "2",
				"stroke-linecap":  "round",
				"stroke-linejoin": "round",
				"aria-hidden":     "true",
			}, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				_, err := io.WriteString(w, `<path d="m15 18-6-6 6-6"></path>`)
				return err
			})); err != nil {
				return err
			}
			if err := renderElement(ctx, w, "span", templ.Attributes{"class": "hidden sm:block"}, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				_, err := io.WriteString(w, "Previous")
				return err
			})); err != nil {
				return err
			}
			return renderChildren(ctx, w, callerChildren)
		})
		return PaginationLink(props).Render(templ.WithChildren(ctx, children), w)
	})
}

func PaginationNext(props PaginationLinkProps) templ.Component {
	props.Size = string(ButtonSizeDefault)
	props.Class = cn("gap-1 px-2.5 sm:pr-2.5", props.Class)
	if props.Attrs == nil {
		props.Attrs = templ.Attributes{}
	}
	if _, ok := props.Attrs["aria-label"]; !ok {
		props.Attrs["aria-label"] = "Go to next page"
	}
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		ctx, callerChildren := childrenFromContext(ctx)
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := renderElement(ctx, w, "span", templ.Attributes{"class": "hidden sm:block"}, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				_, err := io.WriteString(w, "Next")
				return err
			})); err != nil {
				return err
			}
			if err := renderElement(ctx, w, "svg", templ.Attributes{
				"xmlns":           "http://www.w3.org/2000/svg",
				"viewBox":         "0 0 24 24",
				"fill":            "none",
				"stroke":          "currentColor",
				"stroke-width":    "2",
				"stroke-linecap":  "round",
				"stroke-linejoin": "round",
				"aria-hidden":     "true",
			}, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				_, err := io.WriteString(w, `<path d="m9 18 6-6-6-6"></path>`)
				return err
			})); err != nil {
				return err
			}
			return renderChildren(ctx, w, callerChildren)
		})
		return PaginationLink(props).Render(templ.WithChildren(ctx, children), w)
	})
}

func PaginationEllipsis(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "pagination-ellipsis", "flex size-9 items-center justify-center")
		attrs["aria-hidden"] = true
		children := templ.GetChildren(ctx)
		if children == nil {
			children = templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				_, err := io.WriteString(w, `<svg xmlns="http://www.w3.org/2000/svg" class="size-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><circle cx="12" cy="12" r="1"></circle><circle cx="19" cy="12" r="1"></circle><circle cx="5" cy="12" r="1"></circle></svg><span class="sr-only">More pages</span>`)
				return err
			})
		}
		return renderElement(ctx, w, "span", attrs, children)
	})
}
