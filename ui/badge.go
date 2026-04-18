package ui

import (
	"bytes"
	"context"
	"io"

	"github.com/a-h/templ"
)

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
		buf, isBuffer := w.(*bytes.Buffer)
		if !isBuffer {
			buf = templ.GetBuffer()
			defer templ.ReleaseBuffer(buf)
		}

		ctx = templ.InitializeContext(ctx)

		attrs := cloneAttributes(props.Attrs)
		if props.ID != "" {
			attrs["id"] = props.ID
		}

		className := badgeClasses(props.Variant, props.Class)
		if existing, ok := attrs["class"]; ok {
			className = cn(className, templ.Classes(existing).String())
		}
		attrs["class"] = className

		tag := "span"
		if props.Element != "" {
			tag = props.Element
		} else if props.Href != "" {
			tag = "a"
		}

		if tag == "a" {
			attrs["href"] = props.Href
		}

		if _, err := buf.WriteString("<" + tag); err != nil {
			return err
		}
		if err := templ.RenderAttributes(ctx, buf, attrs); err != nil {
			return err
		}
		if _, err := buf.WriteString(">"); err != nil {
			return err
		}
		if props.Label != "" {
			if _, err := buf.WriteString(templ.EscapeString(props.Label)); err != nil {
				return err
			}
		}
		if _, err := buf.WriteString("</" + tag + ">"); err != nil {
			return err
		}

		if !isBuffer {
			_, err := buf.WriteTo(w)
			return err
		}
		return nil
	})
}
