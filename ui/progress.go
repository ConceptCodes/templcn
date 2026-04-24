package ui

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/a-h/templ"
)

type ProgressProps struct {
	DOMProps
	Value float64
	Max   float64
}

func Progress(props ProgressProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		max := props.Max
		if max <= 0 {
			max = 100
		}
		value := props.Value
		if value < 0 {
			value = 0
		}
		if value > max {
			value = max
		}
		attrs := attrsFromDOMProps(props.DOMProps, "progress", "relative h-2 w-full overflow-hidden rounded-full bg-muted")
		attrs["role"] = "progressbar"
		attrs["aria-valuemin"] = "0"
		attrs["aria-valuemax"] = strconv.FormatFloat(max, 'f', -1, 64)
		attrs["aria-valuenow"] = strconv.FormatFloat(value, 'f', -1, 64)

		bar := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			barAttrs := templ.Attributes{
				"class": "h-full w-full flex-1 bg-primary transition-all",
				"style": fmt.Sprintf("transform: translateX(-%s%%);", strconv.FormatFloat(100-(value/max*100), 'f', -1, 64)),
			}
			return renderElement(ctx, w, "div", barAttrs, nil)
		})

		return renderElement(ctx, w, "div", attrs, bar)
	})
}
