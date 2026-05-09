package ui

import (
	"context"
	"io"
	"strconv"

	"github.com/a-h/templ"
)

type InputOTPProps struct {
	DOMProps
	Name         string
	Value        string
	DefaultValue string
	MaxLength    int
	Pattern      string
	Disabled     bool
}

func InputOTP(props InputOTPProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		value := props.Value
		if value == "" {
			value = props.DefaultValue
		}
		attrs := attrsFromDOMProps(props.DOMProps, "input-otp", "flex items-center gap-2")
		if props.Name != "" {
			attrs["data-name"] = props.Name
		}
		if value != "" {
			attrs["data-value"] = value
		}
		if props.DefaultValue != "" {
			attrs["data-default-value"] = props.DefaultValue
		}
		if props.MaxLength > 0 {
			attrs["data-maxlength"] = strconv.Itoa(props.MaxLength)
		}
		if props.Pattern != "" {
			attrs["data-pattern"] = props.Pattern
		}
		if props.Disabled {
			attrs["data-disabled"] = "true"
		}
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if props.Name != "" {
				inputAttrs := templ.Attributes{"type": "hidden", "name": props.Name, "value": value}
				if err := renderVoidElement(ctx, w, "input", inputAttrs); err != nil {
					return err
				}
			}
			return renderChildren(ctx, w, templ.GetChildren(ctx))
		})
		return renderElement(ctx, w, "div", attrs, children)
	})
}

func InputOTPGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "input-otp-group", "flex items-center gap-2"), templ.GetChildren(ctx))
	})
}

type InputOTPSlotProps struct {
	DOMProps
	Index int
	Value string
}

func InputOTPSlot(props InputOTPSlotProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "input-otp-slot", "flex size-10 items-center justify-center rounded-md border border-input bg-background text-sm shadow-xs")
		if props.Index >= 0 {
			attrs["data-index"] = strconv.Itoa(props.Index)
		}
		attrs["tabindex"] = "0"
		attrs["role"] = "textbox"
		attrs["aria-label"] = "Digit " + strconv.Itoa(props.Index+1)
		return renderTextElement(ctx, w, "div", attrs, props.Value)
	})
}

func InputOTPSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "input-otp-separator", "mx-1 h-px w-3 bg-border"), nil)
	})
}
