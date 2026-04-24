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
		className := "relative w-full"
		if props.Orientation == SliderOrientationVertical {
			className = "relative h-40"
		}
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
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
