package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func Table(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "table", attrsFromDOMProps(props, "table", "w-full caption-bottom text-sm"), templ.GetChildren(ctx))
	})
}
func TableHeader(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "thead", attrsFromDOMProps(props, "table-header", "[&_tr]:border-b"), templ.GetChildren(ctx))
	})
}
func TableBody(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "tbody", attrsFromDOMProps(props, "table-body", "[&_tr:last-child]:border-0"), templ.GetChildren(ctx))
	})
}
func TableFooter(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "tfoot", attrsFromDOMProps(props, "table-footer", "border-t bg-muted/50 font-medium [&>tr]:last:border-b-0"), templ.GetChildren(ctx))
	})
}
func TableHead(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "th", attrsFromDOMProps(props, "table-head", "h-12 px-4 text-left align-middle font-medium text-muted-foreground"), templ.GetChildren(ctx))
	})
}
func TableRow(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "tr", attrsFromDOMProps(props, "table-row", "border-b transition-colors hover:bg-muted/50"), templ.GetChildren(ctx))
	})
}
func TableCell(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "td", attrsFromDOMProps(props, "table-cell", "p-4 align-middle"), templ.GetChildren(ctx))
	})
}
func TableCaption(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "caption", attrsFromDOMProps(props, "table-caption", "mt-4 text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}
