# shadcn/ui for Go + templ Implementation Guide

## Goal

Build a Go + `templ` component package that gives developers a usage experience very close to current `shadcn/ui` `new-york` v4:

- same design tokens and global CSS semantics
- same component catalog as the official docs as of `2026-04-18`
- same visual language, spacing, variants, and `data-slot` naming
- similar composition model, but adapted to Go instead of React
- no React runtime anywhere in the final package
- final milestone: recreate the current `ui.shadcn.com` landing page using only our new components and `templ`

This document is the execution guide for the coding agent that will open one PR per component or milestone.

## Source Of Truth

Use these as the canonical references while implementing:

1. Official component inventory: `https://ui.shadcn.com/docs/components`
2. Official v4 registry source: `https://github.com/shadcn-ui/ui/tree/main/apps/v4/registry/new-york-v4/ui`
3. Official manual install CSS and token baseline: `https://ui.shadcn.com/docs/installation/manual`
4. Tailwind v4 migration notes and token model: `https://ui.shadcn.com/docs/tailwind-v4`
5. Base UI / Radix abstraction note: `https://ui.shadcn.com/docs/changelog/2026-01-base-ui`
6. Current homepage structure: `https://ui.shadcn.com/`

## Important Decisions

### 1. Tailwind v4 first

Do not port components on top of the current Tailwind v3-style setup in this repo. The current shadcn v4 codebase assumes:

- `@theme inline`
- OKLCH tokens
- `tw-animate-css`
- `data-slot` attributes on primitives

If we skip the Tailwind v4 migration first, we will spend time translating classes twice.

### 2. Use `templ` + small vanilla JS runtime + `@floating-ui/dom`

We are not using React, Radix, or Base UI directly.

We will implement:

- semantic HTML first
- `dialog` where practical
- a small client runtime for focus management, dismissable layers, typeahead, roving focus, and controlled overlays
- `@floating-ui/dom` for positioning popovers, tooltips, menus, comboboxes, and hover cards

### 3. Keep shadcn naming, adapt the API to Go

React patterns that should **not** be copied literally:

- `asChild`
- React callbacks like `onOpenChange`
- React context APIs exposed directly to end users
- React-only dependencies such as Recharts, `react-day-picker`, Vaul, `cmdk`, `sonner`, and Embla React

Go-native replacements:

- `Element string` and `Href string` instead of `asChild`
- hidden inputs plus DOM custom events instead of React callbacks
- component props structs instead of JSX prop bags

### 4. Preserve `data-slot`

All components should emit the same `data-slot` names used by shadcn v4 whenever practical. This matters for:

- styling parity
- docs parity
- easier diffing against upstream
- easier downstream overrides

## Package Shape

Use a single importable Go package for the public API:

```text
/ui
  accordion.templ
  accordion.go
  alert.templ
  alert.go
  ...
  types.go
  enums.go
  runtime.go
  icons/
  internal/
```

Why a single `ui` package:

- closest Go equivalent to “import the components and use them”
- avoids dozens of per-component imports
- easier discoverability in editors
- simpler examples and docs

Do not keep the current per-component package layout.

## Common Go API Conventions

Every public component should follow these rules.

### Shared base props

Every renderable component accepts:

```go
type DOMProps struct {
	ID      string
	Class   string
	Element string
	Attrs   templ.Attributes
}
```

Rules:

- `Element` is optional; when empty, render the canonical element
- `Attrs` is how users pass `aria-*`, `data-*`, `hx-*`, `onclick`, and other raw attributes
- `Class` always appends to the upstream base classes

### Children

Prefer the standard pattern:

```go
func Button(props ButtonProps, children ...templ.Component) templ.Component
```

For tiny leaf components that are often text-only, allow either children or a text prop, but do not require a text prop.

### State and events

Do not expose Go callback functions.

Interactive components should support:

- controlled initial state via props like `Open`, `Value`, `Checked`
- uncontrolled defaults via props like `DefaultOpen`, `DefaultValue`, `DefaultChecked`
- form submission via hidden inputs when a control is not a native form element
- DOM custom events for advanced integration

Standard custom events to dispatch:

