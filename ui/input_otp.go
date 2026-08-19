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
		ctx, ownChildren := childrenFromContext(ctx)
		value := props.Value
		if value == "" {
			value = props.DefaultValue
		}
		attrs := attrsFromDOMProps(props.DOMProps, "input-otp", "flex items-center gap-2 has-disabled:opacity-50 disabled:cursor-not-allowed")
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
			return renderChildren(ctx, w, ownChildren)
		})
		return renderElement(ctx, w, "div", attrs, children)
	})
}

func InputOTPGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "input-otp-group", "flex items-center"), templ.GetChildren(ctx))
	})
}

type InputOTPSlotProps struct {
	DOMProps
	Index int
	Value string
}

func InputOTPSlot(props InputOTPSlotProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "input-otp-slot", "relative flex h-9 w-9 items-center justify-center border-y border-r border-input text-sm shadow-xs transition-all outline-none first:rounded-l-md first:border-l last:rounded-r-md aria-invalid:border-destructive data-[active=true]:z-10 data-[active=true]:border-ring data-[active=true]:ring-[3px] data-[active=true]:ring-ring/50 data-[active=true]:aria-invalid:border-destructive data-[active=true]:aria-invalid:ring-destructive/20 dark:bg-input/30 dark:data-[active=true]:aria-invalid:ring-destructive/40")
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
		attrs := attrsFromDOMProps(props, "input-otp-separator", "")
		attrs["role"] = "separator"
		children := templ.GetChildren(ctx)
		if children == nil {
			children = templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				_, err := io.WriteString(w, `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" aria-hidden="true"><path d="M5 12h14"></path></svg>`)
				return err
			})
		}
		return renderElement(ctx, w, "div", attrs, children)
	})
}
