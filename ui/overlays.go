package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

type AccordionProps struct {
	DOMProps
	Type         string
	Value        []string
	DefaultValue []string
	Collapsible  bool
}

type AccordionItemProps struct {
	DOMProps
	Value    string
	Disabled bool
}

func Accordion(props AccordionProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "accordion", "grid gap-2")
		if props.Type != "" {
			attrs["data-type"] = props.Type
		}
		if len(props.Value) > 0 {
			attrs["data-value"] = props.Value
		}
		if len(props.DefaultValue) > 0 {
			attrs["data-default-value"] = props.DefaultValue
		}
		if props.Collapsible {
			attrs["data-collapsible"] = "true"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func AccordionItem(props AccordionItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "accordion-item", "rounded-md border border-border")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.Disabled {
			attrs["data-disabled"] = "true"
		}
		return renderElement(ctx, w, "details", attrs, templ.GetChildren(ctx))
	})
}

func AccordionTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "summary", attrsFromDOMProps(props, "accordion-trigger", "flex cursor-pointer items-center justify-between gap-4 px-4 py-3 text-sm font-medium outline-none"), templ.GetChildren(ctx))
	})
}

func AccordionContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "accordion-content", "px-4 pb-4 pt-0 text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

type CollapsibleProps struct {
	DOMProps
	Open        bool
	DefaultOpen bool
	Disabled    bool
}

func Collapsible(props CollapsibleProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "collapsible", "grid gap-2")
		if props.Open {
			attrs["open"] = true
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		if props.Disabled {
			attrs["data-disabled"] = "true"
		}
		return renderElement(ctx, w, "details", attrs, templ.GetChildren(ctx))
	})
}

func CollapsibleTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "summary", attrsFromDOMProps(props, "collapsible-trigger", "cursor-pointer list-none"), templ.GetChildren(ctx))
	})
}

func CollapsibleContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "collapsible-content", "pt-2"), templ.GetChildren(ctx))
	})
}

type DialogProps struct {
	DOMProps
	Open            bool
	DefaultOpen     bool
	Modal           bool
	ShowCloseButton bool
}

func Dialog(props DialogProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "dialog", "fixed inset-0 z-50 grid place-items-center")
		if props.Open {
			attrs["open"] = true
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		if props.Modal {
			attrs["data-modal"] = "true"
		}
		if props.ShowCloseButton {
			attrs["data-show-close-button"] = "true"
		}
		return renderElement(ctx, w, "dialog", attrs, templ.GetChildren(ctx))
	})
}

func DialogTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "dialog-trigger", ""), templ.GetChildren(ctx))
	})
}

func DialogPortal(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dialog-portal", ""), templ.GetChildren(ctx))
	})
}

func DialogOverlay(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dialog-overlay", "fixed inset-0 bg-black/50"), templ.GetChildren(ctx))
	})
}

func DialogContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dialog-content", "relative z-50 w-full max-w-lg rounded-xl border bg-background p-6 shadow-lg"), templ.GetChildren(ctx))
	})
}

func DialogHeader(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dialog-header", "flex flex-col gap-2 text-center sm:text-left"), templ.GetChildren(ctx))
	})
}

func DialogFooter(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dialog-footer", "flex flex-col-reverse gap-2 sm:flex-row sm:justify-end"), templ.GetChildren(ctx))
	})
}

func DialogTitle(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "h2", attrsFromDOMProps(props, "dialog-title", "text-lg font-semibold tracking-tight"), templ.GetChildren(ctx))
	})
}

func DialogDescription(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "p", attrsFromDOMProps(props, "dialog-description", "text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func DialogClose(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "button", attrsFromDOMProps(props, "dialog-close", ""), templ.GetChildren(ctx))
	})
}

type AlertDialogProps struct {
	DialogProps
}

func AlertDialog(props AlertDialogProps) templ.Component    { return Dialog(props.DialogProps) }
func AlertDialogTrigger(props DOMProps) templ.Component     { return DialogTrigger(props) }
func AlertDialogPortal(props DOMProps) templ.Component      { return DialogPortal(props) }
func AlertDialogOverlay(props DOMProps) templ.Component     { return DialogOverlay(props) }
func AlertDialogContent(props DOMProps) templ.Component     { return DialogContent(props) }
func AlertDialogHeader(props DOMProps) templ.Component      { return DialogHeader(props) }
func AlertDialogFooter(props DOMProps) templ.Component      { return DialogFooter(props) }
func AlertDialogTitle(props DOMProps) templ.Component       { return DialogTitle(props) }
func AlertDialogDescription(props DOMProps) templ.Component { return DialogDescription(props) }
func AlertDialogAction(props DOMProps) templ.Component      { return DialogClose(props) }
func AlertDialogCancel(props DOMProps) templ.Component      { return DialogClose(props) }

type SheetProps struct {
	DialogProps
	Side            string
	ShowCloseButton bool
}

func Sheet(props SheetProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "sheet", "fixed inset-0 z-50")
		if props.Side != "" {
			attrs["data-side"] = props.Side
		}
		if props.ShowCloseButton {
			attrs["data-show-close-button"] = "true"
		}
		if props.Open {
			attrs["open"] = true
		}
		return renderElement(ctx, w, "dialog", attrs, templ.GetChildren(ctx))
	})
}

