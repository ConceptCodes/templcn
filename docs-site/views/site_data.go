package views

import (
	"fmt"
	"strings"

	"github.com/a-h/templ"
)

type APIProp struct {
	Name    string
	Type    string
	Default string
	Desc    string
}

type ExampleEntry struct {
	Name    string
	Desc    string
	GoCode  string
	Preview templ.Component
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
	{
		Slug: "accordion", Title: "Accordion", Description: "A vertically stacked set of interactive collapsible panels.",
		Install: "shadcn add button", ManualInstall: "Copy ui/overlays.go into your project.",
		GoUsage: `@ui.Accordion(ui.AccordionProps{Type: "single", Collapsible: true}) {
  @ui.AccordionItem(ui.AccordionItemProps{Value: "item-1"}) {
    @ui.AccordionTrigger(ui.DOMProps{}) { Is it accessible? }
    @ui.AccordionContent(ui.DOMProps{}) { Yes. It adheres to the WAI-ARIA design pattern. }
  }
}`,
		APIProps: []APIProp{
			{Name: "Type", Type: "string", Default: `"single"`, Desc: `"single" allows one panel open at a time; "multiple" allows many.`},
			{Name: "Collapsible", Type: "bool", Default: "false", Desc: "When type is single, allows the open item to be collapsed."},
			{Name: "Value", Type: "[]string", Default: "nil", Desc: "The controlled open state."},
			{Name: "DefaultValue", Type: "[]string", Default: "nil", Desc: "The uncontrolled default open state."},
		},
		Examples: []ExampleEntry{
			{Name: "Single (collapsible)", Desc: "One item open at a time, with collapse allowed.", GoCode: `@ui.Accordion(ui.AccordionProps{Type: "single", Collapsible: true}) {
  @ui.AccordionItem(ui.AccordionItemProps{Value: "item-1"}) {
    @ui.AccordionTrigger(ui.DOMProps{}) { Is it accessible? }
    @ui.AccordionContent(ui.DOMProps{}) { Yes. It adheres to the WAI-ARIA design pattern. }
  }
}`, Preview: AccordionPreview()},
			{Name: "Multiple", Desc: "Multiple items can be open simultaneously.", GoCode: `@ui.Accordion(ui.AccordionProps{Type: "multiple"}) {
  @ui.AccordionItem(ui.AccordionItemProps{Value: "a"}) {
    @ui.AccordionTrigger(ui.DOMProps{}) { Item A }
    @ui.AccordionContent(ui.DOMProps{}) { Content for item A. }
  }
}`, Preview: AccordionMultiplePreview()},
		},
	},
	{
		Slug: "alert", Title: "Alert", Description: "Displays a prominent message to call attention to important information.",
		Install: "shadcn add button", ManualInstall: "Copy ui/foundation.go into your project.",
		GoUsage: `@ui.Alert(ui.AlertProps{}) {
  @ui.AlertTitle(ui.DOMProps{}) { Heads up! }
  @ui.AlertDescription(ui.DOMProps{}) { You can add components using the CLI. }
}`,
		APIProps: []APIProp{
			{Name: "Variant", Type: "AlertVariant", Default: `"default"`, Desc: `"default" or "destructive".`},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "A default informational alert.", GoCode: `@ui.Alert(ui.AlertProps{}) {
  @ui.AlertTitle(ui.DOMProps{}) { Heads up! }
  @ui.AlertDescription(ui.DOMProps{}) { You can add components using the CLI. }
}`, Preview: AlertDefaultPreview()},
			{Name: "Destructive", Desc: "A destructive alert for errors.", GoCode: `@ui.Alert(ui.AlertProps{Variant: ui.AlertVariantDestructive}) {
  @ui.AlertTitle(ui.DOMProps{}) { Error }
  @ui.AlertDescription(ui.DOMProps{}) { Your session has expired. Please log in again. }
}`, Preview: AlertDestructivePreview()},
		},
	},
	{
		Slug: "alert-dialog", Title: "Alert Dialog", Description: "A modal dialog that asks the user to confirm a destructive action.",
		Install: "shadcn add button", ManualInstall: "Copy ui/overlays.go into your project.",
		GoUsage: `@ui.AlertDialog(ui.AlertDialogProps{Open: true}) {
  @ui.AlertDialogContent(ui.DOMProps{}) {
    @ui.AlertDialogHeader(ui.DOMProps{}) {
      @ui.AlertDialogTitle(ui.DOMProps{}) { Are you sure? }
      @ui.AlertDialogDescription(ui.DOMProps{}) { This action cannot be undone. }
    }
    @ui.AlertDialogFooter(ui.DOMProps{}) {
      @ui.AlertDialogCancel(ui.DOMProps{}) { Cancel }
      @ui.AlertDialogAction(ui.DOMProps{}) { Confirm }
    }
  }
}`,
		APIProps: []APIProp{
			{Name: "Open", Type: "bool", Default: "false", Desc: "Controls the open state of the dialog."},
			{Name: "Modal", Type: "bool", Default: "true", Desc: "Whether clicks outside close the dialog."},
		},
		Examples: []ExampleEntry{
			{Name: "Code", Desc: "AlertDialog usage — requires JS for open/close toggling.", GoCode: `@ui.AlertDialog(ui.AlertDialogProps{Open: true}) {
  @ui.AlertDialogContent(ui.DOMProps{}) {
    @ui.AlertDialogTitle(ui.DOMProps{}) { Are you sure? }
    @ui.AlertDialogAction(ui.DOMProps{}) { Continue }
    @ui.AlertDialogCancel(ui.DOMProps{}) { Cancel }
  }
}`},
		},
	},
	{
		Slug: "aspect-ratio", Title: "Aspect Ratio", Description: "Maintains a consistent width-to-height ratio for any content.",
		Install: "shadcn add button", ManualInstall: "Copy ui/foundation.go into your project.",
		GoUsage: `@ui.AspectRatio(ui.AspectRatioProps{Ratio: 16.0/9.0}) {
  <img src="..." alt="..." class="h-full w-full object-cover" />
}`,
		APIProps: []APIProp{
			{Name: "Ratio", Type: "float64", Default: "1", Desc: "The width-to-height ratio (e.g. 16.0/9.0)."},
		},
		Examples: []ExampleEntry{
			{Name: "16:9", Desc: "A 16:9 aspect ratio container.", GoCode: `@ui.AspectRatio(ui.AspectRatioProps{Ratio: 16.0/9.0, DOMProps: ui.DOMProps{Class: "w-full max-w-xs rounded-md overflow-hidden bg-muted"}}) {
  <div class="flex h-full items-center justify-center text-sm">16 : 9</div>
}`, Preview: AspectRatioPreview()},
		},
	},
	{
		Slug: "badge", Title: "Badge", Description: "Small status or metadata labels that can appear as spans or links.",
		Install: "shadcn add button", ManualInstall: "Copy ui/button.go (badge.go) into your project.",
		GoUsage: `@ui.Badge(ui.BadgeProps{Label: "Default"})
@ui.Badge(ui.BadgeProps{Label: "Secondary", Variant: ui.BadgeVariantSecondary})
@ui.Badge(ui.BadgeProps{Label: "Outline", Variant: ui.BadgeVariantOutline})`,
		APIProps: []APIProp{
			{Name: "Label", Type: "string", Default: `""`, Desc: "Text content of the badge."},
			{Name: "Variant", Type: "BadgeVariant", Default: `"default"`, Desc: `"default", "secondary", "outline", "destructive", "ghost", "link".`},
			{Name: "Href", Type: "string", Default: `""`, Desc: "When set, renders as an anchor tag."},
		},
		Examples: []ExampleEntry{
			{Name: "Variants", Desc: "All badge variants side by side.", GoCode: `@ui.Badge(ui.BadgeProps{Label: "Default"})
@ui.Badge(ui.BadgeProps{Label: "Secondary", Variant: ui.BadgeVariantSecondary})
@ui.Badge(ui.BadgeProps{Label: "Outline", Variant: ui.BadgeVariantOutline})
@ui.Badge(ui.BadgeProps{Label: "Destructive", Variant: ui.BadgeVariantDestructive})`, Preview: BadgePreview()},
		},
	},
	{
		Slug: "breadcrumb", Title: "Breadcrumb", Description: "A navigation trail indicating the current page location.",
		Install: "shadcn add button", ManualInstall: "Copy ui/navigation.go into your project.",
		GoUsage: `@ui.Breadcrumb(ui.DOMProps{}) {
  @ui.BreadcrumbList(ui.DOMProps{}) {
    @ui.BreadcrumbItem(ui.DOMProps{}) {
      @ui.BreadcrumbLink(ui.BreadcrumbLinkProps{Href: "/"}) { Home }
    }
    @ui.BreadcrumbSeparator(ui.DOMProps{}) { / }
    @ui.BreadcrumbItem(ui.DOMProps{}) {
      @ui.BreadcrumbPage(ui.DOMProps{}) { Components }
    }
  }
}`,
		APIProps: []APIProp{
			{Name: "BreadcrumbLink.Href", Type: "string", Default: `""`, Desc: "The navigation target."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "A three-level breadcrumb.", GoCode: `@ui.Breadcrumb(ui.DOMProps{}) {
  @ui.BreadcrumbList(ui.DOMProps{}) {
    @ui.BreadcrumbItem(ui.DOMProps{}) {
      @ui.BreadcrumbLink(ui.BreadcrumbLinkProps{Href: "/"}) { Home }
    }
    @ui.BreadcrumbSeparator(ui.DOMProps{}) { / }
    @ui.BreadcrumbItem(ui.DOMProps{}) {
      @ui.BreadcrumbPage(ui.DOMProps{}) { Breadcrumb }
    }
  }
}`, Preview: BreadcrumbPreview()},
		},
	},
	{
		Slug: "button", Title: "Button", Description: "Displays a button or a link styled as a button.",
		Install: "shadcn add button", ManualInstall: "Copy ui/button.go into your project.",
		GoUsage: `@ui.Button(ui.ButtonProps{Label: "Click me"})
@ui.Button(ui.ButtonProps{Label: "Outline", Variant: ui.ButtonVariantOutline})
@ui.Button(ui.ButtonProps{Label: "Open", Href: "/page"})`,
		APIProps: []APIProp{
			{Name: "Label", Type: "string", Default: `""`, Desc: "Text label rendered inside the button."},
			{Name: "Variant", Type: "ButtonVariant", Default: `"default"`, Desc: `"default", "secondary", "outline", "ghost", "link", "destructive".`},
			{Name: "Size", Type: "ButtonSize", Default: `"default"`, Desc: `"xs", "sm", "default", "lg", "icon", "icon-xs", "icon-sm", "icon-lg".`},
			{Name: "Href", Type: "string", Default: `""`, Desc: "When set, renders as an anchor element."},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "Disables the button."},
			{Name: "Type", Type: "string", Default: `"button"`, Desc: `HTML type attribute: "button", "submit", "reset".`},
		},
		Examples: []ExampleEntry{
			{Name: "Variants", Desc: "All button variants.", GoCode: `@ui.Button(ui.ButtonProps{Label: "Default"})
@ui.Button(ui.ButtonProps{Label: "Secondary", Variant: ui.ButtonVariantSecondary})
@ui.Button(ui.ButtonProps{Label: "Outline", Variant: ui.ButtonVariantOutline})
@ui.Button(ui.ButtonProps{Label: "Ghost", Variant: ui.ButtonVariantGhost})
@ui.Button(ui.ButtonProps{Label: "Link", Variant: ui.ButtonVariantLink})
@ui.Button(ui.ButtonProps{Label: "Destructive", Variant: ui.ButtonVariantDestructive})`, Preview: ButtonVariantsPreview()},
			{Name: "Sizes", Desc: "Available size options.", GoCode: `@ui.Button(ui.ButtonProps{Label: "Extra Small", Size: ui.ButtonSizeXS})
@ui.Button(ui.ButtonProps{Label: "Small", Size: ui.ButtonSizeSM})
@ui.Button(ui.ButtonProps{Label: "Default"})
@ui.Button(ui.ButtonProps{Label: "Large", Size: ui.ButtonSizeLG})`, Preview: ButtonSizesPreview()},
			{Name: "Disabled", Desc: "A disabled button.", GoCode: `@ui.Button(ui.ButtonProps{Label: "Disabled", Disabled: true})`, Preview: ButtonDisabledPreview()},
		},
	},
	{
		Slug: "button-group", Title: "Button Group", Description: "Groups buttons into a connected control surface.",
		Install: "shadcn add button", ManualInstall: "Copy ui/forms.go into your project.",
		GoUsage: `@ui.ButtonGroup(ui.ButtonGroupProps{}) {
  @ui.Button(ui.ButtonProps{Label: "Left", Variant: ui.ButtonVariantOutline})
  @ui.ButtonGroupSeparator(ui.DOMProps{})
  @ui.Button(ui.ButtonProps{Label: "Right", Variant: ui.ButtonVariantOutline})
}`,
		APIProps: []APIProp{
			{Name: "Orientation", Type: "ButtonGroupOrientation", Default: `"horizontal"`, Desc: `"horizontal" or "vertical".`},
		},
		Examples: []ExampleEntry{
			{Name: "Horizontal", Desc: "Connected horizontal buttons.", GoCode: `@ui.ButtonGroup(ui.ButtonGroupProps{}) {
  @ui.Button(ui.ButtonProps{Label: "Left", Variant: ui.ButtonVariantOutline})
  @ui.ButtonGroupSeparator(ui.DOMProps{})
  @ui.Button(ui.ButtonProps{Label: "Center", Variant: ui.ButtonVariantOutline})
  @ui.ButtonGroupSeparator(ui.DOMProps{})
  @ui.Button(ui.ButtonProps{Label: "Right", Variant: ui.ButtonVariantOutline})
}`, Preview: ButtonGroupPreview()},
		},
	},
	{
		Slug: "card", Title: "Card", Description: "A flexible surface container for related information.",
		Install: "shadcn add button", ManualInstall: "Copy ui/foundation.go into your project.",
		GoUsage: `@ui.Card(ui.DOMProps{}) {
  @ui.CardHeader(ui.DOMProps{}) {
    @ui.CardTitle(ui.DOMProps{}) { <span>Title</span> }
    @ui.CardDescription(ui.DOMProps{}) { <span>Description</span> }
  }
  @ui.CardContent(ui.DOMProps{}) { Content here }
  @ui.CardFooter(ui.DOMProps{}) { Footer actions }
}`,
		APIProps: []APIProp{
			{Name: "DOMProps.Class", Type: "string", Default: `""`, Desc: "Additional CSS classes applied to the card root."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "A card with header, content, and footer.", GoCode: `@ui.Card(ui.DOMProps{Class: "w-80"}) {
  @ui.CardHeader(ui.DOMProps{}) {
    @ui.CardTitle(ui.DOMProps{}) { <span>Create project</span> }
  }
  @ui.CardContent(ui.DOMProps{}) { ... }
  @ui.CardFooter(ui.DOMProps{}) {
    @ui.Button(ui.ButtonProps{Label: "Deploy"})
  }
}`, Preview: CardPreview()},
		},
	},
	{
		Slug: "checkbox", Title: "Checkbox", Description: "A native checkbox with shadcn styling and accessibility attributes.",
		Install: "shadcn add button", ManualInstall: "Copy ui/forms.go into your project.",
		GoUsage: `@ui.Checkbox(ui.CheckboxProps{Name: "terms"})`,
		APIProps: []APIProp{
			{Name: "Name", Type: "string", Default: `""`, Desc: "HTML name attribute for form submission."},
			{Name: "Checked", Type: "bool", Default: "false", Desc: "Whether the checkbox is checked."},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "Disables the checkbox."},
			{Name: "Required", Type: "bool", Default: "false", Desc: "Marks the checkbox as required."},
			{Name: "Invalid", Type: "bool", Default: "false", Desc: "Adds aria-invalid for error states."},
		},
		Examples: []ExampleEntry{
			{Name: "States", Desc: "Default, disabled and checked states.", GoCode: `@ui.Checkbox(ui.CheckboxProps{Name: "terms"})
@ui.Checkbox(ui.CheckboxProps{Name: "disabled", Disabled: true})
@ui.Checkbox(ui.CheckboxProps{Name: "checked", Checked: true})`, Preview: CheckboxPreview()},
		},
	},
	{
		Slug: "collapsible", Title: "Collapsible", Description: "A lightweight disclosure primitive — lighter than accordion when nesting isn't needed.",
		Install: "shadcn add button", ManualInstall: "Copy ui/overlays.go into your project.",
		GoUsage: `@ui.Collapsible(ui.CollapsibleProps{}) {
  @ui.CollapsibleTrigger(ui.DOMProps{}) { Click to expand }
  @ui.CollapsibleContent(ui.DOMProps{}) { Expanded content here }
}`,
		APIProps: []APIProp{
			{Name: "Open", Type: "bool", Default: "false", Desc: "Controls the open state."},
			{Name: "DefaultOpen", Type: "bool", Default: "false", Desc: "Default open state (uncontrolled)."},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "Disables toggling."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "A collapsible panel with trigger and content.", GoCode: `@ui.Collapsible(ui.CollapsibleProps{}) {
  @ui.CollapsibleTrigger(ui.DOMProps{}) { Click to expand }
  @ui.CollapsibleContent(ui.DOMProps{}) { Hidden content }
}`, Preview: CollapsiblePreview()},
		},
	},
	{
		Slug: "input", Title: "Input", Description: "A styled native text input supporting all HTML input types.",
		Install: "shadcn add button", ManualInstall: "Copy ui/forms.go into your project.",
		GoUsage: `@ui.Input(ui.InputProps{Placeholder: "Email"})
@ui.Input(ui.InputProps{Type: "password", Placeholder: "Password"})`,
		APIProps: []APIProp{
			{Name: "Type", Type: "string", Default: `"text"`, Desc: "HTML input type attribute."},
			{Name: "Name", Type: "string", Default: `""`, Desc: "HTML name for form submission."},
			{Name: "Value", Type: "string", Default: `""`, Desc: "Input value."},
			{Name: "Placeholder", Type: "string", Default: `""`, Desc: "Placeholder text."},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "Disables the input."},
			{Name: "Required", Type: "bool", Default: "false", Desc: "Marks the input as required."},
			{Name: "Invalid", Type: "bool", Default: "false", Desc: "Adds aria-invalid for error states."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Text, password, and disabled inputs.", GoCode: `@ui.Input(ui.InputProps{Placeholder: "Email"})
@ui.Input(ui.InputProps{Type: "password", Placeholder: "Password"})
@ui.Input(ui.InputProps{Placeholder: "Disabled", Disabled: true})`, Preview: InputPreview()},
		},
	},
	{
		Slug: "input-otp", Title: "Input OTP", Description: "One-time code entry with a slot-based UI for PIN and verification flows.",
		Install: "shadcn add button", ManualInstall: "Copy ui/forms.go into your project.",
		GoUsage: `@ui.InputOTP(ui.InputOTPProps{MaxLength: 6}) {
  @ui.InputOTPGroup(ui.DOMProps{}) {
    @ui.InputOTPSlot(ui.InputOTPSlotProps{Index: 0})
    @ui.InputOTPSlot(ui.InputOTPSlotProps{Index: 1})
    @ui.InputOTPSlot(ui.InputOTPSlotProps{Index: 2})
  }
}`,
		APIProps: []APIProp{
			{Name: "MaxLength", Type: "int", Default: "0", Desc: "Total number of OTP digits."},
			{Name: "Pattern", Type: "string", Default: `""`, Desc: "Regex pattern for allowed characters."},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "Disables all slots."},
		},
		Examples: []ExampleEntry{
			{Name: "6-digit OTP", Desc: "A 6-slot OTP input split into two groups.", GoCode: `@ui.InputOTP(ui.InputOTPProps{MaxLength: 6}) {
  @ui.InputOTPGroup(ui.DOMProps{}) {
    @ui.InputOTPSlot(ui.InputOTPSlotProps{Index: 0, Value: "1"})
    @ui.InputOTPSlot(ui.InputOTPSlotProps{Index: 1, Value: "2"})
    @ui.InputOTPSlot(ui.InputOTPSlotProps{Index: 2, Value: "3"})
  }
  @ui.InputOTPSeparator(ui.DOMProps{})
  @ui.InputOTPGroup(ui.DOMProps{}) {
    @ui.InputOTPSlot(ui.InputOTPSlotProps{Index: 3})
    @ui.InputOTPSlot(ui.InputOTPSlotProps{Index: 4})
    @ui.InputOTPSlot(ui.InputOTPSlotProps{Index: 5})
  }
}`, Preview: InputOTPPreview()},
		},
	},
	{
		Slug: "kbd", Title: "Kbd", Description: "A keyboard shortcut token styled as a key badge.",
		Install: "shadcn add button", ManualInstall: "Copy ui/foundation.go into your project.",
		GoUsage: `@ui.Kbd(ui.KbdProps{Text: "⌘K"})`,
		APIProps: []APIProp{
			{Name: "Text", Type: "string", Default: `""`, Desc: "The key text to display."},
			{Name: "Size", Type: "string", Default: `""`, Desc: `"sm", "" (default), or "lg".`},
			{Name: "Inset", Type: "bool", Default: "false", Desc: "Adds left margin for inline usage."},
		},
		Examples: []ExampleEntry{
			{Name: "Keyboard shortcuts", Desc: "Common keyboard shortcut badges.", GoCode: `@ui.Kbd(ui.KbdProps{Text: "⌘K"})
@ui.Kbd(ui.KbdProps{Text: "Ctrl+C"})
@ui.Kbd(ui.KbdProps{Text: "Enter"})`, Preview: KbdPreview()},
		},
	},
	{
		Slug: "label", Title: "Label", Description: "A native label wrapper with shadcn styling.",
		Install: "shadcn add button", ManualInstall: "Copy ui/foundation.go into your project.",
		GoUsage: `@ui.Label(ui.LabelProps{For: "email"}) { <span>Your email</span> }`,
		APIProps: []APIProp{
			{Name: "For", Type: "string", Default: `""`, Desc: "The id of the associated form element."},
		},
		Examples: []ExampleEntry{
			{Name: "With Input", Desc: "A label associated with an input field.", GoCode: `@ui.Label(ui.LabelProps{For: "email"}) { <span>Your email address</span> }
@ui.Input(ui.InputProps{ID: "email", Placeholder: "name@example.com"})`, Preview: LabelPreview()},
		},
	},
	{
		Slug: "native-select", Title: "Native Select", Description: "A styled native HTML select element.",
		Install: "shadcn add button", ManualInstall: "Copy ui/forms.go into your project.",
		GoUsage: `@ui.NativeSelect(ui.NativeSelectProps{}) {
  @ui.NativeSelectOption(ui.DOMProps{}, "apple", "Apple", false)
  @ui.NativeSelectOption(ui.DOMProps{}, "banana", "Banana", true)
}`,
		APIProps: []APIProp{
			{Name: "Name", Type: "string", Default: `""`, Desc: "HTML name for form submission."},
			{Name: "Size", Type: "string", Default: `""`, Desc: `"sm" or "" (default h-9).`},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "Disables the select."},
			{Name: "Required", Type: "bool", Default: "false", Desc: "Marks as required."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "A basic native select.", GoCode: `@ui.NativeSelect(ui.NativeSelectProps{}) {
  @ui.NativeSelectOption(ui.DOMProps{}, "apple", "Apple", false)
  @ui.NativeSelectOption(ui.DOMProps{}, "banana", "Banana", true)
  @ui.NativeSelectOption(ui.DOMProps{}, "cherry", "Cherry", false)
}`, Preview: NativeSelectPreview()},
		},
	},
	{
		Slug: "pagination", Title: "Pagination", Description: "Pagination primitives for navigating between pages of data.",
		Install: "shadcn add button", ManualInstall: "Copy ui/navigation.go into your project.",
		GoUsage: `@ui.Pagination(ui.DOMProps{}) {
  @ui.PaginationContent(ui.DOMProps{}) {
    @ui.PaginationItem(ui.DOMProps{}) {
      @ui.PaginationPrevious(ui.PaginationLinkProps{Href: "?page=1"}) { Previous }
    }
    @ui.PaginationItem(ui.DOMProps{}) {
      @ui.PaginationLink(ui.PaginationLinkProps{Href: "?page=1", IsActive: true}) { 1 }
    }
    @ui.PaginationItem(ui.DOMProps{}) {
      @ui.PaginationNext(ui.PaginationLinkProps{Href: "?page=2"}) { Next }
    }
  }
}`,
		APIProps: []APIProp{
			{Name: "PaginationLink.Href", Type: "string", Default: `""`, Desc: "The link target URL."},
			{Name: "PaginationLink.IsActive", Type: "bool", Default: "false", Desc: "Highlights the current page."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "A standard pagination control.", GoCode: `@ui.Pagination(ui.DOMProps{}) {
  @ui.PaginationContent(ui.DOMProps{}) {
    @ui.PaginationItem(ui.DOMProps{}) {
      @ui.PaginationPrevious(ui.PaginationLinkProps{Href: "#"}) { Previous }
    }
    @ui.PaginationItem(ui.DOMProps{}) {
      @ui.PaginationLink(ui.PaginationLinkProps{Href: "#", IsActive: true}) { 1 }
    }
    @ui.PaginationItem(ui.DOMProps{}) {
      @ui.PaginationNext(ui.PaginationLinkProps{Href: "#"}) { Next }
    }
  }
}`, Preview: PaginationPreview()},
		},
	},
	{
		Slug: "progress", Title: "Progress", Description: "A native progress bar styled with Tailwind, using transform for smooth animation.",
		Install: "shadcn add button", ManualInstall: "Copy ui/forms.go into your project.",
		GoUsage: `@ui.Progress(ui.ProgressProps{Value: 60})`,
		APIProps: []APIProp{
			{Name: "Value", Type: "float64", Default: "0", Desc: "Current progress value."},
			{Name: "Max", Type: "float64", Default: "100", Desc: "Maximum value."},
		},
		Examples: []ExampleEntry{
			{Name: "Various values", Desc: "Progress bars at 20%, 60%, and 100%.", GoCode: `@ui.Progress(ui.ProgressProps{Value: 20})
@ui.Progress(ui.ProgressProps{Value: 60})
@ui.Progress(ui.ProgressProps{Value: 100})`, Preview: ProgressPreview()},
		},
	},
	{
		Slug: "radio-group", Title: "Radio Group", Description: "A group of radio options with accessible markup.",
		Install: "shadcn add button", ManualInstall: "Copy ui/forms.go into your project.",
		GoUsage: `@ui.RadioGroup(ui.RadioGroupProps{Name: "plan"}) {
  @ui.RadioGroupItem(ui.RadioGroupItemProps{Name: "plan", Value: "starter", Label: "Starter"})
  @ui.RadioGroupItem(ui.RadioGroupItemProps{Name: "plan", Value: "pro", Label: "Pro"})
}`,
		APIProps: []APIProp{
			{Name: "Name", Type: "string", Default: `""`, Desc: "HTML name shared by all radio inputs."},
			{Name: "DefaultValue", Type: "string", Default: `""`, Desc: "Pre-selected option value."},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "Disables all options."},
		},
		Examples: []ExampleEntry{
			{Name: "Three options", Desc: "A radio group for plan selection.", GoCode: `@ui.RadioGroup(ui.RadioGroupProps{Name: "plan"}) {
  @ui.RadioGroupItem(ui.RadioGroupItemProps{Name: "plan", Value: "starter", Label: "Starter"})
  @ui.RadioGroupItem(ui.RadioGroupItemProps{Name: "plan", Value: "pro", Label: "Pro"})
  @ui.RadioGroupItem(ui.RadioGroupItemProps{Name: "plan", Value: "enterprise", Label: "Enterprise"})
}`, Preview: RadioGroupPreview()},
		},
	},
	{
		Slug: "separator", Title: "Separator", Description: "A horizontal or vertical visual divider with optional ARIA semantics.",
		Install: "shadcn add button", ManualInstall: "Copy ui/foundation.go into your project.",
		GoUsage: `@ui.Separator(ui.SeparatorProps{})
@ui.Separator(ui.SeparatorProps{Orientation: ui.SeparatorOrientationVertical})`,
		APIProps: []APIProp{
			{Name: "Orientation", Type: "SeparatorOrientation", Default: `"horizontal"`, Desc: `"horizontal" or "vertical".`},
			{Name: "Decorative", Type: "bool", Default: "false", Desc: "When true, adds aria-hidden instead of role=separator."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Horizontal and vertical separators.", GoCode: `@ui.Separator(ui.SeparatorProps{})`, Preview: SeparatorPreview()},
		},
	},
	{
		Slug: "skeleton", Title: "Skeleton", Description: "An animated shimmer placeholder for loading states.",
		Install: "shadcn add button", ManualInstall: "Copy ui/foundation.go into your project.",
		GoUsage: `@ui.Skeleton(ui.DOMProps{Class: "h-4 w-full"})`,
		APIProps: []APIProp{
			{Name: "DOMProps.Class", Type: "string", Default: `""`, Desc: "Width, height, and shape classes (e.g. h-4 w-1/2 rounded-full)."},
		},
		Examples: []ExampleEntry{
			{Name: "User avatar skeleton", Desc: "A typical loading placeholder for a user card.", GoCode: `<div class="flex items-center gap-4">
  @ui.Skeleton(ui.DOMProps{Class: "size-12 rounded-full"})
  <div class="grid gap-2 flex-1">
    @ui.Skeleton(ui.DOMProps{Class: "h-4 w-full"})
    @ui.Skeleton(ui.DOMProps{Class: "h-4 w-3/4"})
  </div>
</div>`, Preview: SkeletonPreview()},
			{Name: "Card skeleton", Desc: "A full card placeholder.", GoCode: `@ui.Card(ui.DOMProps{Class: "w-64"}) {
  @ui.CardHeader(ui.DOMProps{}) {
    @ui.Skeleton(ui.DOMProps{Class: "h-5 w-2/3"})
    @ui.Skeleton(ui.DOMProps{Class: "h-4 w-1/2 mt-1"})
  }
  @ui.CardContent(ui.DOMProps{}) {
    @ui.Skeleton(ui.DOMProps{Class: "h-24 w-full"})
  }
}`, Preview: SkeletonCardPreview()},
		},
	},
	{
		Slug: "slider", Title: "Slider", Description: "A range slider for numeric input with single or dual thumbs.",
		Install: "shadcn add button", ManualInstall: "Copy ui/forms.go into your project.",
		GoUsage: `@ui.Slider(ui.SliderProps{Min: 0, Max: 100, DefaultValue: []float64{40}})`,
		APIProps: []APIProp{
			{Name: "Min", Type: "float64", Default: "0", Desc: "Minimum value."},
			{Name: "Max", Type: "float64", Default: "100", Desc: "Maximum value."},
			{Name: "Step", Type: "float64", Default: "1", Desc: "Step increment."},
			{Name: "DefaultValue", Type: "[]float64", Default: "nil", Desc: "Default thumb positions. Two values = range slider."},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "Disables the slider."},
		},
		Examples: []ExampleEntry{
			{Name: "Single & range", Desc: "Single value and range sliders.", GoCode: `@ui.Slider(ui.SliderProps{Min: 0, Max: 100, DefaultValue: []float64{40}})
@ui.Slider(ui.SliderProps{Min: 0, Max: 100, DefaultValue: []float64{20, 70}})`, Preview: SliderPreview()},
		},
	},
	{
		Slug: "spinner", Title: "Spinner", Description: "A lightweight animated loading indicator.",
		Install: "shadcn add button", ManualInstall: "Copy ui/foundation.go into your project.",
		GoUsage: `@ui.Spinner(ui.SpinnerProps{Size: ui.SpinnerSizeMD, Label: "Loading"})`,
		APIProps: []APIProp{
			{Name: "Size", Type: "SpinnerSize", Default: `"md"`, Desc: `"sm", "md", or "lg".`},
			{Name: "Label", Type: "string", Default: `""`, Desc: "Sets aria-label for accessibility. Hides spinner from screen readers when empty."},
		},
		Examples: []ExampleEntry{
			{Name: "Sizes", Desc: "Small, medium, and large spinners.", GoCode: `@ui.Spinner(ui.SpinnerProps{Size: ui.SpinnerSizeSM})
@ui.Spinner(ui.SpinnerProps{Size: ui.SpinnerSizeMD})
@ui.Spinner(ui.SpinnerProps{Size: ui.SpinnerSizeLG})`, Preview: SpinnerPreview()},
		},
	},
	{
		Slug: "switch", Title: "Switch", Description: "A toggle control styled as a sliding switch.",
		Install: "shadcn add button", ManualInstall: "Copy ui/forms.go into your project.",
		GoUsage: `@ui.Switch(ui.SwitchProps{Name: "notifications", Checked: true})`,
		APIProps: []APIProp{
			{Name: "Name", Type: "string", Default: `""`, Desc: "HTML name for form submission."},
			{Name: "Checked", Type: "bool", Default: "false", Desc: "Whether the switch is on."},
			{Name: "Size", Type: "string", Default: `""`, Desc: `"sm" for a smaller switch.`},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "Disables the switch."},
		},
		Examples: []ExampleEntry{
			{Name: "States", Desc: "Off and on states.", GoCode: `@ui.Switch(ui.SwitchProps{Name: "mode"})
@ui.Switch(ui.SwitchProps{Name: "notifs", Checked: true})`, Preview: SwitchPreview()},
		},
	},
	{
		Slug: "tabs", Title: "Tabs", Description: "Tabbed panels for organizing content into sections.",
		Install: "shadcn add button", ManualInstall: "Copy ui/navigation.go into your project.",
		GoUsage: `@ui.Tabs(ui.TabsProps{DefaultValue: "account"}) {
  @ui.TabsList(ui.TabsListProps{}) {
    @ui.TabsTrigger(ui.TabsTriggerProps{Value: "account", Active: true}) { <span>Account</span> }
    @ui.TabsTrigger(ui.TabsTriggerProps{Value: "password"}) { <span>Password</span> }
  }
  @ui.TabsContent(ui.TabsContentProps{Value: "account", Active: true}) { Account content }
  @ui.TabsContent(ui.TabsContentProps{Value: "password"}) { Password content }
}`,
		APIProps: []APIProp{
			{Name: "DefaultValue", Type: "string", Default: `""`, Desc: "The default active tab value."},
			{Name: "TabsTrigger.Active", Type: "bool", Default: "false", Desc: "Marks the trigger as the active tab."},
			{Name: "TabsContent.Active", Type: "bool", Default: "false", Desc: "Shows this panel when true, hides it when false."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Two-tab layout.", GoCode: `@ui.Tabs(ui.TabsProps{DefaultValue: "account"}) {
  @ui.TabsList(ui.TabsListProps{}) {
    @ui.TabsTrigger(ui.TabsTriggerProps{Value: "account", Active: true}) { <span>Account</span> }
    @ui.TabsTrigger(ui.TabsTriggerProps{Value: "password"}) { <span>Password</span> }
  }
  @ui.TabsContent(ui.TabsContentProps{Value: "account", Active: true}) {
    <p class="text-sm p-4">Account settings.</p>
  }
}`, Preview: TabsPreview()},
		},
	},
	{
		Slug: "textarea", Title: "Textarea", Description: "A styled native multi-line text input.",
		Install: "shadcn add button", ManualInstall: "Copy ui/forms.go into your project.",
		GoUsage: `@ui.Textarea(ui.TextareaProps{Placeholder: "Type your message here."})`,
		APIProps: []APIProp{
			{Name: "Placeholder", Type: "string", Default: `""`, Desc: "Placeholder text."},
			{Name: "Rows", Type: "int", Default: "0", Desc: "Number of visible text rows."},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "Disables the textarea."},
			{Name: "Required", Type: "bool", Default: "false", Desc: "Marks as required."},
		},
		Examples: []ExampleEntry{
			{Name: "Default & Disabled", Desc: "A default and a disabled textarea.", GoCode: `@ui.Textarea(ui.TextareaProps{Placeholder: "Type your message here."})
@ui.Textarea(ui.TextareaProps{Placeholder: "Disabled", Disabled: true})`, Preview: TextareaPreview()},
		},
	},
	{
		Slug: "toggle", Title: "Toggle", Description: "A pressed-state button — on or off.",
		Install: "shadcn add button", ManualInstall: "Copy ui/navigation.go into your project.",
		GoUsage: `@ui.Toggle(ui.ToggleProps{}) { <span>Bold</span> }`,
		APIProps: []APIProp{
			{Name: "Pressed", Type: "bool", Default: "false", Desc: "Whether the toggle is in the on state."},
			{Name: "Variant", Type: "string", Default: `""`, Desc: `"" (default) or "outline".`},
			{Name: "Size", Type: "string", Default: `""`, Desc: `"sm", "" (default), or "lg".`},
			{Name: "Disabled", Type: "bool", Default: "false", Desc: "Disables the toggle."},
		},
		Examples: []ExampleEntry{
			{Name: "States & variants", Desc: "Default, pressed, outline, and disabled.", GoCode: `@ui.Toggle(ui.ToggleProps{}) { <span>Bold</span> }
@ui.Toggle(ui.ToggleProps{Pressed: true}) { <span>Italic</span> }
@ui.Toggle(ui.ToggleProps{Variant: "outline"}) { <span>Outline</span> }`, Preview: TogglePreview()},
		},
	},
	{
		Slug: "typography", Title: "Typography", Description: "A docs-oriented set of semantic HTML typography components.",
		Install: "shadcn add button", ManualInstall: "Copy ui/foundation.go into your project.",
		GoUsage: `@ui.H1(ui.DOMProps{}) { <span>Heading</span> }
@ui.P(ui.DOMProps{}) { <span>Paragraph text</span> }
@ui.Muted(ui.DOMProps{}) { <span>Muted helper</span> }`,
		APIProps: []APIProp{
			{Name: "H1 — H4", Type: "templ component", Default: "-", Desc: "Heading levels 1–4 with typed tracking."},
			{Name: "P", Type: "templ component", Default: "-", Desc: "Body paragraph."},
			{Name: "Lead", Type: "templ component", Default: "-", Desc: "Large introductory paragraph."},
			{Name: "Muted", Type: "templ component", Default: "-", Desc: "De-emphasized text."},
			{Name: "InlineCode", Type: "templ component", Default: "-", Desc: "Inline monospace code span."},
		},
		Examples: []ExampleEntry{
			{Name: "All variants", Desc: "The full typography scale.", GoCode: `@ui.H1(ui.DOMProps{}) { <span>Heading 1</span> }
@ui.H2(ui.DOMProps{}) { <span>Heading 2</span> }
@ui.P(ui.DOMProps{}) { <span>Body paragraph.</span> }
@ui.Muted(ui.DOMProps{}) { <span>Muted text.</span> }`, Preview: TypographyPreview()},
		},
	},
	// ── Stub entries (no live preview available yet — overlay/JS-dependent) ──────
	{Slug: "avatar", Title: "Avatar", Description: "Displays a user image with a graceful fallback to initials.", Install: "shadcn add button", GoUsage: `@ui.Avatar(ui.AvatarProps{Src: "/avatar.jpg", Fallback: "JD"})`, APIProps: []APIProp{{Name: "Src", Type: "string", Default: `""`, Desc: "Image source URL."}, {Name: "Fallback", Type: "string", Default: `""`, Desc: "Text shown when image fails to load."}}},
	{Slug: "calendar", Title: "Calendar", Description: "A month view calendar with flexible selection modes.", Install: "shadcn add button", GoUsage: `@ui.Calendar(ui.CalendarProps{Mode: "single"})`, APIProps: []APIProp{{Name: "Mode", Type: "string", Default: `"single"`, Desc: `"single" or "range".`}}},
	{Slug: "carousel", Title: "Carousel", Description: "A swipeable set of slides with navigation controls.", Install: "shadcn add button", GoUsage: `@ui.Carousel(ui.CarouselProps{}) { ... }`, APIProps: []APIProp{{Name: "Orientation", Type: "string", Default: `"horizontal"`, Desc: "Slide direction."}}},
	{Slug: "chart", Title: "Chart", Description: "A Recharts-style shell for vanilla SVG chart configuration, tooltips, and legend styling.", Install: "shadcn add button", GoUsage: `@ui.ChartContainer(ui.ChartContainerProps{Config: cfg}) { ... }`, APIProps: []APIProp{{Name: "Config", Type: "ChartConfig", Default: "-", Desc: "Series color and label configuration for the vanilla chart runtime."}}},
	{Slug: "combobox", Title: "Combobox", Description: "Searchable selection with single or multi-value support.", Install: "shadcn add button", GoUsage: `@ui.Combobox(ui.ComboboxProps{}) { ... }`, APIProps: []APIProp{{Name: "Multiple", Type: "bool", Default: "false", Desc: "Allow multiple selections."}}},
	{Slug: "command", Title: "Command", Description: "A searchable command palette component.", Install: "shadcn add button", GoUsage: `@ui.Command(ui.CommandProps{}) { ... }`, APIProps: []APIProp{{Name: "DefaultValue", Type: "string", Default: `""`, Desc: "Default search value."}}},
	{Slug: "context-menu", Title: "Context Menu", Description: "A right-click menu with nested submenus.", Install: "shadcn add button", GoUsage: `@ui.ContextMenu(ui.ContextMenuProps{}) { ... }`, APIProps: []APIProp{{Name: "Variant", Type: "string", Default: `""`, Desc: "Item variant."}}},
	{Slug: "data-table", Title: "Data Table", Description: "A server-rendered table pattern for app data with sorting and pagination.", Install: "shadcn add button", GoUsage: `@ui.DataTable(ui.DataTableProps{Columns: cols, Rows: rows}) {}`, APIProps: []APIProp{{Name: "Columns", Type: "[]Column", Default: "nil", Desc: "Column definitions."}, {Name: "Rows", Type: "[][]string", Default: "nil", Desc: "Table row data."}}},
	{Slug: "date-picker", Title: "Date Picker", Description: "A calendar popover for selecting single or range dates.", Install: "shadcn add button", GoUsage: `@ui.DatePicker(ui.DatePickerProps{Mode: "single"})`, APIProps: []APIProp{{Name: "Mode", Type: "string", Default: `"single"`, Desc: `"single" or "range".`}}},
	{Slug: "dialog", Title: "Dialog", Description: "A modal dialog used for forms and transient tasks.", Install: "shadcn add button", GoUsage: `@ui.Dialog(ui.DialogProps{Open: true}) {
  @ui.DialogContent(ui.DOMProps{}) {
    @ui.DialogTitle(ui.DOMProps{}) { Title }
  }
}`, APIProps: []APIProp{{Name: "Open", Type: "bool", Default: "false", Desc: "Controls dialog visibility."}, {Name: "Modal", Type: "bool", Default: "true", Desc: "Traps focus inside the dialog."}}},
	{Slug: "direction", Title: "Direction", Description: "An RTL/LTR direction context provider.", Install: "shadcn add button", GoUsage: `@ui.DirectionProvider(ui.DirectionProviderProps{Direction: "rtl"}) { ... }`, APIProps: []APIProp{{Name: "Direction", Type: "string", Default: `"ltr"`, Desc: `"ltr" or "rtl".`}}},
	{Slug: "drawer", Title: "Drawer", Description: "A slide-over surface from screen edges with mobile-friendly motion.", Install: "shadcn add button", GoUsage: `@ui.Drawer(ui.DrawerProps{}) {
  @ui.DrawerContent(ui.DOMProps{}) { ... }
}`, APIProps: []APIProp{{Name: "Side", Type: "string", Default: `"bottom"`, Desc: `"top", "bottom", "left", "right".`}}},
	{Slug: "dropdown-menu", Title: "Dropdown Menu", Description: "A context-driven popover menu with keyboard navigation.", Install: "shadcn add button", GoUsage: `@ui.DropdownMenu(ui.DropdownMenuProps{}) {
  @ui.DropdownMenuTrigger(ui.DOMProps{}) { Open }
  @ui.DropdownMenuContent(ui.DOMProps{}) {
    @ui.DropdownMenuItem(ui.DropdownMenuItemProps{}) { Item }
  }
}`, APIProps: []APIProp{{Name: "Trigger", Type: "children", Default: "-", Desc: "The trigger element."}}},
	{Slug: "empty", Title: "Empty", Description: "A polished empty-state placeholder helper.", Install: "shadcn add button", GoUsage: `@ui.Empty(ui.DOMProps{}) {
  @ui.EmptyTitle(ui.DOMProps{}) { No results }
  @ui.EmptyDescription(ui.DOMProps{}) { Try adjusting your filters. }
}`, APIProps: []APIProp{}},
	{Slug: "field", Title: "Field", Description: "A form field grouping component with label, description, and error slots.", Install: "shadcn add button", GoUsage: `@ui.Field(ui.DOMProps{}) {
  @ui.FieldLabel(ui.DOMProps{}) { Email }
  @ui.Input(ui.InputProps{Placeholder: "you@example.com"})
  @ui.FieldError(ui.FieldErrorProps{Errors: []string{"Invalid email."}})
}`, APIProps: []APIProp{{Name: "Orientation", Type: "FieldOrientation", Default: `"vertical"`, Desc: `"vertical", "horizontal", "responsive".`}}},
	{Slug: "hover-card", Title: "Hover Card", Description: "A hover/focus-driven preview popover, ideal for user profiles and link previews.", Install: "shadcn add button", GoUsage: `@ui.HoverCard(ui.HoverCardProps{}) {
  @ui.HoverCardTrigger(ui.DOMProps{}) { Hover me }
  @ui.HoverCardContent(ui.DOMProps{}) { Card content }
}`, APIProps: []APIProp{{Name: "OpenDelayMs", Type: "int", Default: "0", Desc: "Delay before showing."}, {Name: "CloseDelayMs", Type: "int", Default: "0", Desc: "Delay before hiding."}}},
	{Slug: "input-group", Title: "Input Group", Description: "Input with leading/trailing addons and icon slots.", Install: "shadcn add button", GoUsage: `@ui.InputGroup(ui.InputGroupProps{}) {
  @ui.InputGroupAddon(ui.InputGroupAddonProps{}) { $ }
  @ui.InputGroupInput(ui.InputGroupInputProps{Placeholder: "0.00"})
}`, APIProps: []APIProp{{Name: "Align", Type: "InputGroupAddonAlign", Default: `"inline-start"`, Desc: "Position of the addon."}}},
	{Slug: "item", Title: "Item", Description: "A generic list/card item surface with header, body, and action slots.", Install: "shadcn add button", GoUsage: `@ui.Item(ui.DOMProps{}) {
  @ui.ItemHeader(ui.DOMProps{}) { Title }
}`, APIProps: []APIProp{}},
	{Slug: "menubar", Title: "Menubar", Description: "A horizontal application menu bar.", Install: "shadcn add button", GoUsage: `@ui.Menubar(ui.MenubarProps{}) {
  @ui.MenubarMenu(ui.DOMProps{}) {
    @ui.MenubarTrigger(ui.DOMProps{}) { File }
    @ui.MenubarContent(ui.DOMProps{}) {
      @ui.MenubarItem(ui.DropdownMenuItemProps{}) { New Tab }
    }
  }
}`, APIProps: []APIProp{}},
	{Slug: "navigation-menu", Title: "Navigation Menu", Description: "A top-level navigation with disclosure viewport.", Install: "shadcn add button", GoUsage: `@ui.NavigationMenu(ui.NavigationMenuProps{}) {
  @ui.NavigationMenuList(ui.DOMProps{}) {
    @ui.NavigationMenuItem(ui.DOMProps{}) {
      @ui.NavigationMenuLink(ui.BreadcrumbLinkProps{Href: "/"}) { Home }
    }
  }
}`, APIProps: []APIProp{}},
	{Slug: "popover", Title: "Popover", Description: "Floating content anchored to a trigger element.", Install: "shadcn add button", GoUsage: `@ui.Popover(ui.PopoverProps{}) {
  @ui.PopoverTrigger(ui.DOMProps{}) { Open popover }
  @ui.PopoverContent(ui.DOMProps{}) { Content here }
}`, APIProps: []APIProp{{Name: "Side", Type: "string", Default: `"bottom"`, Desc: "Preferred placement side."}, {Name: "Align", Type: "string", Default: `"center"`, Desc: "Alignment along the axis."}, {Name: "Modal", Type: "bool", Default: "false", Desc: "Whether to trap focus."}}},
	{Slug: "resizable", Title: "Resizable", Description: "Pointer-driven resize panels.", Install: "shadcn add button", GoUsage: `@ui.ResizablePanelGroup(ui.DOMProps{}) { ... }`, APIProps: []APIProp{{Name: "Direction", Type: "string", Default: `"horizontal"`, Desc: `"horizontal" or "vertical".`}}},
	{Slug: "scroll-area", Title: "Scroll Area", Description: "Styled scrollable containers with custom scrollbars.", Install: "shadcn add button", GoUsage: `@ui.ScrollArea(ui.DOMProps{Class: "h-72"}) {
  // long content
}`, APIProps: []APIProp{{Name: "Orientation", Type: "string", Default: `"vertical"`, Desc: `"vertical" or "both".`}}},
	{Slug: "select", Title: "Select", Description: "A custom accessible select built around a listbox runtime.", Install: "shadcn add button", GoUsage: `@ui.Select(ui.SelectProps{}) {
  @ui.SelectTrigger(ui.DOMProps{}) { Select an option }
  @ui.SelectContent(ui.DOMProps{}) {
    @ui.SelectItem(ui.SelectItemProps{Value: "apple"}) { Apple }
  }
}`, APIProps: []APIProp{{Name: "DefaultOpen", Type: "bool", Default: "false", Desc: "Initial open state."}, {Name: "Required", Type: "bool", Default: "false", Desc: "Marks as required."}}},
	{Slug: "sheet", Title: "Sheet", Description: "A slide-over panel that animates from any screen edge.", Install: "shadcn add button", GoUsage: `@ui.Sheet(ui.SheetProps{Side: "right"}) {
  @ui.SheetContent(ui.DOMProps{}) {
    @ui.SheetTitle(ui.DOMProps{}) { Profile }
  }
}`, APIProps: []APIProp{{Name: "Side", Type: "string", Default: `"right"`, Desc: `"top", "bottom", "left", "right".`}, {Name: "ShowCloseButton", Type: "bool", Default: "false", Desc: "Show a close button."}}},
	{Slug: "sidebar", Title: "Sidebar", Description: "A full app-shell sidebar with navigation, groups, and collapsible menus.", Install: "shadcn add button", GoUsage: `@ui.SidebarProvider(ui.SidebarProviderProps{DefaultOpen: true}) {
  @ui.Sidebar(ui.SidebarProps{}) {
    @ui.SidebarContent(ui.DOMProps{}) {
      @ui.SidebarMenu(ui.DOMProps{}) {
        @ui.SidebarMenuItem(ui.DOMProps{}) {
          @ui.SidebarMenuButton(ui.SidebarMenuButtonProps{IsActive: true}) { Dashboard }
        }
      }
    }
  }
}`, APIProps: []APIProp{{Name: "Side", Type: "string", Default: `"left"`, Desc: `"left" or "right".`}, {Name: "Variant", Type: "string", Default: `"sidebar"`, Desc: `"sidebar", "floating", or "inset".`}, {Name: "Collapsible", Type: "string", Default: `"offcanvas"`, Desc: `"offcanvas", "icon", or "none".`}}},
	{Slug: "sonner", Title: "Sonner", Description: "A toast notification system using the Sonner design.", Install: "shadcn add button", GoUsage: `@ui.Toaster(ui.ToasterProps{Position: "bottom-right", RichColors: true})`, APIProps: []APIProp{{Name: "Theme", Type: "string", Default: `""`, Desc: "Force light or dark mode."}, {Name: "Position", Type: "string", Default: `"bottom-right"`, Desc: "Toast position on screen."}}},
	{Slug: "table", Title: "Table", Description: "Table primitives with shadcn styling for simple data display.", Install: "shadcn add button", GoUsage: `@ui.Table(ui.DOMProps{}) {
  @ui.TableHeader(ui.DOMProps{}) {
    @ui.TableRow(ui.DOMProps{}) {
      @ui.TableHead(ui.DOMProps{}) { Name }
      @ui.TableHead(ui.DOMProps{}) { Status }
    }
  }
  @ui.TableBody(ui.DOMProps{}) {
    @ui.TableRow(ui.DOMProps{}) {
      @ui.TableCell(ui.DOMProps{}) { Alice }
      @ui.TableCell(ui.DOMProps{}) { Active }
    }
  }
}`, APIProps: []APIProp{}},
	{Slug: "toast", Title: "Toast", Description: "Legacy toast primitives for notification messages.", Install: "shadcn add button", GoUsage: `@ui.Toast(ui.ToastProps{Open: true}) {
  @ui.ToastTitle(ui.DOMProps{}) { Scheduled! }
  @ui.ToastDescription(ui.DOMProps{}) { Event created for Friday. }
}`, APIProps: []APIProp{{Name: "Variant", Type: "string", Default: `""`, Desc: "Visual variant (e.g. destructive)."}, {Name: "Duration", Type: "int", Default: "5000", Desc: "Auto-dismiss time in ms."}}},
	{Slug: "toggle-group", Title: "Toggle Group", Description: "A group of related toggle buttons for single or multi-select.", Install: "shadcn add button", GoUsage: `@ui.ToggleGroup(ui.ToggleGroupProps{Type: "single"}) {
  @ui.ToggleGroupItem(ui.ToggleGroupItemProps{Value: "bold"}) { B }
  @ui.ToggleGroupItem(ui.ToggleGroupItemProps{Value: "italic"}) { I }
}`, APIProps: []APIProp{{Name: "Type", Type: "string", Default: `"single"`, Desc: `"single" or "multiple".`}}},
	{Slug: "tooltip", Title: "Tooltip", Description: "A floating hint label for compact icon controls.", Install: "shadcn add button", GoUsage: `@ui.TooltipProvider(ui.TooltipProviderProps{}) {
  @ui.Tooltip(ui.TooltipProps{}) {
    @ui.TooltipTrigger(ui.DOMProps{}) { Hover me }
    @ui.TooltipContent(ui.DOMProps{}) { Helpful hint }
  }
}`, APIProps: []APIProp{{Name: "Side", Type: "string", Default: `"top"`, Desc: `"top", "bottom", "left", "right".`}, {Name: "DelayDuration", Type: "int", Default: "700", Desc: "Hover delay in ms before showing."}}},
}
var blockDocs = []BlockEntry{
	{Slug: "dashboard-01", Title: "A dashboard with sidebar, charts and data table", Description: "A dense app shell that combines a sidebar, KPI cards, charts, and a data table.", Category: "featured", Command: "go run ./cmd/shadcn add dashboard-01", Files: []string{"app/dashboard/page.templ", "components/app-sidebar.templ", "components/chart-area-interactive.templ", "components/data-table.templ", "components/section-cards.templ", "components/site-header.templ"}},
	{Slug: "sidebar-07", Title: "A sidebar that collapses to icons", Description: "A compact application shell with icon-only collapse behavior.", Category: "sidebar", Command: "go run ./cmd/shadcn add sidebar-07", Files: []string{"app/dashboard/page.templ", "components/app-sidebar.templ", "components/team-switcher.templ"}},
	{Slug: "sidebar-03", Title: "A sidebar with submenus", Description: "Sidebar navigation with nested submenu states and a content region.", Category: "sidebar", Command: "go run ./cmd/shadcn add sidebar-03", Files: []string{"app/dashboard/page.templ", "components/app-sidebar.templ"}},
	{Slug: "login-01", Title: "A simple login form", Description: "A centered login form with a compact footprint.", Category: "login", Command: "go run ./cmd/shadcn add login-01", Files: []string{"app/login/page.templ", "components/login-form.templ"}},
	{Slug: "login-03", Title: "A login page with a muted background color", Description: "A basic authentication surface with soft contrast and a small brand lockup.", Category: "login", Command: "go run ./cmd/shadcn add login-03", Files: []string{"app/login/page.templ", "components/login-form.templ"}},
	{Slug: "login-04", Title: "A login page with form and image", Description: "A wider authentication layout that pairs form and image panels.", Category: "login", Command: "go run ./cmd/shadcn add login-04", Files: []string{"app/login/page.templ", "components/login-form.templ"}},
	{Slug: "signup-01", Title: "A signup page with a cover image", Description: "A registration layout with a strong media panel.", Category: "signup", Command: "go run ./cmd/shadcn add signup-01", Files: []string{"app/signup/page.templ", "components/signup-form.templ"}},
	{Slug: "signup-02", Title: "A signup form with social providers", Description: "A signup flow that emphasizes third-party providers and email signup.", Category: "signup", Command: "go run ./cmd/shadcn add signup-02", Files: []string{"app/signup/page.templ", "components/signup-form.templ"}},
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

func init() {
	normalizeComponentDocs()
}

func normalizeComponentDocs() {
	for i := range componentDocs {
		doc := &componentDocs[i]
		if doc.Install == "" || doc.Install == "shadcn add button" || doc.Install == "shadcn add <component>" {
			doc.Install = "shadcn add " + doc.Slug
		}
		if doc.GoUsage == "" {
			doc.GoUsage = defaultGoUsage(doc)
		}
		if doc.ManualInstall == "" {
			doc.ManualInstall = defaultManualInstall(doc)
		}
		if len(doc.APIProps) == 0 {
			doc.APIProps = defaultAPIProps(doc)
		}
		if doc.SourceCode == "" {
			doc.SourceCode = defaultSourceCode(doc)
		}
		if len(doc.Examples) == 0 {
			doc.Examples = defaultExamples(doc)
		}
		for j := range doc.Examples {
			example := &doc.Examples[j]
			if example.Desc == "" {
				example.Desc = defaultExampleDesc(doc, example.Name)
			}
			if example.GoCode == "" {
				example.GoCode = defaultExampleGoCode(doc, example.Name)
			}
			if example.Preview == nil {
				example.Preview = previewForExample(doc.Slug, example.Name)
			}
		}
	}
}

func defaultGoUsage(doc *ComponentDocEntry) string {
	name := strings.ReplaceAll(doc.Title, " ", "")
	if name == "" {
		name = TitleFromSlug(doc.Slug)
		name = strings.ReplaceAll(name, " ", "")
	}
	return fmt.Sprintf("@ui.%s(ui.%sProps{})", name, name)
}

func defaultManualInstall(doc *ComponentDocEntry) string {
	return fmt.Sprintf("Copy `ui/%s.go` into your project and keep the exported API aligned with the upstream registry.", doc.Slug)
}

func defaultAPIProps(doc *ComponentDocEntry) []APIProp {
	props := make([]APIProp, 0, len(doc.API))
	for _, name := range doc.API {
		props = append(props, APIProp{
			Name:    name,
			Type:    inferPropType(name),
			Default: inferPropDefault(name),
			Desc:    inferPropDesc(doc.Title, name),
		})
	}
	return props
}

func inferPropType(name string) string {
	lower := strings.ToLower(name)
	switch {
	case strings.Contains(lower, "open"), strings.Contains(lower, "checked"), strings.Contains(lower, "disabled"), strings.Contains(lower, "required"), strings.Contains(lower, "active"), strings.Contains(lower, "collapsible"):
		return "bool"
	case strings.Contains(lower, "items"), strings.Contains(lower, "values"), strings.Contains(lower, "files"), strings.Contains(lower, "errors"):
		return "[]string"
	case strings.Contains(lower, "max"), strings.Contains(lower, "min"), strings.Contains(lower, "count"), strings.Contains(lower, "index"), strings.Contains(lower, "step"), strings.Contains(lower, "width"), strings.Contains(lower, "height"), strings.Contains(lower, "duration"):
		return "int"
	default:
		return "string"
	}
}

func inferPropDefault(name string) string {
	lower := strings.ToLower(name)
	switch {
	case strings.Contains(lower, "open"), strings.Contains(lower, "checked"), strings.Contains(lower, "disabled"), strings.Contains(lower, "required"), strings.Contains(lower, "active"), strings.Contains(lower, "collapsible"):
		return "false"
	case strings.Contains(lower, "items"), strings.Contains(lower, "values"), strings.Contains(lower, "files"), strings.Contains(lower, "errors"):
		return "nil"
	case strings.Contains(lower, "max"), strings.Contains(lower, "min"), strings.Contains(lower, "count"), strings.Contains(lower, "index"), strings.Contains(lower, "step"), strings.Contains(lower, "width"), strings.Contains(lower, "height"), strings.Contains(lower, "duration"):
		return "0"
	default:
		return "\"\""
	}
}

func inferPropDesc(componentTitle, name string) string {
	return fmt.Sprintf("Controls the %s behavior for %s.", strings.ToLower(name), componentTitle)
}

func defaultSourceCode(doc *ComponentDocEntry) string {
	name := strings.ReplaceAll(doc.Title, " ", "")
	if name == "" {
		name = TitleFromSlug(doc.Slug)
		name = strings.ReplaceAll(name, " ", "")
	}
	return fmt.Sprintf(`func %sPreview() templ.Component {
	return ui.%s(ui.%sProps{})
}`, name, name, name)
}

func defaultExamples(doc *ComponentDocEntry) []ExampleEntry {
	examples := make([]ExampleEntry, 0, len(doc.Example))
	for _, name := range doc.Example {
		examples = append(examples, ExampleEntry{
			Name:    name,
			Desc:    defaultExampleDesc(doc, name),
			GoCode:  defaultExampleGoCode(doc, name),
			Preview: previewForExample(doc.Slug, name),
		})
	}
	return examples
}

func defaultExampleDesc(doc *ComponentDocEntry, name string) string {
	return fmt.Sprintf("%s example for %s.", name, doc.Title)
}

func defaultExampleGoCode(doc *ComponentDocEntry, name string) string {
	componentName := strings.ReplaceAll(doc.Title, " ", "")
	if componentName == "" {
		componentName = TitleFromSlug(doc.Slug)
		componentName = strings.ReplaceAll(componentName, " ", "")
	}
	return fmt.Sprintf(`func %s%s() templ.Component {
	return ui.%s(ui.%sProps{})
}`, componentName, strings.ReplaceAll(name, " ", ""), componentName, componentName)
}

// cliUsageExample returns a Go/templ code snippet for the CLI docs page.
// Kept here (plain Go) so the templ parser does not try to interpret the @ symbols.
func cliUsageExample() string {
	return `package views

import "shadcn/ui"

templ MyPage() {
	@ui.Button(ui.ButtonProps{Label: "Click me"})

	@ui.Card(ui.DOMProps{}) {
		@ui.CardHeader(ui.DOMProps{}) {
			@ui.CardTitle(ui.DOMProps{}) { <span>Hello</span> }
		}
		@ui.CardContent(ui.DOMProps{}) {
			<p>Card body content here.</p>
		}
	}
}`
}

func BlockSourceCode(slug string) string {
	switch slug {
	case "dashboard-01":
		return `templ Dashboard01Preview() {
	@PreviewLayout("Dashboard 01 Preview") {
		<div class="flex h-screen bg-background">
			@ui.Sidebar(ui.SidebarProps{
				Side:        "left",
				Variant:     "inset",
				Collapsible: "icon",
				DOMProps: ui.DOMProps{
					Class: "border-r border-border/70",
				},
			}) {
				@ui.SidebarHeader(ui.DOMProps{Class: "p-4 border-b border-border/70"}) {
					<div class="flex items-center gap-2 font-semibold">
						<span class="flex size-7 items-center justify-center rounded-md bg-primary text-xs text-primary-foreground">D</span>
						<span>Dashboard</span>
					</div>
				}
				@ui.SidebarContent(ui.DOMProps{}) {
					@ui.SidebarGroup(ui.DOMProps{}) {
						@ui.SidebarMenu(ui.DOMProps{}) {
							@ui.SidebarMenuItem(ui.DOMProps{}) { @ui.SidebarMenuButton(ui.SidebarMenuButtonProps{IsActive: true}) { <span>Overview</span> } }
							@ui.SidebarMenuItem(ui.DOMProps{}) { @ui.SidebarMenuButton(ui.SidebarMenuButtonProps{}) { <span>Analytics</span> } }
							@ui.SidebarMenuItem(ui.DOMProps{}) { @ui.SidebarMenuButton(ui.SidebarMenuButtonProps{}) { <span>Reports</span> } }
						}
					}
					@ui.SidebarGroup(ui.DOMProps{Class: "mt-4"}) {
						@ui.SidebarGroupLabel(ui.DOMProps{}) { <span>Settings</span> }
						@ui.SidebarGroupContent(ui.DOMProps{}) {
							@ui.SidebarMenu(ui.DOMProps{}) {
								@ui.SidebarMenuItem(ui.DOMProps{}) { @ui.SidebarMenuButton(ui.SidebarMenuButtonProps{}) { <span>Profile</span> } }
								@ui.SidebarMenuItem(ui.DOMProps{}) { @ui.SidebarMenuButton(ui.SidebarMenuButtonProps{}) { <span>Team</span> } }
							}
						}
					}
				}
			}
			<main class="flex-1 overflow-auto">
				<div class="p-6 space-y-6">
					<div class="flex items-center justify-between">
						<h1 class="text-2xl font-semibold">Dashboard</h1>
						@ui.Button(ui.ButtonProps{Label: "New Report", Variant: ui.ButtonVariantDefault, Size: ui.ButtonSizeSM, DOMProps: ui.DOMProps{Attrs: t.Attributes{"type": "button"}}})
					</div>
					<div class="grid gap-4 md:grid-cols-3">
						@ui.Card(ui.DOMProps{}) {
							@ui.CardHeader(ui.DOMProps{Class: "pb-2"}) {
								@ui.CardDescription(ui.DOMProps{}) { <span>Total Revenue</span> }
								@ui.CardTitle(ui.DOMProps{}) { <span>$45,231.89</span> }
							}
							@ui.CardContent(ui.DOMProps{Class: "pt-0"}) {
								<p class="text-xs text-muted-foreground">+20.1% from last month</p>
							}
						}
						@ui.Card(ui.DOMProps{}) {
							@ui.CardHeader(ui.DOMProps{Class: "pb-2"}) {
								@ui.CardDescription(ui.DOMProps{}) { <span>Subscriptions</span> }
								@ui.CardTitle(ui.DOMProps{}) { <span>+2350</span> }
							}
							@ui.CardContent(ui.DOMProps{Class: "pt-0"}) {
								<p class="text-xs text-muted-foreground">+180.1% from last month</p>
							}
						}
						@ui.Card(ui.DOMProps{}) {
							@ui.CardHeader(ui.DOMProps{Class: "pb-2"}) {
								@ui.CardDescription(ui.DOMProps{}) { <span>Sales</span> }
								@ui.CardTitle(ui.DOMProps{}) { <span>+12,234</span> }
							}
							@ui.CardContent(ui.DOMProps{Class: "pt-0"}) {
								<p class="text-xs text-muted-foreground">+19% from last month</p>
							}
						}
					</div>
					@ui.Card(ui.DOMProps{}) {
						@ui.CardHeader(ui.DOMProps{}) {
							@ui.CardTitle(ui.DOMProps{}) { <span>Recent Activity</span> }
							@ui.CardDescription(ui.DOMProps{}) { <span>Latest transactions and updates</span> }
						}
						@ui.CardContent(ui.DOMProps{}) {
							@ui.Table(ui.DOMProps{}) {
								@ui.TableHeader(ui.DOMProps{}) {
									@ui.TableRow(ui.DOMProps{}) {
										@ui.TableHead(ui.DOMProps{}) { <span>Invoice</span> }
										@ui.TableHead(ui.DOMProps{}) { <span>Status</span> }
										@ui.TableHead(ui.DOMProps{}) { <span>Method</span> }
										@ui.TableHead(ui.DOMProps{}) { <span>Amount</span> }
									}
								}
								@ui.TableBody(ui.DOMProps{}) {
									@ui.TableRow(ui.DOMProps{}) {
										@ui.TableCell(ui.DOMProps{}) { <span>INV001</span> }
										@ui.TableCell(ui.DOMProps{}) { @ui.Badge(ui.BadgeProps{Label: "Paid", Variant: ui.BadgeVariantSecondary}) }
										@ui.TableCell(ui.DOMProps{}) { <span>Credit Card</span> }
										@ui.TableCell(ui.DOMProps{}) { <span>$250.00</span> }
									}
									@ui.TableRow(ui.DOMProps{}) {
										@ui.TableCell(ui.DOMProps{}) { <span>INV002</span> }
										@ui.TableCell(ui.DOMProps{}) { @ui.Badge(ui.BadgeProps{Label: "Pending", Variant: ui.BadgeVariantOutline}) }
										@ui.TableCell(ui.DOMProps{}) { <span>PayPal</span> }
										@ui.TableCell(ui.DOMProps{}) { <span>$150.00</span> }
									}
								}
							}
						}
					}
				</div>
			</main>
		</div>
	}
}`
	case "sidebar-07":
		return `templ Sidebar07Preview() {
	@PreviewLayout("Sidebar 07 Preview") {
		<div class="flex h-screen bg-background">
			@ui.Sidebar(ui.SidebarProps{
				Side:        "left",
				Variant:     "sidebar",
				Collapsible: "icon",
				DOMProps: ui.DOMProps{
					Class: "border-r border-border/70",
				},
			}) {
				@ui.SidebarHeader(ui.DOMProps{Class: "p-4 border-b border-border/70"}) {
					<div class="flex items-center gap-2 font-semibold">
						<span class="flex size-8 items-center justify-center rounded-md bg-primary text-sm text-primary-foreground">S</span>
						<span class="sidebar-text">App</span>
					</div>
				}
				@ui.SidebarContent(ui.DOMProps{}) {
					@ui.SidebarGroup(ui.DOMProps{}) {
						@ui.SidebarMenu(ui.DOMProps{}) {
							@ui.SidebarMenuItem(ui.DOMProps{}) { @ui.SidebarMenuButton(ui.SidebarMenuButtonProps{IsActive: true, Size: "sm"}) { <span class="sidebar-text">Dashboard</span> } }
							@ui.SidebarMenuItem(ui.DOMProps{}) { @ui.SidebarMenuButton(ui.SidebarMenuButtonProps{Size: "sm"}) { <span class="sidebar-text">Projects</span> } }
							@ui.SidebarMenuItem(ui.DOMProps{}) { @ui.SidebarMenuButton(ui.SidebarMenuButtonProps{Size: "sm"}) { <span class="sidebar-text">Team</span> } }
						}
					}
				}
			}
			<main class="flex-1 p-6">
				<div class="flex items-center justify-between mb-6">
					<h1 class="text-2xl font-semibold">Dashboard</h1>
					@ui.Button(ui.ButtonProps{Label: "Collapse Sidebar", Variant: ui.ButtonVariantOutline, Size: ui.ButtonSizeSM, DOMProps: ui.DOMProps{Attrs: t.Attributes{"type": "button"}}})
				</div>
				<div class="grid gap-4 md:grid-cols-2">
					@ui.Card(ui.DOMProps{}) {
						@ui.CardHeader(ui.DOMProps{}) {
							@ui.CardTitle(ui.DOMProps{}) { <span>Quick Stats</span> }
						}
						@ui.CardContent(ui.DOMProps{}) {
							<p class="text-sm text-muted-foreground">A compact sidebar that collapses to icons, maximizing content space.</p>
						}
					}
					@ui.Card(ui.DOMProps{}) {
						@ui.CardHeader(ui.DOMProps{}) {
							@ui.CardTitle(ui.DOMProps{}) { <span>Navigation</span> }
						}
						@ui.CardContent(ui.DOMProps{}) {
							<p class="text-sm text-muted-foreground">Icon-only mode for mobile or space-constrained layouts.</p>
						}
					}
				</div>
			</main>
		</div>
	}
}`
	case "sidebar-03":
		return `templ Sidebar03Preview() {
	@PreviewLayout("Sidebar 03 Preview") {
		<div class="flex h-screen bg-background">
			@ui.Sidebar(ui.SidebarProps{
				Side:        "left",
				Variant:     "inset",
				Collapsible: "none",
				DOMProps: ui.DOMProps{
					Class: "border-r border-border/70",
				},
			}) {
				@ui.SidebarHeader(ui.DOMProps{Class: "p-4 border-b border-border/70"}) {
					<div class="flex items-center gap-2 font-semibold">
						<span class="flex size-7 items-center justify-center rounded-md bg-primary text-xs text-primary-foreground">S</span>
						<span>App</span>
					</div>
				}
				@ui.SidebarContent(ui.DOMProps{}) {
					@ui.SidebarGroup(ui.DOMProps{}) {
						@ui.SidebarGroupLabel(ui.DOMProps{}) { <span>Main</span> }
						@ui.SidebarGroupContent(ui.DOMProps{}) {
							@ui.SidebarMenu(ui.DOMProps{}) {
								@ui.SidebarMenuItem(ui.DOMProps{}) { @ui.SidebarMenuButton(ui.SidebarMenuButtonProps{IsActive: true}) { <span>Dashboard</span> } }
								@ui.SidebarMenuItem(ui.DOMProps{}) {
									@ui.SidebarMenuButton(ui.SidebarMenuButtonProps{}) { <span>Products</span> }
									@ui.SidebarMenuSub(ui.DOMProps{}) {
										@ui.SidebarMenuSubItem(ui.DOMProps{}) { @ui.SidebarMenuSubButton(ui.SidebarMenuButtonProps{Size: "sm"}) { <span>All Products</span> } }
										@ui.SidebarMenuSubItem(ui.DOMProps{}) { @ui.SidebarMenuSubButton(ui.SidebarMenuButtonProps{Size: "sm"}) { <span>Categories</span> } }
									}
								}
								@ui.SidebarMenuItem(ui.DOMProps{}) {
									@ui.SidebarMenuButton(ui.SidebarMenuButtonProps{}) { <span>Settings</span> }
									@ui.SidebarMenuSub(ui.DOMProps{}) {
										@ui.SidebarMenuSubItem(ui.DOMProps{}) { @ui.SidebarMenuSubButton(ui.SidebarMenuButtonProps{Size: "sm"}) { <span>Profile</span> } }
										@ui.SidebarMenuSubItem(ui.DOMProps{}) { @ui.SidebarMenuSubButton(ui.SidebarMenuButtonProps{Size: "sm"}) { <span>Team</span> } }
									}
								}
							}
						}
					}
				}
			}
			<main class="flex-1 p-6">
				<div class="flex items-center justify-between mb-6">
					<h1 class="text-2xl font-semibold">Dashboard</h1>
				</div>
				@ui.Card(ui.DOMProps{}) {
					@ui.CardHeader(ui.DOMProps{}) {
						@ui.CardTitle(ui.DOMProps{}) { <span>Sidebar with Submenus</span> }
						@ui.CardDescription(ui.DOMProps{}) { <span>Nested navigation states for complex applications</span> }
					}
					@ui.CardContent(ui.DOMProps{}) {
						<p class="text-sm text-muted-foreground">This sidebar pattern supports hierarchical navigation with expandable submenu states.</p>
					}
				}
			</main>
		</div>
	}
}`
	case "login-01":
		return `templ Login01Preview() {
	@PreviewLayout("Login 01 Preview") {
		<div class="flex min-h-svh items-center justify-center bg-background p-4">
			<div class="w-full max-w-sm">
				@ui.Card(ui.DOMProps{}) {
					@ui.CardHeader(ui.DOMProps{Class: "space-y-1 text-center"}) {
						@ui.CardTitle(ui.DOMProps{}) { <span>Sign in</span> }
						@ui.CardDescription(ui.DOMProps{}) { <span>Enter your credentials to access your account</span> }
					}
					@ui.CardContent(ui.DOMProps{Class: "space-y-4"}) {
						<div class="space-y-2">
							@ui.Label(ui.LabelProps{For: "email-01"}) { <span>Email</span> }
							@ui.Input(ui.InputProps{
								DOMProps: ui.DOMProps{ID: "email-01"},
								Type: "email",
								Placeholder: "name@example.com",
							})
						</div>
						<div class="space-y-2">
							@ui.Label(ui.LabelProps{For: "password-01"}) { <span>Password</span> }
							@ui.Input(ui.InputProps{
								DOMProps: ui.DOMProps{ID: "password-01"},
								Type: "password",
								Placeholder: "••••••••",
							})
						</div>
							@ui.Button(ui.ButtonProps{
								Label:   "Sign in",
								Variant: ui.ButtonVariantDefault,
								Size:    ui.ButtonSizeDefault,
								DOMProps: ui.DOMProps{
									Class: "w-full",
									Attrs: t.Attributes{"type": "button"},
								},
							})
						<p class="text-center text-sm text-muted-foreground">
							Don't have an account? <a href="#" class="underline underline-offset-4 hover:text-primary">Sign up</a>
						</p>
					}
				}
			</div>
		</div>
	}
}`
	case "login-03":
		return `templ Login03Preview() {
	@PreviewLayout("Login 03 Preview") {
		<div class="flex min-h-svh items-center justify-center bg-muted/50 p-4">
			<div class="w-full max-w-sm">
				<div class="mb-8 text-center">
					<span class="inline-flex size-12 items-center justify-center rounded-lg bg-primary text-primary-foreground text-xl font-semibold">S</span>
					<h1 class="mt-4 text-2xl font-semibold">Welcome back</h1>
					<p class="mt-2 text-sm text-muted-foreground">Sign in to continue to your account</p>
				</div>
				@ui.Card(ui.DOMProps{}) {
					@ui.CardContent(ui.DOMProps{Class: "space-y-4"}) {
						<div class="space-y-2">
							@ui.Label(ui.LabelProps{For: "email-03"}) { <span>Email</span> }
							@ui.Input(ui.InputProps{
								DOMProps: ui.DOMProps{ID: "email-03"},
								Type: "email",
								Placeholder: "name@example.com",
							})
						</div>
						<div class="space-y-2">
							@ui.Label(ui.LabelProps{For: "password-03"}) { <span>Password</span> }
							@ui.Input(ui.InputProps{
								DOMProps: ui.DOMProps{ID: "password-03"},
								Type: "password",
								Placeholder: "••••••••",
							})
						</div>
						<div class="flex items-center justify-between text-sm">
							<label class="flex items-center gap-2">
								@ui.Checkbox(ui.CheckboxProps{DOMProps: ui.DOMProps{ID: "remember-03"}})
								<span>Remember me</span>
							</label>
							<a href="#" class="text-primary hover:underline">Forgot password?</a>
						</div>
						@ui.Button(ui.ButtonProps{
							Label:   "Sign in",
							Variant: ui.ButtonVariantDefault,
							Size:    ui.ButtonSizeDefault,
							DOMProps: ui.DOMProps{
								Class: "w-full",
								Attrs: t.Attributes{"type": "button"},
							},
						})
					}
				}
			</div>
		</div>
	}
}`
	case "login-04":
		return `templ Login04Preview() {
	@PreviewLayout("Login 04 Preview") {
		<div class="flex min-h-svh bg-background">
			<div class="hidden lg:flex lg:w-1/2 items-center justify-center bg-muted/50 p-8">
				<div class="max-w-md space-y-4">
					<div class="flex items-center gap-3">
						<span class="flex size-10 items-center justify-center rounded-lg bg-primary text-primary-foreground text-lg font-semibold">S</span>
						<span class="text-2xl font-semibold">shadcn/ui</span>
					</div>
					<h2 class="text-3xl font-bold">Beautiful components built with Radix UI and Tailwind CSS.</h2>
					<p class="text-muted-foreground">A set of accessible, customizable, and themeable components that you can copy and paste into your apps.</p>
				</div>
			</div>
			<div class="flex w-full lg:w-1/2 items-center justify-center p-8">
				<div class="w-full max-w-sm">
					<div class="mb-8 lg:hidden">
						<div class="flex items-center gap-2">
							<span class="flex size-8 items-center justify-center rounded-lg bg-primary text-primary-foreground text-sm font-semibold">S</span>
							<span class="text-xl font-semibold">shadcn/ui</span>
						</div>
					</div>
					@ui.Card(ui.DOMProps{}) {
						@ui.CardHeader(ui.DOMProps{Class: "space-y-1"}) {
							@ui.CardTitle(ui.DOMProps{}) { <span>Sign in</span> }
							@ui.CardDescription(ui.DOMProps{}) { <span>Enter your email and password to sign in</span> }
						}
						@ui.CardContent(ui.DOMProps{Class: "space-y-4"}) {
							<div class="space-y-2">
								@ui.Label(ui.LabelProps{For: "email-04"}) { <span>Email</span> }
								@ui.Input(ui.InputProps{
									DOMProps: ui.DOMProps{ID: "email-04"},
									Type: "email",
									Placeholder: "name@example.com",
								})
							</div>
							<div class="space-y-2">
								@ui.Label(ui.LabelProps{For: "password-04"}) { <span>Password</span> }
								@ui.Input(ui.InputProps{
									DOMProps: ui.DOMProps{ID: "password-04"},
									Type: "password",
									Placeholder: "••••••••",
								})
							</div>
						@ui.Button(ui.ButtonProps{
							Label:   "Sign in",
							Variant: ui.ButtonVariantDefault,
							Size:    ui.ButtonSizeDefault,
							DOMProps: ui.DOMProps{
								Class: "w-full",
								Attrs: t.Attributes{"type": "button"},
							},
						})
						}
					}
				</div>
			</div>
		</div>
	}
}`
	case "signup-01":
		return `templ Signup01Preview() {
	@PreviewLayout("Signup 01 Preview") {
		<div class="grid min-h-screen lg:grid-cols-2">
			<div class="hidden bg-muted/40 lg:flex lg:flex-col lg:justify-between p-10">
				<div class="space-y-4">
					@ui.Badge(ui.BadgeProps{Label: "New", Variant: ui.BadgeVariantSecondary})
					<h1 class="text-4xl font-semibold tracking-tight">Build your account faster.</h1>
					<p class="max-w-md text-sm leading-6 text-muted-foreground">A cover-image signup layout with a larger marketing panel and a focused form surface.</p>
				</div>
				<div class="rounded-2xl border border-border/70 bg-background p-4 shadow-sm">
					<p class="text-sm font-medium">Trusted by teams shipping Go docs.</p>
					<p class="mt-1 text-sm text-muted-foreground">Teams move from docs to production faster with copyable source and reusable layouts.</p>
				</div>
			</div>
			<div class="flex items-center justify-center p-6">
				@ui.Card(ui.DOMProps{Class: "w-full max-w-md"}) {
					@ui.CardHeader(ui.DOMProps{Class: "space-y-2"}) {
						@ui.CardTitle(ui.DOMProps{}) { <span>Create your account</span> }
						@ui.CardDescription(ui.DOMProps{}) { <span>Sign up with email and password.</span> }
					}
					@ui.CardContent(ui.DOMProps{}) {
						<div class="grid gap-4">
							<div class="grid gap-2">
								@ui.Label(ui.LabelProps{For: "signup-name"}) { <span>Full name</span> }
								@ui.Input(ui.InputProps{Type: "text", Placeholder: "Jane Doe", DOMProps: ui.DOMProps{Attrs: t.Attributes{"id": "signup-name"}}})
							</div>
							<div class="grid gap-2">
								@ui.Label(ui.LabelProps{For: "signup-email"}) { <span>Email</span> }
								@ui.Input(ui.InputProps{Type: "email", Placeholder: "jane@example.com", DOMProps: ui.DOMProps{Attrs: t.Attributes{"id": "signup-email"}}})
							</div>
							<div class="grid gap-2">
								@ui.Label(ui.LabelProps{For: "signup-password"}) { <span>Password</span> }
								@ui.Input(ui.InputProps{Type: "password", Placeholder: "Create a password", DOMProps: ui.DOMProps{Attrs: t.Attributes{"id": "signup-password"}}})
							</div>
							<div class="flex items-center gap-2">
								@ui.Checkbox(ui.CheckboxProps{Checked: true, DOMProps: ui.DOMProps{Attrs: t.Attributes{"id": "signup-terms"}}})
								<label for="signup-terms" class="text-sm text-muted-foreground">I agree to the terms.</label>
							</div>
							@ui.Button(ui.ButtonProps{Label: "Create account"})
						</div>
					}
				}
			</div>
		</div>
	}
}`
	case "signup-02":
		return `templ Signup02Preview() {
	@PreviewLayout("Signup 02 Preview") {
		<div class="grid min-h-screen place-items-center bg-background p-6">
			@ui.Card(ui.DOMProps{Class: "w-full max-w-lg"}) {
				@ui.CardHeader(ui.DOMProps{Class: "space-y-2 text-center"}) {
					@ui.CardTitle(ui.DOMProps{}) { <span>Join the workspace</span> }
					@ui.CardDescription(ui.DOMProps{}) { <span>Use a provider or your email to register.</span> }
				}
				@ui.CardContent(ui.DOMProps{}) {
					<div class="grid gap-4">
						<div class="grid gap-3 sm:grid-cols-2">
							@ui.Button(ui.ButtonProps{Label: "Continue with GitHub", Variant: ui.ButtonVariantOutline})
							@ui.Button(ui.ButtonProps{Label: "Continue with Google", Variant: ui.ButtonVariantOutline})
						</div>
						<div class="flex items-center gap-3">
							@ui.Separator(ui.SeparatorProps{Decorative: true})
							<span class="text-xs uppercase tracking-[0.24em] text-muted-foreground">or</span>
							@ui.Separator(ui.SeparatorProps{Decorative: true})
						</div>
						<div class="grid gap-2">
							@ui.Label(ui.LabelProps{For: "signup2-email"}) { <span>Email</span> }
							@ui.Input(ui.InputProps{Type: "email", Placeholder: "name@company.com", DOMProps: ui.DOMProps{Attrs: t.Attributes{"id": "signup2-email"}}})
						</div>
						<div class="grid gap-2">
							@ui.Label(ui.LabelProps{For: "signup2-password"}) { <span>Password</span> }
							@ui.Input(ui.InputProps{Type: "password", Placeholder: "Create a password", DOMProps: ui.DOMProps{Attrs: t.Attributes{"id": "signup2-password"}}})
						</div>
						@ui.Button(ui.ButtonProps{Label: "Create workspace"})
						<p class="text-center text-xs text-muted-foreground">By continuing you accept the privacy policy and terms.</p>
					</div>
				}
			}
		</div>
	}
}`
	}
	return ""
}
func ChartSourceCode(slug string) string {
	switch slug {
	case "area-basic":
		return `func areaChartBasicPreview() templ.Component {
	return chartFrame("area", map[string]any{
		"title":      "Revenue over time",
		"stacked":    false,
		"showPoints": true,
		"series": []map[string]any{
			{"key": "revenue", "label": "Revenue"},
		},
		"data": []map[string]any{
			{"label": "Jan", "revenue": 68},
			{"label": "Feb", "revenue": 84},
			{"label": "Mar", "revenue": 76},
			{"label": "Apr", "revenue": 122},
			{"label": "May", "revenue": 108},
			{"label": "Jun", "revenue": 142},
			{"label": "Jul", "revenue": 156},
		},
	})
}`
	case "area-stacked":
		return `func areaChartStackedPreview() templ.Component {
	return chartFrame("area", map[string]any{
		"title":   "Stacked revenue and subscriptions",
		"stacked": true,
		"series": []map[string]any{
			{"key": "desktop", "label": "Desktop"},
			{"key": "mobile", "label": "Mobile"},
		},
		"data": []map[string]any{
			{"label": "Jan", "desktop": 42, "mobile": 24},
			{"label": "Feb", "desktop": 55, "mobile": 26},
			{"label": "Mar", "desktop": 63, "mobile": 31},
			{"label": "Apr", "desktop": 74, "mobile": 39},
			{"label": "May", "desktop": 71, "mobile": 42},
			{"label": "Jun", "desktop": 88, "mobile": 49},
			{"label": "Jul", "desktop": 96, "mobile": 57},
		},
	})
}`
	case "area-interactive":
		return `func areaChartInteractivePreview() templ.Component {
	return chartFrame("area", map[string]any{
		"title":         "Traffic trend with focus state",
		"stacked":       false,
		"showPoints":    true,
		"showLastPoint": true,
		"series": []map[string]any{
			{"key": "traffic", "label": "Traffic"},
		},
		"data": []map[string]any{
			{"label": "Mon", "traffic": 92},
			{"label": "Tue", "traffic": 116},
			{"label": "Wed", "traffic": 112},
			{"label": "Thu", "traffic": 138},
			{"label": "Fri", "traffic": 146},
			{"label": "Sat", "traffic": 162},
			{"label": "Sun", "traffic": 154},
		},
	})
}`
	case "bar-vertical":
		return `func barChartVerticalPreview() templ.Component {
	return chartFrame("bar", map[string]any{
		"orientation": "vertical",
		"series": []map[string]any{
			{"key": "sales", "label": "Sales"},
		},
		"data": []map[string]any{
			{"label": "Mon", "sales": 24},
			{"label": "Tue", "sales": 36},
			{"label": "Wed", "sales": 48},
			{"label": "Thu", "sales": 31},
			{"label": "Fri", "sales": 58},
			{"label": "Sat", "sales": 74},
		},
	})
}`
	case "bar-horizontal":
		return `func barChartHorizontalPreview() templ.Component {
	return chartFrame("bar", map[string]any{
		"orientation": "horizontal",
		"series": []map[string]any{
			{"key": "completion", "label": "Completion"},
		},
		"data": []map[string]any{
			{"label": "Marketing", "completion": 84},
			{"label": "Engineering", "completion": 92},
			{"label": "Sales", "completion": 66},
			{"label": "Support", "completion": 78},
			{"label": "Design", "completion": 58},
		},
	})
}`
	case "bar-multiple":
		return `func barChartMultiplePreview() templ.Component {
	return chartFrame("bar", map[string]any{
		"orientation": "vertical",
		"series": []map[string]any{
			{"key": "desktop", "label": "Desktop"},
			{"key": "mobile", "label": "Mobile"},
		},
		"data": []map[string]any{
			{"label": "Jan", "desktop": 24, "mobile": 14},
			{"label": "Feb", "desktop": 32, "mobile": 18},
			{"label": "Mar", "desktop": 40, "mobile": 22},
			{"label": "Apr", "desktop": 38, "mobile": 26},
			{"label": "May", "desktop": 50, "mobile": 30},
			{"label": "Jun", "desktop": 56, "mobile": 34},
		},
	})
}`
	case "line-linear":
		return `func lineChartLinearPreview() templ.Component {
	return chartFrame("line", map[string]any{
		"series": []map[string]any{
			{"key": "visits", "label": "Visits"},
		},
		"showPoints": true,
		"data": []map[string]any{
			{"label": "Jan", "visits": 92},
			{"label": "Feb", "visits": 114},
			{"label": "Mar", "visits": 104},
			{"label": "Apr", "visits": 136},
			{"label": "May", "visits": 122},
			{"label": "Jun", "visits": 158},
			{"label": "Jul", "visits": 174},
		},
	})
}`
	case "line-step":
		return `func lineChartStepPreview() templ.Component {
	return chartFrame("line", map[string]any{
		"series": []map[string]any{
			{"key": "sessions", "label": "Sessions"},
		},
		"step":       true,
		"showPoints": true,
		"data": []map[string]any{
			{"label": "Mon", "sessions": 8},
			{"label": "Tue", "sessions": 13},
			{"label": "Wed", "sessions": 13},
			{"label": "Thu", "sessions": 19},
			{"label": "Fri", "sessions": 19},
			{"label": "Sat", "sessions": 26},
			{"label": "Sun", "sessions": 31},
		},
	})
}`
	case "line-multiple":
		return `func lineChartMultiplePreview() templ.Component {
	return chartFrame("line", map[string]any{
		"series": []map[string]any{
			{"key": "desktop", "label": "Desktop"},
			{"key": "mobile", "label": "Mobile"},
		},
		"showPoints": true,
		"data": []map[string]any{
			{"label": "Jan", "desktop": 88, "mobile": 64},
			{"label": "Feb", "desktop": 96, "mobile": 70},
			{"label": "Mar", "desktop": 104, "mobile": 78},
			{"label": "Apr", "desktop": 118, "mobile": 82},
			{"label": "May", "desktop": 126, "mobile": 94},
			{"label": "Jun", "desktop": 144, "mobile": 104},
		},
	})
}`
	case "pie-basic":
		return `func pieChartBasicPreview() templ.Component {
	return chartFrame("pie", map[string]any{
		"innerRadius": 0.0,
		"data": []map[string]any{
			{"label": "Desktop", "value": 46},
			{"label": "Mobile", "value": 32},
			{"label": "Tablet", "value": 22},
		},
	})
}`
	case "pie-donut":
		return `func pieChartDonutPreview() templ.Component {
	return chartFrame("pie", map[string]any{
		"innerRadius": 0.62,
		"data": []map[string]any{
			{"label": "Desktop", "value": 54},
			{"label": "Mobile", "value": 28},
			{"label": "Tablet", "value": 18},
		},
	})
}`
	case "pie-label":
		return `func pieChartLabelPreview() templ.Component {
	return chartFrame("pie", map[string]any{
		"innerRadius":   0.68,
		"centerLabel":   "72%",
		"centerCaption": "Completion",
		"data": []map[string]any{
			{"label": "Complete", "value": 72},
			{"label": "Remaining", "value": 28},
		},
	})
}`
	case "radial-basic":
		return `func radialChartBasicPreview() templ.Component {
	return chartFrame("radial", map[string]any{
		"value":         64,
		"max":           100,
		"centerLabel":   "64%",
		"centerCaption": "Progress",
	})
}`
	case "radial-label":
		return `func radialChartLabelPreview() templ.Component {
	return chartFrame("radial", map[string]any{
		"value":         84,
		"max":           100,
		"centerLabel":   "84",
		"centerCaption": "Score",
	})
}`
	case "radial-shape":
		return `func radialChartShapePreview() templ.Component {
	return chartFrame("radial", map[string]any{
		"segments": []map[string]any{
			{"label": "Core", "value": 42},
			{"label": "Growth", "value": 26},
			{"label": "Ops", "value": 18},
			{"label": "R&D", "value": 14},
		},
		"centerLabel":   "100",
		"centerCaption": "Units",
	})
}`
	}
	return ""
}
