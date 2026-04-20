package views

import (
	"fmt"
	"strings"
)

type APIProp struct {
	Name    string
	Type    string
	Default string
	Desc    string
}

type ExampleEntry struct {
	Name string
	Desc string
}

type ComponentDocEntry struct {
	Slug          string
	Title         string
	Description   string
	Install       string
	Usage         string
	API           []string       // KEEP existing — deprecated but keep for compat
	Example       []string       // KEEP existing — deprecated but keep for compat
	GoUsage       string         // NEW: Go templ usage snippet
	ManualInstall string         // NEW: manual install steps
	APIProps      []APIProp      // NEW: structured API props
	Examples      []ExampleEntry // NEW: named examples with descriptions
	SourceCode    string
}

type BlockEntry struct {
	Slug        string
	Title       string
	Description string
	Category    string
	Command     string
	Files       []string
}

var componentDocs = []ComponentDocEntry{
	{Slug: "accordion", Title: "Accordion", Description: "A vertically stacked set of collapsible panels.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Accordion(...)", API: []string{"type", "value", "defaultValue", "collapsible"}, Example: []string{"Accordion", "AccordionItem", "AccordionTrigger", "AccordionContent"}},
	{Slug: "alert", Title: "Alert", Description: "Displays a prominent message to the user.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Alert(...)", API: []string{"variant"}, Example: []string{"Alert", "AlertTitle", "AlertDescription"}, SourceCode: `type AlertVariant string

const (
	AlertVariantDefault     AlertVariant = "default"
	AlertVariantDestructive AlertVariant = "destructive"
)

type AlertProps struct {
	DOMProps
	Variant AlertVariant
}

func alertClasses(variant AlertVariant, className string) string {
	if variant == "" {
		variant = AlertVariantDefault
	}

	base := "relative w-full rounded-xl border px-4 py-3 text-sm grid gap-2"
	switch variant {
	case AlertVariantDestructive:
		return cn(base, "border-destructive/50 text-destructive dark:border-destructive [&_svg]:text-destructive", className)
	default:
		return cn(base, "bg-background text-foreground", className)
	}
}

func Alert(props AlertProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "alert", alertClasses(props.Variant, ""))
		attrs["role"] = "alert"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func AlertTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "h5", attrsFromDOMProps(props, "alert-title", "mb-1 font-medium leading-none tracking-tight"), templ.GetChildren(ctx))
	})
}

func AlertDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "alert-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}`},
	{Slug: "alert-dialog", Title: "Alert Dialog", Description: "A modal dialog that asks the user to confirm a destructive action.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.AlertDialog(...)", API: []string{"open", "defaultOpen", "modal"}, Example: []string{"AlertDialog", "AlertDialogAction", "AlertDialogCancel"}},
	{Slug: "aspect-ratio", Title: "Aspect Ratio", Description: "Maintains a consistent width-to-height ratio.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.AspectRatio(...)", API: []string{"ratio"}, Example: []string{"AspectRatio"}},
	{Slug: "avatar", Title: "Avatar", Description: "Displays a user image with a graceful fallback.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Avatar(...)", API: []string{"size"}, Example: []string{"Avatar", "AvatarImage", "AvatarFallback"}},
	{Slug: "badge", Title: "Badge", Description: "Small status or metadata labels.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Badge(...)", API: []string{"variant"}, Example: []string{"Badge"}, SourceCode: `package ui

import (
	"bytes"
	"context"
	"io"

	"github.com/a-h/templ"
)

var badgeVariantClasses = map[BadgeVariant]string{
	BadgeVariantDefault:     "bg-primary text-primary-foreground [a&]:hover:bg-primary/90",
	BadgeVariantSecondary:   "bg-secondary text-secondary-foreground [a&]:hover:bg-secondary/90",
	BadgeVariantDestructive: "bg-destructive text-white focus-visible:ring-destructive/20 dark:bg-destructive/60 dark:focus-visible:ring-destructive/40 [a&]:hover:bg-destructive/90",
	BadgeVariantOutline:     "border-border text-foreground [a&]:hover:bg-accent [a&]:hover:text-accent-foreground",
	BadgeVariantGhost:       "[a&]:hover:bg-accent [a&]:hover:text-accent-foreground",
	BadgeVariantLink:        "text-primary underline-offset-4 [a&]:hover:underline",
}

func badgeClasses(variant BadgeVariant, className string) string {
	if variant == "" {
		variant = BadgeVariantDefault
	}

	base := "inline-flex w-fit shrink-0 items-center justify-center gap-1 overflow-hidden rounded-full border border-transparent px-2 py-0.5 text-xs font-medium whitespace-nowrap transition-[color,box-shadow] focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 aria-invalid:border-destructive aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 [&>svg]:pointer-events-none [&>svg]:size-3"
	return cn(base, badgeVariantClasses[variant], className)
}

func Badge(props BadgeProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		buf, isBuffer := w.(*bytes.Buffer)
		if !isBuffer {
			buf = templ.GetBuffer()
			defer templ.ReleaseBuffer(buf)
		}

		ctx = templ.InitializeContext(ctx)

		attrs := cloneAttributes(props.Attrs)
		if props.ID != "" {
			attrs["id"] = props.ID
		}

		className := badgeClasses(props.Variant, props.Class)
		if existing, ok := attrs["class"]; ok {
			className = cn(className, templ.Classes(existing).String())
		}
		attrs["class"] = className

		tag := "span"
		if props.Element != "" {
			tag = props.Element
		} else if props.Href != "" {
			tag = "a"
		}

		if tag == "a" {
			attrs["href"] = props.Href
		}

		if _, err := buf.WriteString("<" + tag); err != nil {
			return err
		}
		if err := templ.RenderAttributes(ctx, buf, attrs); err != nil {
			return err
		}
		if _, err := buf.WriteString(">"); err != nil {
			return err
		}
		if props.Label != "" {
			if _, err := buf.WriteString(templ.EscapeString(props.Label)); err != nil {
				return err
			}
		}
		if _, err := buf.WriteString("</" + tag + ">"); err != nil {
			return err
		}

		if !isBuffer {
			_, err := buf.WriteTo(w)
			return err
		}
		return nil
	})
}`},
	{Slug: "breadcrumb", Title: "Breadcrumb", Description: "A navigation trail for the current location.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Breadcrumb(...)", API: []string{"href"}, Example: []string{"Breadcrumb", "BreadcrumbLink", "BreadcrumbPage"}},
	{Slug: "button", Title: "Button", Description: "Displays a button or button-like component.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Button(...)", API: []string{"variant", "size", "href", "disabled"}, Example: []string{"Button", "ButtonGroup"}, SourceCode: `package ui

import (
	"bytes"
	"context"
	"io"
	"sort"

	"github.com/a-h/templ"
)

var buttonVariantClasses = map[ButtonVariant]string{
	ButtonVariantDefault:     "bg-primary text-primary-foreground hover:bg-primary/90",
	ButtonVariantDestructive: "bg-destructive text-white hover:bg-destructive/90 focus-visible:ring-destructive/20 dark:bg-destructive/60 dark:focus-visible:ring-destructive/40",
	ButtonVariantOutline:     "border bg-background shadow-xs hover:bg-accent hover:text-accent-foreground dark:border-input dark:bg-input/30 dark:hover:bg-input/50",
	ButtonVariantSecondary:   "bg-secondary text-secondary-foreground hover:bg-secondary/80",
	ButtonVariantGhost:       "hover:bg-accent hover:text-accent-foreground dark:hover:bg-accent/50",
	ButtonVariantLink:        "text-primary underline-offset-4 hover:underline",
}

var buttonSizeClasses = map[ButtonSize]string{
	ButtonSizeDefault: "h-9 px-4 py-2 has-[>svg]:px-3",
	ButtonSizeXS:      "h-6 gap-1 rounded-md px-2 text-xs has-[>svg]:px-1.5 [&_svg:not([class*='size-'])]:size-3",
	ButtonSizeSM:      "h-8 gap-1.5 rounded-md px-3 has-[>svg]:px-2.5",
	ButtonSizeLG:      "h-10 rounded-md px-6 has-[>svg]:px-4",
	ButtonSizeIcon:    "size-9",
	ButtonSizeIconXS:  "size-6 rounded-md [&_svg:not([class*='size-'])]:size-3",
	ButtonSizeIconSM:  "size-8",
	ButtonSizeIconLG:  "size-10",
}

func buttonClasses(variant ButtonVariant, size ButtonSize, className string) string {
	if variant == "" {
		variant = ButtonVariantDefault
	}
	if size == "" {
		size = ButtonSizeDefault
	}

	base := "inline-flex shrink-0 items-center justify-center gap-2 rounded-md text-sm font-medium whitespace-nowrap transition-all outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:pointer-events-none disabled:opacity-50 aria-invalid:border-destructive aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 [&_svg]:pointer-events-none [&_svg]:shrink-0 [&_svg:not([class*='size-'])]:size-4"
	return cn(base, buttonVariantClasses[variant], buttonSizeClasses[size], className)
}

func cloneAttributes(attrs templ.Attributes) templ.Attributes {
	if len(attrs) == 0 {
		return templ.Attributes{}
	}

	out := make(templ.Attributes, len(attrs))
	for k, v := range attrs {
		out[k] = v
	}
	return out
}

func Button(props ButtonProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		buf, isBuffer := w.(*bytes.Buffer)
		if !isBuffer {
			buf = templ.GetBuffer()
			defer templ.ReleaseBuffer(buf)
		}

		ctx = templ.InitializeContext(ctx)

		attrs := cloneAttributes(props.Attrs)
		if props.ID != "" {
			attrs["id"] = props.ID
		}

		className := buttonClasses(props.Variant, props.Size, props.Class)
		if existing, ok := attrs["class"]; ok {
			className = cn(className, templ.Classes(existing).String())
		}
		attrs["class"] = className

		tag := "button"
		if props.Element != "" {
			tag = props.Element
		} else if props.Href != "" {
			tag = "a"
		}

		if tag == "a" {
			attrs["href"] = props.Href
		} else {
			if props.Type == "" {
				attrs["type"] = "button"
			} else {
				attrs["type"] = props.Type
			}
			if props.Disabled {
				attrs["disabled"] = true
			}
		}

		if _, err := buf.WriteString("<" + tag); err != nil {
			return err
		}
		if err := templ.RenderAttributes(ctx, buf, attrs); err != nil {
			return err
		}
		if _, err := buf.WriteString(">"); err != nil {
			return err
		}
		if props.Label != "" {
			if _, err := buf.WriteString(templ.EscapeString(props.Label)); err != nil {
				return err
			}
		}
		if _, err := buf.WriteString("</" + tag + ">"); err != nil {
			return err
		}

		if !isBuffer {
			_, err := buf.WriteTo(w)
			return err
		}
		return nil
	})
}

func ButtonVariantNames() []string {
	names := make([]string, 0, len(buttonVariantClasses))
	for name := range buttonVariantClasses {
		names = append(names, string(name))
	}
	sort.Strings(names)
	return names
}`},
	{Slug: "button-group", Title: "Button Group", Description: "Groups buttons into a single control surface.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.ButtonGroup(...)", API: []string{"orientation"}, Example: []string{"ButtonGroup", "ButtonGroupText", "ButtonGroupSeparator"}},
	{Slug: "calendar", Title: "Calendar", Description: "A month view calendar with flexible selection modes.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Calendar(...)", API: []string{"mode", "selected", "month", "locale"}, Example: []string{"Calendar", "CalendarDayButton"}},
	{Slug: "card", Title: "Card", Description: "A flexible container for related information.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Card(...)", API: []string{"none"}, Example: []string{"Card", "CardHeader", "CardContent"}, SourceCode: `func cardWrapper(props DOMProps, slot, className string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, slot, className), templ.GetChildren(ctx))
	})
}

func Card(props DOMProps) templ.Component {
	return cardWrapper(props, "card", "rounded-xl border bg-card text-card-foreground shadow")
}
func CardHeader(props DOMProps) templ.Component {
	return cardWrapper(props, "card-header", "flex flex-col gap-1.5 p-6")
}
func CardTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "h3", attrsFromDOMProps(props, "card-title", "font-semibold leading-none tracking-tight"), templ.GetChildren(ctx))
	})
}
func CardDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "p", attrsFromDOMProps(props, "card-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}
func CardAction(props DOMProps) templ.Component {
	return cardWrapper(props, "card-action", "ml-auto flex items-center")
}
func CardContent(props DOMProps) templ.Component {
	return cardWrapper(props, "card-content", "p-6 pt-0")
}
func CardFooter(props DOMProps) templ.Component {
	return cardWrapper(props, "card-footer", "flex items-center p-6 pt-0")
}`},
	{Slug: "carousel", Title: "Carousel", Description: "A swipeable set of slides with navigation controls.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Carousel(...)", API: []string{"orientation", "loop", "align"}, Example: []string{"Carousel", "CarouselContent", "CarouselItem"}},
	{Slug: "chart", Title: "Chart", Description: "A shell for chart configuration, tooltips, and legend styling.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.ChartContainer(...)", API: []string{"config", "series"}, Example: []string{"ChartContainer", "ChartTooltip", "ChartLegend"}},
	{Slug: "checkbox", Title: "Checkbox", Description: "A native checkbox with shadcn styling.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Checkbox(...)", API: []string{"checked", "defaultChecked", "required"}, Example: []string{"Checkbox"}, SourceCode: `type CheckboxProps struct {
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
}`},
	{Slug: "collapsible", Title: "Collapsible", Description: "A lighter disclosure primitive than accordion.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Collapsible(...)", API: []string{"open", "defaultOpen"}, Example: []string{"Collapsible", "CollapsibleTrigger", "CollapsibleContent"}},
	{Slug: "combobox", Title: "Combobox", Description: "Searchable selection with single or multi-value support.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Combobox(...)", API: []string{"multiple", "value", "values", "filter"}, Example: []string{"Combobox", "ComboboxItem", "ComboboxChips"}},
	{Slug: "command", Title: "Command", Description: "A searchable command palette component.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Command(...)", API: []string{"value", "defaultValue", "open"}, Example: []string{"Command", "CommandInput", "CommandItem"}},
	{Slug: "context-menu", Title: "Context Menu", Description: "A right-click menu with nested submenus.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.ContextMenu(...)", API: []string{"variant", "inset"}, Example: []string{"ContextMenu", "ContextMenuItem", "ContextMenuSub"}},
	{Slug: "data-table", Title: "Data Table", Description: "A server-rendered table pattern for app data.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.DataTable(...)", API: []string{"columns", "rows", "page", "filters"}, Example: []string{"DataTable", "DataTableToolbar", "DataTablePagination"}},
	{Slug: "date-picker", Title: "Date Picker", Description: "A calendar popover for single or range dates.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.DatePicker(...)", API: []string{"mode", "value", "range"}, Example: []string{"DatePicker", "DateRangePicker"}},
	{Slug: "dialog", Title: "Dialog", Description: "A modal dialog used for forms and transient tasks.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Dialog(...)", API: []string{"open", "defaultOpen", "modal"}, Example: []string{"Dialog", "DialogTitle", "DialogDescription"}},
	{Slug: "direction", Title: "Direction", Description: "A small RTL/LTR direction provider.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.DirectionProvider(...)", API: []string{"direction"}, Example: []string{"DirectionProvider"}},
	{Slug: "drawer", Title: "Drawer", Description: "A slide-over surface with mobile-friendly motion.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Drawer(...)", API: []string{"side", "dismissible", "snapPoints"}, Example: []string{"Drawer", "DrawerContent"}},
	{Slug: "dropdown-menu", Title: "Dropdown Menu", Description: "A context-driven popover menu.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.DropdownMenu(...)", API: []string{"sideOffset", "variant"}, Example: []string{"DropdownMenu", "DropdownMenuContent", "DropdownMenuItem"}},
	{Slug: "empty", Title: "Empty", Description: "A polished empty state helper.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Empty(...)", API: []string{"variant"}, Example: []string{"Empty", "EmptyTitle", "EmptyDescription"}},
	{Slug: "field", Title: "Field", Description: "A form field grouping helper.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Field(...)", API: []string{"orientation", "invalid", "disabled"}, Example: []string{"FieldSet", "FieldLabel", "FieldError"}},
	{Slug: "hover-card", Title: "Hover Card", Description: "A hover/focus-driven preview popover.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.HoverCard(...)", API: []string{"openDelayMs", "closeDelayMs"}, Example: []string{"HoverCard", "HoverCardContent"}},
	{Slug: "input", Title: "Input", Description: "A native text input styled for shadcn themes.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Input(...)", API: []string{"type", "name", "value"}, Example: []string{"Input"}, SourceCode: `type InputProps struct {
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
}`},
	{Slug: "input-group", Title: "Input Group", Description: "Input with addons, buttons, and labels.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.InputGroup(...)", API: []string{"align"}, Example: []string{"InputGroup", "InputGroupAddon", "InputGroupInput"}},
	{Slug: "input-otp", Title: "Input OTP", Description: "One-time code entry with slot-based UI.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.InputOTP(...)", API: []string{"maxlength", "pattern"}, Example: []string{"InputOTP", "InputOTPSlot"}},
	{Slug: "item", Title: "Item", Description: "A generic list/card item surface.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Item(...)", API: []string{"variant", "size"}, Example: []string{"Item", "ItemHeader", "ItemActions"}},
	{Slug: "kbd", Title: "Kbd", Description: "A keyboard shortcut token.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Kbd(...)", API: []string{"size", "inset"}, Example: []string{"Kbd"}},
	{Slug: "label", Title: "Label", Description: "A native label wrapper.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Label(...)", API: []string{"for"}, Example: []string{"Label"}},
	{Slug: "menubar", Title: "Menubar", Description: "A horizontal application menu.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Menubar(...)", API: []string{"orientation"}, Example: []string{"Menubar", "MenubarItem", "MenubarSub"}},
	{Slug: "native-select", Title: "Native Select", Description: "A styled native select element.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.NativeSelect(...)", API: []string{"size", "required"}, Example: []string{"NativeSelect", "NativeSelectOption"}},
	{Slug: "navigation-menu", Title: "Navigation Menu", Description: "A top-level navigation pattern with a viewport.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.NavigationMenu(...)", API: []string{"value", "orientation"}, Example: []string{"NavigationMenu", "NavigationMenuTrigger"}},
	{Slug: "pagination", Title: "Pagination", Description: "Pagination primitives for list navigation.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Pagination(...)", API: []string{"href", "isActive"}, Example: []string{"Pagination", "PaginationLink"}},
	{Slug: "popover", Title: "Popover", Description: "Floating content anchored to a trigger.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Popover(...)", API: []string{"side", "align", "modal"}, Example: []string{"Popover", "PopoverContent"}},
	{Slug: "progress", Title: "Progress", Description: "A native progress bar.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Progress(...)", API: []string{"value", "max"}, Example: []string{"Progress"}},
	{Slug: "radio-group", Title: "Radio Group", Description: "A group of radio options.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.RadioGroup(...)", API: []string{"name", "value", "required"}, Example: []string{"RadioGroup", "RadioGroupItem"}},
	{Slug: "resizable", Title: "Resizable", Description: "Pointer-driven resize panels.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.ResizablePanelGroup(...)", API: []string{"direction", "defaultSize"}, Example: []string{"ResizablePanelGroup", "ResizableHandle"}},
	{Slug: "scroll-area", Title: "Scroll Area", Description: "Styled scrollable containers and bars.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.ScrollArea(...)", API: []string{"orientation"}, Example: []string{"ScrollArea", "ScrollBar"}},
	{Slug: "select", Title: "Select", Description: "A custom select built around a listbox runtime.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Select(...)", API: []string{"open", "defaultOpen", "required"}, Example: []string{"Select", "SelectTrigger", "SelectContent"}},
	{Slug: "separator", Title: "Separator", Description: "A horizontal or vertical divider.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Separator(...)", API: []string{"orientation"}, Example: []string{"Separator"}},
	{Slug: "sheet", Title: "Sheet", Description: "A slide-over panel.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Sheet(...)", API: []string{"side", "showCloseButton"}, Example: []string{"Sheet", "SheetContent"}},
	{Slug: "sidebar", Title: "Sidebar", Description: "A full app shell sidebar with navigation controls.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Sidebar(...)", API: []string{"variant", "collapsible", "side"}, Example: []string{"SidebarProvider", "Sidebar", "SidebarMenuButton"}},
	{Slug: "skeleton", Title: "Skeleton", Description: "A placeholder shimmer block.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Skeleton(...)", API: []string{"none"}, Example: []string{"Skeleton"}},
	{Slug: "slider", Title: "Slider", Description: "A range slider for numeric input.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Slider(...)", API: []string{"min", "max", "step"}, Example: []string{"Slider"}},
	{Slug: "sonner", Title: "Sonner", Description: "A toast experience with shadcn defaults.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Toaster(...)", API: []string{"theme", "position"}, Example: []string{"Toaster"}},
	{Slug: "spinner", Title: "Spinner", Description: "A lightweight loading spinner.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Spinner(...)", API: []string{"size", "label"}, Example: []string{"Spinner"}},
	{Slug: "switch", Title: "Switch", Description: "A checkbox-like binary control.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Switch(...)", API: []string{"size", "checked", "defaultChecked"}, Example: []string{"Switch"}},
	{Slug: "table", Title: "Table", Description: "Table primitives with shadcn table styling.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Table(...)", API: []string{"none"}, Example: []string{"Table", "TableRow", "TableCell"}},
	{Slug: "tabs", Title: "Tabs", Description: "Tab buttons and associated panels.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Tabs(...)", API: []string{"value", "defaultValue", "orientation"}, Example: []string{"Tabs", "TabsList", "TabsTrigger"}},
	{Slug: "textarea", Title: "Textarea", Description: "A native multi-line input.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Textarea(...)", API: []string{"rows", "placeholder"}, Example: []string{"Textarea"}},
	{Slug: "toast", Title: "Toast", Description: "Legacy toast primitives.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Toast(...)", API: []string{"variant", "duration"}, Example: []string{"Toast", "ToastProvider"}},
	{Slug: "toggle", Title: "Toggle", Description: "A pressed-state button.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Toggle(...)", API: []string{"variant", "size"}, Example: []string{"Toggle"}},
	{Slug: "toggle-group", Title: "Toggle Group", Description: "A group of toggle buttons.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.ToggleGroup(...)", API: []string{"type", "variant", "size"}, Example: []string{"ToggleGroup", "ToggleGroupItem"}},
	{Slug: "tooltip", Title: "Tooltip", Description: "A floating hint for compact controls.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.Tooltip(...)", API: []string{"delayDuration", "side", "align"}, Example: []string{"Tooltip", "TooltipContent"}},
	{Slug: "typography", Title: "Typography", Description: "A docs-oriented prose helper layer.", Install: "go get github.com/shadcn-ui/ui", Usage: "@ui.H1(...)", API: []string{"none"}, Example: []string{"H1", "P", "InlineCode"}},
}