- `shadcn-templ:open-change`
- `shadcn-templ:value-change`
- `shadcn-templ:checked-change`
- `shadcn-templ:select`
- `shadcn-templ:carousel-change`
- `shadcn-templ:sidebar-toggle`
- `shadcn-templ:toast-dismiss`

### Accessibility

Every component PR must include:

- keyboard support
- correct roles and aria attributes
- focus-visible states matching shadcn
- RTL-safe behavior where applicable

## Styling Baseline

The package must adopt the current shadcn v4 token model. This is the baseline CSS contract we should replicate in our own global stylesheet.

```css
@import "tailwindcss";
@import "tw-animate-css";

@custom-variant dark (&:is(.dark *));

@theme inline {
  --color-background: var(--background);
  --color-foreground: var(--foreground);
  --color-card: var(--card);
  --color-card-foreground: var(--card-foreground);
  --color-popover: var(--popover);
  --color-popover-foreground: var(--popover-foreground);
  --color-primary: var(--primary);
  --color-primary-foreground: var(--primary-foreground);
  --color-secondary: var(--secondary);
  --color-secondary-foreground: var(--secondary-foreground);
  --color-muted: var(--muted);
  --color-muted-foreground: var(--muted-foreground);
  --color-accent: var(--accent);
  --color-accent-foreground: var(--accent-foreground);
  --color-destructive: var(--destructive);
  --color-destructive-foreground: var(--destructive-foreground);
  --color-border: var(--border);
  --color-input: var(--input);
  --color-ring: var(--ring);
  --color-chart-1: var(--chart-1);
  --color-chart-2: var(--chart-2);
  --color-chart-3: var(--chart-3);
  --color-chart-4: var(--chart-4);
  --color-chart-5: var(--chart-5);
  --radius-sm: calc(var(--radius) * 0.6);
  --radius-md: calc(var(--radius) * 0.8);
  --radius-lg: var(--radius);
  --radius-xl: calc(var(--radius) * 1.4);
  --radius-2xl: calc(var(--radius) * 1.8);
  --radius-3xl: calc(var(--radius) * 2.2);
  --radius-4xl: calc(var(--radius) * 2.6);
  --color-sidebar: var(--sidebar);
  --color-sidebar-foreground: var(--sidebar-foreground);
  --color-sidebar-primary: var(--sidebar-primary);
  --color-sidebar-primary-foreground: var(--sidebar-primary-foreground);
  --color-sidebar-accent: var(--sidebar-accent);
  --color-sidebar-accent-foreground: var(--sidebar-accent-foreground);
  --color-sidebar-border: var(--sidebar-border);
  --color-sidebar-ring: var(--sidebar-ring);
}

:root {
  --radius: 0.625rem;
  --background: oklch(1 0 0);
  --foreground: oklch(0.145 0 0);
  --card: oklch(1 0 0);
  --card-foreground: oklch(0.145 0 0);
  --popover: oklch(1 0 0);
  --popover-foreground: oklch(0.145 0 0);
  --primary: oklch(0.205 0 0);
  --primary-foreground: oklch(0.985 0 0);
  --secondary: oklch(0.97 0 0);
  --secondary-foreground: oklch(0.205 0 0);
  --muted: oklch(0.97 0 0);
  --muted-foreground: oklch(0.556 0 0);
  --accent: oklch(0.97 0 0);
  --accent-foreground: oklch(0.205 0 0);
  --destructive: oklch(0.577 0.245 27.325);
  --border: oklch(0.922 0 0);
  --input: oklch(0.922 0 0);
  --ring: oklch(0.708 0 0);
  --chart-1: oklch(0.646 0.222 41.116);
  --chart-2: oklch(0.6 0.118 184.704);
  --chart-3: oklch(0.398 0.07 227.392);
  --chart-4: oklch(0.828 0.189 84.429);
  --chart-5: oklch(0.769 0.188 70.08);
  --sidebar: oklch(0.985 0 0);
  --sidebar-foreground: oklch(0.145 0 0);
  --sidebar-primary: oklch(0.205 0 0);
  --sidebar-primary-foreground: oklch(0.985 0 0);
  --sidebar-accent: oklch(0.97 0 0);
  --sidebar-accent-foreground: oklch(0.205 0 0);
  --sidebar-border: oklch(0.922 0 0);
  --sidebar-ring: oklch(0.708 0 0);
}

.dark {
  --background: oklch(0.145 0 0);
  --foreground: oklch(0.985 0 0);
  --card: oklch(0.205 0 0);
  --card-foreground: oklch(0.985 0 0);
  --popover: oklch(0.205 0 0);
  --popover-foreground: oklch(0.985 0 0);
  --primary: oklch(0.922 0 0);
  --primary-foreground: oklch(0.205 0 0);
  --secondary: oklch(0.269 0 0);
  --secondary-foreground: oklch(0.985 0 0);
  --muted: oklch(0.269 0 0);
  --muted-foreground: oklch(0.708 0 0);
  --accent: oklch(0.269 0 0);
  --accent-foreground: oklch(0.985 0 0);
  --destructive: oklch(0.704 0.191 22.216);
  --border: oklch(1 0 0 / 10%);
  --input: oklch(1 0 0 / 15%);
  --ring: oklch(0.556 0 0);
  --chart-1: oklch(0.488 0.243 264.376);
  --chart-2: oklch(0.696 0.17 162.48);
  --chart-3: oklch(0.769 0.188 70.08);
  --chart-4: oklch(0.627 0.265 303.9);
  --chart-5: oklch(0.645 0.246 16.439);
  --sidebar: oklch(0.205 0 0);
  --sidebar-foreground: oklch(0.985 0 0);
  --sidebar-primary: oklch(0.488 0.243 264.376);
  --sidebar-primary-foreground: oklch(0.985 0 0);
  --sidebar-accent: oklch(0.269 0 0);
  --sidebar-accent-foreground: oklch(0.985 0 0);
  --sidebar-border: oklch(1 0 0 / 10%);
  --sidebar-ring: oklch(0.556 0 0);
}

@layer base {
  * {
    @apply border-border outline-ring/50;
  }

  body {
    @apply bg-background text-foreground;
  }
}
```

