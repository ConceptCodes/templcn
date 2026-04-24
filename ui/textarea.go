package ui

import (
	"context"
	"io"
	"strconv"

	"github.com/a-h/templ"
)

type TextareaProps struct {
	DOMProps
	Name        string
	Value       string
	Placeholder string
	Rows        int
	Disabled    bool
	Required    bool
	Invalid     bool
}

func Textarea(props TextareaProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "textarea", "flex min-h-20 w-full rounded-md border border-input bg-transparent px-3 py-2 text-base shadow-xs placeholder:text-muted-foreground focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 outline-none disabled:cursor-not-allowed disabled:opacity-50 md:text-sm")
		if props.Name != "" {
			attrs["name"] = props.Name
		}
		if props.Placeholder != "" {
			attrs["placeholder"] = props.Placeholder
		}
		if props.Rows > 0 {
			attrs["rows"] = strconv.Itoa(props.Rows)
		}
		if props.Disabled {
			attrs["disabled"] = true
		}
		if props.Required {
			attrs["required"] = true
		}
		if props.Invalid {
			attrs["aria-invalid"] = "true"
		}

		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if props.Value == "" {
				return nil
			}
			_, err := io.WriteString(w, templ.EscapeString(props.Value))
			return err
		})
		return renderElement(ctx, w, "textarea", attrs, children)
	})
}
