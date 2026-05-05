package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type BadgeVariant string

const (
	BadgeVariantDefault     BadgeVariant = "default"
	BadgeVariantSecondary   BadgeVariant = "secondary"
	BadgeVariantDestructive BadgeVariant = "destructive"
	BadgeVariantOutline     BadgeVariant = "outline"
	BadgeVariantGhost       BadgeVariant = "ghost"
	BadgeVariantLink        BadgeVariant = "link"
)

type BadgeProps struct {
	DOMProps
	Label   string
	Variant BadgeVariant
	Href    string
}

var badgeVariantClasses = map[BadgeVariant]string{
	BadgeVariantDefault:     "bg-primary text-primary-foreground [a&]:hover:bg-primary/90",
	BadgeVariantSecondary:   "bg-secondary text-secondary-foreground [a&]:hover:bg-secondary/90",
	BadgeVariantDestructive: "bg-destructive text-white focus-visible:ring-destructive/20 dark:bg-destructive/60 dark:focus-visible:ring-destructive/40 [a&]:hover:bg-destructive/90",
	BadgeVariantOutline:     "border-border text-foreground [a&]:hover:bg-accent [a&]:hover:text-accent-foreground",
	BadgeVariantGhost:       "[a&]:hover:bg-accent [a&]:hover:text-accent-foreground",
	BadgeVariantLink:        "text-primary underline-offset-4 [a&]:hover:underline",
}

func badgeClasses(variant BadgeVariant, className string) string {
	if variant == "" {
		variant = BadgeVariantDefault
	}

	base := "inline-flex w-fit shrink-0 items-center justify-center gap-1 overflow-hidden rounded-full border border-transparent px-2 py-0.5 text-xs font-medium whitespace-nowrap transition-[color,box-shadow] focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 aria-invalid:border-destructive aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 [&>svg]:pointer-events-none [&>svg]:size-3"
	return cn(base, badgeVariantClasses[variant], className)
}

func Badge(props BadgeProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		ctx, slottedChildren := childrenFromContext(ctx)
		variant := props.Variant
		if variant == "" {
			variant = BadgeVariantDefault
		}
		attrs := attrsFromDOMProps(props.DOMProps, "badge", badgeClasses(variant, ""))
		attrs["data-variant"] = string(variant)

		tag := "span"
		if props.Element != "" {
			tag = props.Element
		} else if props.Href != "" {
			tag = "a"
		}

		if tag == "a" {
			attrs["href"] = props.Href
		}

		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if props.Label != "" {
				if _, err := io.WriteString(w, templ.EscapeString(props.Label)); err != nil {
					return err
				}
			}
			if props.Label == "" {
				if err := renderChildren(ctx, w, slottedChildren); err != nil {
					return err
				}
			}
			return nil
		})
		return renderElement(ctx, w, tag, attrs, children)
	})
}