Implementation notes:

- add the pointer cursor exception if we want pre-v4 button behavior, otherwise match current upstream
- sidebar tokens are required before `Sidebar`
- chart tokens are required before `Chart`

## Runtime Modules

Build one small browser runtime instead of per-component ad-hoc scripts.

Required modules:

1. `portal`
2. `dismissable-layer`
3. `focus-scope`
4. `presence`
5. `floating`
6. `roving-focus`
7. `typeahead`
8. `selectable-listbox`
9. `toast-store`
10. `carousel`
11. `resizable`
12. `sidebar`

Guidelines:

- use `@floating-ui/dom` only where positioning is required
- do not require hydration for static components
- initialize behavior by `data-slot` and `data-ui` markers
- keep all behavior progressive; markup should remain usable if JS fails

## PR Order

The agent opening PRs should use this sequence.

### Foundation PRs

1. Tailwind v4 migration, `globals.css`, `cn` helper, `templ` package consolidation
2. shared enums and base props
3. runtime shell, portal root, floating layer, custom events
4. icon strategy

### Low-complexity component PRs

`Button`, `Badge`, `Alert`, `AspectRatio`, `Card`, `Separator`, `Skeleton`, `Spinner`, `Kbd`, `Label`, `Typography`, `Progress`

### Form and field component PRs

`Input`, `Textarea`, `Checkbox`, `Switch`, `RadioGroup`, `ButtonGroup`, `InputGroup`, `NativeSelect`, `Field`, `Form`, `Slider`, `InputOTP`

### Disclosure and overlay PRs

`Accordion`, `Collapsible`, `Dialog`, `AlertDialog`, `Sheet`, `Drawer`, `Popover`, `HoverCard`, `Tooltip`

### Menu and navigation PRs

`Breadcrumb`, `Pagination`, `Tabs`, `Toggle`, `ToggleGroup`, `DropdownMenu`, `ContextMenu`, `Menubar`, `NavigationMenu`, `Direction`

### Data and collection PRs

`Avatar`, `Empty`, `Item`, `Table`, `ScrollArea`, `Resizable`, `Carousel`, `Calendar`, `Select`, `Combobox`, `Command`, `DatePicker`, `DataTable`, `Chart`

### App-shell and feedback PRs

`Sidebar`, `Sonner`, `Toast`

### Final PR

Rebuild the current `ui.shadcn.com` landing page using only our new `ui` package and `templ`.

