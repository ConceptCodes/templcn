package views

import (
	"fmt"
	"strings"
)

type ComponentDocEntry struct {
	Slug        string
	Title       string
	Description string
	Install     string
	Usage       string
	API         []string
	Example     []string
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
	{Slug: "accordion", Title: "Accordion", Description: "A vertically stacked set of collapsible panels.", Install: "npx shadcn@latest add accordion", Usage: "@ui.Accordion(...)", API: []string{"type", "value", "defaultValue", "collapsible"}, Example: []string{"Accordion", "AccordionItem", "AccordionTrigger", "AccordionContent"}},
	{Slug: "alert", Title: "Alert", Description: "Displays a prominent message to the user.", Install: "npx shadcn@latest add alert", Usage: "@ui.Alert(...)", API: []string{"variant"}, Example: []string{"Alert", "AlertTitle", "AlertDescription"}},
	{Slug: "alert-dialog", Title: "Alert Dialog", Description: "A modal dialog that asks the user to confirm a destructive action.", Install: "npx shadcn@latest add alert-dialog", Usage: "@ui.AlertDialog(...)", API: []string{"open", "defaultOpen", "modal"}, Example: []string{"AlertDialog", "AlertDialogAction", "AlertDialogCancel"}},
	{Slug: "aspect-ratio", Title: "Aspect Ratio", Description: "Maintains a consistent width-to-height ratio.", Install: "npx shadcn@latest add aspect-ratio", Usage: "@ui.AspectRatio(...)", API: []string{"ratio"}, Example: []string{"AspectRatio"}},
	{Slug: "avatar", Title: "Avatar", Description: "Displays a user image with a graceful fallback.", Install: "npx shadcn@latest add avatar", Usage: "@ui.Avatar(...)", API: []string{"size"}, Example: []string{"Avatar", "AvatarImage", "AvatarFallback"}},
	{Slug: "badge", Title: "Badge", Description: "Small status or metadata labels.", Install: "npx shadcn@latest add badge", Usage: "@ui.Badge(...)", API: []string{"variant"}, Example: []string{"Badge"}},
	{Slug: "breadcrumb", Title: "Breadcrumb", Description: "A navigation trail for the current location.", Install: "npx shadcn@latest add breadcrumb", Usage: "@ui.Breadcrumb(...)", API: []string{"href"}, Example: []string{"Breadcrumb", "BreadcrumbLink", "BreadcrumbPage"}},
	{Slug: "button", Title: "Button", Description: "Displays a button or button-like component.", Install: "npx shadcn@latest add button", Usage: "@ui.Button(...)", API: []string{"variant", "size", "href", "disabled"}, Example: []string{"Button", "ButtonGroup"}},
	{Slug: "button-group", Title: "Button Group", Description: "Groups buttons into a single control surface.", Install: "npx shadcn@latest add button-group", Usage: "@ui.ButtonGroup(...)", API: []string{"orientation"}, Example: []string{"ButtonGroup", "ButtonGroupText", "ButtonGroupSeparator"}},
	{Slug: "calendar", Title: "Calendar", Description: "A month view calendar with flexible selection modes.", Install: "npx shadcn@latest add calendar", Usage: "@ui.Calendar(...)", API: []string{"mode", "selected", "month", "locale"}, Example: []string{"Calendar", "CalendarDayButton"}},
	{Slug: "card", Title: "Card", Description: "A flexible container for related information.", Install: "npx shadcn@latest add card", Usage: "@ui.Card(...)", API: []string{"none"}, Example: []string{"Card", "CardHeader", "CardContent"}},
	{Slug: "carousel", Title: "Carousel", Description: "A swipeable set of slides with navigation controls.", Install: "npx shadcn@latest add carousel", Usage: "@ui.Carousel(...)", API: []string{"orientation", "loop", "align"}, Example: []string{"Carousel", "CarouselContent", "CarouselItem"}},
	{Slug: "chart", Title: "Chart", Description: "A shell for chart configuration, tooltips, and legend styling.", Install: "npx shadcn@latest add chart", Usage: "@ui.ChartContainer(...)", API: []string{"config", "series"}, Example: []string{"ChartContainer", "ChartTooltip", "ChartLegend"}},
	{Slug: "checkbox", Title: "Checkbox", Description: "A native checkbox with shadcn styling.", Install: "npx shadcn@latest add checkbox", Usage: "@ui.Checkbox(...)", API: []string{"checked", "defaultChecked", "required"}, Example: []string{"Checkbox"}},
	{Slug: "collapsible", Title: "Collapsible", Description: "A lighter disclosure primitive than accordion.", Install: "npx shadcn@latest add collapsible", Usage: "@ui.Collapsible(...)", API: []string{"open", "defaultOpen"}, Example: []string{"Collapsible", "CollapsibleTrigger", "CollapsibleContent"}},
	{Slug: "combobox", Title: "Combobox", Description: "Searchable selection with single or multi-value support.", Install: "npx shadcn@latest add combobox", Usage: "@ui.Combobox(...)", API: []string{"multiple", "value", "values", "filter"}, Example: []string{"Combobox", "ComboboxItem", "ComboboxChips"}},
	{Slug: "command", Title: "Command", Description: "A searchable command palette component.", Install: "npx shadcn@latest add command", Usage: "@ui.Command(...)", API: []string{"value", "defaultValue", "open"}, Example: []string{"Command", "CommandInput", "CommandItem"}},
	{Slug: "context-menu", Title: "Context Menu", Description: "A right-click menu with nested submenus.", Install: "npx shadcn@latest add context-menu", Usage: "@ui.ContextMenu(...)", API: []string{"variant", "inset"}, Example: []string{"ContextMenu", "ContextMenuItem", "ContextMenuSub"}},
	{Slug: "data-table", Title: "Data Table", Description: "A server-rendered table pattern for app data.", Install: "npx shadcn@latest add data-table", Usage: "@ui.DataTable(...)", API: []string{"columns", "rows", "page", "filters"}, Example: []string{"DataTable", "DataTableToolbar", "DataTablePagination"}},
	{Slug: "date-picker", Title: "Date Picker", Description: "A calendar popover for single or range dates.", Install: "npx shadcn@latest add date-picker", Usage: "@ui.DatePicker(...)", API: []string{"mode", "value", "range"}, Example: []string{"DatePicker", "DateRangePicker"}},
	{Slug: "dialog", Title: "Dialog", Description: "A modal dialog used for forms and transient tasks.", Install: "npx shadcn@latest add dialog", Usage: "@ui.Dialog(...)", API: []string{"open", "defaultOpen", "modal"}, Example: []string{"Dialog", "DialogTitle", "DialogDescription"}},
	{Slug: "direction", Title: "Direction", Description: "A small RTL/LTR direction provider.", Install: "npx shadcn@latest add direction", Usage: "@ui.DirectionProvider(...)", API: []string{"direction"}, Example: []string{"DirectionProvider"}},
	{Slug: "drawer", Title: "Drawer", Description: "A slide-over surface with mobile-friendly motion.", Install: "npx shadcn@latest add drawer", Usage: "@ui.Drawer(...)", API: []string{"side", "dismissible", "snapPoints"}, Example: []string{"Drawer", "DrawerContent"}},
	{Slug: "dropdown-menu", Title: "Dropdown Menu", Description: "A context-driven popover menu.", Install: "npx shadcn@latest add dropdown-menu", Usage: "@ui.DropdownMenu(...)", API: []string{"sideOffset", "variant"}, Example: []string{"DropdownMenu", "DropdownMenuContent", "DropdownMenuItem"}},
	{Slug: "empty", Title: "Empty", Description: "A polished empty state helper.", Install: "npx shadcn@latest add empty", Usage: "@ui.Empty(...)", API: []string{"variant"}, Example: []string{"Empty", "EmptyTitle", "EmptyDescription"}},
	{Slug: "field", Title: "Field", Description: "A form field grouping helper.", Install: "npx shadcn@latest add field", Usage: "@ui.Field(...)", API: []string{"orientation", "invalid", "disabled"}, Example: []string{"FieldSet", "FieldLabel", "FieldError"}},
	{Slug: "hover-card", Title: "Hover Card", Description: "A hover/focus-driven preview popover.", Install: "npx shadcn@latest add hover-card", Usage: "@ui.HoverCard(...)", API: []string{"openDelayMs", "closeDelayMs"}, Example: []string{"HoverCard", "HoverCardContent"}},
	{Slug: "input", Title: "Input", Description: "A native text input styled for shadcn themes.", Install: "npx shadcn@latest add input", Usage: "@ui.Input(...)", API: []string{"type", "name", "value"}, Example: []string{"Input"}},
	{Slug: "input-group", Title: "Input Group", Description: "Input with addons, buttons, and labels.", Install: "npx shadcn@latest add input-group", Usage: "@ui.InputGroup(...)", API: []string{"align"}, Example: []string{"InputGroup", "InputGroupAddon", "InputGroupInput"}},
	{Slug: "input-otp", Title: "Input OTP", Description: "One-time code entry with slot-based UI.", Install: "npx shadcn@latest add input-otp", Usage: "@ui.InputOTP(...)", API: []string{"maxlength", "pattern"}, Example: []string{"InputOTP", "InputOTPSlot"}},
	{Slug: "item", Title: "Item", Description: "A generic list/card item surface.", Install: "npx shadcn@latest add item", Usage: "@ui.Item(...)", API: []string{"variant", "size"}, Example: []string{"Item", "ItemHeader", "ItemActions"}},
	{Slug: "kbd", Title: "Kbd", Description: "A keyboard shortcut token.", Install: "npx shadcn@latest add kbd", Usage: "@ui.Kbd(...)", API: []string{"size", "inset"}, Example: []string{"Kbd"}},
	{Slug: "label", Title: "Label", Description: "A native label wrapper.", Install: "npx shadcn@latest add label", Usage: "@ui.Label(...)", API: []string{"for"}, Example: []string{"Label"}},
	{Slug: "menubar", Title: "Menubar", Description: "A horizontal application menu.", Install: "npx shadcn@latest add menubar", Usage: "@ui.Menubar(...)", API: []string{"orientation"}, Example: []string{"Menubar", "MenubarItem", "MenubarSub"}},
	{Slug: "native-select", Title: "Native Select", Description: "A styled native select element.", Install: "npx shadcn@latest add native-select", Usage: "@ui.NativeSelect(...)", API: []string{"size", "required"}, Example: []string{"NativeSelect", "NativeSelectOption"}},
	{Slug: "navigation-menu", Title: "Navigation Menu", Description: "A top-level navigation pattern with a viewport.", Install: "npx shadcn@latest add navigation-menu", Usage: "@ui.NavigationMenu(...)", API: []string{"value", "orientation"}, Example: []string{"NavigationMenu", "NavigationMenuTrigger"}},
	{Slug: "pagination", Title: "Pagination", Description: "Pagination primitives for list navigation.", Install: "npx shadcn@latest add pagination", Usage: "@ui.Pagination(...)", API: []string{"href", "isActive"}, Example: []string{"Pagination", "PaginationLink"}},
	{Slug: "popover", Title: "Popover", Description: "Floating content anchored to a trigger.", Install: "npx shadcn@latest add popover", Usage: "@ui.Popover(...)", API: []string{"side", "align", "modal"}, Example: []string{"Popover", "PopoverContent"}},
	{Slug: "progress", Title: "Progress", Description: "A native progress bar.", Install: "npx shadcn@latest add progress", Usage: "@ui.Progress(...)", API: []string{"value", "max"}, Example: []string{"Progress"}},
	{Slug: "radio-group", Title: "Radio Group", Description: "A group of radio options.", Install: "npx shadcn@latest add radio-group", Usage: "@ui.RadioGroup(...)", API: []string{"name", "value", "required"}, Example: []string{"RadioGroup", "RadioGroupItem"}},
	{Slug: "resizable", Title: "Resizable", Description: "Pointer-driven resize panels.", Install: "npx shadcn@latest add resizable", Usage: "@ui.ResizablePanelGroup(...)", API: []string{"direction", "defaultSize"}, Example: []string{"ResizablePanelGroup", "ResizableHandle"}},
	{Slug: "scroll-area", Title: "Scroll Area", Description: "Styled scrollable containers and bars.", Install: "npx shadcn@latest add scroll-area", Usage: "@ui.ScrollArea(...)", API: []string{"orientation"}, Example: []string{"ScrollArea", "ScrollBar"}},
	{Slug: "select", Title: "Select", Description: "A custom select built around a listbox runtime.", Install: "npx shadcn@latest add select", Usage: "@ui.Select(...)", API: []string{"open", "defaultOpen", "required"}, Example: []string{"Select", "SelectTrigger", "SelectContent"}},
	{Slug: "separator", Title: "Separator", Description: "A horizontal or vertical divider.", Install: "npx shadcn@latest add separator", Usage: "@ui.Separator(...)", API: []string{"orientation"}, Example: []string{"Separator"}},
	{Slug: "sheet", Title: "Sheet", Description: "A slide-over panel.", Install: "npx shadcn@latest add sheet", Usage: "@ui.Sheet(...)", API: []string{"side", "showCloseButton"}, Example: []string{"Sheet", "SheetContent"}},
	{Slug: "sidebar", Title: "Sidebar", Description: "A full app shell sidebar with navigation controls.", Install: "npx shadcn@latest add sidebar", Usage: "@ui.Sidebar(...)", API: []string{"variant", "collapsible", "side"}, Example: []string{"SidebarProvider", "Sidebar", "SidebarMenuButton"}},
	{Slug: "skeleton", Title: "Skeleton", Description: "A placeholder shimmer block.", Install: "npx shadcn@latest add skeleton", Usage: "@ui.Skeleton(...)", API: []string{"none"}, Example: []string{"Skeleton"}},
	{Slug: "slider", Title: "Slider", Description: "A range slider for numeric input.", Install: "npx shadcn@latest add slider", Usage: "@ui.Slider(...)", API: []string{"min", "max", "step"}, Example: []string{"Slider"}},
	{Slug: "sonner", Title: "Sonner", Description: "A toast experience with shadcn defaults.", Install: "npx shadcn@latest add sonner", Usage: "@ui.Toaster(...)", API: []string{"theme", "position"}, Example: []string{"Toaster"}},
	{Slug: "spinner", Title: "Spinner", Description: "A lightweight loading spinner.", Install: "npx shadcn@latest add spinner", Usage: "@ui.Spinner(...)", API: []string{"size", "label"}, Example: []string{"Spinner"}},
	{Slug: "switch", Title: "Switch", Description: "A checkbox-like binary control.", Install: "npx shadcn@latest add switch", Usage: "@ui.Switch(...)", API: []string{"size", "checked", "defaultChecked"}, Example: []string{"Switch"}},
	{Slug: "table", Title: "Table", Description: "Table primitives with shadcn table styling.", Install: "npx shadcn@latest add table", Usage: "@ui.Table(...)", API: []string{"none"}, Example: []string{"Table", "TableRow", "TableCell"}},
	{Slug: "tabs", Title: "Tabs", Description: "Tab buttons and associated panels.", Install: "npx shadcn@latest add tabs", Usage: "@ui.Tabs(...)", API: []string{"value", "defaultValue", "orientation"}, Example: []string{"Tabs", "TabsList", "TabsTrigger"}},
	{Slug: "textarea", Title: "Textarea", Description: "A native multi-line input.", Install: "npx shadcn@latest add textarea", Usage: "@ui.Textarea(...)", API: []string{"rows", "placeholder"}, Example: []string{"Textarea"}},
	{Slug: "toast", Title: "Toast", Description: "Legacy toast primitives.", Install: "npx shadcn@latest add toast", Usage: "@ui.Toast(...)", API: []string{"variant", "duration"}, Example: []string{"Toast", "ToastProvider"}},
	{Slug: "toggle", Title: "Toggle", Description: "A pressed-state button.", Install: "npx shadcn@latest add toggle", Usage: "@ui.Toggle(...)", API: []string{"variant", "size"}, Example: []string{"Toggle"}},
	{Slug: "toggle-group", Title: "Toggle Group", Description: "A group of toggle buttons.", Install: "npx shadcn@latest add toggle-group", Usage: "@ui.ToggleGroup(...)", API: []string{"type", "variant", "size"}, Example: []string{"ToggleGroup", "ToggleGroupItem"}},
	{Slug: "tooltip", Title: "Tooltip", Description: "A floating hint for compact controls.", Install: "npx shadcn@latest add tooltip", Usage: "@ui.Tooltip(...)", API: []string{"delayDuration", "side", "align"}, Example: []string{"Tooltip", "TooltipContent"}},
	{Slug: "typography", Title: "Typography", Description: "A docs-oriented prose helper layer.", Install: "npx shadcn@latest add typography", Usage: "@ui.H1(...)", API: []string{"none"}, Example: []string{"H1", "P", "InlineCode"}},
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
