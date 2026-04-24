package ui

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/a-h/templ"
)

type AspectRatioProps struct {
	DOMProps
	Ratio float64
}

func AspectRatio(props AspectRatioProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		ratio := props.Ratio
		if ratio <= 0 {
			ratio = 1
		}

		attrs := attrsFromDOMProps(props.DOMProps, "aspect-ratio", "relative w-full overflow-hidden")
		style := fmt.Sprintf("aspect-ratio: %s;", strconv.FormatFloat(ratio, 'f', -1, 64))
		if existing, ok := attrs["style"]; ok {
			style = cn(templ.Classes(existing).String(), style)
		}
		attrs["style"] = style
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