## Component Catalog And Public API

Unless stated otherwise, all components accept `DOMProps`.

### Foundation

| Component | Exports | Specialized props | Notes |
| --- | --- | --- | --- |
| Button | `Button` | `Variant: default|destructive|outline|secondary|ghost|link`, `Size: default|xs|sm|lg|icon|icon-xs|icon-sm|icon-lg`, `Type`, `Disabled`, `Href`, `Leading`, `Trailing` | Use `Element` or `Href` instead of `asChild`. Match upstream variant names exactly. |
| Badge | `Badge` | `Variant: default|secondary|destructive|outline|ghost|link`, `Href` | Same visual variants as upstream. |
| Alert | `Alert`, `AlertTitle`, `AlertDescription` | `Variant: default|destructive` | Role must be `alert`. |
| AspectRatio | `AspectRatio` | `Ratio float64` | Pure layout helper. |
| Card | `Card`, `CardHeader`, `CardTitle`, `CardDescription`, `CardAction`, `CardContent`, `CardFooter` | none | Preserve upstream spacing and slots. |
| Separator | `Separator` | `Orientation: horizontal|vertical`, `Decorative bool` | Can be fully native. |
| Skeleton | `Skeleton` | none | Single wrapper div with upstream class contract. |
| Spinner | `Spinner` | `Size`, `Label` | Ship as SVG/CSS spinner, not an icon package dependency. |
| Kbd | `Kbd` | `Size`, `Inset bool` | Match current simple kbd styling. |
| Typography | `H1`, `H2`, `H3`, `H4`, `P`, `Blockquote`, `InlineCode`, `Lead`, `Large`, `Small`, `Muted`, `List`, `TableProse` | minimal | This is a docs convenience layer, not a primitive. Keep it lightweight. |
| Label | `Label` | `For` | Native `<label>` wrapper. |

### Form And Inputs

| Component | Exports | Specialized props | Notes |
| --- | --- | --- | --- |
| Input | `Input` | `Type`, `Name`, `Value`, `Placeholder`, `Disabled`, `Required`, `Invalid` | Native input with upstream classes. |
| Textarea | `Textarea` | `Name`, `Value`, `Placeholder`, `Rows`, `Disabled`, `Required`, `Invalid` | Native textarea. |
| Checkbox | `Checkbox` | `Name`, `Checked`, `DefaultChecked`, `Disabled`, `Required`, `Value` | Must emit hidden input when needed for forms. Support indeterminate state later if necessary. |
| Switch | `Switch` | `Size: default|sm`, `Name`, `Checked`, `DefaultChecked`, `Disabled`, `Required`, `Value` | Model as checkbox-like control with custom UI. |
| RadioGroup | `RadioGroup`, `RadioGroupItem` | root: `Name`, `Value`, `DefaultValue`, `Disabled`, `Required`; item: `Value`, `Disabled` | Arrow-key navigation required. |
| NativeSelect | `NativeSelect`, `NativeSelectOption`, `NativeSelectOptGroup` | `Size: default|sm`, native select props | Exact shadcn styles, native semantics. |
| Field | `FieldSet`, `FieldLegend`, `FieldGroup`, `Field`, `FieldContent`, `FieldLabel`, `FieldTitle`, `FieldDescription`, `FieldSeparator`, `FieldError` | `Orientation: vertical|horizontal|responsive`, `Invalid`, `Disabled`, `Errors []string` | This replaces most ad-hoc form markup. |
| Form | `Form`, `FormItem`, `FormLabel`, `FormControl`, `FormDescription`, `FormMessage` | `Name`, `ID`, `DescribedBy`, `Invalid`, `Message` | Do not port React Hook Form internals. Make this a thin accessibility helper for server-rendered forms. |
| ButtonGroup | `ButtonGroup`, `ButtonGroupText`, `ButtonGroupSeparator` | `Orientation: horizontal|vertical` | Supports buttons, selects, and inputs. |
| InputGroup | `InputGroup`, `InputGroupAddon`, `InputGroupButton`, `InputGroupText`, `InputGroupInput`, `InputGroupTextarea` | addon `Align: inline-start|inline-end|block-start|block-end`; button `Size: xs|sm|icon-xs|icon-sm` | Important for parity with newer docs. |
| InputOTP | `InputOTP`, `InputOTPGroup`, `InputOTPSlot`, `InputOTPSeparator` | `Name`, `Value`, `DefaultValue`, `MaxLength`, `Pattern`, `Disabled` | Build our own OTP runtime instead of `input-otp`. |
| Slider | `Slider` | `Name`, `Value []float64`, `DefaultValue []float64`, `Min`, `Max`, `Step`, `Disabled`, `Orientation` | Range and multi-thumb support required. |
| Progress | `Progress` | `Value`, `Max` | Native semantics preferred. |

