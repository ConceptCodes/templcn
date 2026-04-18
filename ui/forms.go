package ui

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/a-h/templ"
)

type InputProps struct {
	DOMProps
	Type        string
	Name        string
	Value       string
	Placeholder string
	Disabled    bool
	Required    bool
	Invalid     bool
}

func Input(props InputProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "input", "flex h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-base shadow-xs transition-[color,box-shadow] outline-none file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:cursor-not-allowed disabled:opacity-50 md:text-sm")
		if props.Type == "" {
			props.Type = "text"
		}
		attrs["type"] = props.Type
		if props.Name != "" {
			attrs["name"] = props.Name
		}
		if props.Value != "" {
			attrs["value"] = props.Value
		}
		if props.Placeholder != "" {
			attrs["placeholder"] = props.Placeholder
		}
		if props.Disabled {
			attrs["disabled"] = true
		}
		if props.Required {
			attrs["required"] = true
		}
		if props.Invalid {
			attrs["aria-invalid"] = "true"
		}
		return renderVoidElement(ctx, w, "input", attrs)
	})
}

type TextareaProps struct {
	DOMProps
	Name        string
	Value       string
	Placeholder string
	Rows        int
	Disabled    bool
	Required    bool
	Invalid     bool
}

func Textarea(props TextareaProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "textarea", "flex min-h-20 w-full rounded-md border border-input bg-transparent px-3 py-2 text-base shadow-xs placeholder:text-muted-foreground focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 outline-none disabled:cursor-not-allowed disabled:opacity-50 md:text-sm")
		if props.Name != "" {
			attrs["name"] = props.Name
		}
		if props.Placeholder != "" {
			attrs["placeholder"] = props.Placeholder
		}
		if props.Rows > 0 {
			attrs["rows"] = strconv.Itoa(props.Rows)
		}
		if props.Disabled {
			attrs["disabled"] = true
		}
		if props.Required {
			attrs["required"] = true
		}
		if props.Invalid {
			attrs["aria-invalid"] = "true"
		}

		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if props.Value == "" {
				return nil
			}
			_, err := io.WriteString(w, templ.EscapeString(props.Value))
			return err
		})
		return renderElement(ctx, w, "textarea", attrs, children)
	})
}

type CheckboxProps struct {
	DOMProps
	Name           string
	Value          string
	Checked        bool
	DefaultChecked bool
	Disabled       bool
	Required       bool
	Invalid        bool
}

func Checkbox(props CheckboxProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "checkbox", "h-4 w-4 shrink-0 rounded-sm border border-input bg-background shadow-xs outline-none transition-all focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:cursor-not-allowed disabled:opacity-50 checked:border-primary checked:bg-primary checked:text-primary-foreground")
		attrs["type"] = "checkbox"
		if props.Name != "" {
			attrs["name"] = props.Name
		}
		if props.Value != "" {
			attrs["value"] = props.Value
		}
		if props.Checked {
			attrs["checked"] = true
		}
		if props.DefaultChecked {
			attrs["defaultChecked"] = true
		}
		if props.Disabled {
			attrs["disabled"] = true
		}
		if props.Required {
			attrs["required"] = true
		}
		if props.Invalid {
			attrs["aria-invalid"] = "true"
		}
		return renderVoidElement(ctx, w, "input", attrs)
	})
}

type SwitchProps struct {
	DOMProps
	Name           string
	Value          string
	Checked        bool
	DefaultChecked bool
	Disabled       bool
	Required       bool
	Size           string
}

func Switch(props SwitchProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		sizeClass := "h-5 w-9"
		if props.Size == "sm" {
			sizeClass = "h-4 w-7"
		}
		attrs := attrsFromDOMProps(props.DOMProps, "switch", cn("peer appearance-none shrink-0 rounded-full border border-transparent bg-input shadow-xs outline-none transition-all focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:cursor-not-allowed disabled:opacity-50 checked:bg-primary", sizeClass, "after:pointer-events-none after:block after:rounded-full after:bg-background after:shadow-sm after:transition-transform after:content-[''] checked:after:translate-x-4 rtl:checked:after:-translate-x-4"))
		attrs["type"] = "checkbox"
		attrs["role"] = "switch"
		if props.Name != "" {
			attrs["name"] = props.Name
		}
		if props.Value != "" {
			attrs["value"] = props.Value
		}
		if props.Checked {
			attrs["checked"] = true
		}
		if props.DefaultChecked {
			attrs["defaultChecked"] = true
		}
		if props.Disabled {
			attrs["disabled"] = true
		}
		if props.Required {
			attrs["required"] = true
		}
		return renderVoidElement(ctx, w, "input", attrs)
	})
}

type RadioGroupProps struct {
	DOMProps
	Name         string
	Value        string
	DefaultValue string
	Disabled     bool
	Required     bool
}