var blockDocs = []BlockEntry{
	{Slug: "dashboard-01", Title: "A dashboard with sidebar, charts and data table", Description: "A dense app shell that combines a sidebar, KPI cards, charts, and a data table.", Category: "featured", Command: "npx shadcn add dashboard-01", Files: []string{"app/dashboard/page.tsx", "components/app-sidebar.tsx", "components/chart-area-interactive.tsx", "components/data-table.tsx", "components/section-cards.tsx", "components/site-header.tsx"}},
	{Slug: "sidebar-07", Title: "A sidebar that collapses to icons", Description: "A compact application shell with icon-only collapse behavior.", Category: "sidebar", Command: "npx shadcn add sidebar-07", Files: []string{"app/dashboard/page.tsx", "components/app-sidebar.tsx", "components/team-switcher.tsx"}},
	{Slug: "sidebar-03", Title: "A sidebar with submenus", Description: "Sidebar navigation with nested submenu states and a content region.", Category: "sidebar", Command: "npx shadcn add sidebar-03", Files: []string{"app/dashboard/page.tsx", "components/app-sidebar.tsx"}},
	{Slug: "login-01", Title: "A simple login form", Description: "A centered login form with a compact footprint.", Category: "login", Command: "npx shadcn add login-01", Files: []string{"app/login/page.tsx", "components/login-form.tsx"}},
	{Slug: "login-03", Title: "A login page with a muted background color", Description: "A basic authentication surface with soft contrast and a small brand lockup.", Category: "login", Command: "npx shadcn add login-03", Files: []string{"app/login/page.tsx", "components/login-form.tsx"}},
	{Slug: "login-04", Title: "A login page with form and image", Description: "A wider authentication layout that pairs form and image panels.", Category: "login", Command: "npx shadcn add login-04", Files: []string{"app/login/page.tsx", "components/login-form.tsx"}},
	{Slug: "signup-01", Title: "A signup page with a cover image", Description: "A registration layout with a strong media panel.", Category: "signup", Command: "npx shadcn add signup-01", Files: []string{"app/signup/page.tsx", "components/signup-form.tsx"}},
	{Slug: "signup-02", Title: "A signup form with social providers", Description: "A signup flow that emphasizes third-party providers and email signup.", Category: "signup", Command: "npx shadcn add signup-02", Files: []string{"app/signup/page.tsx", "components/signup-form.tsx"}},
}