### Disclosure And Overlays

| Component | Exports | Specialized props | Notes |
| --- | --- | --- | --- |
| Accordion | `Accordion`, `AccordionItem`, `AccordionTrigger`, `AccordionContent` | root: `Type: single|multiple`, `Value []string`, `DefaultValue []string`, `Collapsible`; item: `Value`, `Disabled` | Use disclosure + height animation. |
| Collapsible | `Collapsible`, `CollapsibleTrigger`, `CollapsibleContent` | `Open`, `DefaultOpen`, `Disabled` | Lighter than accordion. |
| Dialog | `Dialog`, `DialogTrigger`, `DialogPortal`, `DialogOverlay`, `DialogContent`, `DialogHeader`, `DialogFooter`, `DialogTitle`, `DialogDescription`, `DialogClose` | `Open`, `DefaultOpen`, `Modal`, `ShowCloseButton` | Use native `<dialog>` if it does not break composition; otherwise custom dialog runtime. |
| AlertDialog | `AlertDialog`, `AlertDialogTrigger`, `AlertDialogPortal`, `AlertDialogOverlay`, `AlertDialogContent`, `AlertDialogHeader`, `AlertDialogFooter`, `AlertDialogTitle`, `AlertDialogDescription`, `AlertDialogAction`, `AlertDialogCancel` | same as `Dialog` | Confirm/cancel semantics and destructive styling. |
| Sheet | `Sheet`, `SheetTrigger`, `SheetPortal`, `SheetOverlay`, `SheetContent`, `SheetHeader`, `SheetFooter`, `SheetTitle`, `SheetDescription`, `SheetClose` | `Open`, `DefaultOpen`, `Side: top|right|bottom|left`, `ShowCloseButton` | Slide-over variant of dialog. |
| Drawer | `Drawer`, `DrawerTrigger`, `DrawerPortal`, `DrawerOverlay`, `DrawerContent`, `DrawerHeader`, `DrawerFooter`, `DrawerTitle`, `DrawerDescription`, `DrawerClose` | `Open`, `DefaultOpen`, `Side`, `Dismissible`, `SnapPoints []string`, `DefaultSnapPoint` | Vaul behavior should be approximated with our runtime. |
| Popover | `Popover`, `PopoverTrigger`, `PopoverAnchor`, `PopoverContent`, `PopoverHeader`, `PopoverTitle`, `PopoverDescription` | `Open`, `DefaultOpen`, `Side`, `Align`, `SideOffset`, `AlignOffset`, `Modal` | Use `@floating-ui/dom`. |
| HoverCard | `HoverCard`, `HoverCardTrigger`, `HoverCardContent` | `OpenDelayMs`, `CloseDelayMs`, `Side`, `Align`, `SideOffset` | Hover/focus driven popover. |
| Tooltip | `TooltipProvider`, `Tooltip`, `TooltipTrigger`, `TooltipContent` | provider `DelayDuration`; content `Side`, `Align`, `SideOffset` | Use Floating UI and arrow rendering. |

### Menus And Navigation

