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

type carouselRenderState struct {
	orientation string
}

type carouselRenderStateKey struct{}

func Carousel(props CarouselProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "carousel", "relative")
		attrs["role"] = "region"
		attrs["aria-roledescription"] = "carousel"
		attrs["data-index"] = props.StartIndex
		orientation := props.Orientation
		if orientation == "" {
			orientation = "horizontal"
		}
		attrs["data-orientation"] = orientation
		if props.Loop {
			attrs["data-loop"] = "true"
		}
		if props.Align != "" {
			attrs["data-align"] = props.Align
		}
		if props.StartIndex > 0 {
			attrs["data-start-index"] = props.StartIndex
		}
		ctx = context.WithValue(ctx, carouselRenderStateKey{}, carouselRenderState{orientation: orientation})
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func CarouselContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		orientation := carouselOrientationFromContext(ctx)
		innerClass := "flex -ml-4"
		if orientation == "vertical" {
			innerClass = "flex -mt-4 flex-col"
		}
		attrs := attrsFromDOMProps(props, "carousel-content", "overflow-hidden")
		attrs["aria-live"] = "polite"
		inner := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return renderElement(ctx, w, "div", templ.Attributes{"class": innerClass}, templ.GetChildren(ctx))
		})
		return renderElement(ctx, w, "div", attrs, inner)
	})
}
func CarouselItem(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "min-w-0 shrink-0 grow-0 basis-full pl-4"
		if carouselOrientationFromContext(ctx) == "vertical" {
			className = "min-w-0 shrink-0 grow-0 basis-full pt-4"
		}
		attrs := attrsFromDOMProps(props, "carousel-item", className)
		attrs["role"] = "group"
		attrs["aria-roledescription"] = "slide"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func CarouselPrevious(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := cn(buttonClasses(ButtonVariantOutline, ButtonSizeIcon, ""), "absolute size-8 rounded-full top-1/2 -left-12 -translate-y-1/2")
		if carouselOrientationFromContext(ctx) == "vertical" {
			className = cn(buttonClasses(ButtonVariantOutline, ButtonSizeIcon, ""), "absolute size-8 rounded-full -top-12 left-1/2 -translate-x-1/2 rotate-90")
		}
		attrs := attrsFromDOMProps(props, "carousel-previous", className)
		attrs["type"] = "button"
		attrs["aria-label"] = "Previous slide"
		children := templ.GetChildren(ctx)
		if children == nil {
			children = templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				_, err := io.WriteString(w, `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="m12 19-7-7 7-7"></path><path d="M19 12H5"></path></svg><span class="sr-only">Previous slide</span>`)
				return err
			})
		}
		return renderElement(ctx, w, "button", attrs, children)
	})
}

func CarouselNext(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := cn(buttonClasses(ButtonVariantOutline, ButtonSizeIcon, ""), "absolute size-8 rounded-full top-1/2 -right-12 -translate-y-1/2")
		if carouselOrientationFromContext(ctx) == "vertical" {
			className = cn(buttonClasses(ButtonVariantOutline, ButtonSizeIcon, ""), "absolute size-8 rounded-full -bottom-12 left-1/2 -translate-x-1/2 rotate-90")
		}
		attrs := attrsFromDOMProps(props, "carousel-next", className)
		attrs["type"] = "button"
		attrs["aria-label"] = "Next slide"
		children := templ.GetChildren(ctx)
		if children == nil {
			children = templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				_, err := io.WriteString(w, `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M5 12h14"></path><path d="m12 5 7 7-7 7"></path></svg><span class="sr-only">Next slide</span>`)
				return err
			})
		}
		return renderElement(ctx, w, "button", attrs, children)
	})
}

func carouselOrientationFromContext(ctx context.Context) string {
	state, ok := ctx.Value(carouselRenderStateKey{}).(carouselRenderState)
	if !ok || state.orientation == "" {
		return "horizontal"
	}
	return state.orientation
}