func FindComponentDoc(slug string) (ComponentDocEntry, bool) {
	for _, doc := range componentDocs {
		if doc.Slug == slug {
			return doc, true
		}
	}
	return ComponentDocEntry{}, false
}

func ComponentDocBefore(slug string) (ComponentDocEntry, bool) {
	for i, doc := range componentDocs {
		if doc.Slug == slug && i > 0 {
			return componentDocs[i-1], true
		}
	}
	return ComponentDocEntry{}, false
}

func ComponentDocAfter(slug string) (ComponentDocEntry, bool) {
	for i, doc := range componentDocs {
		if doc.Slug == slug && i < len(componentDocs)-1 {
			return componentDocs[i+1], true
		}
	}
	return ComponentDocEntry{}, false
}

func ComponentIndex() []ComponentDocEntry {
	return componentDocs
}

func Blocks() []BlockEntry {
	return blockDocs
}

func BlocksByCategory(category string) []BlockEntry {
	out := make([]BlockEntry, 0, len(blockDocs))
	for _, block := range blockDocs {
		if block.Category == category {
			out = append(out, block)
		}
	}
	return out
}

func FindBlockDoc(slug string) (BlockEntry, bool) {
	for _, block := range blockDocs {
		if block.Slug == slug {
			return block, true
		}
	}
	return BlockEntry{}, false
}

