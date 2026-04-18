package ui

import (
	"context"
	"io"
	"strconv"

	"github.com/a-h/templ"
)

type AvatarProps struct {
	DOMProps
	Size string
}

func Avatar(props AvatarProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		sizeClass := "size-10"
		switch props.Size {
		case "sm":
			sizeClass = "size-8"
		case "lg":
			sizeClass = "size-12"
		}
		return renderElement(ctx, w, "div", attrsFromDOMProps(props.DOMProps, "avatar", cn("relative flex shrink-0 overflow-hidden rounded-full", sizeClass)), templ.GetChildren(ctx))
	})
}

type AvatarImageProps struct {
	DOMProps
	Src string
	Alt string
}

func AvatarImage(props AvatarImageProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "avatar-image", "size-full object-cover")
		if props.Src != "" {
			attrs["src"] = props.Src
		}
		if props.Alt != "" {
			attrs["alt"] = props.Alt
		}
		return renderVoidElement(ctx, w, "img", attrs)
	})
}

func AvatarFallback(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "avatar-fallback", "flex size-full items-center justify-center rounded-full bg-muted text-xs"), templ.GetChildren(ctx))
	})
}

func AvatarBadge(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "avatar-badge", "absolute bottom-0 right-0 inline-flex size-2.5 rounded-full border border-background bg-primary"), templ.GetChildren(ctx))
	})
}

func AvatarGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "avatar-group", "flex -space-x-2"), templ.GetChildren(ctx))
	})
}

type AvatarGroupCountProps struct {
	DOMProps
	Count int
}

func AvatarGroupCount(props AvatarGroupCountProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		label := "+"
		if props.Count > 0 {
			label += strconv.Itoa(props.Count)
		}
		return renderTextElement(ctx, w, "span", attrsFromDOMProps(props.DOMProps, "avatar-group-count", "inline-flex size-8 items-center justify-center rounded-full bg-muted text-xs font-medium"), label)
	})
}

func Empty(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "empty", "grid place-items-center gap-4 rounded-xl border border-dashed p-8 text-center"), templ.GetChildren(ctx))
	})
}

func EmptyHeader(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "empty-header", "grid gap-1"), templ.GetChildren(ctx))
	})
}
func EmptyMedia(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "empty-media", "flex size-12 items-center justify-center rounded-full bg-muted"), templ.GetChildren(ctx))
	})
}
func EmptyTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "h3", attrsFromDOMProps(props, "empty-title", "text-base font-semibold"), templ.GetChildren(ctx))
	})
}
func EmptyDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "p", attrsFromDOMProps(props, "empty-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}
func EmptyContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "empty-content", ""), templ.GetChildren(ctx))
	})
}

type ItemVariant string

const (
	ItemVariantDefault ItemVariant = "default"
	ItemVariantOutline ItemVariant = "outline"
	ItemVariantMuted   ItemVariant = "muted"
)

type ItemSize string

const (
	ItemSizeDefault ItemSize = "default"
	ItemSizeSM      ItemSize = "sm"
)

type ItemProps struct {
	DOMProps
	Variant ItemVariant
	Size    ItemSize
}

func ItemGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "item-group", "grid gap-2"), templ.GetChildren(ctx))
	})
}

func ItemSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "hr", attrsFromDOMProps(props, "item-separator", "my-2 border-border"), nil)
	})
}

func Item(props ItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "grid gap-2 rounded-xl border bg-card p-4"
		switch props.Variant {
		case ItemVariantOutline:
			className = cn(className, "border-input")
		case ItemVariantMuted:
			className = cn(className, "bg-muted")
		}
		if props.Size == ItemSizeSM {
			className = cn(className, "p-3")
		}
		return renderElement(ctx, w, "div", attrsFromDOMProps(props.DOMProps, "item", className), templ.GetChildren(ctx))
	})
}

func ItemMedia(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "item-media", "flex size-10 items-center justify-center rounded-md bg-muted"), templ.GetChildren(ctx))
	})
}

func ItemContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "item-content", "grid gap-1"), templ.GetChildren(ctx))
	})
}

func ItemTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "h4", attrsFromDOMProps(props, "item-title", "font-medium leading-none"), templ.GetChildren(ctx))
	})
}

func ItemDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "p", attrsFromDOMProps(props, "item-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func ItemActions(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "item-actions", "ml-auto flex items-center gap-2"), templ.GetChildren(ctx))
	})
}

func ItemHeader(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "item-header", "flex items-start gap-3"), templ.GetChildren(ctx))
	})
}

