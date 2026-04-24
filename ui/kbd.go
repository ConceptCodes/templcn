package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type KbdProps struct {
	DOMProps
	Text  string
	Inset bool
	Size  string
}

func Kbd(props KbdProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		sizeClass := "text-[0.7rem]"
		switch props.Size {
		case "sm":
			sizeClass = "text-[0.65rem]"
		case "lg":
			sizeClass = "text-[0.75rem]"
		}

		className := cn(
			"inline-flex items-center rounded border border-border bg-muted px-1.5 py-0.5 font-mono font-medium text-foreground shadow-sm",
			sizeClass,
		)
		if props.Inset {
			className = cn(className, "ms-1")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "kbd", className)
		return renderTextElement(ctx, w, "kbd", attrs, props.Text)
	})
}
