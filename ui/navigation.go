package ui

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func Breadcrumb(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "nav", attrsFromDOMProps(props, "breadcrumb", ""), templ.GetChildren(ctx))
	})
}

func BreadcrumbList(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "ol", attrsFromDOMProps(props, "breadcrumb-list", "flex flex-wrap items-center gap-1.5 text-sm text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func BreadcrumbItem(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "li", attrsFromDOMProps(props, "breadcrumb-item", "inline-flex items-center gap-1.5"), templ.GetChildren(ctx))
	})
}

type BreadcrumbLinkProps struct {
	DOMProps
	Href string
}

func BreadcrumbLink(props BreadcrumbLinkProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "breadcrumb-link", "transition-colors hover:text-foreground")
		if props.Href != "" {
			attrs["href"] = props.Href
		}
		return renderElement(ctx, w, "a", attrs, templ.GetChildren(ctx))
	})
}

func BreadcrumbPage(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "breadcrumb-page", "font-normal text-foreground"), templ.GetChildren(ctx))
	})
}

func BreadcrumbSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "li", attrsFromDOMProps(props, "breadcrumb-separator", "mx-1 text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func BreadcrumbEllipsis(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "breadcrumb-ellipsis", "flex h-9 w-9 items-center justify-center"), templ.GetChildren(ctx))
	})
}

func Pagination(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "nav", attrsFromDOMProps(props, "pagination", "mx-auto flex w-full justify-center"), templ.GetChildren(ctx))
	})
}

func PaginationContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "ul", attrsFromDOMProps(props, "pagination-content", "flex flex-row items-center gap-1"), templ.GetChildren(ctx))
	})
}

func PaginationItem(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "li", attrsFromDOMProps(props, "pagination-item", ""), templ.GetChildren(ctx))
	})
}

type PaginationLinkProps struct {
	DOMProps
	Href     string
	IsActive bool
	Size     string
}

func PaginationLink(props PaginationLinkProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "inline-flex h-9 min-w-9 items-center justify-center rounded-md border border-input bg-background px-3 text-sm shadow-xs transition-colors hover:bg-accent hover:text-accent-foreground"
		if props.IsActive {
			className = cn(className, "bg-primary text-primary-foreground hover:bg-primary/90")
		}
		if props.Size == "sm" {
			className = cn(className, "h-8 min-w-8 px-2")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "pagination-link", className)
		if props.Href != "" {
			attrs["href"] = props.Href
		}
		return renderElement(ctx, w, "a", attrs, templ.GetChildren(ctx))
	})
}

func PaginationPrevious(props PaginationLinkProps) templ.Component { return PaginationLink(props) }
func PaginationNext(props PaginationLinkProps) templ.Component     { return PaginationLink(props) }

func PaginationEllipsis(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "pagination-ellipsis", "inline-flex h-9 w-9 items-center justify-center"), templ.GetChildren(ctx))
	})
}

type TabsProps struct {
	DOMProps
	Value        string
	DefaultValue string
	Orientation  string
}

type TabsListProps struct {
	DOMProps
	Variant string
}

