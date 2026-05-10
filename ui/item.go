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
		attrs := attrsFromDOMProps(props, "item-group", "group/item-group flex flex-col")
		attrs["role"] = "list"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func ItemSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "item-separator", "my-0 shrink-0 bg-border data-[orientation=horizontal]:h-px data-[orientation=horizontal]:w-full")
		attrs["data-orientation"] = SeparatorOrientationHorizontal
		attrs["role"] = "separator"
		return renderVoidElement(ctx, w, "hr", attrs)
	})
}

func Item(props ItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "group/item flex flex-wrap items-center rounded-md border border-transparent text-sm transition-colors duration-100 outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 [a]:transition-colors [a]:hover:bg-accent/50"
		switch props.Variant {
		case ItemVariantOutline:
			className = cn(className, "border-border")
		case ItemVariantMuted:
			className = cn(className, "bg-muted/50")
		default:
			className = cn(className, "bg-transparent")
		}
		if props.Size == ItemSizeSM {
			className = cn(className, "gap-2.5 px-4 py-3")
		} else {
			className = cn(className, "gap-4 p-4")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "item", className)
		if props.Variant != "" {
			attrs["data-variant"] = props.Variant
		} else {
			attrs["data-variant"] = "default"
		}
		if props.Size != "" {
			attrs["data-size"] = props.Size
		} else {
			attrs["data-size"] = "default"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func ItemMedia(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		variant, _ := props.Attrs["data-variant"].(string)
		className := "flex shrink-0 items-center justify-center gap-2 bg-transparent group-has-[[data-slot=item-description]]/item:translate-y-0.5 group-has-[[data-slot=item-description]]/item:self-start [&_svg]:pointer-events-none"
		if variant == "icon" {
			className = "flex size-8 shrink-0 items-center justify-center gap-2 rounded-sm border bg-muted group-has-[[data-slot=item-description]]/item:translate-y-0.5 group-has-[[data-slot=item-description]]/item:self-start [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4"
		} else if variant == "image" {
			className = "flex size-10 shrink-0 items-center justify-center gap-2 overflow-hidden rounded-sm group-has-[[data-slot=item-description]]/item:translate-y-0.5 group-has-[[data-slot=item-description]]/item:self-start [&_img]:size-full [&_img]:object-cover [&_svg]:pointer-events-none"
		}
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "item-media", className), templ.GetChildren(ctx))
	})
}

func ItemContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "item-content", "flex flex-1 flex-col gap-1 [&+[data-slot=item-content]]:flex-none"), templ.GetChildren(ctx))
	})
}

func ItemTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "item-title", "flex w-fit items-center gap-2 text-sm leading-snug font-medium"), templ.GetChildren(ctx))
	})
}

func ItemDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "p", attrsFromDOMProps(props, "item-description", "line-clamp-2 text-sm leading-normal font-normal text-balance text-muted-foreground [&>a]:underline [&>a]:underline-offset-4 [&>a:hover]:text-primary"), templ.GetChildren(ctx))
	})
}

func ItemActions(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "item-actions", "flex items-center gap-2"), templ.GetChildren(ctx))
	})
}

func ItemHeader(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "item-header", "flex basis-full items-center justify-between gap-2"), templ.GetChildren(ctx))
	})
}

func ItemFooter(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "item-footer", "flex basis-full items-center justify-between gap-2"), templ.GetChildren(ctx))
	})
}
