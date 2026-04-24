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

func ComboboxValue(props DOMProps) templ.Component             { return SelectValue(props) }
func ComboboxTrigger(props SelectTriggerProps) templ.Component { return SelectTrigger(props) }
func ComboboxClear(props DOMProps) templ.Component             { return DialogTrigger(props) }
func ComboboxInput(props InputProps) templ.Component           { return Input(props) }
func ComboboxContent(props SelectContentProps) templ.Component { return SelectContent(props) }
func ComboboxList(props DOMProps) templ.Component              { return SelectGroup(props) }
func ComboboxItem(props DropdownMenuItemProps) templ.Component { return DropdownMenuItem(props) }
func ComboboxGroup(props DOMProps) templ.Component             { return SelectGroup(props) }
func ComboboxLabel(props DOMProps) templ.Component             { return SelectLabel(props) }
func ComboboxCollection(props DOMProps) templ.Component        { return SelectGroup(props) }
func ComboboxEmpty(props DOMProps) templ.Component             { return Empty(props) }
func ComboboxSeparator(props DOMProps) templ.Component         { return SelectSeparator(props) }
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
func ComboboxChipsInput(props InputProps) templ.Component { return Input(props) }
