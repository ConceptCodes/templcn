package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type CarouselProps struct {
	DOMProps
	Orientation string
	Loop        bool
	Align       string
	StartIndex  int
}

func Carousel(props CarouselProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "carousel", "relative")
		attrs["role"] = "region"
		attrs["aria-roledescription"] = "carousel"
		attrs["data-index"] = props.StartIndex
		if props.Orientation != "" {
			attrs["data-orientation"] = props.Orientation
		}
		if props.Loop {
			attrs["data-loop"] = "true"
		}
		if props.Align != "" {
			attrs["data-align"] = props.Align
		}
		if props.StartIndex > 0 {
			attrs["data-start-index"] = props.StartIndex
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func CarouselContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "carousel-content", "flex")
		attrs["aria-live"] = "polite"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func CarouselItem(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "carousel-item", "min-w-0 shrink-0 grow-0 basis-full")
		attrs["role"] = "group"
		attrs["aria-roledescription"] = "slide"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func CarouselPrevious(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "carousel-previous", "absolute left-2 top-1/2 -translate-y-1/2 rounded-full border bg-background p-2 shadow")
		attrs["type"] = "button"
		attrs["aria-label"] = "Previous slide"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func CarouselNext(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "carousel-next", "absolute right-2 top-1/2 -translate-y-1/2 rounded-full border bg-background p-2 shadow")
		attrs["type"] = "button"
		attrs["aria-label"] = "Next slide"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