func SheetTrigger(props DOMProps) templ.Component     { return DialogTrigger(props) }
func SheetPortal(props DOMProps) templ.Component      { return DialogPortal(props) }
func SheetOverlay(props DOMProps) templ.Component     { return DialogOverlay(props) }
func SheetHeader(props DOMProps) templ.Component      { return DialogHeader(props) }
func SheetFooter(props DOMProps) templ.Component      { return DialogFooter(props) }
func SheetTitle(props DOMProps) templ.Component       { return DialogTitle(props) }
func SheetDescription(props DOMProps) templ.Component { return DialogDescription(props) }
func SheetClose(props DOMProps) templ.Component       { return DialogClose(props) }

func SheetContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "sheet-content", "relative z-50 h-full w-full max-w-lg bg-background p-6 shadow-lg"), templ.GetChildren(ctx))
	})
}

type DrawerProps struct {
	SheetProps
	Dismissible      bool
	SnapPoints       []string
	DefaultSnapPoint string
}

func Drawer(props DrawerProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "drawer", "fixed inset-0 z-50")
		if props.Side != "" {
			attrs["data-side"] = props.Side
		}
		if props.Dismissible {
			attrs["data-dismissible"] = "true"
		}
		if len(props.SnapPoints) > 0 {
			attrs["data-snap-points"] = props.SnapPoints
		}
		if props.DefaultSnapPoint != "" {
			attrs["data-default-snap-point"] = props.DefaultSnapPoint
		}
		return renderElement(ctx, w, "dialog", attrs, templ.GetChildren(ctx))
	})
}

func DrawerTrigger(props DOMProps) templ.Component     { return DialogTrigger(props) }
func DrawerPortal(props DOMProps) templ.Component      { return DialogPortal(props) }
func DrawerOverlay(props DOMProps) templ.Component     { return DialogOverlay(props) }
func DrawerHeader(props DOMProps) templ.Component      { return DialogHeader(props) }
func DrawerFooter(props DOMProps) templ.Component      { return DialogFooter(props) }
func DrawerTitle(props DOMProps) templ.Component       { return DialogTitle(props) }
func DrawerDescription(props DOMProps) templ.Component { return DialogDescription(props) }
func DrawerClose(props DOMProps) templ.Component       { return DialogClose(props) }

func DrawerContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "drawer-content", "relative z-50 h-full w-full bg-background p-6 shadow-lg"), templ.GetChildren(ctx))
	})
}

type PopoverProps struct {
	DOMProps
	Open        bool
	DefaultOpen bool
	Side        string
	Align       string
	SideOffset  string
	AlignOffset string
	Modal       bool
}

func Popover(props PopoverProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "popover", "relative inline-block")
		if props.Open {
			attrs["data-open"] = "true"
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
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
		if props.AlignOffset != "" {
			attrs["data-align-offset"] = props.AlignOffset
		}
		if props.Modal {
			attrs["data-modal"] = "true"
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func PopoverTrigger(props DOMProps) templ.Component { return DialogTrigger(props) }
func PopoverAnchor(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "popover-anchor", ""), templ.GetChildren(ctx))
	})
}
func PopoverContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "popover-content", "z-50 rounded-md border bg-popover p-4 text-popover-foreground shadow-md outline-none"), templ.GetChildren(ctx))
	})
}
func PopoverHeader(props DOMProps) templ.Component      { return DialogHeader(props) }
func PopoverTitle(props DOMProps) templ.Component       { return DialogTitle(props) }
func PopoverDescription(props DOMProps) templ.Component { return DialogDescription(props) }

type HoverCardProps struct {
	DOMProps
	OpenDelayMs  int
	CloseDelayMs int
	Side         string
	Align        string
	SideOffset   string
}

func HoverCard(props HoverCardProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "hover-card", "relative inline-block")
		if props.OpenDelayMs > 0 {
			attrs["data-open-delay"] = props.OpenDelayMs
		}
		if props.CloseDelayMs > 0 {
			attrs["data-close-delay"] = props.CloseDelayMs
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

func HoverCardTrigger(props DOMProps) templ.Component { return DialogTrigger(props) }
func HoverCardContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "hover-card-content", "z-50 w-80 rounded-md border bg-popover p-4 text-popover-foreground shadow-md"), templ.GetChildren(ctx))
	})
}

type TooltipProviderProps struct {
	DOMProps
	DelayDuration int
}

func TooltipProvider(props TooltipProviderProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "tooltip-provider", "")
		if props.DelayDuration > 0 {
			attrs["data-delay-duration"] = props.DelayDuration
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

type TooltipProps struct {
	DOMProps
	Open        bool
	DefaultOpen bool
	Side        string
	Align       string
	SideOffset  string
	AlignOffset string
}

func Tooltip(props TooltipProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "tooltip", "inline-block")
		if props.Open {
			attrs["data-open"] = "true"
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
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
		if props.AlignOffset != "" {
			attrs["data-align-offset"] = props.AlignOffset
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func TooltipTrigger(props DOMProps) templ.Component { return DialogTrigger(props) }

func TooltipContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "tooltip-content", "z-50 rounded-md bg-primary px-3 py-1.5 text-xs text-primary-foreground shadow-md"), templ.GetChildren(ctx))
	})
}