| Component | Exports | Specialized props | Notes |
| --- | --- | --- | --- |
| Breadcrumb | `Breadcrumb`, `BreadcrumbList`, `BreadcrumbItem`, `BreadcrumbLink`, `BreadcrumbPage`, `BreadcrumbSeparator`, `BreadcrumbEllipsis` | link `Href`; separator custom child optional | Straightforward native markup. |
| Pagination | `Pagination`, `PaginationContent`, `PaginationItem`, `PaginationLink`, `PaginationPrevious`, `PaginationNext`, `PaginationEllipsis` | link `IsActive`, `Href`, `Size` | Composition over data source. |
| Tabs | `Tabs`, `TabsList`, `TabsTrigger`, `TabsContent` | root `Value`, `DefaultValue`, `Orientation`; list `Variant: default|line` | Keyboard navigation required. |
| Toggle | `Toggle` | `Variant: default|outline`, `Size: default|sm|lg`, `Pressed`, `DefaultPressed`, `Disabled`, `Value` | Button-like pressed state. |
| ToggleGroup | `ToggleGroup`, `ToggleGroupItem` | root `Type: single|multiple`, `Value []string`, `DefaultValue []string`, `Variant`, `Size`, `Spacing` | Reuse toggle variants. |
| DropdownMenu | `DropdownMenu`, `DropdownMenuTrigger`, `DropdownMenuPortal`, `DropdownMenuContent`, `DropdownMenuGroup`, `DropdownMenuItem`, `DropdownMenuCheckboxItem`, `DropdownMenuRadioGroup`, `DropdownMenuRadioItem`, `DropdownMenuLabel`, `DropdownMenuSeparator`, `DropdownMenuShortcut`, `DropdownMenuSub`, `DropdownMenuSubTrigger`, `DropdownMenuSubContent` | item `Inset`, `Variant: default|destructive`; content `SideOffset` | This is one of the core runtime-heavy components. |
| ContextMenu | `ContextMenu`, `ContextMenuTrigger`, `ContextMenuPortal`, `ContextMenuContent`, `ContextMenuItem`, `ContextMenuCheckboxItem`, `ContextMenuRadioItem`, `ContextMenuLabel`, `ContextMenuSeparator`, `ContextMenuShortcut`, `ContextMenuGroup`, `ContextMenuSub`, `ContextMenuSubContent`, `ContextMenuSubTrigger`, `ContextMenuRadioGroup` | same menu props as dropdown | Triggered on right click / keyboard context menu. |
| Menubar | `Menubar`, `MenubarMenu`, `MenubarTrigger`, `MenubarPortal`, `MenubarContent`, `MenubarGroup`, `MenubarItem`, `MenubarCheckboxItem`, `MenubarRadioGroup`, `MenubarRadioItem`, `MenubarLabel`, `MenubarSeparator`, `MenubarShortcut`, `MenubarSub`, `MenubarSubTrigger`, `MenubarSubContent` | same menu props | Horizontal application menu behavior. |
| NavigationMenu | `NavigationMenu`, `NavigationMenuList`, `NavigationMenuItem`, `NavigationMenuTrigger`, `NavigationMenuContent`, `NavigationMenuLink`, `NavigationMenuIndicator`, `NavigationMenuViewport` | `Value`, `DefaultValue`, `Orientation`, `DelayDuration` | Keep upstream trigger styling and viewport motion. |
| Direction | `DirectionProvider` | `Direction: ltr|rtl` | Small wrapper for RTL support. |

### Data, Lists, And Composite Controls

