package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type BubbleProps struct {
	DOMProps
	Variant string
	Align   string
}

func BubbleGroup(props DOMProps) templ.Component {
	return bubbleBlock("bubble-group", "flex min-w-0 flex-col gap-2", props)
}
func Bubble(props BubbleProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		variant, align := props.Variant, props.Align
		if variant == "" {
			variant = "default"
		}
		if align == "" {
			align = "start"
		}
		className := "group/bubble relative flex w-fit max-w-[80%] min-w-0 flex-col gap-1 data-[align=end]:self-end"
		attrs := attrsFromDOMProps(props.DOMProps, "bubble", className)
		attrs["data-variant"], attrs["data-align"] = variant, align
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func BubbleContent(props DOMProps) templ.Component {
	return bubbleBlock("bubble-content", "w-fit max-w-full min-w-0 overflow-hidden rounded-xl border border-transparent px-3 py-2 text-sm leading-relaxed", props)
}

type BubbleReactionsProps struct {
	DOMProps
	Align string
	Side  string
}

func BubbleReactions(props BubbleReactionsProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		align, side := props.Align, props.Side
		if align == "" {
			align = "end"
		}
		if side == "" {
			side = "bottom"
		}
		attrs := attrsFromDOMProps(props.DOMProps, "bubble-reactions", "absolute z-10 flex w-fit shrink-0 items-center justify-center gap-1 rounded-full bg-muted px-1.5 py-0.5 text-sm")
		attrs["data-align"], attrs["data-side"] = align, side
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func bubbleBlock(slot, className string, props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, slot, className), templ.GetChildren(ctx))
	})
}
