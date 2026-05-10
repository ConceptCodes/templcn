package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type FieldOrientation string

const (
	FieldOrientationVertical   FieldOrientation = "vertical"
	FieldOrientationHorizontal FieldOrientation = "horizontal"
	FieldOrientationResponsive FieldOrientation = "responsive"
)

type FieldSetProps struct {
	DOMProps
	Orientation FieldOrientation
	Invalid     bool
	Disabled    bool
}

func FieldSet(props FieldSetProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "field-set", "flex flex-col gap-6 has-[>[data-slot=checkbox-group]]:gap-3 has-[>[data-slot=radio-group]]:gap-3")
		if props.Invalid {
			attrs["data-invalid"] = "true"
		}
		if props.Disabled {
			attrs["disabled"] = true
		}
		return renderElement(ctx, w, "fieldset", attrs, templ.GetChildren(ctx))
	})
}

func FieldLegend(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "field-legend", "mb-3 font-medium data-[variant=legend]:text-base data-[variant=label]:text-sm")
		if _, ok := attrs["data-variant"]; !ok {
			attrs["data-variant"] = "legend"
		}
		return renderElement(ctx, w, "legend", attrs, templ.GetChildren(ctx))
	})
}

func FieldGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "field-group", "group/field-group @container/field-group flex w-full flex-col gap-7 data-[slot=checkbox-group]:gap-3 [&>[data-slot=field-group]]:gap-4"), templ.GetChildren(ctx))
	})
}

func Field(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "group/field flex w-full gap-3 data-[invalid=true]:text-destructive flex-col [&>*]:w-full [&>.sr-only]:w-auto"
		if orientation, ok := props.Attrs["data-orientation"].(string); ok {
			switch FieldOrientation(orientation) {
			case FieldOrientationHorizontal:
				className = "group/field flex w-full gap-3 data-[invalid=true]:text-destructive flex-row items-center [&>[data-slot=field-label]]:flex-auto has-[>[data-slot=field-content]]:items-start has-[>[data-slot=field-content]]:[&>[role=checkbox],[role=radio]]:mt-px"
			case FieldOrientationResponsive:
				className = "group/field flex w-full gap-3 data-[invalid=true]:text-destructive flex-col @md/field-group:flex-row @md/field-group:items-center [&>*]:w-full @md/field-group:[&>*]:w-auto [&>.sr-only]:w-auto @md/field-group:[&>[data-slot=field-label]]:flex-auto @md/field-group:has-[>[data-slot=field-content]]:items-start @md/field-group:has-[>[data-slot=field-content]]:[&>[role=checkbox],[role=radio]]:mt-px"
			}
		}
		attrs := attrsFromDOMProps(props, "field", className)
		attrs["role"] = "group"
		if _, ok := attrs["data-orientation"]; !ok {
			attrs["data-orientation"] = "vertical"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func FieldContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "field-content", "group/field-content flex flex-1 flex-col gap-1.5 leading-snug"), templ.GetChildren(ctx))
	})
}

func FieldLabel(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "label", attrsFromDOMProps(props, "field-label", "group/field-label peer/field-label flex w-fit gap-2 leading-snug group-data-[disabled=true]/field:opacity-50 has-[>[data-slot=field]]:w-full has-[>[data-slot=field]]:flex-col has-[>[data-slot=field]]:rounded-md has-[>[data-slot=field]]:border [&>*]:data-[slot=field]:p-4 has-data-[state=checked]:border-primary has-data-[state=checked]:bg-primary/5 dark:has-data-[state=checked]:bg-primary/10"), templ.GetChildren(ctx))
	})
}

func FieldTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "field-label", "flex w-fit items-center gap-2 text-sm leading-snug font-medium group-data-[disabled=true]/field:opacity-50"), templ.GetChildren(ctx))
	})
}

func FieldDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "p", attrsFromDOMProps(props, "field-description", "text-sm leading-normal font-normal text-muted-foreground group-has-[[data-orientation=horizontal]]/field:text-balance last:mt-0 nth-last-2:-mt-1 [[data-variant=legend]+&]:-mt-1.5 [&>a]:underline [&>a]:underline-offset-4 [&>a:hover]:text-primary"), templ.GetChildren(ctx))
	})
}

func FieldSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "field-separator", "relative -my-2 h-5 text-sm group-data-[variant=outline]/field-group:-mb-2")
		children := templ.GetChildren(ctx)
		attrs["data-content"] = children != nil
		content := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := Separator(SeparatorProps{DOMProps: DOMProps{Class: "absolute inset-0 top-1/2"}}).Render(ctx, w); err != nil {
				return err
			}
			if children == nil {
				return nil
			}
			return renderElement(ctx, w, "span", templ.Attributes{
				"class":     "relative mx-auto block w-fit bg-background px-2 text-muted-foreground",
				"data-slot": "field-separator-content",
			}, children)
		})
		return renderElement(ctx, w, "div", attrs, content)
	})
}

type FieldErrorProps struct {
	DOMProps
	Errors []string
}

func FieldError(props FieldErrorProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		text := ""
		if len(props.Errors) > 0 {
			text = props.Errors[0]
		}
		attrs := attrsFromDOMProps(props.DOMProps, "field-error", "text-sm font-normal text-destructive")
		attrs["role"] = "alert"
		return renderTextElement(ctx, w, "div", attrs, text)
	})
}