func ItemFooter(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "item-footer", "flex items-center gap-2 pt-2"), templ.GetChildren(ctx))
	})
}

func Table(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "table", attrsFromDOMProps(props, "table", "w-full caption-bottom text-sm"), templ.GetChildren(ctx))
	})
}
func TableHeader(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "thead", attrsFromDOMProps(props, "table-header", "[&_tr]:border-b"), templ.GetChildren(ctx))
	})
}
func TableBody(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "tbody", attrsFromDOMProps(props, "table-body", "[&_tr:last-child]:border-0"), templ.GetChildren(ctx))
	})
}
func TableFooter(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "tfoot", attrsFromDOMProps(props, "table-footer", "border-t bg-muted/50 font-medium [&>tr]:last:border-b-0"), templ.GetChildren(ctx))
	})
}
func TableHead(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "th", attrsFromDOMProps(props, "table-head", "h-12 px-4 text-left align-middle font-medium text-muted-foreground"), templ.GetChildren(ctx))
	})
}
func TableRow(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "tr", attrsFromDOMProps(props, "table-row", "border-b transition-colors hover:bg-muted/50"), templ.GetChildren(ctx))
	})
}
func TableCell(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "td", attrsFromDOMProps(props, "table-cell", "p-4 align-middle"), templ.GetChildren(ctx))
	})
}
func TableCaption(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "caption", attrsFromDOMProps(props, "table-caption", "mt-4 text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func ScrollArea(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "scroll-area", "relative overflow-hidden"), templ.GetChildren(ctx))
	})
}

type ScrollBarProps struct {
	DOMProps
	Orientation string
}

func ScrollBar(props ScrollBarProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "flex touch-none select-none transition-colors"
		if props.Orientation == "horizontal" {
			className = cn(className, "h-2 w-full flex-col border-t border-t-transparent p-px")
		} else {
			className = cn(className, "h-full w-2 border-l border-l-transparent p-px")
		}
		return renderElement(ctx, w, "div", attrsFromDOMProps(props.DOMProps, "scrollbar", className), templ.GetChildren(ctx))
	})
}

type ResizablePanelGroupProps struct {
	DOMProps
	Direction string
}

func ResizablePanelGroup(props ResizablePanelGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "resizable-panel-group", "flex min-h-0 min-w-0")
		if props.Direction != "" {
			attrs["data-direction"] = props.Direction
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

type ResizablePanelProps struct {
	DOMProps
	DefaultSize float64
	MinSize     float64
	MaxSize     float64
	Collapsible bool
}

func ResizablePanel(props ResizablePanelProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "resizable-panel", "min-h-0 min-w-0")
		if props.DefaultSize > 0 {
			attrs["data-default-size"] = props.DefaultSize
		}
		if props.MinSize > 0 {
			attrs["data-min-size"] = props.MinSize
		}
		if props.MaxSize > 0 {
			attrs["data-max-size"] = props.MaxSize
		}
		if props.Collapsible {
			attrs["data-collapsible"] = "true"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

type ResizableHandleProps struct {
	DOMProps
	WithHandle bool
}

func ResizableHandle(props ResizableHandleProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "resizable-handle", "relative flex items-center justify-center bg-border")
		if props.WithHandle {
			attrs["data-with-handle"] = "true"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

type CarouselProps struct {
	DOMProps
	Orientation string
	Loop        bool
	Align       string
	StartIndex  int
}

func Carousel(props CarouselProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "carousel", "relative")
		if props.Orientation != "" {
			attrs["data-orientation"] = props.Orientation
		}
		if props.Loop {
			attrs["data-loop"] = "true"
		}
		if props.Align != "" {
			attrs["data-align"] = props.Align
		}
		if props.StartIndex > 0 {
			attrs["data-start-index"] = props.StartIndex
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func CarouselContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "carousel-content", "flex"), templ.GetChildren(ctx))
	})
}
func CarouselItem(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "carousel-item", "min-w-0 shrink-0 grow-0 basis-full"), templ.GetChildren(ctx))
	})
}

func CarouselPrevious(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "carousel-previous", "absolute left-2 top-1/2 -translate-y-1/2 rounded-full border bg-background p-2 shadow"), templ.GetChildren(ctx))
	})
}

func CarouselNext(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "carousel-next", "absolute right-2 top-1/2 -translate-y-1/2 rounded-full border bg-background p-2 shadow"), templ.GetChildren(ctx))
	})
}

