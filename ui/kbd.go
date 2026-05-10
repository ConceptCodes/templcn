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
		sizeClass := "text-xs"
		switch props.Size {
		case "sm":
			sizeClass = "text-[0.65rem]"
		case "lg":
			sizeClass = "text-[0.75rem]"
		}

		className := cn(
			"pointer-events-none inline-flex h-5 w-fit min-w-5 items-center justify-center gap-1 rounded-sm bg-muted px-1 font-sans font-medium text-muted-foreground select-none [&_svg:not([class*='size-'])]:size-3 [[data-slot=tooltip-content]_&]:bg-background/20 [[data-slot=tooltip-content]_&]:text-background dark:[[data-slot=tooltip-content]_&]:bg-background/10",
			sizeClass,
		)
		if props.Inset {
			className = cn(className, "ms-1")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "kbd", className)
		return renderTextElement(ctx, w, "kbd", attrs, props.Text)
	})
}

func KbdGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "kbd", attrsFromDOMProps(props, "kbd-group", "inline-flex items-center gap-1"), templ.GetChildren(ctx))
	})
}
