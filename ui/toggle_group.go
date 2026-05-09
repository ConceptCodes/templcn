package ui

import (
	"context"
	"io"
	"strings"

	"github.com/a-h/templ"
)

type ToggleGroupProps struct {
	DOMProps
	Type         string
	Value        []string
	DefaultValue []string
	Variant      string
	Size         string
	Spacing      string
}

type ToggleGroupItemProps struct {
	DOMProps
	Value    string
	Disabled bool
}

type toggleGroupRenderState struct {
	values map[string]bool
}

type toggleGroupRenderStateKey struct{}

func ToggleGroup(props ToggleGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		values := props.Value
		if len(values) == 0 {
			values = props.DefaultValue
		}
		selected := make(map[string]bool, len(values))
		for _, value := range values {
			selected[value] = true
		}
		attrs := attrsFromDOMProps(props.DOMProps, "toggle-group", "inline-flex gap-1")
		attrs["role"] = "group"
		if props.Type != "" {
			attrs["data-type"] = props.Type
		}
		if len(values) > 0 {
			attrs["data-value"] = strings.Join(values, " ")
		}
		if len(props.DefaultValue) > 0 {
			attrs["data-default-value"] = props.DefaultValue
		}
		if props.Variant != "" {
			attrs["data-variant"] = props.Variant
		}
		if props.Size != "" {
			attrs["data-size"] = props.Size
		}
		if props.Spacing != "" {
			attrs["data-spacing"] = props.Spacing
		}
		ctx = context.WithValue(ctx, toggleGroupRenderStateKey{}, toggleGroupRenderState{values: selected})
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func ToggleGroupItem(props ToggleGroupItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "toggle-group-item", "rounded-md")
		state, _ := ctx.Value(toggleGroupRenderStateKey{}).(toggleGroupRenderState)
		pressed := props.Value != "" && state.values[props.Value]
		if _, ok := attrs["type"]; !ok {
			attrs["type"] = "button"
		}
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		attrs["data-state"] = map[bool]string{true: "on", false: "off"}[pressed]
		attrs["aria-pressed"] = map[bool]string{true: "true", false: "false"}[pressed]
		if props.Disabled {
			attrs["disabled"] = true
			attrs["aria-disabled"] = "true"
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