func Tabs(props TabsProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "tabs", "grid gap-2")
		attrs["data-tabs-root"] = "true"
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.DefaultValue != "" {
			attrs["data-default-value"] = props.DefaultValue
		}
		if props.Orientation != "" {
			attrs["data-orientation"] = props.Orientation
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func TabsList(props TabsListProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "inline-flex h-9 items-center justify-center rounded-lg bg-muted p-1 text-muted-foreground"
		if props.Variant == "line" {
			className = "inline-flex h-10 items-center gap-4 border-b border-border bg-transparent p-0 text-muted-foreground"
		}
		attrs := attrsFromDOMProps(props.DOMProps, "tabs-list", className)
		attrs["role"] = "tablist"
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

type TabsTriggerProps struct {
	DOMProps
	Value    string
	Active   bool
	Disabled bool
}

func TabsTrigger(props TabsTriggerProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "tabs-trigger", "inline-flex items-center justify-center whitespace-nowrap rounded-md px-3 py-1.5 text-sm font-medium ring-offset-background transition-all focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 data-[state=active]:bg-background data-[state=active]:text-foreground data-[state=active]:shadow")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.Active {
			attrs["data-state"] = "active"
			attrs["aria-selected"] = "true"
			attrs["tabindex"] = "0"
		} else {
			attrs["data-state"] = "inactive"
			attrs["aria-selected"] = "false"
			attrs["tabindex"] = "-1"
		}
		attrs["role"] = "tab"
		if props.Disabled {
			attrs["disabled"] = true
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

type TabsContentProps struct {
	DOMProps
	Value  string
	Active bool
}

func TabsContent(props TabsContentProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "tabs-content", "mt-2 outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		attrs["role"] = "tabpanel"
		if props.Active {
			attrs["data-state"] = "active"
		} else {
			attrs["data-state"] = "inactive"
			attrs["hidden"] = true
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

type ToggleProps struct {
	DOMProps
	Variant        string
	Size           string
	Pressed        bool
	DefaultPressed bool
	Disabled       bool
	Value          string
}

func Toggle(props ToggleProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "inline-flex items-center justify-center rounded-md text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 data-[state=on]:bg-accent data-[state=on]:text-accent-foreground"
		if props.Variant == "outline" {
			className = cn(className, "border border-input bg-transparent shadow-xs")
		}
		switch props.Size {
		case "sm":
			className = cn(className, "h-8 px-3")
		case "lg":
			className = cn(className, "h-10 px-4")
		default:
			className = cn(className, "h-9 px-3")
		}
		attrs := attrsFromDOMProps(props.DOMProps, "toggle", className)
		if props.Pressed {
			attrs["aria-pressed"] = "true"
		}
		if props.DefaultPressed {
			attrs["data-default-pressed"] = "true"
		}
		if props.Disabled {
			attrs["disabled"] = true
		}
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

type ToggleGroupProps struct {
	DOMProps
	Type         string
	Value        []string
	DefaultValue []string
	Variant      string
	Size         string
	Spacing      string
}

type ToggleGroupItemProps struct {
	DOMProps
	Value    string
	Disabled bool
}

func ToggleGroup(props ToggleGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "toggle-group", "inline-flex gap-1")
		if props.Type != "" {
			attrs["data-type"] = props.Type
		}
		if len(props.Value) > 0 {
			attrs["data-value"] = props.Value
		}
		if len(props.DefaultValue) > 0 {
			attrs["data-default-value"] = props.DefaultValue
		}
		if props.Variant != "" {
			attrs["data-variant"] = props.Variant
		}
		if props.Size != "" {
			attrs["data-size"] = props.Size
		}
		if props.Spacing != "" {
			attrs["data-spacing"] = props.Spacing
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func ToggleGroupItem(props ToggleGroupItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "toggle-group-item", "rounded-md")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.Disabled {
			attrs["disabled"] = true
		}
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}

func menuContainer(attrs templ.Attributes, children templ.Component, ctx context.Context, w io.Writer) error {
	return renderElement(ctx, w, "div", attrs, children)
}

type DropdownMenuProps struct {
	DOMProps
	Open        bool
	DefaultOpen bool
	Modal       bool
	Side        string
	Align       string
	SideOffset  string
	AlignOffset string
}

func DropdownMenu(props DropdownMenuProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "dropdown-menu", "relative inline-block")
		attrs["data-dropdown-menu-root"] = "true"
		if props.Open {
			attrs["data-open"] = "true"
		}
		if props.DefaultOpen {
			attrs["data-default-open"] = "true"
		}
		if props.Modal {
			attrs["data-modal"] = "true"
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
		return menuContainer(attrs, templ.GetChildren(ctx), ctx, w)
	})
}

func DropdownMenuTrigger(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "dropdown-menu-trigger", "")
		attrs["data-dropdown-menu-trigger"] = "true"
		attrs["aria-haspopup"] = "menu"
		return renderElement(ctx, w, "button", attrs, templ.GetChildren(ctx))
	})
}
func DropdownMenuPortal(props DOMProps) templ.Component { return DialogPortal(props) }

func DropdownMenuContent(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props, "dropdown-menu-content", "z-50 min-w-32 rounded-md border bg-popover p-1 text-popover-foreground shadow-md")
		attrs["data-dropdown-menu-content"] = "true"
		attrs["hidden"] = true
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func DropdownMenuGroup(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dropdown-menu-group", "grid gap-1"), templ.GetChildren(ctx))
	})
}

