package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type ItemVariant string

const (
	ItemVariantDefault ItemVariant = "default"
	ItemVariantOutline ItemVariant = "outline"
	ItemVariantMuted   ItemVariant = "muted"
)

type ItemSize string

const (
	ItemSizeDefault ItemSize = "default"
	ItemSizeSM      ItemSize = "sm"
)

type ItemProps struct {
	DOMProps
	Variant ItemVariant
	Size    ItemSize
}

func ItemGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "item-group", "grid gap-2"), templ.GetChildren(ctx))
	})
}

func ItemSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "hr", attrsFromDOMProps(props, "item-separator", "my-2 border-border"), nil)
	})
}

func Item(props ItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "grid gap-2 rounded-xl border bg-card p-4"
		switch props.Variant {
		case ItemVariantOutline:
			className = cn(className, "border-input")
		case ItemVariantMuted:
			className = cn(className, "bg-muted")
		}
		if props.Size == ItemSizeSM {
			className = cn(className, "p-3")
		}
		return renderElement(ctx, w, "div", attrsFromDOMProps(props.DOMProps, "item", className), templ.GetChildren(ctx))
	})
}

func ItemMedia(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "item-media", "flex size-10 items-center justify-center rounded-md bg-muted"), templ.GetChildren(ctx))
	})
}

func ItemContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "item-content", "grid gap-1"), templ.GetChildren(ctx))
	})
}

func ItemTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "h4", attrsFromDOMProps(props, "item-title", "font-medium leading-none"), templ.GetChildren(ctx))
	})
}

func ItemDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "p", attrsFromDOMProps(props, "item-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func ItemActions(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "item-actions", "ml-auto flex items-center gap-2"), templ.GetChildren(ctx))
	})
}

func ItemHeader(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "item-header", "flex items-start gap-3"), templ.GetChildren(ctx))
	})
}

func ItemFooter(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "item-footer", "flex items-center gap-2 pt-2"), templ.GetChildren(ctx))
	})
}
