package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type ComboboxProps struct {
	DOMProps
	Name         string
	Value        string
	Values       []string
	DefaultValue string
	Multiple     bool
	Placeholder  string
	ShowTrigger  bool
	ShowClear    bool
	Filter       string
	Open         bool
	DefaultOpen  bool
	AnchorID     string
	Side         string
	Align        string
	SideOffset   string
}

func Combobox(props ComboboxProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "combobox", "relative")
		if props.Name != "" {
			attrs["data-name"] = props.Name
		}
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if len(props.Values) > 0 {
			attrs["data-values"] = props.Values
		}
		if props.DefaultValue != "" {
			attrs["data-default-value"] = props.DefaultValue
		}
		if props.Multiple {
			attrs["data-multiple"] = "true"
		}
		if props.Placeholder != "" {
			attrs["data-placeholder"] = props.Placeholder
		}
		if props.ShowTrigger {
			attrs["data-show-trigger"] = "true"
		}
		if props.ShowClear {
			attrs["data-show-clear"] = "true"
		}
		if props.Filter != "" {
			attrs["data-filter"] = props.Filter
		}
		if props.Open {
			attrs["data-open"] = "true"
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		if props.AnchorID != "" {
			attrs["data-anchor-id"] = props.AnchorID
		}
		if props.Side != "" {
			attrs["data-side"] = props.Side
		}
		if props.Align != "" {
			attrs["data-align"] = props.Align
		}
		if props.SideOffset != "" {
			attrs["data-side-offset"] = props.SideOffset
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func ComboboxValue(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "combobox-value", ""), templ.GetChildren(ctx))
	})
}
func ComboboxTrigger(props SelectTriggerProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "flex h-9 w-full items-center justify-between rounded-md border border-input bg-background px-3 py-2 text-sm shadow-xs outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:cursor-not-allowed disabled:opacity-50"
		if props.Size == "sm" {
			className = cn(className, "h-8")
		}
		return renderElement(ctx, w, "button", attrsFromDOMProps(props.DOMProps, "combobox-trigger", className), templ.GetChildren(ctx))
	})
}
func ComboboxClear(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "combobox-clear", ""), templ.GetChildren(ctx))
	})
}
func ComboboxInput(props InputProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "combobox-input", inputClasses)
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
func ComboboxContent(props SelectContentProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "combobox-content", "z-50 min-w-32 rounded-md border bg-popover p-1 text-popover-foreground shadow-md")
		if props.Position != "" {
			attrs["data-position"] = props.Position
		}
		if props.Align != "" {
			attrs["data-align"] = props.Align
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func ComboboxList(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "combobox-list", "grid gap-1"), templ.GetChildren(ctx))
	})
}
func ComboboxItem(props DropdownMenuItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "relative flex cursor-pointer select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground"
		if props.Inset {
			className = cn(className, "pl-8")
		}
		if props.Variant == "destructive" {
			className = cn(className, "text-destructive")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "combobox-item", className)
		attrs["role"] = "option"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
func ComboboxGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "combobox-group", "grid gap-1"), templ.GetChildren(ctx))
	})
}
func ComboboxLabel(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "combobox-label", "px-2 py-1.5 text-sm font-semibold"), templ.GetChildren(ctx))
	})
}
func ComboboxCollection(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "combobox-collection", "grid gap-1"), templ.GetChildren(ctx))
	})
}
func ComboboxEmpty(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "combobox-empty", "py-6 text-center text-sm"), templ.GetChildren(ctx))
	})
}
func ComboboxSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "combobox-separator", "-mx-1 my-1 h-px bg-border"), nil)
	})
}
func ComboboxChips(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "combobox-chips", "flex flex-wrap gap-2"), templ.GetChildren(ctx))
	})
}
func ComboboxChip(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "combobox-chip", "inline-flex items-center rounded-md bg-muted px-2 py-1 text-xs"), templ.GetChildren(ctx))
	})
}
func ComboboxChipsInput(props InputProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "combobox-chips-input", inputClasses)
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
		return renderVoidElement(ctx, w, "input", attrs)
	})
}
