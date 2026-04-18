package ui

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/a-h/templ"
)

type AlertVariant string

const (
	AlertVariantDefault     AlertVariant = "default"
	AlertVariantDestructive AlertVariant = "destructive"
)

type AlertProps struct {
	DOMProps
	Variant AlertVariant
}

func alertClasses(variant AlertVariant, className string) string {
	if variant == "" {
		variant = AlertVariantDefault
	}

	base := "relative w-full rounded-xl border px-4 py-3 text-sm grid gap-2"
	switch variant {
	case AlertVariantDestructive:
		return cn(base, "border-destructive/50 text-destructive dark:border-destructive [&_svg]:text-destructive", className)
	default:
		return cn(base, "bg-background text-foreground", className)
	}
}

func Alert(props AlertProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "alert", alertClasses(props.Variant, ""))
		attrs["role"] = "alert"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func AlertTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "h5", attrsFromDOMProps(props, "alert-title", "mb-1 font-medium leading-none tracking-tight"), templ.GetChildren(ctx))
	})
}

func AlertDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "alert-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

type AspectRatioProps struct {
	DOMProps
	Ratio float64
}

func AspectRatio(props AspectRatioProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		ratio := props.Ratio
		if ratio <= 0 {
			ratio = 1
		}

		attrs := attrsFromDOMProps(props.DOMProps, "aspect-ratio", "relative w-full overflow-hidden")
		style := fmt.Sprintf("aspect-ratio: %s;", strconv.FormatFloat(ratio, 'f', -1, 64))
		if existing, ok := attrs["style"]; ok {
			style = cn(templ.Classes(existing).String(), style)
		}
		attrs["style"] = style
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

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

type SeparatorOrientation string

const (
	SeparatorOrientationHorizontal SeparatorOrientation = "horizontal"
	SeparatorOrientationVertical   SeparatorOrientation = "vertical"
)

type SeparatorProps struct {
	DOMProps
	Orientation SeparatorOrientation
	Decorative  bool
}

func Separator(props SeparatorProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		orientation := props.Orientation
		if orientation == "" {
			orientation = SeparatorOrientationHorizontal
		}

		className := "shrink-0 bg-border"
		if orientation == SeparatorOrientationHorizontal {
			className = cn(className, "h-px w-full")
		} else {
			className = cn(className, "h-full w-px")
		}

		attrs := attrsFromDOMProps(props.DOMProps, "separator", className)
		if !props.Decorative {
			attrs["role"] = "separator"
		} else {
			attrs["aria-hidden"] = "true"
		}
		return renderVoidElement(ctx, w, "hr", attrs)
	})
}

func Skeleton(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "skeleton", "animate-pulse rounded-md bg-muted"), nil)
	})
}

type SpinnerSize string

const (
	SpinnerSizeSM SpinnerSize = "sm"
	SpinnerSizeMD SpinnerSize = "md"
	SpinnerSizeLG SpinnerSize = "lg"
)

type SpinnerProps struct {
	DOMProps
	Size  SpinnerSize
	Label string
}

func Spinner(props SpinnerProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		size := props.Size
		if size == "" {
			size = SpinnerSizeMD
		}

		dimensions := map[SpinnerSize]string{
			SpinnerSizeSM: "size-4",
			SpinnerSizeMD: "size-5",
			SpinnerSizeLG: "size-6",
		}[size]

		attrs := attrsFromDOMProps(props.DOMProps, "spinner", cn("inline-flex items-center justify-center animate-spin text-muted-foreground", dimensions))
		if props.Label != "" {
			attrs["role"] = "status"
			attrs["aria-label"] = props.Label
		} else {
			attrs["aria-hidden"] = "true"
		}

		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			return renderVoidElement(ctx, w, "svg", templ.Attributes{
				"viewBox":     "0 0 24 24",
				"fill":        "none",
				"aria-hidden": "true",
			})
		})

		return renderElement(ctx, w, "div", attrs, children)
	})
}

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

func typographyComponent(tag, slot, className string, props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, tag, attrsFromDOMProps(props, slot, className), templ.GetChildren(ctx))
	})
}

func H1(props DOMProps) templ.Component {
	return typographyComponent("h1", "h1", "scroll-m-20 text-4xl font-bold tracking-tight lg:text-5xl", props)
}
func H2(props DOMProps) templ.Component {
	return typographyComponent("h2", "h2", "scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight first:mt-0", props)
}
func H3(props DOMProps) templ.Component {
	return typographyComponent("h3", "h3", "scroll-m-20 text-2xl font-semibold tracking-tight", props)
}
func H4(props DOMProps) templ.Component {
	return typographyComponent("h4", "h4", "scroll-m-20 text-xl font-semibold tracking-tight", props)
}
func P(props DOMProps) templ.Component {
	return typographyComponent("p", "p", "leading-7 [&:not(:first-child)]:mt-6", props)
}
func Blockquote(props DOMProps) templ.Component {
	return typographyComponent("blockquote", "blockquote", "mt-6 border-l-2 border-border pl-6 italic", props)
}
func InlineCode(props DOMProps) templ.Component {
	return typographyComponent("code", "inline-code", "relative rounded bg-muted px-[0.3rem] py-[0.2rem] font-mono text-sm font-semibold", props)
}
func Lead(props DOMProps) templ.Component {
	return typographyComponent("p", "lead", "text-xl text-muted-foreground", props)
}
func Large(props DOMProps) templ.Component {
	return typographyComponent("div", "large", "text-lg font-semibold", props)
}
func Small(props DOMProps) templ.Component {
	return typographyComponent("small", "small", "text-sm font-medium leading-none", props)
}
func Muted(props DOMProps) templ.Component {
	return typographyComponent("p", "muted", "text-sm text-muted-foreground", props)
}
func List(props DOMProps) templ.Component {
	return typographyComponent("ul", "list", "my-6 ml-6 list-disc [&>li]:mt-2", props)
}
func TableProse(props DOMProps) templ.Component {
	return typographyComponent("div", "table-prose", "my-6 w-full overflow-y-auto", props)
}

type LabelProps struct {
	DOMProps
	For string
}

func Label(props LabelProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "label", "text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70")
		if props.For != "" {
			attrs["for"] = props.For
		}
		return renderElement(ctx, w, "label", attrs, templ.GetChildren(ctx))
	})
}