| Component | Exports | Specialized props | Notes |
| --- | --- | --- | --- |
| Avatar | `Avatar`, `AvatarImage`, `AvatarFallback`, `AvatarBadge`, `AvatarGroup`, `AvatarGroupCount` | `Size: sm|default|lg` | No runtime needed beyond image fallback. |
| Empty | `Empty`, `EmptyHeader`, `EmptyMedia`, `EmptyTitle`, `EmptyDescription`, `EmptyContent` | media `Variant: default|icon` | Newer shadcn helper component. |
| Item | `ItemGroup`, `ItemSeparator`, `Item`, `ItemMedia`, `ItemContent`, `ItemTitle`, `ItemDescription`, `ItemActions`, `ItemHeader`, `ItemFooter` | item `Variant: default|outline|muted`, `Size: default|sm`; media `Variant: default|icon|image` | Very useful for lists and cards. |
| Table | `Table`, `TableHeader`, `TableBody`, `TableFooter`, `TableHead`, `TableRow`, `TableCell`, `TableCaption` | native table attrs | Keep it simple and identical to shadcn styles. |
| DataTable | `DataTable`, `DataTableToolbar`, `DataTableFilters`, `DataTablePagination`, `DataTableColumnHeader` | `Columns`, `Rows`, `Page`, `PageSize`, `Sort`, `Filters`, `Empty` | This cannot be TanStack React parity. Treat it as a Go-native server-rendered table pattern built from `Table`, `Input`, `Button`, `DropdownMenu`, `Pagination`. |
| ScrollArea | `ScrollArea`, `ScrollBar` | `Orientation` on scroll bar | Can be mostly CSS/native, but custom scrollbar styling required. |
| Resizable | `ResizablePanelGroup`, `ResizablePanel`, `ResizableHandle` | group `Direction`; panel `DefaultSize`, `MinSize`, `MaxSize`, `Collapsible`; handle `WithHandle` | Implement with pointer events, not React dependency. |
| Carousel | `Carousel`, `CarouselContent`, `CarouselItem`, `CarouselPrevious`, `CarouselNext` | root `Orientation`, `Loop`, `Align`, `StartIndex`; nav buttons reuse `Button` props | Use Embla core or a tiny snap-based runtime. Do not use Embla React. |
| Calendar | `Calendar`, `CalendarDayButton` | `Mode: single|multiple|range`, `Selected`, `Range`, `Month`, `DefaultMonth`, `ShowOutsideDays`, `ShowWeekNumber`, `CaptionLayout`, `ButtonVariant`, `Locale`, `DisabledDates` | Do not use `react-day-picker`. Build the month grid in Go and enhance keyboard interaction with JS. |
| DatePicker | `DatePicker`, `DateRangePicker` | `Mode`, `Name`, `Value`, `Range`, `Placeholder`, `Open`, `DefaultOpen` | Composite of `Button`, `Popover`, and `Calendar`. |
| Select | `Select`, `SelectGroup`, `SelectValue`, `SelectTrigger`, `SelectContent`, `SelectLabel`, `SelectItem`, `SelectSeparator`, `SelectScrollUpButton`, `SelectScrollDownButton` | root `Name`, `Value`, `DefaultValue`, `Open`, `DefaultOpen`, `Disabled`, `Required`; trigger `Size: default|sm`; content `Position: item-aligned|popper`, `Align` | Requires listbox runtime and hidden input. |
| Combobox | `Combobox`, `ComboboxValue`, `ComboboxTrigger`, `ComboboxClear`, `ComboboxInput`, `ComboboxContent`, `ComboboxList`, `ComboboxItem`, `ComboboxGroup`, `ComboboxLabel`, `ComboboxCollection`, `ComboboxEmpty`, `ComboboxSeparator`, `ComboboxChips`, `ComboboxChip`, `ComboboxChipsInput` | `Name`, `Value`, `Values`, `DefaultValue`, `Multiple`, `Placeholder`, `ShowTrigger`, `ShowClear`, `Filter`, `Open`, `DefaultOpen`, `AnchorID`, `Side`, `Align`, `SideOffset` | This is our replacement for Base UI combobox. Build carefully. |
| Command | `Command`, `CommandDialog`, `CommandInput`, `CommandList`, `CommandEmpty`, `CommandGroup`, `CommandItem`, `CommandShortcut`, `CommandSeparator` | `Value`, `DefaultValue`, `Open`, `Placeholder`, `Title`, `Description`, `ShowCloseButton` | This replaces `cmdk` with a small searchable list runtime. |
| Chart | `ChartContainer`, `ChartStyle`, `ChartTooltip`, `ChartTooltipContent`, `ChartLegend`, `ChartLegendContent` | `Config`, `Library`, `Series`, `InitialWidth`, `InitialHeight` | This is not strict Recharts parity. Implement the shadcn chart shell and tooltip/legend styling around a pluggable vanilla chart adapter. |

### App Shell And Feedback