func RadioGroup(props RadioGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "radio-group", "grid gap-2")
		attrs["role"] = "radiogroup"
		if props.Name != "" {
			attrs["data-name"] = props.Name
		}
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.DefaultValue != "" {
			attrs["data-default-value"] = props.DefaultValue
		}
		if props.Disabled {
			attrs["data-disabled"] = "true"
		}
		if props.Required {
			attrs["aria-required"] = "true"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

type RadioGroupItemProps struct {
	DOMProps
	Name     string
	Value    string
	Disabled bool
	Label    string
}

func RadioGroupItem(props RadioGroupItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "flex items-center gap-2 text-sm"
		attrs := attrsFromDOMProps(props.DOMProps, "radio-group-item", className)
		children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			inputAttrs := templ.Attributes{
				"type":  "radio",
				"class": "h-4 w-4 border border-input text-primary focus-visible:ring-[3px] focus-visible:ring-ring/50",
				"value": props.Value,
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
			if props.Label != "" {
				_, err := io.WriteString(w, templ.EscapeString(props.Label))
				return err
			}
			return renderChildren(ctx, w, templ.GetChildren(ctx))
		})
		return renderElement(ctx, w, "label", attrs, children)
	})
}

type NativeSelectProps struct {
	DOMProps
	Name     string
	Value    string
	Size     string
	Disabled bool
	Required bool
	Invalid  bool
}

func NativeSelect(props NativeSelectProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		sizeClass := "h-9"
		if props.Size == "sm" {
			sizeClass = "h-8"
		}
		attrs := attrsFromDOMProps(props.DOMProps, "native-select", cn("flex w-full appearance-none rounded-md border border-input bg-background px-3 py-2 text-sm shadow-xs outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:cursor-not-allowed disabled:opacity-50", sizeClass))
		if props.Name != "" {
			attrs["name"] = props.Name
		}
		if props.Value != "" {
			attrs["value"] = props.Value
		}
		if props.Disabled {
			attrs["disabled"] = true
		}
		if props.Required {
			attrs["required"] = true
		}
		if props.Invalid {
			attrs["aria-invalid"] = "true"
		}
		return renderElement(ctx, w, "select", attrs, templ.GetChildren(ctx))
	})
}

func NativeSelectOption(props DOMProps, value, text string, selected bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "native-select-option", "")
		attrs["value"] = value
		if selected {
			attrs["selected"] = true
		}
		return renderTextElement(ctx, w, "option", attrs, text)
	})
}

func NativeSelectOptGroup(props DOMProps, label string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "native-select-optgroup", "")
		if label != "" {
			attrs["label"] = label
		}
		return renderElement(ctx, w, "optgroup", attrs, templ.GetChildren(ctx))
	})
}

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
		orientation := props.Orientation
		if orientation == "" {
			orientation = FieldOrientationVertical
		}
		attrs := attrsFromDOMProps(props.DOMProps, "fieldset", cn("grid gap-2 rounded-lg border border-border p-4", map[FieldOrientation]string{
			FieldOrientationVertical:   "grid",
			FieldOrientationHorizontal: "grid gap-4 md:grid-cols-[180px_1fr]",
			FieldOrientationResponsive: "grid gap-4 md:grid-cols-[180px_1fr]",
		}[orientation]))
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
		return renderElement(ctx, w, "legend", attrsFromDOMProps(props, "field-legend", "px-1 text-sm font-medium"), templ.GetChildren(ctx))
	})
}

func FieldGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "field-group", "grid gap-4"), templ.GetChildren(ctx))
	})
}

func Field(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "field", "grid gap-2"), templ.GetChildren(ctx))
	})
}

func FieldContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "field-content", "grid gap-1.5"), templ.GetChildren(ctx))
	})
}

func FieldLabel(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "label", attrsFromDOMProps(props, "field-label", "text-sm font-medium"), templ.GetChildren(ctx))
	})
}

func FieldTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "field-title", "text-sm font-medium"), templ.GetChildren(ctx))
	})
}

func FieldDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "field-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func FieldSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderVoidElement(ctx, w, "hr", attrsFromDOMProps(props, "field-separator", "my-2 border-border"))
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
		attrs := attrsFromDOMProps(props.DOMProps, "field-error", "text-sm text-destructive")
		return renderTextElement(ctx, w, "p", attrs, text)
	})
}

type FormProps struct {
	DOMProps
	Name        string
	IDValue     string
	DescribedBy string
	Invalid     bool
	Message     string
}

func Form(props FormProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "form", "grid gap-4")
		if props.IDValue != "" {
			attrs["id"] = props.IDValue
		}
		if props.Name != "" {
			attrs["name"] = props.Name
		}
		if props.DescribedBy != "" {
			attrs["aria-describedby"] = props.DescribedBy
		}
		if props.Invalid {
			attrs["aria-invalid"] = "true"
		}
		if props.Message != "" {
			attrs["data-message"] = props.Message
		}
		tag := "form"
		if props.Element != "" {
			tag = props.Element
		}
		return renderElement(ctx, w, tag, attrs, templ.GetChildren(ctx))
	})
}