type CalendarProps struct {
	DOMProps
	Mode            string
	Selected        string
	Range           string
	Month           string
	DefaultMonth    string
	ShowOutsideDays bool
	ShowWeekNumber  bool
	CaptionLayout   string
	ButtonVariant   string
	Locale          string
	DisabledDates   string
}

func Calendar(props CalendarProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "calendar", "grid gap-2 rounded-lg border p-3")
		if props.Mode != "" {
			attrs["data-mode"] = props.Mode
		}
		if props.Selected != "" {
			attrs["data-selected"] = props.Selected
		}
		if props.Range != "" {
			attrs["data-range"] = props.Range
		}
		if props.Month != "" {
			attrs["data-month"] = props.Month
		}
		if props.DefaultMonth != "" {
			attrs["data-default-month"] = props.DefaultMonth
		}
		if props.ShowOutsideDays {
			attrs["data-show-outside-days"] = "true"
		}
		if props.ShowWeekNumber {
			attrs["data-show-week-number"] = "true"
		}
		if props.CaptionLayout != "" {
			attrs["data-caption-layout"] = props.CaptionLayout
		}
		if props.ButtonVariant != "" {
			attrs["data-button-variant"] = props.ButtonVariant
		}
		if props.Locale != "" {
			attrs["data-locale"] = props.Locale
		}
		if props.DisabledDates != "" {
			attrs["data-disabled-dates"] = props.DisabledDates
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func CalendarDayButton(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "calendar-day-button", "inline-flex size-8 items-center justify-center rounded-md text-sm hover:bg-accent hover:text-accent-foreground"), templ.GetChildren(ctx))
	})
}

type DatePickerProps struct {
	DOMProps
	Mode        string
	Name        string
	Value       string
	Range       string
	Placeholder string
	Open        bool
	DefaultOpen bool
}

func DatePicker(props DatePickerProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "date-picker", "grid gap-2")
		if props.Mode != "" {
			attrs["data-mode"] = props.Mode
		}
		if props.Name != "" {
			attrs["data-name"] = props.Name
		}
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.Range != "" {
			attrs["data-range"] = props.Range
		}
		if props.Placeholder != "" {
			attrs["data-placeholder"] = props.Placeholder
		}
		if props.Open {
			attrs["data-open"] = "true"
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

type DateRangePickerProps struct {
	DatePickerProps
}

func DateRangePicker(props DateRangePickerProps) templ.Component {
	return DatePicker(props.DatePickerProps)
}

type SelectProps struct {
	DOMProps
	Name         string
	Value        string
	DefaultValue string
	Open         bool
	DefaultOpen  bool
	Disabled     bool
	Required     bool
}

type SelectTriggerProps struct {
	DOMProps
	Size string
}

type SelectContentProps struct {
	DOMProps
	Position string
	Align    string
}

func Select(props SelectProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "select", "relative")
		if props.Name != "" {
			attrs["data-name"] = props.Name
		}
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.DefaultValue != "" {
			attrs["data-default-value"] = props.DefaultValue
		}
		if props.Open {
			attrs["data-open"] = "true"
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		if props.Disabled {
			attrs["data-disabled"] = "true"
		}
		if props.Required {
			attrs["data-required"] = "true"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func SelectGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "select-group", "grid gap-1"), templ.GetChildren(ctx))
	})
}
func SelectValue(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "select-value", ""), templ.GetChildren(ctx))
	})
}
func SelectTrigger(props SelectTriggerProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "flex h-9 w-full items-center justify-between rounded-md border border-input bg-background px-3 py-2 text-sm shadow-xs outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:cursor-not-allowed disabled:opacity-50"
		if props.Size == "sm" {
			className = cn(className, "h-8")
		}
		return renderElement(ctx, w, "button", attrsFromDOMProps(props.DOMProps, "select-trigger", className), templ.GetChildren(ctx))
	})
}

