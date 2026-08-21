# templcn Component Matrix & API Reference

All components are rendered as `templ.Component` functions taking Go configuration structs and optionally accepting children inside `{ ... }`.

Every component struct embeds `ui.DOMProps`:
```go
type DOMProps struct {
    ID      string
    Class   string
    Element string           // override default HTML tag (e.g. "a", "button", "div")
    Attrs   templ.Attributes // arbitrary HTML / HTMX / data attributes
}
```

---

## Catalog

### 1. Accordion
- **Install**: `templcn add accordion`
- **Runtime**: Required (`assets/runtime.js`)
- **Key Types**:
  - `AccordionProps`: `Type` ("single" | "multiple"), `Collapsible` bool, `Value` string, `DefaultValue` string
  - `AccordionItemProps`: `Value` string, `Disabled` bool
- **Exported Functions**: `Accordion`, `AccordionItem`, `AccordionTrigger`, `AccordionContent`
- **Example**:
  ```templ
  @ui.Accordion(ui.AccordionProps{Type: "single", Collapsible: true}) {
      @ui.AccordionItem(ui.AccordionItemProps{Value: "item-1"}) {
          @ui.AccordionTrigger(ui.DOMProps{}) { Heading 1 }
          @ui.AccordionContent(ui.DOMProps{}) { Content 1 }
      }
  }
  ```

---

### 2. Alert & Alert Dialog
- **Install**: `templcn add alert alert-dialog`
- **Key Types**:
  - `AlertProps`: `Variant` (`AlertVariantDefault`, `AlertVariantDestructive`)
  - `AlertDialogProps`: `Open` bool, `DefaultOpen` bool
- **Exported Functions**:
  - Alert: `Alert`, `AlertTitle`, `AlertDescription`
  - Alert Dialog: `AlertDialog`, `AlertDialogTrigger`, `AlertDialogContent`, `AlertDialogHeader`, `AlertDialogFooter`, `AlertDialogTitle`, `AlertDialogDescription`, `AlertDialogAction`, `AlertDialogCancel`

---

### 3. Avatar
- **Install**: `templcn add avatar`
- **Key Types**:
  - `AvatarProps`: `DOMProps`
  - `AvatarImageProps`: `Src` string, `Alt` string
  - `AvatarFallbackProps`: `DOMProps`
- **Exported Functions**: `Avatar`, `AvatarImage`, `AvatarFallback`
- **Example**:
  ```templ
  @ui.Avatar(ui.DOMProps{Class: "size-10"}) {
      @ui.AvatarImage(ui.AvatarImageProps{Src: "https://github.com/shadcn.png", Alt: "@shadcn"})
      @ui.AvatarFallback(ui.DOMProps{}) { CN }
  }
  ```

---

### 4. Badge
- **Install**: `templcn add badge`
- **Key Types**:
  - `BadgeVariant`: `BadgeVariantDefault`, `BadgeVariantSecondary`, `BadgeVariantDestructive`, `BadgeVariantOutline`, `BadgeVariantGhost`
  - `BadgeProps`: `DOMProps`, `Label` string, `Variant` BadgeVariant
- **Exported Functions**: `Badge`

---

### 5. Button & Button Group
- **Install**: `templcn add button button-group`
- **Key Types**:
  - `ButtonVariant`: `ButtonVariantDefault`, `ButtonVariantDestructive`, `ButtonVariantOutline`, `ButtonVariantSecondary`, `ButtonVariantGhost`, `ButtonVariantLink`
  - `ButtonSize`: `ButtonSizeDefault`, `ButtonSizeXS`, `ButtonSizeSM`, `ButtonSizeLG`, `ButtonSizeIcon`, `ButtonSizeIconXS`, `ButtonSizeIconSM`, `ButtonSizeIconLG`
  - `ButtonProps`: `DOMProps`, `Label` string, `Variant` ButtonVariant, `Size` ButtonSize, `Type` string, `Disabled` bool, `Href` string, `Leading` templ.Component, `Trailing` templ.Component
  - `ButtonGroupProps`: `DOMProps`, `Orientation` string ("horizontal" | "vertical")
- **Exported Functions**: `Button`, `ButtonGroup`

---

### 6. Card
- **Install**: `templcn add card`
- **Exported Functions**: `Card`, `CardHeader`, `CardTitle`, `CardDescription`, `CardAction`, `CardContent`, `CardFooter`

---

### 7. Checkbox & Switch
- **Install**: `templcn add checkbox switch`
- **Runtime**: Required for Switch/Checkbox toggle events
- **Key Types**:
  - `CheckboxProps`: `DOMProps`, `Checked` bool, `DefaultChecked` bool, `Required` bool, `Disabled` bool, `Name` string, `Value` string
  - `SwitchProps`: `DOMProps`, `Checked` bool, `DefaultChecked` bool, `Disabled` bool, `Name` string, `Value` string
- **Exported Functions**: `Checkbox`, `Switch`

---

