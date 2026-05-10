package ui

import (
	"context"
	"io"
	"strconv"

	"github.com/a-h/templ"
)

type SliderOrientation string

const (
	SliderOrientationHorizontal SliderOrientation = "horizontal"
	SliderOrientationVertical   SliderOrientation = "vertical"
)

type SliderProps struct {
	DOMProps
	Name         string
	Value        []float64
	DefaultValue []float64
	Min          float64
	Max          float64
	Step         float64
	Disabled     bool
	Orientation  SliderOrientation
}

func Slider(props SliderProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		if props.Step == 0 {
			props.Step = 1
		}
		if props.Max == 0 {
			props.Max = 100
		}
		className := "relative flex w-full touch-none items-center select-none data-[disabled]:opacity-50 data-[orientation=vertical]:h-full data-[orientation=vertical]:min-h-44 data-[orientation=vertical]:w-auto data-[orientation=vertical]:flex-col"
		attrs := attrsFromDOMProps(props.DOMProps, "slider", className)
		if props.Name != "" {
			attrs["data-name"] = props.Name
		}
		if props.Min != 0 {
			attrs["data-min"] = strconv.FormatFloat(props.Min, 'f', -1, 64)
		}
		attrs["data-max"] = strconv.FormatFloat(props.Max, 'f', -1, 64)
		attrs["data-step"] = strconv.FormatFloat(props.Step, 'f', -1, 64)
		if props.Disabled {
			attrs["data-disabled"] = "true"
		}
		if props.Orientation == SliderOrientationVertical {
			attrs["data-orientation"] = "vertical"
		} else {
			attrs["data-orientation"] = "horizontal"
		}
		value := props.Min
		if len(props.DefaultValue) > 0 {
			value = props.DefaultValue[0]
		}
		if len(props.Value) > 0 {
			value = props.Value[0]
		}
		attrs["data-value"] = strconv.FormatFloat(value, 'f', -1, 64)
		percent := 0.0
		if props.Max > props.Min {
			percent = ((value - props.Min) / (props.Max - props.Min)) * 100
		}
		if percent < 0 {
			percent = 0
		}
		if percent > 100 {
			percent = 100
		}
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			track := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				rangeAttrs := templ.Attributes{
					"data-slot":        "slider-range",
					"data-orientation": attrs["data-orientation"],
					"class":            "absolute bg-primary data-[orientation=horizontal]:h-full data-[orientation=vertical]:w-full",
				}
				if props.Orientation == SliderOrientationVertical {
					rangeAttrs["style"] = "height: " + strconv.FormatFloat(percent, 'f', -1, 64) + "%;"
				} else {
					rangeAttrs["style"] = "width: " + strconv.FormatFloat(percent, 'f', -1, 64) + "%;"
				}
				return renderElement(ctx, w, "span", rangeAttrs, nil)
			})
			if err := renderElement(ctx, w, "span", templ.Attributes{
				"data-slot":        "slider-track",
				"data-orientation": attrs["data-orientation"],
				"class":            "relative grow overflow-hidden rounded-full bg-muted data-[orientation=horizontal]:h-1.5 data-[orientation=horizontal]:w-full data-[orientation=vertical]:h-full data-[orientation=vertical]:w-1.5",
			}, track); err != nil {
				return err
			}
			thumbCount := len(props.Value)
			if thumbCount == 0 {
				thumbCount = len(props.DefaultValue)
			}
			if thumbCount == 0 {
				thumbCount = 1
			}
			for i := 0; i < thumbCount; i++ {
				if err := renderElement(ctx, w, "span", templ.Attributes{
					"data-slot": "slider-thumb",
					"class":     "block size-4 shrink-0 rounded-full border border-primary bg-white shadow-sm ring-ring/50 transition-[color,box-shadow] hover:ring-4 focus-visible:ring-4 focus-visible:outline-hidden disabled:pointer-events-none disabled:opacity-50",
				}, nil); err != nil {
					return err
				}
			}
			inputAttrs := templ.Attributes{
				"type":          "range",
				"role":          "slider",
				"min":           strconv.FormatFloat(props.Min, 'f', -1, 64),
				"max":           strconv.FormatFloat(props.Max, 'f', -1, 64),
				"step":          strconv.FormatFloat(props.Step, 'f', -1, 64),
				"value":         strconv.FormatFloat(value, 'f', -1, 64),
				"aria-valuemin": strconv.FormatFloat(props.Min, 'f', -1, 64),
				"aria-valuemax": strconv.FormatFloat(props.Max, 'f', -1, 64),
				"aria-valuenow": strconv.FormatFloat(value, 'f', -1, 64),
				"class":         "sr-only",
			}
			if props.Name != "" {
				inputAttrs["name"] = props.Name
			}
			if props.Disabled {
				inputAttrs["disabled"] = true
			}
			if err := renderVoidElement(ctx, w, "input", inputAttrs); err != nil {
				return err
			}
			return renderChildren(ctx, w, templ.GetChildren(ctx))
		})
		return renderElement(ctx, w, "div", attrs, children)
	})
}