type DropdownMenuItemProps struct {
	DOMProps
	Inset   bool
	Variant string
}

func DropdownMenuItem(props DropdownMenuItemProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		className := "relative flex cursor-pointer select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground"
		if props.Inset {
			className = cn(className, "pl-8")
		}
		if props.Variant == "destructive" {
			className = cn(className, "text-destructive")
		}
		return renderElement(ctx, w, "button", attrsFromDOMProps(props.DOMProps, "dropdown-menu-item", className), templ.GetChildren(ctx))
	})
}

func DropdownMenuCheckboxItem(props DropdownMenuItemProps) templ.Component {
	return DropdownMenuItem(props)
}

type DropdownMenuRadioGroupProps struct {
	DOMProps
	Value string
}

func DropdownMenuRadioGroup(props DropdownMenuRadioGroupProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "dropdown-menu-radio-group", "grid gap-1")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}

func DropdownMenuRadioItem(props DropdownMenuItemProps) templ.Component {
	return DropdownMenuItem(props)
}
func DropdownMenuLabel(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dropdown-menu-label", "px-2 py-1.5 text-sm font-semibold"), templ.GetChildren(ctx))
	})
}
func DropdownMenuSeparator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dropdown-menu-separator", "-mx-1 my-1 h-px bg-border"), nil)
	})
}
func DropdownMenuShortcut(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "span", attrsFromDOMProps(props, "dropdown-menu-shortcut", "ml-auto text-xs tracking-widest text-muted-foreground"), templ.GetChildren(ctx))
	})
}

func DropdownMenuSub(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "dropdown-menu-sub", "relative"), templ.GetChildren(ctx))
	})
}
func DropdownMenuSubTrigger(props DropdownMenuItemProps) templ.Component {
	return DropdownMenuItem(props)
}
func DropdownMenuSubContent(props DOMProps) templ.Component { return DropdownMenuContent(props) }

type ContextMenuProps struct{ DOMProps }

func ContextMenu(props ContextMenuProps) templ.Component {
	return DropdownMenu(DropdownMenuProps{DOMProps: props.DOMProps})
}
func ContextMenuTrigger(props DOMProps) templ.Component           { return DialogTrigger(props) }
func ContextMenuPortal(props DOMProps) templ.Component            { return DialogPortal(props) }
func ContextMenuContent(props DOMProps) templ.Component           { return DropdownMenuContent(props) }
func ContextMenuItem(props DropdownMenuItemProps) templ.Component { return DropdownMenuItem(props) }
func ContextMenuCheckboxItem(props DropdownMenuItemProps) templ.Component {
	return DropdownMenuCheckboxItem(props)
}
func ContextMenuRadioItem(props DropdownMenuItemProps) templ.Component {
	return DropdownMenuRadioItem(props)
}
func ContextMenuRadioGroup(props DropdownMenuRadioGroupProps) templ.Component {
	return DropdownMenuRadioGroup(props)
}
func ContextMenuLabel(props DOMProps) templ.Component      { return DropdownMenuLabel(props) }
func ContextMenuSeparator(props DOMProps) templ.Component  { return DropdownMenuSeparator(props) }
func ContextMenuShortcut(props DOMProps) templ.Component   { return DropdownMenuShortcut(props) }
func ContextMenuGroup(props DOMProps) templ.Component      { return DropdownMenuGroup(props) }
func ContextMenuSub(props DOMProps) templ.Component        { return DropdownMenuSub(props) }
func ContextMenuSubContent(props DOMProps) templ.Component { return DropdownMenuSubContent(props) }
func ContextMenuSubTrigger(props DropdownMenuItemProps) templ.Component {
	return DropdownMenuSubTrigger(props)
}

