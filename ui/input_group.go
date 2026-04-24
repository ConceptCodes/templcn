package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type InputGroupProps struct {
	DOMProps
}

func InputGroup(props InputGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props.DOMProps, "input-group", "flex rounded-md shadow-xs ring-1 ring-inset ring-input"), templ.GetChildren(ctx))
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
	return Input(InputProps{
		DOMProps:    props.DOMProps,
		Type:        props.Type,
		Name:        props.Name,
		Value:       props.Value,
		Placeholder: props.Placeholder,
		Disabled:    props.Disabled,
		Required:    props.Required,
		Invalid:     props.Invalid,
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
	return Textarea(TextareaProps{
		DOMProps:    props.DOMProps,
		Name:        props.Name,
		Value:       props.Value,
		Placeholder: props.Placeholder,
		Rows:        props.Rows,
		Disabled:    props.Disabled,
		Required:    props.Required,
		Invalid:     props.Invalid,
	})
}
