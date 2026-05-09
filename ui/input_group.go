package ui

import (
	"context"
	"io"
	"strconv"

	"github.com/a-h/templ"
)

type InputGroupProps struct {
	DOMProps
}

func InputGroup(props InputGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "input-group", "flex rounded-md shadow-xs ring-1 ring-inset ring-input")
		attrs["role"] = "group"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

type InputGroupAddonAlign string

const (
	InputGroupAddonInlineStart InputGroupAddonAlign = "inline-start"
	InputGroupAddonInlineEnd   InputGroupAddonAlign = "inline-end"
	InputGroupAddonBlockStart  InputGroupAddonAlign = "block-start"
	InputGroupAddonBlockEnd    InputGroupAddonAlign = "block-end"
)

type InputGroupAddonProps struct {
	DOMProps
	Align InputGroupAddonAlign
}

func InputGroupAddon(props InputGroupAddonProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props.DOMProps, "input-group-addon", "inline-flex items-center rounded-md border-0 px-3 py-2 text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

type InputGroupButtonProps struct {
	DOMProps
	Size ButtonSize
}

func InputGroupButton(props InputGroupButtonProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := buttonClasses(ButtonVariantDefault, props.Size, "rounded-none border-0 shadow-none")
		attrs := attrsFromDOMProps(props.DOMProps, "input-group-button", className)
		if _, ok := attrs["type"]; !ok {
			attrs["type"] = "button"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

type InputGroupTextProps struct {
	DOMProps
}

func InputGroupText(props InputGroupTextProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props.DOMProps, "input-group-text", "px-3 py-2 text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

type InputGroupInputProps struct {
	DOMProps
	Type        string
	Name        string
	Value       string
	Placeholder string
	Disabled    bool
	Required    bool
	Invalid     bool
}

func InputGroupInput(props InputGroupInputProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "input-group-control", inputClasses)
		if props.Type == "" {
			props.Type = "text"
		}
		attrs["type"] = props.Type
		if props.Name != "" {
			attrs["name"] = props.Name
		}
		if props.Value != "" {
			attrs["value"] = props.Value
		}
		if props.Placeholder != "" {
			attrs["placeholder"] = props.Placeholder
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
		return renderVoidElement(ctx, w, "input", attrs)
	})
}

type InputGroupTextareaProps struct {
	DOMProps
	Name        string
	Value       string
	Placeholder string
	Rows        int
	Disabled    bool
	Required    bool
	Invalid     bool
}

func InputGroupTextarea(props InputGroupTextareaProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "input-group-control", textareaClasses)
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
