package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type TabsProps struct {
	DOMProps
	Value        string
	DefaultValue string
	Orientation  string
}

type TabsListProps struct {
	DOMProps
	Variant string
}

func Tabs(props TabsProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "tabs", "grid gap-2")
		attrs["data-tabs-root"] = "true"
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.DefaultValue != "" {
			attrs["data-default-value"] = props.DefaultValue
		}
		if props.Orientation != "" {
			attrs["data-orientation"] = props.Orientation
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func TabsList(props TabsListProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "inline-flex h-9 items-center justify-center rounded-lg bg-muted p-1 text-muted-foreground"
		if props.Variant == "line" {
			className = "inline-flex h-10 items-center gap-4 border-b border-border bg-transparent p-0 text-muted-foreground"
		}
		attrs := attrsFromDOMProps(props.DOMProps, "tabs-list", className)
		attrs["role"] = "tablist"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

type TabsTriggerProps struct {
	DOMProps
	Value    string
	Active   bool
	Disabled bool
}

func TabsTrigger(props TabsTriggerProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "tabs-trigger", "inline-flex items-center justify-center whitespace-nowrap rounded-md px-3 py-1.5 text-sm font-medium ring-offset-background transition-all focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 data-[state=active]:bg-background data-[state=active]:text-foreground data-[state=active]:shadow")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.Active {
			attrs["data-state"] = "active"
			attrs["aria-selected"] = "true"
			attrs["tabindex"] = "0"
		} else {
			attrs["data-state"] = "inactive"
			attrs["aria-selected"] = "false"
			attrs["tabindex"] = "-1"
		}
		attrs["role"] = "tab"
		if props.Disabled {
			attrs["disabled"] = true
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

type TabsContentProps struct {
	DOMProps
	Value  string
	Active bool
}

func TabsContent(props TabsContentProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "tabs-content", "mt-2 outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		attrs["role"] = "tabpanel"
		if props.Active {
			attrs["data-state"] = "active"
		} else {
			attrs["data-state"] = "inactive"
			attrs["hidden"] = true
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
