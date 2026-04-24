package ui

import (
	"context"
	"io"
	"strconv"

	"github.com/a-h/templ"
)

type AvatarProps struct {
	DOMProps
	Size string
}

func Avatar(props AvatarProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		sizeClass := "size-10"
		switch props.Size {
		case "sm":
			sizeClass = "size-8"
		case "lg":
			sizeClass = "size-12"
		}
		return renderElement(ctx, w, "div", attrsFromDOMProps(props.DOMProps, "avatar", cn("relative flex shrink-0 overflow-hidden rounded-full", sizeClass)), templ.GetChildren(ctx))
	})
}

type AvatarImageProps struct {
	DOMProps
	Src string
	Alt string
}

func AvatarImage(props AvatarImageProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "avatar-image", "size-full object-cover")
		if props.Src != "" {
			attrs["src"] = props.Src
		}
		if props.Alt != "" {
			attrs["alt"] = props.Alt
		}
		return renderVoidElement(ctx, w, "img", attrs)
	})
}

func AvatarFallback(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "avatar-fallback", "flex size-full items-center justify-center rounded-full bg-muted text-xs"), templ.GetChildren(ctx))
	})
}

func AvatarBadge(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "avatar-badge", "absolute bottom-0 right-0 inline-flex size-2.5 rounded-full border border-background bg-primary"), templ.GetChildren(ctx))
	})
}

func AvatarGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "avatar-group", "flex -space-x-2"), templ.GetChildren(ctx))
	})
}

type AvatarGroupCountProps struct {
	DOMProps
	Count int
}

func AvatarGroupCount(props AvatarGroupCountProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		label := "+"
		if props.Count > 0 {
			label += strconv.Itoa(props.Count)
		}
		return renderTextElement(ctx, w, "span", attrsFromDOMProps(props.DOMProps, "avatar-group-count", "inline-flex size-8 items-center justify-center rounded-full bg-muted text-xs font-medium"), label)
	})
}
