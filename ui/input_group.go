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
		attrs := attrsFromDOMProps(props.DOMProps, "input-group", "group/input-group relative flex h-9 min-w-0 w-full items-center rounded-md border border-input shadow-xs transition-[color,box-shadow] outline-none has-[>textarea]:h-auto dark:bg-input/30 has-[>[data-align=inline-start]]:[&>input]:pl-2 has-[>[data-align=inline-end]]:[&>input]:pr-2 has-[>[data-align=block-start]]:h-auto has-[>[data-align=block-start]]:flex-col has-[>[data-align=block-start]]:[&>input]:pb-3 has-[>[data-align=block-end]]:h-auto has-[>[data-align=block-end]]:flex-col has-[>[data-align=block-end]]:[&>input]:pt-3 has-[[data-slot=input-group-control]:focus-visible]:border-ring has-[[data-slot=input-group-control]:focus-visible]:ring-[3px] has-[[data-slot=input-group-control]:focus-visible]:ring-ring/50 has-[[data-slot][aria-invalid=true]]:border-destructive has-[[data-slot][aria-invalid=true]]:ring-destructive/20 dark:has-[[data-slot][aria-invalid=true]]:ring-destructive/40")
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
		align := props.Align
		if align == "" {
			align = InputGroupAddonInlineStart
		}
		className := "flex h-auto cursor-text items-center justify-center gap-2 py-1.5 text-sm font-medium text-muted-foreground select-none group-data-[disabled=true]/input-group:opacity-50 [&>kbd]:rounded-[calc(var(--radius)-5px)] [&>svg:not([class*='size-'])]:size-4"
		switch align {
		case InputGroupAddonInlineEnd:
			className = cn(className, "order-last pr-3 has-[>button]:mr-[-0.45rem] has-[>kbd]:mr-[-0.35rem]")
		case InputGroupAddonBlockStart:
			className = cn(className, "order-first w-full justify-start px-3 pt-3 group-has-[>input]/input-group:pt-2.5 [.border-b]:pb-3")
		case InputGroupAddonBlockEnd:
			className = cn(className, "order-last w-full justify-start px-3 pb-3 group-has-[>input]/input-group:pb-2.5 [.border-t]:pt-3")
		default:
			className = cn(className, "order-first pl-3 has-[>button]:ml-[-0.45rem] has-[>kbd]:ml-[-0.35rem]")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "input-group-addon", className)
		attrs["role"] = "group"
		attrs["data-align"] = align
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

type InputGroupButtonProps struct {
	DOMProps
	Size ButtonSize
}

func InputGroupButton(props InputGroupButtonProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		size := props.Size
		if size == "" {
			size = ButtonSizeXS
		}
		className := "flex items-center gap-2 text-sm shadow-none"
		switch size {
		case ButtonSizeSM:
			className = cn(className, "h-8 gap-1.5 rounded-md px-2.5 has-[>svg]:px-2.5")
		case ButtonSizeIconXS:
			className = cn(className, "size-6 rounded-[calc(var(--radius)-5px)] p-0 has-[>svg]:p-0")
		case ButtonSizeIconSM:
			className = cn(className, "size-8 p-0 has-[>svg]:p-0")
		default:
			className = cn(className, "h-6 gap-1 rounded-[calc(var(--radius)-5px)] px-2 has-[>svg]:px-2 [&>svg:not([class*='size-'])]:size-3.5")
		}
		className = cn(buttonClasses(ButtonVariantGhost, ButtonSizeDefault, ""), className)
		attrs := attrsFromDOMProps(props.DOMProps, "input-group-button", className)
		attrs["data-size"] = size
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
		return renderElement(ctx, w, "span", attrsFromDOMProps(props.DOMProps, "input-group-text", "flex items-center gap-2 text-sm text-muted-foreground [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4"), templ.GetChildren(ctx))
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
		attrs := attrsFromDOMProps(props.DOMProps, "input-group-control", cn(inputClasses, "flex-1 rounded-none border-0 bg-transparent shadow-none focus-visible:ring-0 dark:bg-transparent"))
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
		attrs := attrsFromDOMProps(props.DOMProps, "input-group-control", cn(textareaClasses, "flex-1 resize-none rounded-none border-0 bg-transparent py-3 shadow-none focus-visible:ring-0 dark:bg-transparent"))
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