func FormItem(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "form-item", "grid gap-2"), templ.GetChildren(ctx))
	})
}

func FormLabel(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "label", attrsFromDOMProps(props, "form-label", "text-sm font-medium leading-none"), templ.GetChildren(ctx))
	})
}

func FormControl(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "form-control", "grid gap-1"), templ.GetChildren(ctx))
	})
}

func FormDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "p", attrsFromDOMProps(props, "form-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func FormMessage(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "form-message", "text-sm text-destructive")
		return renderElement(ctx, w, "p", attrs, templ.GetChildren(ctx))
	})
}

type ButtonGroupOrientation string

const (
	ButtonGroupOrientationHorizontal ButtonGroupOrientation = "horizontal"
	ButtonGroupOrientationVertical   ButtonGroupOrientation = "vertical"
)

type ButtonGroupProps struct {
	DOMProps
	Orientation ButtonGroupOrientation
}

func ButtonGroup(props ButtonGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		orientation := props.Orientation
		if orientation == "" {
			orientation = ButtonGroupOrientationHorizontal
		}
		className := "inline-flex"
		if orientation == ButtonGroupOrientationVertical {
			className = "inline-flex flex-col"
		}
		return renderElement(ctx, w, "div", attrsFromDOMProps(props.DOMProps, "button-group", className), templ.GetChildren(ctx))
	})
}

func ButtonGroupText(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "button-group-text", "px-3 py-2 text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func ButtonGroupSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "button-group-separator", "w-px self-stretch bg-border"), nil)
	})
}

type InputGroupProps struct {
	DOMProps
}

func InputGroup(props InputGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props.DOMProps, "input-group", "flex rounded-md shadow-xs ring-1 ring-inset ring-input"), templ.GetChildren(ctx))
	})
}

type InputGroupAddonAlign string

const (
	InputGroupAddonInlineStart InputGroupAddonAlign = "inline-start"
	InputGroupAddonInlineEnd   InputGroupAddonAlign = "inline-end"
	InputGroupAddonBlockStart  InputGroupAddonAlign = "block-start"
	InputGroupAddonBlockEnd    InputGroupAddonAlign = "block-end"
)

type InputGroupAddonProps struct {
	DOMProps
	Align InputGroupAddonAlign
}

func InputGroupAddon(props InputGroupAddonProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props.DOMProps, "input-group-addon", "inline-flex items-center rounded-md border-0 px-3 py-2 text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

type InputGroupButtonProps struct {
	DOMProps
	Size ButtonSize
}

func InputGroupButton(props InputGroupButtonProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := buttonClasses(ButtonVariantDefault, props.Size, "rounded-none border-0 shadow-none")
		attrs := attrsFromDOMProps(props.DOMProps, "input-group-button", className)
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

type InputGroupTextProps struct {
	DOMProps
}

func InputGroupText(props InputGroupTextProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props.DOMProps, "input-group-text", "px-3 py-2 text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

type InputGroupInputProps struct {
	DOMProps
	Type        string
	Name        string
	Value       string
	Placeholder string
	Disabled    bool
	Required    bool
	Invalid     bool
}

func InputGroupInput(props InputGroupInputProps) templ.Component {
	return Input(InputProps{
		DOMProps:    props.DOMProps,
		Type:        props.Type,
		Name:        props.Name,
		Value:       props.Value,
		Placeholder: props.Placeholder,
		Disabled:    props.Disabled,
		Required:    props.Required,
		Invalid:     props.Invalid,
	})
}

type InputGroupTextareaProps struct {
	DOMProps
	Name        string
	Value       string
	Placeholder string
	Rows        int
	Disabled    bool
	Required    bool
	Invalid     bool
}

func InputGroupTextarea(props InputGroupTextareaProps) templ.Component {
	return Textarea(TextareaProps{
		DOMProps:    props.DOMProps,
		Name:        props.Name,
		Value:       props.Value,
		Placeholder: props.Placeholder,
		Rows:        props.Rows,
		Disabled:    props.Disabled,
		Required:    props.Required,
		Invalid:     props.Invalid,
	})
}

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
		attrs := attrsFromDOMProps(props.DOMProps, "input-otp", "flex items-center gap-2")
		if props.Name != "" {
			attrs["data-name"] = props.Name
		}
		if props.Value != "" {
			attrs["data-value"] = props.Value
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
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
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
		return renderTextElement(ctx, w, "div", attrs, props.Value)
	})
}

func InputOTPSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "input-otp-separator", "mx-1 h-px w-3 bg-border"), nil)
	})
}

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