func SelectContent(props SelectContentProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "select-content", "z-50 min-w-32 rounded-md border bg-popover p-1 text-popover-foreground shadow-md")
		if props.Position != "" {
			attrs["data-position"] = props.Position
		}
		if props.Align != "" {
			attrs["data-align"] = props.Align
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
func SelectLabel(props DOMProps) templ.Component             { return DropdownMenuLabel(props) }
func SelectItem(props DropdownMenuItemProps) templ.Component { return DropdownMenuItem(props) }
func SelectSeparator(props DOMProps) templ.Component         { return DropdownMenuSeparator(props) }
func SelectScrollUpButton(props DOMProps) templ.Component    { return DialogTrigger(props) }
func SelectScrollDownButton(props DOMProps) templ.Component  { return DialogTrigger(props) }

type ComboboxProps struct {
	DOMProps
	Name         string
	Value        string
	Values       []string
	DefaultValue string
	Multiple     bool
	Placeholder  string
	ShowTrigger  bool
	ShowClear    bool
	Filter       string
	Open         bool
	DefaultOpen  bool
	AnchorID     string
	Side         string
	Align        string
	SideOffset   string
}

func Combobox(props ComboboxProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "combobox", "relative")
		if props.Name != "" {
			attrs["data-name"] = props.Name
		}
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if len(props.Values) > 0 {
			attrs["data-values"] = props.Values
		}
		if props.DefaultValue != "" {
			attrs["data-default-value"] = props.DefaultValue
		}
		if props.Multiple {
			attrs["data-multiple"] = "true"
		}
		if props.Placeholder != "" {
			attrs["data-placeholder"] = props.Placeholder
		}
		if props.ShowTrigger {
			attrs["data-show-trigger"] = "true"
		}
		if props.ShowClear {
			attrs["data-show-clear"] = "true"
		}
		if props.Filter != "" {
			attrs["data-filter"] = props.Filter
		}
		if props.Open {
			attrs["data-open"] = "true"
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		if props.AnchorID != "" {
			attrs["data-anchor-id"] = props.AnchorID
		}
		if props.Side != "" {
			attrs["data-side"] = props.Side
		}
		if props.Align != "" {
			attrs["data-align"] = props.Align
		}
		if props.SideOffset != "" {
			attrs["data-side-offset"] = props.SideOffset
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func ComboboxValue(props DOMProps) templ.Component             { return SelectValue(props) }
func ComboboxTrigger(props SelectTriggerProps) templ.Component { return SelectTrigger(props) }
func ComboboxClear(props DOMProps) templ.Component             { return DialogTrigger(props) }
func ComboboxInput(props InputProps) templ.Component           { return Input(props) }
func ComboboxContent(props SelectContentProps) templ.Component { return SelectContent(props) }
func ComboboxList(props DOMProps) templ.Component              { return SelectGroup(props) }
func ComboboxItem(props DropdownMenuItemProps) templ.Component { return DropdownMenuItem(props) }
func ComboboxGroup(props DOMProps) templ.Component             { return SelectGroup(props) }
func ComboboxLabel(props DOMProps) templ.Component             { return SelectLabel(props) }
func ComboboxCollection(props DOMProps) templ.Component        { return SelectGroup(props) }
func ComboboxEmpty(props DOMProps) templ.Component             { return Empty(props) }
func ComboboxSeparator(props DOMProps) templ.Component         { return SelectSeparator(props) }
func ComboboxChips(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "combobox-chips", "flex flex-wrap gap-2"), templ.GetChildren(ctx))
	})
}
func ComboboxChip(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "combobox-chip", "inline-flex items-center rounded-md bg-muted px-2 py-1 text-xs"), templ.GetChildren(ctx))
	})
}
func ComboboxChipsInput(props InputProps) templ.Component { return Input(props) }

type CommandProps struct {
	DOMProps
	Value           string
	DefaultValue    string
	Open            bool
	Placeholder     string
	Title           string
	Description     string
	ShowCloseButton bool
}

func Command(props CommandProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "command", "grid gap-2 rounded-lg border bg-popover p-2 text-popover-foreground shadow-md")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.DefaultValue != "" {
			attrs["data-default-value"] = props.DefaultValue
		}
		if props.Open {
			attrs["data-open"] = "true"
		}
		if props.Placeholder != "" {
			attrs["data-placeholder"] = props.Placeholder
		}
		if props.Title != "" {
			attrs["data-title"] = props.Title
		}
		if props.Description != "" {
			attrs["data-description"] = props.Description
		}
		if props.ShowCloseButton {
			attrs["data-show-close-button"] = "true"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func CommandDialog(props CommandProps) templ.Component        { return Command(props) }
func CommandInput(props InputProps) templ.Component           { return Input(props) }
func CommandList(props DOMProps) templ.Component              { return SelectGroup(props) }
func CommandEmpty(props DOMProps) templ.Component             { return Empty(props) }
func CommandGroup(props DOMProps) templ.Component             { return SelectGroup(props) }
func CommandItem(props DropdownMenuItemProps) templ.Component { return DropdownMenuItem(props) }
func CommandShortcut(props DOMProps) templ.Component          { return DropdownMenuShortcut(props) }
func CommandSeparator(props DOMProps) templ.Component         { return DropdownMenuSeparator(props) }

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