func BlockCategories() []string {
	seen := map[string]struct{}{}
	categories := make([]string, 0, len(blockDocs))
	for _, block := range blockDocs {
		if _, ok := seen[block.Category]; ok {
			continue
		}
		seen[block.Category] = struct{}{}
		categories = append(categories, block.Category)
	}
	return categories
}

func TitleFromSlug(slug string) string {
	parts := strings.Split(slug, "-")
	for i, part := range parts {
		switch part {
		case "otp":
			parts[i] = "OTP"
		case "ui":
			parts[i] = "UI"
		default:
			if len(part) == 0 {
				continue
			}
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}
	return strings.Join(parts, " ")
}

func ComponentLinkClass(active, slug string) string {
	base := "rounded-md px-3 py-2 transition-colors hover:bg-accent hover:text-accent-foreground"
	if active == "component-"+slug {
		return base + " bg-accent text-accent-foreground"
	}
	return base
}

func SectionLinkClass(active, section string) string {
	base := "rounded-md px-3 py-2 transition-colors hover:bg-accent hover:text-accent-foreground"
	if active == section {
		return base + " bg-accent text-accent-foreground"
	}
	return base
}

func BlockCategoryCountLabel(category string) string {
	return fmt.Sprintf("%d layouts available", len(BlocksByCategory(category)))
}

func ChartNavLinkClass(active, chart string) string {
	base := "px-4 py-2 text-sm font-medium transition-colors hover:text-foreground whitespace-nowrap"
	if active == "charts-"+chart {
		return base + " text-foreground border-b-2 border-primary"
	}
	return base + " text-muted-foreground"
}
