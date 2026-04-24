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
	return cardWrapper(props, "card", "rounded-xl border bg-card text-card-foreground shadow")
}
func CardHeader(props DOMProps) templ.Component {
	return cardWrapper(props, "card-header", "flex flex-col gap-1.5 p-6")
}
func CardTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "h3", attrsFromDOMProps(props, "card-title", "font-semibold leading-none tracking-tight"), templ.GetChildren(ctx))
	})
}
func CardDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "p", attrsFromDOMProps(props, "card-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}
func CardAction(props DOMProps) templ.Component {
	return cardWrapper(props, "card-action", "ml-auto flex items-center")
}
func CardContent(props DOMProps) templ.Component {
	return cardWrapper(props, "card-content", "p-6 pt-0")
}
func CardFooter(props DOMProps) templ.Component {
	return cardWrapper(props, "card-footer", "flex items-center p-6 pt-0")
}