### 8. Collapsible
- **Install**: `templcn add collapsible`
- **Exported Functions**: `Collapsible`, `CollapsibleTrigger`, `CollapsibleContent`

---

### 9. Dialog, Drawer & Sheet
- **Install**: `templcn add dialog drawer sheet`
- **Runtime**: Required
- **Exported Functions**:
  - Dialog: `Dialog`, `DialogTrigger`, `DialogContent`, `DialogHeader`, `DialogFooter`, `DialogTitle`, `DialogDescription`, `DialogClose`
  - Sheet: `Sheet`, `SheetTrigger`, `SheetContent` (Side: "top" | "right" | "bottom" | "left"), `SheetHeader`, `SheetFooter`, `SheetTitle`, `SheetDescription`
  - Drawer: `Drawer`, `DrawerTrigger`, `DrawerContent`, `DrawerHeader`, `DrawerFooter`, `DrawerTitle`, `DrawerDescription`

---

### 10. Dropdown Menu, Context Menu & Menubar
- **Install**: `templcn add dropdown-menu context-menu menubar`
- **Runtime**: Required (`@floating-ui/dom` bundled in `assets/runtime.js`)
- **Exported Functions**:
  - `DropdownMenu`, `DropdownMenuTrigger`, `DropdownMenuContent`, `DropdownMenuItem`, `DropdownMenuLabel`, `DropdownMenuSeparator`, `DropdownMenuShortcut`, `DropdownMenuGroup`, `DropdownMenuSub`, `DropdownMenuSubTrigger`, `DropdownMenuSubContent`, `DropdownMenuCheckboxItem`, `DropdownMenuRadioGroup`, `DropdownMenuRadioItem`

---

### 11. Form, Field & Inputs
- **Install**: `templcn add form field input input-group input-otp textarea native-select`
- **Key Types**:
  - `InputProps`: `DOMProps`, `Type` string, `Name` string, `Value` string, `Placeholder` string, `Disabled` bool, `Required` bool, `ReadOnly` bool
  - `TextareaProps`: `DOMProps`, `Name` string, `Value` string, `Placeholder` string, `Rows` int
  - `InputOTPProps`: `MaxLength` int, `Pattern` string
  - `FieldProps`: `Orientation` ("vertical" | "horizontal"), `Invalid` bool
- **Exported Functions**: `Input`, `InputGroup`, `InputGroupAddon`, `InputOTP`, `InputOTPSlot`, `InputOTPSeparator`, `Textarea`, `Label`, `Field`, `FieldLabel`, `FieldDescription`, `FieldError`, `Form`

---

### 12. Select & Combobox
- **Install**: `templcn add select combobox command`
- **Runtime**: Required
- **Exported Functions**:
  - Select: `Select`, `SelectTrigger`, `SelectValue`, `SelectContent`, `SelectGroup`, `SelectLabel`, `SelectItem`, `SelectSeparator`
  - Combobox: `Combobox`, `ComboboxTrigger`, `ComboboxContent`, `ComboboxInput`, `ComboboxList`, `ComboboxEmpty`, `ComboboxGroup`, `ComboboxItem`

---

### 13. Popover, Tooltip & Hover Card
- **Install**: `templcn add popover tooltip hover-card`
- **Runtime**: Required
- **Exported Functions**:
  - Popover: `Popover`, `PopoverTrigger`, `PopoverContent`, `PopoverAnchor`
  - Tooltip: `TooltipProvider`, `Tooltip`, `TooltipTrigger`, `TooltipContent`
  - Hover Card: `HoverCard`, `HoverCardTrigger`, `HoverCardContent`

---

### 14. Tabs
- **Install**: `templcn add tabs`
- **Runtime**: Required
- **Key Types**:
  - `TabsProps`: `Value` string, `DefaultValue` string, `Orientation` ("horizontal" | "vertical")
  - `TabsTriggerProps`: `Value` string, `Disabled` bool
  - `TabsContentProps`: `Value` string
- **Exported Functions**: `Tabs`, `TabsList`, `TabsTrigger`, `TabsContent`

---

### 15. Table & Data Table
- **Install**: `templcn add table data-table`
- **Exported Functions**: `Table`, `TableHeader`, `TableBody`, `TableFooter`, `TableHead`, `TableRow`, `TableCell`, `TableCaption`, `DataTable`

---

### 16. Sonner & Toast
- **Install**: `templcn add sonner toast`
- **Runtime**: Required
- **Exported Functions**: `Sonner`, `Toaster`, `Toast`, `ToastTitle`, `ToastDescription`, `ToastAction`, `ToastClose`

---

### 17. Chart
- **Install**: `templcn add chart`
- **Runtime**: Required (`assets/tanstack-runtime.js` or chart runtime)
- **Supported Types**: Area, Bar, Line, Pie, Radial, Scatter with Tooltip, Legend, Cartesian Grid, and Axes.
- **Exported Functions**: `ChartContainer`, `ChartTooltip`, `ChartTooltipContent`, `ChartLegend`, `ChartLegendContent`