type MenubarProps struct{ DOMProps }

func Menubar(props MenubarProps) templ.Component {
	return DropdownMenu(DropdownMenuProps{DOMProps: props.DOMProps})
}
func MenubarMenu(props DOMProps) templ.Component              { return DropdownMenuGroup(props) }
func MenubarTrigger(props DOMProps) templ.Component           { return DialogTrigger(props) }
func MenubarPortal(props DOMProps) templ.Component            { return DialogPortal(props) }
func MenubarContent(props DOMProps) templ.Component           { return DropdownMenuContent(props) }
func MenubarGroup(props DOMProps) templ.Component             { return DropdownMenuGroup(props) }
func MenubarItem(props DropdownMenuItemProps) templ.Component { return DropdownMenuItem(props) }
func MenubarCheckboxItem(props DropdownMenuItemProps) templ.Component {
	return DropdownMenuCheckboxItem(props)
}
func MenubarRadioGroup(props DropdownMenuRadioGroupProps) templ.Component {
	return DropdownMenuRadioGroup(props)
}
func MenubarRadioItem(props DropdownMenuItemProps) templ.Component {
	return DropdownMenuRadioItem(props)
}
func MenubarLabel(props DOMProps) templ.Component     { return DropdownMenuLabel(props) }
func MenubarSeparator(props DOMProps) templ.Component { return DropdownMenuSeparator(props) }
func MenubarShortcut(props DOMProps) templ.Component  { return DropdownMenuShortcut(props) }
func MenubarSub(props DOMProps) templ.Component       { return DropdownMenuSub(props) }
func MenubarSubTrigger(props DropdownMenuItemProps) templ.Component {
	return DropdownMenuSubTrigger(props)
}
func MenubarSubContent(props DOMProps) templ.Component { return DropdownMenuSubContent(props) }

type NavigationMenuProps struct {
	DOMProps
	Value         string
	DefaultValue  string
	Orientation   string
	DelayDuration int
}

func NavigationMenu(props NavigationMenuProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "navigation-menu", "relative")
		if props.Value != "" {
			attrs["data-value"] = props.Value
		}
		if props.DefaultValue != "" {
			attrs["data-default-value"] = props.DefaultValue
		}
		if props.Orientation != "" {
			attrs["data-orientation"] = props.Orientation
		}
		if props.DelayDuration > 0 {
			attrs["data-delay-duration"] = props.DelayDuration
		}
		return renderElement(ctx, w, "nav", attrs, templ.GetChildren(ctx))
	})
}

func NavigationMenuList(props DOMProps) templ.Component            { return DropdownMenuGroup(props) }
func NavigationMenuItem(props DOMProps) templ.Component            { return DropdownMenuGroup(props) }
func NavigationMenuTrigger(props DOMProps) templ.Component         { return DialogTrigger(props) }
func NavigationMenuContent(props DOMProps) templ.Component         { return DropdownMenuContent(props) }
func NavigationMenuLink(props BreadcrumbLinkProps) templ.Component { return BreadcrumbLink(props) }
func NavigationMenuIndicator(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "navigation-menu-indicator", ""), templ.GetChildren(ctx))
	})
}
func NavigationMenuViewport(props DOMProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return renderElement(ctx, w, "div", attrsFromDOMProps(props, "navigation-menu-viewport", ""), templ.GetChildren(ctx))
	})
}

type DirectionProviderProps struct {
	DOMProps
	Direction string
}

func DirectionProvider(props DirectionProviderProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		attrs := attrsFromDOMProps(props.DOMProps, "direction-provider", "")
		if props.Direction != "" {
			attrs["dir"] = props.Direction
		}
		return renderElement(ctx, w, "div", attrs, templ.GetChildren(ctx))
	})
}
