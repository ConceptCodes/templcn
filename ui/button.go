package ui

import (
	"bytes"
	"context"
	"io"
	"sort"

	"github.com/a-h/templ"
)

type ButtonVariant string

const (
	ButtonVariantDefault     ButtonVariant = "default"
	ButtonVariantDestructive ButtonVariant = "destructive"
	ButtonVariantOutline     ButtonVariant = "outline"
	ButtonVariantSecondary   ButtonVariant = "secondary"
	ButtonVariantGhost       ButtonVariant = "ghost"
	ButtonVariantLink        ButtonVariant = "link"
)

type ButtonSize string

const (
	ButtonSizeDefault ButtonSize = "default"
	ButtonSizeXS      ButtonSize = "xs"
	ButtonSizeSM      ButtonSize = "sm"
	ButtonSizeLG      ButtonSize = "lg"
	ButtonSizeIcon    ButtonSize = "icon"
	ButtonSizeIconXS  ButtonSize = "icon-xs"
	ButtonSizeIconSM  ButtonSize = "icon-sm"
	ButtonSizeIconLG  ButtonSize = "icon-lg"
)

type ButtonProps struct {
	DOMProps
	Label    string
	Variant  ButtonVariant
	Size     ButtonSize
	Type     string
	Disabled bool
	Href     string
}

var buttonVariantClasses = map[ButtonVariant]string{
	ButtonVariantDefault:     "bg-primary text-primary-foreground hover:bg-primary/90",
	ButtonVariantDestructive: "bg-destructive text-white hover:bg-destructive/90 focus-visible:ring-destructive/20 dark:bg-destructive/60 dark:focus-visible:ring-destructive/40",
	ButtonVariantOutline:     "border bg-background shadow-xs hover:bg-accent hover:text-accent-foreground dark:border-input dark:bg-input/30 dark:hover:bg-input/50",
	ButtonVariantSecondary:   "bg-secondary text-secondary-foreground hover:bg-secondary/80",
	ButtonVariantGhost:       "hover:bg-accent hover:text-accent-foreground dark:hover:bg-accent/50",
	ButtonVariantLink:        "text-primary underline-offset-4 hover:underline",
}

var buttonSizeClasses = map[ButtonSize]string{
	ButtonSizeDefault: "h-9 px-4 py-2 has-[>svg]:px-3",
	ButtonSizeXS:      "h-6 gap-1 rounded-md px-2 text-xs has-[>svg]:px-1.5 [&_svg:not([class*='size-'])]:size-3",
	ButtonSizeSM:      "h-8 gap-1.5 rounded-md px-3 has-[>svg]:px-2.5",
	ButtonSizeLG:      "h-10 rounded-md px-6 has-[>svg]:px-4",
	ButtonSizeIcon:    "size-9",
	ButtonSizeIconXS:  "size-6 rounded-md [&_svg:not([class*='size-'])]:size-3",
	ButtonSizeIconSM:  "size-8",
	ButtonSizeIconLG:  "size-10",
}

func buttonClasses(variant ButtonVariant, size ButtonSize, className string) string {
	if variant == "" {
		variant = ButtonVariantDefault
	}
	if size == "" {
		size = ButtonSizeDefault
	}

	base := "inline-flex shrink-0 items-center justify-center gap-2 rounded-md text-sm font-medium whitespace-nowrap transition-all outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:pointer-events-none disabled:opacity-50 aria-invalid:border-destructive aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4"
	return cn(base, buttonVariantClasses[variant], buttonSizeClasses[size], className)
}

func Button(props ButtonProps) templ.Component {
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

		className := buttonClasses(props.Variant, props.Size, props.Class)
		if existing, ok := attrs["class"]; ok {
			className = cn(className, templ.Classes(existing).String())
		}
		attrs["class"] = className

		tag := "button"
		if props.Element != "" {
			tag = props.Element
		} else if props.Href != "" {
			tag = "a"
		}

		if tag == "a" {
			attrs["href"] = props.Href
		} else {
			if props.Type == "" {
				attrs["type"] = "button"
			} else {
				attrs["type"] = props.Type
			}
			if props.Disabled {
				attrs["disabled"] = true
			}
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

func ButtonVariantNames() []string {
	names := make([]string, 0, len(buttonVariantClasses))
	for name := range buttonVariantClasses {
		names = append(names, string(name))
	}
	sort.Strings(names)
	return names
}
