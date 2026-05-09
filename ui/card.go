package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func cardWrapper(props DOMProps, slot, className string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, slot, className), templ.GetChildren(ctx))
	})
}

func Card(props DOMProps) templ.Component {
	return cardWrapper(props, "card", "flex min-w-0 flex-col gap-6 rounded-xl border bg-card py-6 text-card-foreground shadow-sm")
}
func CardHeader(props DOMProps) templ.Component {
	return cardWrapper(props, "card-header", "@container/card-header grid min-w-0 auto-rows-min grid-rows-[auto_auto] items-start gap-2 px-6 has-data-[slot=card-action]:grid-cols-[1fr_auto] [.border-b]:pb-6")
}
func CardTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "card-title", "leading-none font-semibold"), templ.GetChildren(ctx))
	})
}
func CardDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "card-description", "text-muted-foreground text-sm"), templ.GetChildren(ctx))
	})
}
func CardAction(props DOMProps) templ.Component {
	return cardWrapper(props, "card-action", "col-start-2 row-span-2 row-start-1 self-start justify-self-end")
}
func CardContent(props DOMProps) templ.Component {
	return cardWrapper(props, "card-content", "min-w-0 px-6")
}
func CardFooter(props DOMProps) templ.Component {
	return cardWrapper(props, "card-footer", "flex items-center px-6 [.border-t]:pt-6")
}
