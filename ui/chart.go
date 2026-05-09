package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type ChartContainerProps struct {
	DOMProps
	Config        string
	Library       string
	Series        string
	InitialWidth  int
	InitialHeight int
}

func ChartContainer(props ChartContainerProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "chart-container", "relative w-full")
		attrs["role"] = "img"
		if props.Config != "" {
			attrs["data-config"] = props.Config
		}
		if props.Library != "" {
			attrs["data-library"] = props.Library
		}
		if props.Series != "" {
			attrs["data-series"] = props.Series
		}
		if props.InitialWidth > 0 {
			attrs["data-initial-width"] = props.InitialWidth
		}
		if props.InitialHeight > 0 {
			attrs["data-initial-height"] = props.InitialHeight
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func ChartStyle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "style", attrsFromDOMProps(props, "chart-style", ""), templ.GetChildren(ctx))
	})
}
func ChartTooltip(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "chart-tooltip", ""), templ.GetChildren(ctx))
	})
}
func ChartTooltipContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "chart-tooltip-content", "rounded-md border bg-background p-2 text-xs shadow"), templ.GetChildren(ctx))
	})
}
func ChartLegend(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "chart-legend", "flex flex-wrap gap-2"), templ.GetChildren(ctx))
	})
}
func ChartLegendContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "chart-legend-content", ""), templ.GetChildren(ctx))
	})
}