| Component | Exports | Specialized props | Notes |
| --- | --- | --- | --- |
| Sidebar | `SidebarProvider`, `Sidebar`, `SidebarTrigger`, `SidebarRail`, `SidebarInset`, `SidebarInput`, `SidebarHeader`, `SidebarFooter`, `SidebarSeparator`, `SidebarContent`, `SidebarGroup`, `SidebarGroupLabel`, `SidebarGroupAction`, `SidebarGroupContent`, `SidebarMenu`, `SidebarMenuItem`, `SidebarMenuButton`, `SidebarMenuAction`, `SidebarMenuBadge`, `SidebarMenuSkeleton`, `SidebarMenuSub`, `SidebarMenuSubItem`, `SidebarMenuSubButton` | provider `DefaultOpen`, `Open`; sidebar `Side`, `Variant: sidebar|floating|inset`, `Collapsible: offcanvas|icon|none`; menu button `IsActive`, `Size` | Sidebar gets its own runtime, tokens, cookie persistence, and keyboard shortcut. Do not expose a React-style hook API. |
| Sonner | `Toaster`, `Toast` helpers | `Theme`, `Position`, `RichColors`, `Duration` | Do not use React Sonner. Build a small toast store and renderer, but keep the shadcn visual defaults. |
| Toast | `ToastProvider`, `ToastViewport`, `Toast`, `ToastTitle`, `ToastDescription`, `ToastAction`, `ToastClose` | `Open`, `DefaultOpen`, `Variant`, `Duration` | Legacy compatibility. Implement after `Sonner`, not before. |

## Variant Names That Must Match Upstream

These should not be renamed:

- `Button`: `default`, `destructive`, `outline`, `secondary`, `ghost`, `link`
- `Button` sizes: `default`, `xs`, `sm`, `lg`, `icon`, `icon-xs`, `icon-sm`, `icon-lg`
- `Badge`: `default`, `secondary`, `destructive`, `outline`, `ghost`, `link`
- `Alert`: `default`, `destructive`
- `Toggle`: `default`, `outline`
- `TabsList`: `default`, `line`
- `Field`: `vertical`, `horizontal`, `responsive`
- `Item`: `default`, `outline`, `muted`
- `ItemMedia`: `default`, `icon`, `image`
- `EmptyMedia`: `default`, `icon`
- `Sidebar`: `sidebar`, `floating`, `inset`
- `Sidebar` collapsible: `offcanvas`, `icon`, `none`

## What Counts As Parity

### Strict parity targets

- visual appearance
- spacing
- radius
- dark mode
- data-slot names
- variant names
- keyboard support
- basic composition surface

### Adapted parity targets

These cannot be React API clones and should be explicitly treated as Go-native adaptations:

- `Combobox`
- `Command`
- `Calendar`
- `DatePicker`
- `DataTable`
- `Chart`
- `Drawer`
- `Carousel`
- `Sonner`
- `Toast`

## Acceptance Criteria For Each PR

Every component PR should include:

1. component implementation in the `ui` package
2. one example page or preview snippet
3. dark mode verification
4. keyboard interaction verification where applicable
5. form submission verification for custom controls where applicable
6. notes on any intentional divergence from upstream

## Landing Page Recreation

This is the final milestone after the component library is usable.

Target: recreate the current `ui.shadcn.com` home page in `templ`, using only our own package components.

Minimum required sections:

1. top navigation with logo, docs links, search trigger, GitHub, Twitter, theme toggle
2. hero with the current core message
3. primary CTA and secondary CTA
4. tabbed preview switcher for `Mail`, `Dashboard`, `Cards`, `Tasks`, `Playground`, `Forms`, `Music`, `Authentication`
5. main preview surface using our components, not copied HTML fragments
6. footer attribution block

Requirements:

- do not ship the landing page until the components used on the page are already implemented
- use our `Tabs`, `Button`, `Card`, `Input`, `Avatar`, `Sidebar`, `Table`, `Badge`, `DropdownMenu`, and other package components wherever relevant
- do not hardcode one-off marketing-only styles when a reusable component should exist

## Final Developer Experience Target

After this work, a Go developer familiar with shadcn/ui should be able to:

```go
import "your-module/ui"

templ Example() {
  @ui.Card(ui.DOMProps{}) {
    @ui.CardHeader(ui.DOMProps{}) {
      @ui.CardTitle(ui.DOMProps{}) { <span>Account</span> }
      @ui.CardDescription(ui.DOMProps{}) { <span>Manage your settings.</span> }
    }
    @ui.CardContent(ui.DOMProps{}) {
      @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline}) {
        <span>Save</span>
      }
    }
  }
}
```

That experience does not need to look exactly like React JSX, but it should feel recognizably shadcn:

- same names
- same variants
- same tokens
- same compositional structure
- minimal surprises for someone who already knows the original library
