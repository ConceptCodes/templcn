package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type SelectProps struct {
	DOMProps
	Name         string
	Value        string
	DefaultValue string
	Open         bool
	DefaultOpen  bool
	Disabled     bool
	Required     bool
}

type SelectTriggerProps struct {
	DOMProps
	Size string
}

type SelectContentProps struct {
	DOMProps
	Position string
	Align    string
}

type selectRenderState struct {
	open  bool
	value string
}

type selectRenderStateKey struct{}

func Select(props SelectProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		open := props.Open || props.DefaultOpen
		value := props.Value
		if value == "" {
			value = props.DefaultValue
		}
		attrs := attrsFromDOMProps(props.DOMProps, "select", "relative")
		attrs["data-state"] = openState(open)
		if props.Name != "" {
			attrs["data-name"] = props.Name
		}
		if value != "" {
			attrs["data-value"] = value
		}
		if props.DefaultValue != "" {
			attrs["data-default-value"] = props.DefaultValue
		}
		if open {
			attrs["data-open"] = "true"
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		if props.Disabled {
			attrs["data-disabled"] = "true"
		}
		if props.Required {
			attrs["data-required"] = "true"
		}
		ctx = context.WithValue(ctx, selectRenderStateKey{}, selectRenderState{open: open, value: value})
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if props.Name != "" {
				inputAttrs := templ.Attributes{"type": "hidden", "name": props.Name, "value": value}
				if err := renderVoidElement(ctx, w, "input", inputAttrs); err != nil {
					return err
				}
			}
			return renderChildren(ctx, w, templ.GetChildren(ctx))
		})
		return renderElement(ctx, w, "div", attrs, children)
	})
}

func SelectGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "select-group", "grid gap-1"), templ.GetChildren(ctx))
	})
}
func SelectValue(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "select-value", ""), templ.GetChildren(ctx))
	})
}
func SelectTrigger(props SelectTriggerProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "flex h-9 w-full items-center justify-between rounded-md border border-input bg-background px-3 py-2 text-sm shadow-xs outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:cursor-not-allowed disabled:opacity-50"
		if props.Size == "sm" {
			className = cn(className, "h-8")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "select-trigger", className)
		attrs["type"] = "button"
		attrs["role"] = "combobox"
		attrs["aria-haspopup"] = "listbox"
		if _, ok := attrs["aria-expanded"]; !ok {
			attrs["aria-expanded"] = "false"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func SelectContent(props SelectContentProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "select-content", "z-50 min-w-32 rounded-md border bg-popover p-1 text-popover-foreground shadow-md")
		attrs["role"] = "listbox"
		if _, ok := attrs["data-state"]; !ok {
			attrs["data-state"] = selectStateFromContext(ctx)
		}
		if selectStateFromContext(ctx) == "closed" {
			attrs["hidden"] = true
		}
		if props.Position != "" {
			attrs["data-position"] = props.Position
		}
		if props.Align != "" {
			attrs["data-align"] = props.Align
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func SelectLabel(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "select-label", "px-2 py-1.5 text-sm font-semibold"), templ.GetChildren(ctx))
	})
}
func SelectItem(props DropdownMenuItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "relative flex cursor-pointer select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground"
		if props.Inset {
			className = cn(className, "pl-8")
		}
		if props.Variant == "destructive" {
			className = cn(className, "text-destructive")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "select-item", className)
		attrs["type"] = "button"
		attrs["role"] = "option"
		value := props.Value
		if value == "" {
			if attrValue, ok := attrs["data-value"].(string); ok {
				value = attrValue
			}
		}
		if value != "" {
			attrs["data-value"] = value
		}
		selected := value != "" && value == selectValueFromContext(ctx)
		attrs["aria-selected"] = map[bool]string{true: "true", false: "false"}[selected]
		attrs["data-state"] = map[bool]string{true: "checked", false: "unchecked"}[selected]
		if props.Disabled {
			attrs["disabled"] = true
			attrs["aria-disabled"] = "true"
			attrs["data-disabled"] = "true"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
func SelectSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "select-separator", "-mx-1 my-1 h-px bg-border"), nil)
	})
}
func SelectScrollUpButton(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "select-scroll-up-button", "flex cursor-default items-center justify-center py-1")
		attrs["type"] = "button"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
func SelectScrollDownButton(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "select-scroll-down-button", "flex cursor-default items-center justify-center py-1")
		attrs["type"] = "button"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func selectStateFromContext(ctx context.Context) string {
	state, ok := ctx.Value(selectRenderStateKey{}).(selectRenderState)
	if !ok {
		return "closed"
	}
	return openState(state.open)
}

func selectValueFromContext(ctx context.Context) string {
	state, ok := ctx.Value(selectRenderStateKey{}).(selectRenderState)
	if !ok {
		return ""
	}
	return state.value
}
