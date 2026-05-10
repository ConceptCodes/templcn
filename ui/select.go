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
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "select-group", ""), templ.GetChildren(ctx))
	})
}
func SelectValue(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "select-value", ""), templ.GetChildren(ctx))
	})
}
func SelectTrigger(props SelectTriggerProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		size := props.Size
		if size == "" {
			size = "default"
		}
		className := "flex w-fit items-center justify-between gap-2 rounded-md border border-input bg-transparent px-3 py-2 text-sm whitespace-nowrap shadow-xs transition-[color,box-shadow] outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:cursor-not-allowed disabled:opacity-50 aria-invalid:border-destructive aria-invalid:ring-destructive/20 data-[placeholder]:text-muted-foreground data-[size=default]:h-9 data-[size=sm]:h-8 *:data-[slot=select-value]:line-clamp-1 *:data-[slot=select-value]:flex *:data-[slot=select-value]:items-center *:data-[slot=select-value]:gap-2 dark:bg-input/30 dark:hover:bg-input/50 dark:aria-invalid:ring-destructive/40 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4 [&_svg:not([class*='text-'])]:text-muted-foreground"
		attrs := attrsFromDOMProps(props.DOMProps, "select-trigger", className)
		attrs["type"] = "button"
		attrs["role"] = "combobox"
		attrs["data-size"] = size
		attrs["aria-haspopup"] = "listbox"
		if _, ok := attrs["aria-expanded"]; !ok {
			attrs["aria-expanded"] = "false"
		}
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := renderChildren(ctx, w, templ.GetChildren(ctx)); err != nil {
				return err
			}
			_, err := io.WriteString(w, `<svg xmlns="http://www.w3.org/2000/svg" class="size-4 opacity-50" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="m6 9 6 6 6-6"></path></svg>`)
			return err
		})
		return renderElement(ctx, w, "button", attrs, children)
	})
}

func SelectContent(props SelectContentProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "select-content", "relative z-50 max-h-(--radix-select-content-available-height) min-w-[8rem] origin-(--radix-select-content-transform-origin) overflow-x-hidden overflow-y-auto rounded-md border bg-popover p-1 text-popover-foreground shadow-md data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=closed]:zoom-out-95 data-[state=open]:animate-in data-[state=open]:fade-in-0 data-[state=open]:zoom-in-95 data-[position=popper]:data-[side=bottom]:translate-y-1 data-[position=popper]:data-[side=left]:-translate-x-1 data-[position=popper]:data-[side=right]:translate-x-1 data-[position=popper]:data-[side=top]:-translate-y-1")
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
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "select-label", "px-2 py-1.5 text-xs text-muted-foreground"), templ.GetChildren(ctx))
	})
}
func SelectItem(props DropdownMenuItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "relative flex w-full cursor-default items-center gap-2 rounded-sm py-1.5 pr-8 pl-2 text-sm outline-hidden select-none focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4 [&_svg:not([class*='text-'])]:text-muted-foreground *:[span]:last:flex *:[span]:last:items-center *:[span]:last:gap-2"
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
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if _, err := io.WriteString(w, `<span data-slot="select-item-indicator" class="absolute right-2 flex size-3.5 items-center justify-center">`); err != nil {
				return err
			}
			if selected {
				if _, err := io.WriteString(w, `<svg xmlns="http://www.w3.org/2000/svg" class="size-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M20 6 9 17l-5-5"></path></svg>`); err != nil {
					return err
				}
			}
			if _, err := io.WriteString(w, `</span><span>`); err != nil {
				return err
			}
			if err := renderChildren(ctx, w, templ.GetChildren(ctx)); err != nil {
				return err
			}
			_, err := io.WriteString(w, `</span>`)
			return err
		})
		return renderElement(ctx, w, "button", attrs, children)
	})
}
func SelectSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "select-separator", "pointer-events-none -mx-1 my-1 h-px bg-border"), nil)
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
