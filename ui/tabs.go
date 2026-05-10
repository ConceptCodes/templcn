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

type tabsRenderState struct {
	value string
}

type tabsRenderStateKey struct{}

func Tabs(props TabsProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		value := props.Value
		if value == "" {
			value = props.DefaultValue
		}
		attrs := attrsFromDOMProps(props.DOMProps, "tabs", "group/tabs flex gap-2 data-[orientation=horizontal]:flex-col")
		attrs["data-tabs-root"] = "true"
		if value != "" {
			attrs["data-value"] = value
		}
		if props.DefaultValue != "" {
			attrs["data-default-value"] = props.DefaultValue
		}
		orientation := props.Orientation
		if orientation == "" {
			orientation = "horizontal"
		}
		attrs["data-orientation"] = orientation
		ctx = context.WithValue(ctx, tabsRenderStateKey{}, tabsRenderState{value: value})
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func TabsList(props TabsListProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "group/tabs-list inline-flex w-fit items-center justify-center rounded-lg bg-muted p-[3px] text-muted-foreground group-data-[orientation=horizontal]/tabs:h-9 group-data-[orientation=vertical]/tabs:h-fit group-data-[orientation=vertical]/tabs:flex-col data-[variant=line]:rounded-none"
		if props.Variant == "line" {
			className = "group/tabs-list inline-flex w-fit items-center justify-center gap-1 rounded-lg p-[3px] text-muted-foreground group-data-[orientation=horizontal]/tabs:h-9 group-data-[orientation=vertical]/tabs:h-fit group-data-[orientation=vertical]/tabs:flex-col data-[variant=line]:rounded-none bg-transparent"
		}
		attrs := attrsFromDOMProps(props.DOMProps, "tabs-list", className)
		attrs["role"] = "tablist"
		if props.Variant == "line" {
			attrs["data-variant"] = "line"
		} else {
			attrs["data-variant"] = "default"
		}
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
		attrs := attrsFromDOMProps(props.DOMProps, "tabs-trigger", "relative inline-flex h-[calc(100%-1px)] flex-1 items-center justify-center gap-1.5 rounded-md border border-transparent px-2 py-1 text-sm font-medium whitespace-nowrap text-foreground/60 transition-all group-data-[orientation=vertical]/tabs:w-full group-data-[orientation=vertical]/tabs:justify-start hover:text-foreground focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 focus-visible:outline-1 focus-visible:outline-ring disabled:pointer-events-none disabled:opacity-50 group-data-[variant=default]/tabs-list:data-[state=active]:shadow-sm group-data-[variant=line]/tabs-list:data-[state=active]:shadow-none dark:text-muted-foreground dark:hover:text-foreground [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4 group-data-[variant=line]/tabs-list:bg-transparent group-data-[variant=line]/tabs-list:data-[state=active]:bg-transparent dark:group-data-[variant=line]/tabs-list:data-[state=active]:border-transparent dark:group-data-[variant=line]/tabs-list:data-[state=active]:bg-transparent data-[state=active]:bg-background data-[state=active]:text-foreground dark:data-[state=active]:border-input dark:data-[state=active]:bg-input/30 dark:data-[state=active]:text-foreground after:absolute after:bg-foreground after:opacity-0 after:transition-opacity group-data-[orientation=horizontal]/tabs:after:inset-x-0 group-data-[orientation=horizontal]/tabs:after:bottom-[-5px] group-data-[orientation=horizontal]/tabs:after:h-0.5 group-data-[orientation=vertical]/tabs:after:inset-y-0 group-data-[orientation=vertical]/tabs:after:-right-1 group-data-[orientation=vertical]/tabs:after:w-0.5 group-data-[variant=line]/tabs-list:data-[state=active]:after:opacity-100")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		active := props.Active || (props.Value != "" && props.Value == tabsValueFromContext(ctx))
		if active {
			attrs["data-state"] = "active"
			attrs["aria-selected"] = "true"
			attrs["tabindex"] = "0"
		} else {
			attrs["data-state"] = "inactive"
			attrs["aria-selected"] = "false"
			attrs["tabindex"] = "-1"
		}
		attrs["role"] = "tab"
		attrs["type"] = "button"
		if props.Disabled {
			attrs["disabled"] = true
			attrs["aria-disabled"] = "true"
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
		attrs := attrsFromDOMProps(props.DOMProps, "tabs-content", "flex-1 outline-none")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		attrs["role"] = "tabpanel"
		active := props.Active || (props.Value != "" && props.Value == tabsValueFromContext(ctx))
		if active {
			attrs["data-state"] = "active"
		} else {
			attrs["data-state"] = "inactive"
			attrs["hidden"] = true
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func tabsValueFromContext(ctx context.Context) string {
	state, ok := ctx.Value(tabsRenderStateKey{}).(tabsRenderState)
	if !ok {
		return ""
	}
	return state.value
}
