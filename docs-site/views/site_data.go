package views

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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
	Slug        string
	Title       string
	Description string
	Install     string
	Usage       string
	API         []string       // KEEP existing — deprecated but keep for compat
	Example     []string       // KEEP existing — deprecated but keep for compat
	GoUsage     string         // NEW: Go templ usage snippet
	APIProps    []APIProp      // NEW: structured API props
	Examples    []ExampleEntry // NEW: named examples with descriptions
	SourceCode  string
	Composition []string
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
		Install: "templcn add accordion",
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
		Install: "templcn add alert",
		GoUsage: `@ui.Alert(ui.AlertProps{}) {
  @ui.AlertTitle(ui.DOMProps{}) { Heads up! }
  @ui.AlertDescription(ui.DOMProps{}) { You can add components to your app using the CLI. }
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
		Slug: "alert-dialog", Title: "Alert Dialog", Description: "A modal dialog that interrupts the user with important content and expects a response.",
		Install: "templcn add alert-dialog",
		GoUsage: `@ui.AlertDialog(ui.AlertDialogProps{}) {
  @ui.AlertDialogTrigger(ui.DOMProps{}) {
    @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline}) { Open }
  }
  @ui.AlertDialogContent(ui.DOMProps{}) {
    @ui.AlertDialogHeader(ui.DOMProps{}) {
      @ui.AlertDialogTitle(ui.DOMProps{}) { Are you absolutely sure? }
      @ui.AlertDialogDescription(ui.DOMProps{}) {
        This action cannot be undone. This will permanently delete your account
        and remove your data from our servers.
      }
    }
    @ui.AlertDialogFooter(ui.DOMProps{}) {
      @ui.AlertDialogCancel(ui.DOMProps{}) { Cancel }
      @ui.AlertDialogAction(ui.DOMProps{}) { Continue }
    }
  }
}`,
		APIProps: []APIProp{
			{Name: "Open", Type: "bool", Default: "false", Desc: "Controls the open state of the dialog."},
			{Name: "Modal", Type: "bool", Default: "true", Desc: "Whether clicks outside close the dialog."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Alert dialog with confirmation and cancellation actions.", GoCode: `@ui.AlertDialog(ui.AlertDialogProps{}) {
  @ui.AlertDialogTrigger(ui.DOMProps{}) {
    @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline}) { Show Dialog }
  }
  @ui.AlertDialogContent(ui.DOMProps{}) {
    @ui.AlertDialogHeader(ui.DOMProps{}) {
      @ui.AlertDialogTitle(ui.DOMProps{}) { Are you absolutely sure? }
      @ui.AlertDialogDescription(ui.DOMProps{}) { This action cannot be undone. }
    }
    @ui.AlertDialogFooter(ui.DOMProps{}) {
      @ui.AlertDialogCancel(ui.DOMProps{}) { Cancel }
      @ui.AlertDialogAction(ui.DOMProps{}) { Continue }
    }
  }
}`, Preview: alertDialogPreview()},
		},
	},
	{
		Slug: "aspect-ratio", Title: "Aspect Ratio", Description: "Displays content within a desired ratio.",
		Install: "templcn add aspect-ratio",
		GoUsage: `@ui.AspectRatio(ui.AspectRatioProps{Ratio: 16.0 / 9.0, DOMProps: ui.DOMProps{Class: "bg-muted rounded-md overflow-hidden"}}) {
  <img src="https://images.unsplash.com/photo-1588345921523-c2dcdb7f1dcd?w=800&dpr=2&q=80" alt="Photo by Drew Beamer" class="h-full w-full object-cover"/>
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
		Slug: "avatar", Title: "Avatar", Description: "An image element with a fallback for representing the user.",
		Install: "templcn add avatar",
		GoUsage: `@ui.Avatar(ui.AvatarProps{}) {
  @ui.AvatarImage(ui.AvatarImageProps{Src: "https://github.com/shadcn.png", Alt: "@shadcn"})
  @ui.AvatarFallback(ui.DOMProps{}) { CN }
}`,
		APIProps: []APIProp{
			{Name: "Size", Type: "string", Default: `"default"`, Desc: `"sm", "default", or "lg".`},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Avatar with image and initials fallback.", GoCode: `@ui.Avatar(ui.AvatarProps{}) {
  @ui.AvatarImage(ui.AvatarImageProps{Src: "https://github.com/shadcn.png", Alt: "@shadcn"})
  @ui.AvatarFallback(ui.DOMProps{}) { CN }
}`, Preview: avatarPreview()},
		},
	},
	{
		Slug: "badge", Title: "Badge", Description: "Displays a badge or a component that looks like a badge.",
		Install: "templcn add badge",
		GoUsage: `@ui.Badge(ui.BadgeProps{Label: "Badge"})`,
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
		Slug: "breadcrumb", Title: "Breadcrumb", Description: "Displays the path to the current resource using a hierarchy of links.",
		Install: "templcn add breadcrumb",
		GoUsage: `@ui.Breadcrumb(ui.DOMProps{}) {
  @ui.BreadcrumbList(ui.DOMProps{}) {
    @ui.BreadcrumbItem(ui.DOMProps{}) {
      @ui.BreadcrumbLink(ui.BreadcrumbLinkProps{Href: "/"}) { Home }
    }
    @ui.BreadcrumbSeparator(ui.DOMProps{})
    @ui.BreadcrumbItem(ui.DOMProps{}) {
      @ui.BreadcrumbLink(ui.BreadcrumbLinkProps{Href: "/components"}) { Components }
    }
    @ui.BreadcrumbSeparator(ui.DOMProps{})
    @ui.BreadcrumbItem(ui.DOMProps{}) {
      @ui.BreadcrumbPage(ui.DOMProps{}) { Breadcrumb }
    }
  }
}`,
		APIProps: []APIProp{
			{Name: "BreadcrumbLink.Href", Type: "string", Default: `""`, Desc: "The navigation target."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "A three-level breadcrumb navigation trail.", GoCode: `@ui.Breadcrumb(ui.DOMProps{}) {
  @ui.BreadcrumbList(ui.DOMProps{}) {
    @ui.BreadcrumbItem(ui.DOMProps{}) {
      @ui.BreadcrumbLink(ui.BreadcrumbLinkProps{Href: "/"}) { Home }
    }
    @ui.BreadcrumbSeparator(ui.DOMProps{})
    @ui.BreadcrumbItem(ui.DOMProps{}) {
      @ui.BreadcrumbPage(ui.DOMProps{}) { Breadcrumb }
    }
  }
}`, Preview: BreadcrumbPreview()},
		},
	},
	{
		Slug: "button", Title: "Button", Description: "Displays a button or a component that looks like a button.",
		Install: "templcn add button",
		GoUsage: `@ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline}) {
  Button
}`,
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
		Install: "templcn add button-group",
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
		Slug: "calendar", Title: "Calendar", Description: "A date field component that allows users to enter and edit dates.",
		Install: "templcn add calendar",
		GoUsage: `@ui.Calendar(ui.CalendarProps{
  DOMProps: ui.DOMProps{Class: "rounded-md border shadow"},
  Mode: "single",
  DefaultMonth: "2026-05",
  Selected: "2026-05-07",
})`,
		APIProps: []APIProp{
			{Name: "Mode", Type: "string", Default: `"single"`, Desc: `"single", "multiple", or "range".`},
			{Name: "CaptionLayout", Type: "string", Default: `"label"`, Desc: `"label" or "dropdown" month navigation.`},
			{Name: "ShowWeekNumber", Type: "bool", Default: `false`, Desc: "Shows ISO week numbers."},
			{Name: "HideOutsideDays", Type: "bool", Default: `false`, Desc: "Hides days outside the current month."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Single date selection calendar.", GoCode: `@ui.Calendar(ui.CalendarProps{
  DOMProps: ui.DOMProps{Class: "rounded-md border shadow"},
  Mode: "single",
  DefaultMonth: "2026-05",
  Selected: "2026-05-07",
})`, Preview: calendarPreview()},
		},
	},
	{
		Slug: "card", Title: "Card", Description: "Displays a card with header, content, and footer.",
		Install: "templcn add card",
		GoUsage: `@ui.Card(ui.DOMProps{Class: "w-[350px]"}) {
  @ui.CardHeader(ui.DOMProps{}) {
    @ui.CardTitle(ui.DOMProps{}) { Create project }
    @ui.CardDescription(ui.DOMProps{}) { Deploy your new project in one-click. }
  }
  @ui.CardContent(ui.DOMProps{}) {
    <p class="text-sm text-muted-foreground">Your project will be deployed immediately.</p>
  }
  @ui.CardFooter(ui.DOMProps{Class: "flex justify-between"}) {
    @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline, Label: "Cancel"})
    @ui.Button(ui.ButtonProps{Label: "Deploy"})
  }
}`,
		APIProps: []APIProp{
			{Name: "DOMProps.Class", Type: "string", Default: `""`, Desc: "Additional CSS classes applied to the card root."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "A card with header, content, and footer actions.", GoCode: `@ui.Card(ui.DOMProps{Class: "w-80"}) {
  @ui.CardHeader(ui.DOMProps{}) {
    @ui.CardTitle(ui.DOMProps{}) { <span>Create project</span> }
  }
  @ui.CardContent(ui.DOMProps{}) { <span>Deploy your app with one command.</span> }
  @ui.CardFooter(ui.DOMProps{}) {
    @ui.Button(ui.ButtonProps{Label: "Deploy"})
  }
}`, Preview: CardPreview()},
		},
	},
	{
		Slug: "carousel", Title: "Carousel", Description: "A carousel with motion and swipe built using native scroll behavior.",
		Install: "templcn add carousel",
		GoUsage: `@ui.Carousel(ui.CarouselProps{DOMProps: ui.DOMProps{Class: "w-full max-w-xs"}}) {
  @ui.CarouselContent(ui.DOMProps{}) {
    @ui.CarouselItem(ui.DOMProps{}) {
      <div class="p-1">
        @ui.Card(ui.DOMProps{}) {
          @ui.CardContent(ui.DOMProps{Class: "flex aspect-square items-center justify-center p-6"}) {
            <span class="text-4xl font-semibold">1</span>
          }
        }
      </div>
    }
  }
  @ui.CarouselPrevious(ui.DOMProps{})
  @ui.CarouselNext(ui.DOMProps{})
}`,
		APIProps: []APIProp{
			{Name: "Orientation", Type: "string", Default: `"horizontal"`, Desc: "Slide direction."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "A swipeable slide carousel.", GoCode: `@ui.Carousel(ui.CarouselProps{DOMProps: ui.DOMProps{Class: "w-full max-w-xs"}}) {
  @ui.CarouselContent(ui.DOMProps{}) {
    @ui.CarouselItem(ui.DOMProps{}) { <div class="grid h-28 place-items-center rounded-md border bg-muted/30">Slide 1</div> }
  }
  @ui.CarouselPrevious(ui.DOMProps{})
  @ui.CarouselNext(ui.DOMProps{})
}`, Preview: carouselPreview()},
		},
	},
	{
		Slug: "chart", Title: "Chart", Description: "Beautiful charts built with TanStack Charts, server rendering, and SVG marks.",
		Install: "templcn add chart",
		GoUsage: `@ui.ChartContainer(ui.ChartContainerProps{
  Engine: "tanstack",
  Data: ` + "`" + `{"type":"line","data":[{"label":"Jan","revenue":42},{"label":"Feb","revenue":58}],"series":[{"key":"revenue","label":"Revenue"}]}` + "`" + `,
  Config: ` + "`" + `{"revenue":{"label":"Revenue","color":"var(--chart-1)"}}` + "`" + `,
  InitialHeight: 240,
})`,
		APIProps: []APIProp{
			{Name: "Engine", Type: "string", Default: `"tanstack"`, Desc: "Charting engine identifier."},
			{Name: "Data", Type: "string", Default: `""`, Desc: "JSON data payload."},
			{Name: "Config", Type: "string", Default: `""`, Desc: "Series colors and labels configuration."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Interactive chart container.", GoCode: `@ui.ChartContainer(ui.ChartContainerProps{
  Engine: "tanstack",
  Data: ` + "`" + `{"type":"line","data":[{"label":"Jan","revenue":42}],"series":[{"key":"revenue","label":"Revenue"}]}` + "`" + `,
  InitialHeight: 240,
})`, Preview: chartPreview()},
		},
	},
	{
		Slug: "checkbox", Title: "Checkbox", Description: "A control that allows the user to toggle between checked and not checked.",
		Install: "templcn add checkbox",
		GoUsage: `<div class="flex items-center space-x-2">
  @ui.Checkbox(ui.CheckboxProps{ID: "terms", Name: "terms"})
  @ui.Label(ui.LabelProps{For: "terms"}) { Accept terms and conditions }
</div>`,
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
		Slug: "collapsible", Title: "Collapsible", Description: "An interactive component which expands/collapses a panel.",
		Install: "templcn add collapsible",
		GoUsage: `@ui.Collapsible(ui.CollapsibleProps{DOMProps: ui.DOMProps{Class: "w-[350px] space-y-2"}}) {
  <div class="flex items-center justify-between space-x-4 px-4">
    <h4 class="text-sm font-semibold">@peduarte starred 3 repositories</h4>
    @ui.CollapsibleTrigger(ui.DOMProps{}) {
      @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantGhost, Size: ui.ButtonSizeSM}) { Toggle }
    }
  </div>
  @ui.CollapsibleContent(ui.DOMProps{Class: "space-y-2"}) {
    <div class="rounded-md border px-4 py-2 font-mono text-sm shadow-sm">
      @radix-ui/primitives
    </div>
  }
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
		Slug: "combobox", Title: "Combobox", Description: "Autocomplete input and command palette with a list of suggestions.",
		Install: "templcn add combobox",
		GoUsage: `@ui.Combobox(ui.ComboboxProps{Name: "framework", DefaultValue: "templ"}) {
  @ui.ComboboxTrigger(ui.SelectTriggerProps{DOMProps: ui.DOMProps{Class: "w-[200px]"}}) {
    @ui.ComboboxValue(ui.DOMProps{}) { Select framework... }
  }
  @ui.ComboboxContent(ui.SelectContentProps{DOMProps: ui.DOMProps{Class: "w-[200px] p-0"}}) {
    @ui.ComboboxInput(ui.InputProps{Placeholder: "Search framework..."})
    @ui.ComboboxList(ui.DOMProps{}) {
      @ui.ComboboxItem(ui.DropdownMenuItemProps{Value: "next"}) { Next.js }
      @ui.ComboboxItem(ui.DropdownMenuItemProps{Value: "templ"}) { templ }
      @ui.ComboboxItem(ui.DropdownMenuItemProps{Value: "go"}) { Go }
    }
  }
}`,
		APIProps: []APIProp{
			{Name: "Multiple", Type: "bool", Default: "false", Desc: "Allow multiple selections."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "A searchable combobox.", GoCode: `@ui.Combobox(ui.ComboboxProps{Name: "framework"}) {
  @ui.ComboboxTrigger(ui.SelectTriggerProps{}) {
    @ui.ComboboxValue(ui.DOMProps{}) { Select framework }
  }
  @ui.ComboboxContent(ui.SelectContentProps{}) {
    @ui.ComboboxInput(ui.InputProps{Placeholder: "Search..."})
    @ui.ComboboxList(ui.DOMProps{}) {
      @ui.ComboboxItem(ui.DropdownMenuItemProps{Value: "templ"}) { templ }
    }
  }
}`, Preview: examplePreview("Combobox")},
		},
	},
	{
		Slug: "command", Title: "Command", Description: "Fast, composable, unstyled command menu for Go and templ.",
		Install: "templcn add command",
		GoUsage: `@ui.Command(ui.CommandProps{DOMProps: ui.DOMProps{Class: "rounded-lg border shadow-md md:min-w-[450px]"}}) {
  @ui.CommandInput(ui.InputProps{Placeholder: "Type a command or search..."})
  @ui.CommandList(ui.DOMProps{}) {
    @ui.CommandEmpty(ui.DOMProps{}) { No results found. }
    @ui.CommandGroup(ui.DOMProps{}) {
      @ui.CommandItem(ui.DropdownMenuItemProps{Value: "calendar"}) { Calendar }
      @ui.CommandItem(ui.DropdownMenuItemProps{Value: "search-emoji"}) { Search Emoji }
      @ui.CommandItem(ui.DropdownMenuItemProps{Value: "calculator"}) { Calculator }
    }
  }
}`,
		APIProps: []APIProp{
			{Name: "DefaultValue", Type: "string", Default: `""`, Desc: "Default search value."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Command palette with input and item groups.", GoCode: `@ui.Command(ui.CommandProps{}) {
  @ui.CommandInput(ui.InputProps{Placeholder: "Search commands"})
  @ui.CommandList(ui.DOMProps{}) {
    @ui.CommandGroup(ui.DOMProps{}) {
      @ui.CommandItem(ui.DropdownMenuItemProps{Value: "settings"}) { Settings }
    }
  }
}`, Preview: commandPreview()},
		},
	},
	{
		Slug: "context-menu", Title: "Context Menu", Description: "Displays a menu to the user — triggered by a right-click or long-press.",
		Install: "templcn add context-menu",
		GoUsage: `@ui.ContextMenu(ui.ContextMenuProps{}) {
  @ui.ContextMenuTrigger(ui.DOMProps{Class: "flex h-[150px] w-[300px] items-center justify-center rounded-md border border-dashed text-sm"}) {
    Right click here
  }
  @ui.ContextMenuContent(ui.DOMProps{Class: "w-64"}) {
    @ui.ContextMenuItem(ui.DropdownMenuItemProps{Value: "back"}) { Back }
    @ui.ContextMenuItem(ui.DropdownMenuItemProps{Value: "forward", Disabled: true}) { Forward }
    @ui.ContextMenuItem(ui.DropdownMenuItemProps{Value: "reload"}) { Reload }
  }
}`,
		APIProps: []APIProp{
			{Name: "Variant", Type: "string", Default: `""`, Desc: "Item variant."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Right click trigger area with contextual options.", GoCode: `@ui.ContextMenu(ui.ContextMenuProps{}) {
  @ui.ContextMenuTrigger(ui.DOMProps{}) { Right click here }
  @ui.ContextMenuContent(ui.DOMProps{}) {
    @ui.ContextMenuItem(ui.DropdownMenuItemProps{Value: "reload"}) { Reload }
  }
}`, Preview: examplePreview("Context Menu")},
		},
	},
	{
		Slug: "data-table", Title: "Data Table", Description: "Powerful table and datagrids built using Go rendering and TanStack Table.",
		Install: "templcn add data-table",
		GoUsage: `@ui.DataTable(ui.DataTableProps{
  Engine: "tanstack",
  Columns: []ui.DataTableColumn{
    {Key: "name", Header: "Name", Sortable: true},
    {Key: "status", Header: "Status"},
    {Key: "email", Header: "Email"},
    {Key: "amount", Header: "Amount"},
  },
  Rows: []ui.DataTableRow{
    {ID: "1", Cells: []ui.DataTableCell{{Value: "Ada Lovelace"}, {Value: "Success"}, {Value: "ada@example.com"}, {Value: "$250.00"}}},
  },
})`,
		APIProps: []APIProp{
			{Name: "Columns", Type: "[]DataTableColumn", Default: "nil", Desc: "Column definitions."},
			{Name: "Rows", Type: "[]DataTableRow", Default: "nil", Desc: "Table row data."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Interactive sortable data table.", GoCode: `@ui.DataTable(ui.DataTableProps{
  Engine: "tanstack",
  Columns: []ui.DataTableColumn{
    {Key: "name", Header: "Name", Sortable: true},
    {Key: "status", Header: "Status"},
  },
  Rows: []ui.DataTableRow{
    {ID: "1", Cells: []ui.DataTableCell{{Value: "Ada"}, {Value: "Active"}}},
  },
})`, Preview: dataTablePreview()},
		},
	},
	{
		Slug: "date-picker", Title: "Date Picker", Description: "A date picker component with range and presets.",
		Install: "templcn add date-picker",
		GoUsage: `@ui.DatePicker(ui.DatePickerProps{Name: "date", Value: "2026-05-08"}) {
  @ui.Calendar(ui.CalendarProps{
    DOMProps: ui.DOMProps{Class: "rounded-md border"},
    Mode: "single",
    Selected: "2026-05-08",
  })
}`,
		APIProps: []APIProp{
			{Name: "Mode", Type: "string", Default: `"single"`, Desc: `"single" or "range".`},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "A date picker input with embedded calendar popover.", GoCode: `@ui.DatePicker(ui.DatePickerProps{Name: "date", Value: "2026-05-08"})`, Preview: datePickerPreview()},
		},
	},
	{
		Slug: "dialog", Title: "Dialog", Description: "A window overlaid on either the primary window or another dialog window.",
		Install: "templcn add dialog",
		GoUsage: `@ui.Dialog(ui.DialogProps{}) {
  @ui.DialogTrigger(ui.DOMProps{}) {
    @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline}) { Edit Profile }
  }
  @ui.DialogContent(ui.DOMProps{Class: "sm:max-w-[425px]"}) {
    @ui.DialogHeader(ui.DOMProps{}) {
      @ui.DialogTitle(ui.DOMProps{}) { Edit profile }
      @ui.DialogDescription(ui.DOMProps{}) {
        Make changes to your profile here. Click save when you're done.
      }
    }
    <div class="grid gap-4 py-4">
      <div class="grid grid-cols-4 items-center gap-4">
        @ui.Label(ui.LabelProps{For: "name", DOMProps: ui.DOMProps{Class: "text-right"}}) { Name }
        @ui.Input(ui.InputProps{ID: "name", Value: "Pedro Duarte", DOMProps: ui.DOMProps{Class: "col-span-3"}})
      </div>
    </div>
    @ui.DialogFooter(ui.DOMProps{}) {
      @ui.Button(ui.ButtonProps{Type: "submit", Label: "Save changes"})
    }
  }
}`,
		APIProps: []APIProp{
			{Name: "Open", Type: "bool", Default: "false", Desc: "Controls dialog visibility."},
			{Name: "Modal", Type: "bool", Default: "true", Desc: "Traps focus inside the dialog."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Modal dialog with form inputs and trigger.", GoCode: `@ui.Dialog(ui.DialogProps{}) {
  @ui.DialogTrigger(ui.DOMProps{}) { @ui.Button(ui.ButtonProps{Label: "Open Dialog"}) }
  @ui.DialogContent(ui.DOMProps{}) {
    @ui.DialogHeader(ui.DOMProps{}) {
      @ui.DialogTitle(ui.DOMProps{}) { Dialog Title }
    }
  }
}`, Preview: dialogPreview()},
		},
	},
	{
		Slug: "direction", Title: "Direction", Description: "Sets the reading direction of content (LTR or RTL).",
		Install: "templcn add direction",
		GoUsage: `@ui.DirectionProvider(ui.DirectionProviderProps{Direction: "rtl"}) {
  @ui.Card(ui.DOMProps{}) {
    @ui.CardContent(ui.DOMProps{Class: "p-6"}) {
      <p class="text-sm">RTL layout content with mirrored text direction and alignment.</p>
    }
  }
}`,
		APIProps: []APIProp{
			{Name: "Direction", Type: "string", Default: `"ltr"`, Desc: `"ltr" or "rtl".`},
		},
		Examples: []ExampleEntry{
			{Name: "RTL", Desc: "Right-to-left layout direction wrapper.", GoCode: `@ui.DirectionProvider(ui.DirectionProviderProps{Direction: "rtl"}) {
  @ui.Card(ui.DOMProps{}) {
    @ui.CardContent(ui.DOMProps{}) { RTL content }
  }
}`, Preview: examplePreview("Direction Provider")},
		},
	},
	{
		Slug: "drawer", Title: "Drawer", Description: "A drawer component for mobile-friendly bottom/side slide-over surfaces.",
		Install: "templcn add drawer",
		GoUsage: `@ui.Drawer(ui.DrawerProps{SheetProps: ui.SheetProps{Side: "bottom"}}) {
  @ui.DrawerTrigger(ui.DOMProps{}) {
    @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline}) { Open Drawer }
  }
  @ui.DrawerContent(ui.DOMProps{}) {
    @ui.DrawerHeader(ui.DOMProps{}) {
      @ui.DrawerTitle(ui.DOMProps{}) { Move Goal }
      @ui.DrawerDescription(ui.DOMProps{}) { Set your daily activity goal. }
    }
    @ui.DrawerFooter(ui.DOMProps{}) {
      @ui.Button(ui.ButtonProps{Label: "Submit"})
      @ui.DrawerClose(ui.DOMProps{}) {
        @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline, Label: "Cancel"})
      }
    }
  }
}`,
		APIProps: []APIProp{
			{Name: "Side", Type: "string", Default: `"bottom"`, Desc: `"top", "bottom", "left", "right".`},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Bottom slide-over drawer surface.", GoCode: `@ui.Drawer(ui.DrawerProps{SheetProps: ui.SheetProps{Side: "bottom"}}) {
  @ui.DrawerTrigger(ui.DOMProps{}) { @ui.Button(ui.ButtonProps{Label: "Open Drawer"}) }
  @ui.DrawerContent(ui.DOMProps{}) {
    @ui.DrawerHeader(ui.DOMProps{}) {
      @ui.DrawerTitle(ui.DOMProps{}) { Drawer Title }
    }
  }
}`, Preview: drawerPreview()},
		},
	},
	{
		Slug: "dropdown-menu", Title: "Dropdown Menu", Description: "Displays a menu to the user — such as a set of actions or functions — triggered by a button.",
		Install: "templcn add dropdown-menu",
		GoUsage: `@ui.DropdownMenu(ui.DropdownMenuProps{}) {
  @ui.DropdownMenuTrigger(ui.DOMProps{}) {
    @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline}) { Open }
  }
  @ui.DropdownMenuContent(ui.DOMProps{Class: "w-56"}) {
    @ui.DropdownMenuLabel(ui.DOMProps{}) { My Account }
    @ui.DropdownMenuSeparator(ui.DOMProps{})
    @ui.DropdownMenuItem(ui.DropdownMenuItemProps{}) { Profile }
    @ui.DropdownMenuItem(ui.DropdownMenuItemProps{}) { Billing }
    @ui.DropdownMenuItem(ui.DropdownMenuItemProps{}) { Team }
    @ui.DropdownMenuItem(ui.DropdownMenuItemProps{}) { Subscription }
  }
}`,
		APIProps: []APIProp{
			{Name: "Side", Type: "string", Default: `"bottom"`, Desc: "Menu placement side."},
			{Name: "Align", Type: "string", Default: `"start"`, Desc: "Menu alignment."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Dropdown menu with labels, separators, and actions.", GoCode: `@ui.DropdownMenu(ui.DropdownMenuProps{}) {
  @ui.DropdownMenuTrigger(ui.DOMProps{}) { @ui.Button(ui.ButtonProps{Label: "Open"}) }
  @ui.DropdownMenuContent(ui.DOMProps{}) {
    @ui.DropdownMenuItem(ui.DropdownMenuItemProps{}) { Profile }
  }
}`, Preview: dropdownMenuPreview()},
		},
	},
	{
		Slug: "empty", Title: "Empty", Description: "A placeholder helper for empty datasets, searches, and lists.",
		Install: "templcn add empty",
		GoUsage: `@ui.Empty(ui.DOMProps{Class: "p-8 text-center"}) {
  @ui.EmptyTitle(ui.DOMProps{}) { No components installed }
  @ui.EmptyDescription(ui.DOMProps{}) { Run templcn add to install your first component. }
}`,
		APIProps: []APIProp{},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Empty state card with title and instructions.", GoCode: `@ui.Empty(ui.DOMProps{}) {
  @ui.EmptyTitle(ui.DOMProps{}) { No components installed }
  @ui.EmptyDescription(ui.DOMProps{}) { Run templcn add button. }
}`, Preview: emptyPreview()},
		},
	},
	{
		Slug: "field", Title: "Field", Description: "A form field grouping component with label, description, and error slots.",
		Install: "templcn add field",
		GoUsage: `@ui.Field(ui.DOMProps{}) {
  @ui.FieldLabel(ui.DOMProps{}) { Email }
  @ui.Input(ui.InputProps{Type: "email", Placeholder: "you@example.com"})
  @ui.FieldDescription(ui.DOMProps{}) { We'll never share your email with anyone else. }
  @ui.FieldError(ui.FieldErrorProps{Errors: []string{"Invalid email address."}})
}`,
		APIProps: []APIProp{
			{Name: "Orientation", Type: "FieldOrientation", Default: `"vertical"`, Desc: `"vertical", "horizontal", "responsive".`},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Form field with input, label, description, and validation error.", GoCode: `@ui.Field(ui.DOMProps{}) {
  @ui.FieldLabel(ui.DOMProps{}) { Email }
  @ui.Input(ui.InputProps{Placeholder: "you@example.com"})
}`, Preview: examplePreview("Field")},
		},
	},
	{
		Slug: "hover-card", Title: "Hover Card", Description: "For sighted users to preview content available behind a link.",
		Install: "templcn add hover-card",
		GoUsage: `@ui.HoverCard(ui.HoverCardProps{}) {
  @ui.HoverCardTrigger(ui.DOMProps{}) {
    @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantLink}) { @nextjs }
  }
  @ui.HoverCardContent(ui.DOMProps{Class: "w-80"}) {
    <div class="flex justify-between space-x-4">
      @ui.Avatar(ui.AvatarProps{}) {
        @ui.AvatarImage(ui.AvatarImageProps{Src: "https://github.com/vercel.png"})
        @ui.AvatarFallback(ui.DOMProps{}) { VC }
      }
      <div class="space-y-1">
        <h4 class="text-sm font-semibold">@nextjs</h4>
        <p class="text-sm">The React Framework – created and maintained by @vercel.</p>
      </div>
    </div>
  }
}`,
		APIProps: []APIProp{
			{Name: "OpenDelayMs", Type: "int", Default: "0", Desc: "Delay before showing."},
			{Name: "CloseDelayMs", Type: "int", Default: "0", Desc: "Delay before hiding."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Hover card displaying user bio on hover.", GoCode: `@ui.HoverCard(ui.HoverCardProps{}) {
  @ui.HoverCardTrigger(ui.DOMProps{}) { Hover me }
  @ui.HoverCardContent(ui.DOMProps{}) { Card content }
}`, Preview: hoverCardPreview()},
		},
	},
	{
		Slug: "input", Title: "Input", Description: "Displays a form input field or a component that looks like an input field.",
		Install: "templcn add input",
		GoUsage: `@ui.Input(ui.InputProps{Type: "email", Placeholder: "Email"})`,
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
		Slug: "input-group", Title: "Input Group", Description: "Input with leading/trailing addons and icon slots.",
		Install: "templcn add input-group",
		GoUsage: `@ui.InputGroup(ui.InputGroupProps{}) {
  @ui.InputGroupAddon(ui.InputGroupAddonProps{}) { https:// }
  @ui.InputGroupInput(ui.InputGroupInputProps{Placeholder: "example.com"})
}`,
		APIProps: []APIProp{
			{Name: "Align", Type: "InputGroupAddonAlign", Default: `"inline-start"`, Desc: "Position of the addon."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Input with a prefix addon.", GoCode: `@ui.InputGroup(ui.InputGroupProps{}) {
  @ui.InputGroupAddon(ui.InputGroupAddonProps{}) { https:// }
  @ui.InputGroupInput(ui.InputGroupInputProps{Placeholder: "example.com"})
}`, Preview: inputGroupPreview()},
		},
	},
	{
		Slug: "input-otp", Title: "Input OTP", Description: "Accessible one-time password component with copy paste functionality.",
		Install: "templcn add input-otp",
		GoUsage: `@ui.InputOTP(ui.InputOTPProps{MaxLength: 6}) {
  @ui.InputOTPGroup(ui.DOMProps{}) {
    @ui.InputOTPSlot(ui.InputOTPSlotProps{Index: 0})
    @ui.InputOTPSlot(ui.InputOTPSlotProps{Index: 1})
    @ui.InputOTPSlot(ui.InputOTPSlotProps{Index: 2})
  }
  @ui.InputOTPSeparator(ui.DOMProps{})
  @ui.InputOTPGroup(ui.DOMProps{}) {
    @ui.InputOTPSlot(ui.InputOTPSlotProps{Index: 3})
    @ui.InputOTPSlot(ui.InputOTPSlotProps{Index: 4})
    @ui.InputOTPSlot(ui.InputOTPSlotProps{Index: 5})
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
		Slug: "item", Title: "Item", Description: "A generic list/card item surface with header, body, and action slots.",
		Install: "templcn add item",
		GoUsage: `@ui.Item(ui.DOMProps{Class: "flex items-center justify-between p-4 border rounded-lg"}) {
  @ui.ItemHeader(ui.DOMProps{}) {
    <div class="font-medium">Component item</div>
    <div class="text-sm text-muted-foreground">button.go</div>
  }
  @ui.ItemActions(ui.DOMProps{}) {
    @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline, Size: ui.ButtonSizeSM, Label: "View"})
  }
}`,
		APIProps: []APIProp{},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Item with header and action button.", GoCode: `@ui.Item(ui.DOMProps{}) {
  @ui.ItemHeader(ui.DOMProps{}) { Title }
}`, Preview: itemPreview()},
		},
	},
	{
		Slug: "kbd", Title: "Kbd", Description: "A keyboard shortcut token styled as a key badge.",
		Install: "templcn add kbd",
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
		Slug: "label", Title: "Label", Description: "Renders an accessible label associated with controls.",
		Install: "templcn add label",
		GoUsage: `<div class="grid w-full max-w-sm items-center gap-1.5">
  @ui.Label(ui.LabelProps{For: "email"}) { Email }
  @ui.Input(ui.InputProps{ID: "email", Type: "email", Placeholder: "Email"})
</div>`,
		APIProps: []APIProp{
			{Name: "For", Type: "string", Default: `""`, Desc: "The id of the associated form element."},
		},
		Examples: []ExampleEntry{
			{Name: "With Input", Desc: "A label associated with an input field.", GoCode: `@ui.Label(ui.LabelProps{For: "email"}) { <span>Your email address</span> }
@ui.Input(ui.InputProps{ID: "email", Placeholder: "name@example.com"})`, Preview: LabelPreview()},
		},
	},
	{
		Slug: "menubar", Title: "Menubar", Description: "A visually persistent menu common in desktop applications that provides quick access to a consistent set of commands.",
		Install: "templcn add menubar",
		GoUsage: `@ui.Menubar(ui.MenubarProps{}) {
  @ui.MenubarMenu(ui.DOMProps{}) {
    @ui.MenubarTrigger(ui.DOMProps{}) { File }
    @ui.MenubarContent(ui.DOMProps{}) {
      @ui.MenubarItem(ui.DropdownMenuItemProps{}) { New Tab }
      @ui.MenubarItem(ui.DropdownMenuItemProps{}) { New Window }
      @ui.MenubarSeparator(ui.DOMProps{})
      @ui.MenubarItem(ui.DropdownMenuItemProps{}) { Share }
      @ui.MenubarSeparator(ui.DOMProps{})
      @ui.MenubarItem(ui.DropdownMenuItemProps{}) { Print }
    }
  }
}`,
		APIProps: []APIProp{},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Desktop-style menu bar.", GoCode: `@ui.Menubar(ui.MenubarProps{}) {
  @ui.MenubarMenu(ui.DOMProps{}) {
    @ui.MenubarTrigger(ui.DOMProps{}) { File }
    @ui.MenubarContent(ui.DOMProps{}) {
      @ui.MenubarItem(ui.DropdownMenuItemProps{}) { New Tab }
    }
  }
}`, Preview: menubarPreview()},
		},
	},
	{
		Slug: "native-select", Title: "Native Select", Description: "A styled native HTML select element.",
		Install: "templcn add native-select",
		GoUsage: `@ui.NativeSelect(ui.NativeSelectProps{Name: "framework"}) {
  @ui.NativeSelectOption(ui.DOMProps{}, "templ", "templ", false)
  @ui.NativeSelectOption(ui.DOMProps{}, "next", "Next.js", false)
  @ui.NativeSelectOption(ui.DOMProps{}, "svelte", "SvelteKit", false)
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
		Slug: "navigation-menu", Title: "Navigation Menu", Description: "A collection of links for navigating websites.",
		Install: "templcn add navigation-menu",
		GoUsage: `@ui.NavigationMenu(ui.NavigationMenuProps{}) {
  @ui.NavigationMenuList(ui.DOMProps{}) {
    @ui.NavigationMenuItem(ui.DOMProps{}) {
      @ui.NavigationMenuLink(ui.BreadcrumbLinkProps{Href: "/docs"}) {
        Documentation
      }
    }
    @ui.NavigationMenuItem(ui.DOMProps{}) {
      @ui.NavigationMenuLink(ui.BreadcrumbLinkProps{Href: "/components"}) {
        Components
      }
    }
  }
}`,
		APIProps: []APIProp{},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "A responsive navigation bar with links.", GoCode: `@ui.NavigationMenu(ui.NavigationMenuProps{}) {
  @ui.NavigationMenuList(ui.DOMProps{}) {
    @ui.NavigationMenuItem(ui.DOMProps{}) {
      @ui.NavigationMenuLink(ui.BreadcrumbLinkProps{Href: "/"}) { Home }
    }
  }
}`, Preview: navigationMenuPreview()},
		},
	},
	{
		Slug: "pagination", Title: "Pagination", Description: "Pagination with page navigation, next and previous links.",
		Install: "templcn add pagination",
		GoUsage: `@ui.Pagination(ui.DOMProps{}) {
  @ui.PaginationContent(ui.DOMProps{}) {
    @ui.PaginationItem(ui.DOMProps{}) {
      @ui.PaginationPrevious(ui.PaginationLinkProps{Href: "#"}) { Previous }
    }
    @ui.PaginationItem(ui.DOMProps{}) {
      @ui.PaginationLink(ui.PaginationLinkProps{Href: "#", IsActive: true}) { 1 }
    }
    @ui.PaginationItem(ui.DOMProps{}) {
      @ui.PaginationLink(ui.PaginationLinkProps{Href: "#"}) { 2 }
    }
    @ui.PaginationItem(ui.DOMProps{}) {
      @ui.PaginationEllipsis(ui.DOMProps{})
    }
    @ui.PaginationItem(ui.DOMProps{}) {
      @ui.PaginationNext(ui.PaginationLinkProps{Href: "#"}) { Next }
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
		Slug: "popover", Title: "Popover", Description: "Displays rich content in a portal, triggered by a button.",
		Install: "templcn add popover",
		GoUsage: `@ui.Popover(ui.PopoverProps{}) {
  @ui.PopoverTrigger(ui.DOMProps{}) {
    @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline}) { Open popover }
  }
  @ui.PopoverContent(ui.DOMProps{Class: "w-80"}) {
    <div class="grid gap-4">
      <div class="space-y-2">
        <h4 class="font-medium leading-none">Dimensions</h4>
        <p class="text-sm text-muted-foreground">Set the dimensions for the layer.</p>
      </div>
    </div>
  }
}`,
		APIProps: []APIProp{
			{Name: "Side", Type: "string", Default: `"bottom"`, Desc: "Preferred placement side."},
			{Name: "Align", Type: "string", Default: `"center"`, Desc: "Alignment along the axis."},
			{Name: "Modal", Type: "bool", Default: "false", Desc: "Whether to trap focus."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Anchored popover panel with custom content.", GoCode: `@ui.Popover(ui.PopoverProps{}) {
  @ui.PopoverTrigger(ui.DOMProps{}) { Open }
  @ui.PopoverContent(ui.DOMProps{}) { Popover content }
}`, Preview: popoverPreview()},
		},
	},
	{
		Slug: "progress", Title: "Progress", Description: "Displays an indicator showing the completion progress of a task, typically displayed as a progress bar.",
		Install: "templcn add progress",
		GoUsage: `@ui.Progress(ui.ProgressProps{Value: 60, DOMProps: ui.DOMProps{Class: "w-[60%]"}})`,
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
		Slug: "radio-group", Title: "Radio Group", Description: "A set of checkable buttons—known as radio buttons—where no more than one button can be checked at a time.",
		Install: "templcn add radio-group",
		GoUsage: `@ui.RadioGroup(ui.RadioGroupProps{Name: "option", DefaultValue: "option-one"}) {
  <div class="flex items-center space-x-2">
    @ui.RadioGroupItem(ui.RadioGroupItemProps{ID: "r1", Name: "option", Value: "option-one", Label: "Option One"})
    @ui.Label(ui.LabelProps{For: "r1"}) { Option One }
  </div>
  <div class="flex items-center space-x-2">
    @ui.RadioGroupItem(ui.RadioGroupItemProps{ID: "r2", Name: "option", Value: "option-two", Label: "Option Two"})
    @ui.Label(ui.LabelProps{For: "r2"}) { Option Two }
  </div>
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
		Slug: "resizable", Title: "Resizable", Description: "Accessible resizable panel groups and layouts with keyboard support.",
		Install: "templcn add resizable",
		GoUsage: `@ui.ResizablePanelGroup(ui.ResizablePanelGroupProps{DOMProps: ui.DOMProps{Class: "max-w-md rounded-lg border md:min-w-[450px]"}}) {
  @ui.ResizablePanel(ui.ResizablePanelProps{DefaultSize: 50}) {
    <div class="flex h-[200px] items-center justify-center p-6">
      <span class="font-semibold">One</span>
    </div>
  }
  @ui.ResizableHandle(ui.ResizableHandleProps{})
  @ui.ResizablePanel(ui.ResizablePanelProps{DefaultSize: 50}) {
    <div class="flex h-[200px] items-center justify-center p-6">
      <span class="font-semibold">Two</span>
    </div>
  }
}`,
		APIProps: []APIProp{
			{Name: "Direction", Type: "string", Default: `"horizontal"`, Desc: `"horizontal" or "vertical".`},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Two resizable panels with a divider handle.", GoCode: `@ui.ResizablePanelGroup(ui.ResizablePanelGroupProps{}) {
  @ui.ResizablePanel(ui.ResizablePanelProps{DefaultSize: 50}) { <span>One</span> }
  @ui.ResizableHandle(ui.ResizableHandleProps{})
  @ui.ResizablePanel(ui.ResizablePanelProps{DefaultSize: 50}) { <span>Two</span> }
}`, Preview: resizablePreview()},
		},
	},
	{
		Slug: "scroll-area", Title: "Scroll Area", Description: "Augments native scroll functionality for custom, cross-browser styling.",
		Install: "templcn add scroll-area",
		GoUsage: `@ui.ScrollArea(ui.DOMProps{Class: "h-72 w-48 rounded-md border p-4"}) {
  <h4 class="mb-4 text-sm font-medium leading-none">Tags</h4>
  <div class="text-sm">v1.0.0</div>
  @ui.Separator(ui.SeparatorProps{DOMProps: ui.DOMProps{Class: "my-2"}})
  <div class="text-sm">v1.1.0</div>
  @ui.Separator(ui.SeparatorProps{DOMProps: ui.DOMProps{Class: "my-2"}})
  <div class="text-sm">v1.2.0</div>
}`,
		APIProps: []APIProp{
			{Name: "Orientation", Type: "string", Default: `"vertical"`, Desc: `"vertical" or "both".`},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Scrollable view with custom styled scrollbars.", GoCode: `@ui.ScrollArea(ui.DOMProps{Class: "h-32"}) {
  <p>Scrollable content</p>
}`, Preview: scrollAreaPreview()},
		},
	},
	{
		Slug: "select", Title: "Select", Description: "Displays a list of options for the user to pick from—triggered by a button.",
		Install: "templcn add select",
		GoUsage: `@ui.Select(ui.SelectProps{Name: "fruit"}) {
  @ui.SelectTrigger(ui.SelectTriggerProps{DOMProps: ui.DOMProps{Class: "w-[180px]"}}) {
    @ui.SelectValue(ui.DOMProps{}) { Select a fruit }
  }
  @ui.SelectContent(ui.SelectContentProps{}) {
    @ui.SelectGroup(ui.DOMProps{}) {
      @ui.SelectLabel(ui.DOMProps{}) { Fruits }
      @ui.SelectItem(ui.DropdownMenuItemProps{Value: "apple"}) { Apple }
      @ui.SelectItem(ui.DropdownMenuItemProps{Value: "banana"}) { Banana }
      @ui.SelectItem(ui.DropdownMenuItemProps{Value: "blueberry"}) { Blueberry }
      @ui.SelectItem(ui.DropdownMenuItemProps{Value: "grapes"}) { Grapes }
      @ui.SelectItem(ui.DropdownMenuItemProps{Value: "pineapple"}) { Pineapple }
    }
  }
}`,
		APIProps: []APIProp{
			{Name: "DefaultOpen", Type: "bool", Default: "false", Desc: "Initial open state."},
			{Name: "Required", Type: "bool", Default: "false", Desc: "Marks as required."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Select trigger with grouped options list.", GoCode: `@ui.Select(ui.SelectProps{Name: "fruit"}) {
  @ui.SelectTrigger(ui.SelectTriggerProps{DOMProps: ui.DOMProps{Class: "w-[180px]"}}) {
    @ui.SelectValue(ui.DOMProps{}) { Select a fruit }
  }
  @ui.SelectContent(ui.SelectContentProps{}) {
    @ui.SelectItem(ui.DropdownMenuItemProps{Value: "apple"}) { Apple }
    @ui.SelectItem(ui.DropdownMenuItemProps{Value: "banana"}) { Banana }
  }
}`, Preview: selectPreview()},
		},
	},
	{
		Slug: "separator", Title: "Separator", Description: "Visually or semantically separates content.",
		Install: "templcn add separator",
		GoUsage: `<div>
  <div class="space-y-1">
    <h4 class="text-sm font-medium leading-none">Radix Primitives</h4>
    <p class="text-sm text-muted-foreground">An open-source UI component library.</p>
  </div>
  @ui.Separator(ui.SeparatorProps{DOMProps: ui.DOMProps{Class: "my-4"}})
  <div class="flex h-5 items-center space-x-4 text-sm">
    <div>Blog</div>
    @ui.Separator(ui.SeparatorProps{Orientation: ui.SeparatorOrientationVertical})
    <div>Docs</div>
    @ui.Separator(ui.SeparatorProps{Orientation: ui.SeparatorOrientationVertical})
    <div>Source</div>
  </div>
</div>`,
		APIProps: []APIProp{
			{Name: "Orientation", Type: "SeparatorOrientation", Default: `"horizontal"`, Desc: `"horizontal" or "vertical".`},
			{Name: "Decorative", Type: "bool", Default: "false", Desc: "When true, adds aria-hidden instead of role=separator."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Horizontal and vertical separators.", GoCode: `@ui.Separator(ui.SeparatorProps{})`, Preview: SeparatorPreview()},
		},
	},
	{
		Slug: "sheet", Title: "Sheet", Description: "Extends the Dialog component to display content that complements the main screen.",
		Install: "templcn add sheet",
		GoUsage: `@ui.Sheet(ui.SheetProps{Side: "right"}) {
  @ui.SheetTrigger(ui.DOMProps{}) {
    @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline}) { Open Sheet }
  }
  @ui.SheetContent(ui.DOMProps{}) {
    @ui.SheetHeader(ui.DOMProps{}) {
      @ui.SheetTitle(ui.DOMProps{}) { Edit profile }
      @ui.SheetDescription(ui.DOMProps{}) {
        Make changes to your profile here. Click save when you're done.
      }
    }
    <div class="grid gap-4 py-4">
      <div class="grid grid-cols-4 items-center gap-4">
        @ui.Label(ui.LabelProps{For: "name", DOMProps: ui.DOMProps{Class: "text-right"}}) { Name }
        @ui.Input(ui.InputProps{ID: "name", Value: "Pedro Duarte", DOMProps: ui.DOMProps{Class: "col-span-3"}})
      </div>
    </div>
    @ui.SheetFooter(ui.DOMProps{}) {
      @ui.Button(ui.ButtonProps{Type: "submit", Label: "Save changes"})
    }
  }
}`,
		APIProps: []APIProp{
			{Name: "Side", Type: "string", Default: `"right"`, Desc: `"top", "bottom", "left", "right".`},
			{Name: "ShowCloseButton", Type: "bool", Default: "false", Desc: "Show a close button."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Right slide-over panel with header and actions.", GoCode: `@ui.Sheet(ui.SheetProps{Side: "right"}) {
  @ui.SheetTrigger(ui.DOMProps{}) { @ui.Button(ui.ButtonProps{Label: "Open Sheet"}) }
  @ui.SheetContent(ui.DOMProps{}) {
    @ui.SheetHeader(ui.DOMProps{}) {
      @ui.SheetTitle(ui.DOMProps{}) { Edit profile }
    }
  }
}`, Preview: sheetPreview()},
		},
	},
	{
		Slug: "sidebar", Title: "Sidebar", Description: "A composable, themeable and customizable sidebar component.",
		Install: "templcn add sidebar",
		GoUsage: `@ui.SidebarProvider(ui.SidebarProviderProps{}) {
  @ui.Sidebar(ui.SidebarProps{}) {
    @ui.SidebarContent(ui.DOMProps{}) {
      @ui.SidebarGroup(ui.DOMProps{}) {
        @ui.SidebarGroupLabel(ui.DOMProps{}) { Application }
        @ui.SidebarGroupContent(ui.DOMProps{}) {
          @ui.SidebarMenu(ui.DOMProps{}) {
            @ui.SidebarMenuItem(ui.DOMProps{}) {
              @ui.SidebarMenuButton(ui.SidebarMenuButtonProps{IsActive: true}) {
                <span>Dashboard</span>
              }
            }
          }
        }
      }
    }
  }
  <main class="flex-1 p-6">
    @ui.SidebarTrigger(ui.DOMProps{})
  </main>
}`,
		APIProps: []APIProp{
			{Name: "Side", Type: "string", Default: `"left"`, Desc: `"left" or "right".`},
			{Name: "Variant", Type: "string", Default: `"sidebar"`, Desc: `"sidebar", "floating", or "inset".`},
			{Name: "Collapsible", Type: "string", Default: `"offcanvas"`, Desc: `"offcanvas", "icon", or "none".`},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Application layout sidebar with menus and triggers.", GoCode: `@ui.SidebarProvider(ui.SidebarProviderProps{DefaultOpen: true}) {
  @ui.Sidebar(ui.SidebarProps{}) {
    @ui.SidebarContent(ui.DOMProps{}) {
      @ui.SidebarMenu(ui.DOMProps{}) {
        @ui.SidebarMenuItem(ui.DOMProps{}) {
          @ui.SidebarMenuButton(ui.SidebarMenuButtonProps{IsActive: true}) { Dashboard }
        }
      }
    }
  }
}`, Preview: sidebarPreview()},
		},
	},
	{
		Slug: "skeleton", Title: "Skeleton", Description: "Use to show a placeholder while content is loading.",
		Install: "templcn add skeleton",
		GoUsage: `<div class="flex items-center space-x-4">
  @ui.Skeleton(ui.DOMProps{Class: "size-12 rounded-full"})
  <div class="space-y-2">
    @ui.Skeleton(ui.DOMProps{Class: "h-4 w-[250px]"})
    @ui.Skeleton(ui.DOMProps{Class: "h-4 w-[200px]"})
  </div>
</div>`,
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
		Slug: "slider", Title: "Slider", Description: "An input where the user selects a value from within a given range.",
		Install: "templcn add slider",
		GoUsage: `@ui.Slider(ui.SliderProps{
  DefaultValue: []float64{50},
  Max: 100,
  Step: 1,
  DOMProps: ui.DOMProps{Class: "w-[60%]"},
})`,
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
		Slug: "sonner", Title: "Sonner", Description: "An opinionated toast component for Go and templ.",
		Install: "templcn add sonner",
		GoUsage: `// 1. Add Toaster to your root layout:
@ui.Toaster(ui.ToasterProps{})

// 2. Trigger toasts with a button or script:
@ui.Button(ui.ButtonProps{
  Variant: ui.ButtonVariantOutline,
  DOMProps: ui.DOMProps{
    Attrs: templ.Attributes{
      "onclick": "toast('Event has been created', { description: 'Sunday, December 03, 2023 at 9:00 AM', action: { label: 'Undo', onClick: () => console.log('Undo') } })",
    },
  },
}) {
  Show Toast
}`,
		APIProps: []APIProp{
			{Name: "Theme", Type: "string", Default: `""`, Desc: "Force light or dark mode."},
			{Name: "Position", Type: "string", Default: `"bottom-right"`, Desc: "Toast position on screen."},
			{Name: "RichColors", Type: "bool", Default: "false", Desc: "Enable rich status colors."},
		},
		Examples: []ExampleEntry{
			{
				Name: "Default",
				Desc: "A default toast with title and description.",
				GoCode: `@ui.Button(ui.ButtonProps{
  Variant: ui.ButtonVariantOutline,
  DOMProps: ui.DOMProps{
    Attrs: templ.Attributes{
      "onclick": "toast('Event has been created', { description: 'Sunday, December 03, 2023 at 9:00 AM', action: { label: 'Undo' } })",
    },
  },
}) {
  Show Toast
}`,
				Preview: sonnerPreview(),
			},
			{
				Name: "Success",
				Desc: "A success toast with green accent.",
				GoCode: `@ui.Button(ui.ButtonProps{
  Variant: ui.ButtonVariantOutline,
  DOMProps: ui.DOMProps{
    Attrs: templ.Attributes{
      "onclick": "toast.success('Event has been created')",
    },
  },
}) {
  Success
}`,
				Preview: sonnerPreview(),
			},
			{
				Name: "Error",
				Desc: "A destructive error toast.",
				GoCode: `@ui.Button(ui.ButtonProps{
  Variant: ui.ButtonVariantOutline,
  DOMProps: ui.DOMProps{
    Attrs: templ.Attributes{
      "onclick": "toast.error('Event has not been created')",
    },
  },
}) {
  Error
}`,
				Preview: sonnerPreview(),
			},
			{
				Name: "Action",
				Desc: "A toast with an action button.",
				GoCode: `@ui.Button(ui.ButtonProps{
  Variant: ui.ButtonVariantOutline,
  DOMProps: ui.DOMProps{
    Attrs: templ.Attributes{
      "onclick": "toast('Event created', { action: { label: 'Undo', onClick: () => console.log('Undo') } })",
    },
  },
}) {
  Action
}`,
				Preview: sonnerPreview(),
			},
		},
	},
	{
		Slug: "spinner", Title: "Spinner", Description: "A lightweight animated loading indicator.",
		Install: "templcn add spinner",
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
		Slug: "switch", Title: "Switch", Description: "A control that allows the user to toggle between checked and not checked.",
		Install: "templcn add switch",
		GoUsage: `<div class="flex items-center space-x-2">
  @ui.Switch(ui.SwitchProps{ID: "airplane-mode", Name: "airplane-mode"})
  @ui.Label(ui.LabelProps{For: "airplane-mode"}) { Airplane Mode }
</div>`,
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
		Slug: "table", Title: "Table", Description: "A responsive table component.",
		Install: "templcn add table",
		GoUsage: `@ui.Table(ui.DOMProps{}) {
  @ui.TableCaption(ui.DOMProps{}) { A list of your recent invoices. }
  @ui.TableHeader(ui.DOMProps{}) {
    @ui.TableRow(ui.DOMProps{}) {
      @ui.TableHead(ui.DOMProps{Class: "w-[100px]"}) { Invoice }
      @ui.TableHead(ui.DOMProps{}) { Status }
      @ui.TableHead(ui.DOMProps{}) { Method }
      @ui.TableHead(ui.DOMProps{Class: "text-right"}) { Amount }
    }
  }
  @ui.TableBody(ui.DOMProps{}) {
    @ui.TableRow(ui.DOMProps{}) {
      @ui.TableCell(ui.DOMProps{Class: "font-medium"}) { INV001 }
      @ui.TableCell(ui.DOMProps{}) { Paid }
      @ui.TableCell(ui.DOMProps{}) { Credit Card }
      @ui.TableCell(ui.DOMProps{Class: "text-right"}) { $250.00 }
    }
  }
}`,
		APIProps: []APIProp{},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "A basic responsive data table.", GoCode: `@ui.Table(ui.DOMProps{}) {
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
}`, Preview: tablePreview()},
		},
	},
	{
		Slug: "tabs", Title: "Tabs", Description: "A set of layered sections of content—known as tab panels—that are displayed one at a time.",
		Install: "templcn add tabs",
		GoUsage: `@ui.Tabs(ui.TabsProps{DefaultValue: "account", DOMProps: ui.DOMProps{Class: "w-[400px]"}}) {
  @ui.TabsList(ui.TabsListProps{}) {
    @ui.TabsTrigger(ui.TabsTriggerProps{Value: "account", Active: true}) { Account }
    @ui.TabsTrigger(ui.TabsTriggerProps{Value: "password"}) { Password }
  }
  @ui.TabsContent(ui.TabsContentProps{Value: "account", Active: true}) {
    <p class="text-sm text-muted-foreground p-4">Make changes to your account here.</p>
  }
  @ui.TabsContent(ui.TabsContentProps{Value: "password"}) {
    <p class="text-sm text-muted-foreground p-4">Change your password here.</p>
  }
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
		Slug: "textarea", Title: "Textarea", Description: "Displays a form textarea or a component that looks like a textarea.",
		Install: "templcn add textarea",
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
		Slug: "toast", Title: "Toast", Description: "A succinct message that is displayed temporarily.",
		Install: "templcn add toast",
		GoUsage: `@ui.Toast(ui.ToastProps{Open: true}) {
  @ui.ToastTitle(ui.DOMProps{}) { Scheduled: Catch up }
  @ui.ToastDescription(ui.DOMProps{}) { Friday, February 10, 2026 at 5:57 PM }
  @ui.ToastClose(ui.DOMProps{})
}`,
		APIProps: []APIProp{
			{Name: "Variant", Type: "string", Default: `""`, Desc: "Visual variant (e.g. destructive)."},
			{Name: "Duration", Type: "int", Default: "5000", Desc: "Auto-dismiss time in ms."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "A toast notification popup.", GoCode: `@ui.Toast(ui.ToastProps{Open: true}) {
  @ui.ToastTitle(ui.DOMProps{}) { Scheduled! }
  @ui.ToastDescription(ui.DOMProps{}) { Event created for Friday. }
}`, Preview: toastPreview()},
		},
	},
	{
		Slug: "toggle", Title: "Toggle", Description: "A two-state button that can be either on or off.",
		Install: "templcn add toggle",
		GoUsage: `@ui.Toggle(ui.ToggleProps{AriaLabel: "Toggle italic"}) {
  Italic
}`,
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
		Slug: "toggle-group", Title: "Toggle Group", Description: "A set of two-state buttons that can be toggled on or off.",
		Install: "templcn add toggle-group",
		GoUsage: `@ui.ToggleGroup(ui.ToggleGroupProps{Type: "multiple"}) {
  @ui.ToggleGroupItem(ui.ToggleGroupItemProps{Value: "bold", AriaLabel: "Toggle bold"}) { B }
  @ui.ToggleGroupItem(ui.ToggleGroupItemProps{Value: "italic", AriaLabel: "Toggle italic"}) { I }
  @ui.ToggleGroupItem(ui.ToggleGroupItemProps{Value: "underline", AriaLabel: "Toggle underline"}) { U }
}`,
		APIProps: []APIProp{
			{Name: "Type", Type: "string", Default: `"single"`, Desc: `"single" or "multiple".`},
		},
		Examples: []ExampleEntry{
			{Name: "Single & Multiple", Desc: "Toggle groups for single or multi selection.", GoCode: `@ui.ToggleGroup(ui.ToggleGroupProps{Type: "single"}) {
  @ui.ToggleGroupItem(ui.ToggleGroupItemProps{Value: "bold"}) { B }
  @ui.ToggleGroupItem(ui.ToggleGroupItemProps{Value: "italic"}) { I }
}`, Preview: toggleGroupPreview()},
		},
	},
	{
		Slug: "tooltip", Title: "Tooltip", Description: "A popup that displays information related to an element when the element receives keyboard focus or the mouse hovers over it.",
		Install: "templcn add tooltip",
		GoUsage: `@ui.TooltipProvider(ui.TooltipProviderProps{}) {
  @ui.Tooltip(ui.TooltipProps{}) {
    @ui.TooltipTrigger(ui.DOMProps{}) {
      @ui.Button(ui.ButtonProps{Variant: ui.ButtonVariantOutline}) { Hover }
    }
    @ui.TooltipContent(ui.DOMProps{}) {
      <p>Add to library</p>
    }
  }
}`,
		APIProps: []APIProp{
			{Name: "Side", Type: "string", Default: `"top"`, Desc: `"top", "bottom", "left", "right".`},
			{Name: "DelayDuration", Type: "int", Default: "700", Desc: "Hover delay in ms before showing."},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Floating tooltip on trigger hover/focus.", GoCode: `@ui.TooltipProvider(ui.TooltipProviderProps{}) {
  @ui.Tooltip(ui.TooltipProps{}) {
    @ui.TooltipTrigger(ui.DOMProps{}) { Hover me }
    @ui.TooltipContent(ui.DOMProps{}) { Helpful hint }
  }
}`, Preview: tooltipPreview()},
		},
	},
	{
		Slug: "typography", Title: "Typography", Description: "Styles for headings, paragraphs, lists...etc",
		Install: "templcn add typography",
		GoUsage: `@ui.H1(ui.DOMProps{}) { Taxing Laughter: The Joke Tax Chronicles }
@ui.P(ui.DOMProps{}) {
  The king, seeing how much happier his subjects were, realized the error of
  his ways and repealed the joke tax.
}`,
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
	{
		Slug: "attachment", Title: "Attachment", Description: "A compact file attachment surface with media, status, content, and actions.",
		Install: "templcn add attachment",
		GoUsage: `@ui.Attachment(ui.AttachmentProps{State: "done"}) {
  @ui.AttachmentMedia(ui.DOMProps{}) { <span>PDF</span> }
  @ui.AttachmentContent(ui.DOMProps{}) {
    @ui.AttachmentTitle(ui.DOMProps{}) { report.pdf }
    @ui.AttachmentDescription(ui.DOMProps{}) { 2.4 MB }
  }
}`,
		APIProps: []APIProp{
			{Name: "State", Type: "string", Default: `"done"`, Desc: `"idle", "uploading", "processing", "error", or "done".`},
			{Name: "Size", Type: "string", Default: `"default"`, Desc: `"default", "sm", or "xs".`},
			{Name: "Orientation", Type: "string", Default: `"horizontal"`, Desc: `"horizontal" or "vertical".`},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Completed file upload attachment badge.", GoCode: `@ui.Attachment(ui.AttachmentProps{State: "done"}) {
  @ui.AttachmentContent(ui.DOMProps{}) { report.pdf }
}`, Preview: examplePreview("Attachment")},
		},
	},
	{
		Slug: "bubble", Title: "Bubble", Description: "A chat bubble with variants, alignment, content, and reactions.",
		Install: "templcn add bubble",
		GoUsage: `@ui.Bubble(ui.BubbleProps{Align: "end"}) {
  @ui.BubbleContent(ui.DOMProps{}) { Hello from templ! }
}`,
		APIProps: []APIProp{
			{Name: "Variant", Type: "string", Default: `"default"`, Desc: "Visual bubble variant."},
			{Name: "Align", Type: "string", Default: `"start"`, Desc: `"start" or "end".`},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Aligned message chat bubble.", GoCode: `@ui.Bubble(ui.BubbleProps{Align: "end"}) {
  @ui.BubbleContent(ui.DOMProps{}) { Hello! }
}`, Preview: examplePreview("Bubble")},
		},
	},
	{
		Slug: "marker", Title: "Marker", Description: "A compact inline marker for labels, separators, and metadata.",
		Install: "templcn add marker",
		GoUsage: `@ui.Marker(ui.MarkerProps{Variant: "separator"}) {
  @ui.MarkerContent(ui.DOMProps{}) { Or continue with }
}`,
		APIProps: []APIProp{
			{Name: "Variant", Type: "string", Default: `"default"`, Desc: `"default", "separator", or "border".`},
		},
		Examples: []ExampleEntry{
			{Name: "Separator", Desc: "Inline text divider badge marker.", GoCode: `@ui.Marker(ui.MarkerProps{Variant: "separator"}) {
  @ui.MarkerContent(ui.DOMProps{}) { Or continue with }
}`, Preview: examplePreview("Marker")},
		},
	},
	{
		Slug: "message", Title: "Message", Description: "Message layout primitives for avatars, headers, content, and footers.",
		Install: "templcn add message",
		GoUsage: `@ui.Message(ui.MessageProps{}) {
  @ui.MessageAvatar(ui.DOMProps{}) { <span>JD</span> }
  @ui.MessageContent(ui.DOMProps{}) { <span>Welcome to the workspace!</span> }
}`,
		APIProps: []APIProp{
			{Name: "Align", Type: "string", Default: `"start"`, Desc: `"start" or "end".`},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Chat message item with avatar and content.", GoCode: `@ui.Message(ui.MessageProps{}) {
  @ui.MessageAvatar(ui.DOMProps{}) { <span>JD</span> }
  @ui.MessageContent(ui.DOMProps{}) { <span>Welcome!</span> }
}`, Preview: examplePreview("Message")},
		},
	},
	{
		Slug: "message-scroller", Title: "Message Scroller", Description: "A scrollable message viewport with content, items, and scroll controls.",
		Install: "templcn add message-scroller",
		GoUsage: `@ui.MessageScroller(ui.MessageScrollerProps{}) {
  @ui.MessageScrollerViewport(ui.DOMProps{}) {
    @ui.MessageScrollerContent(ui.DOMProps{}) {
      @ui.MessageScrollerItem(ui.MessageScrollerItemProps{}) { Welcome! }
      @ui.MessageScrollerItem(ui.MessageScrollerItemProps{}) { Messages go here. }
    }
  }
  @ui.MessageScrollerButton(ui.MessageScrollerButtonProps{Direction: "end"})
}`,
		APIProps: []APIProp{
			{Name: "Direction", Type: "string", Default: `"end"`, Desc: `"start" or "end" for the scroll button.`},
		},
		Examples: []ExampleEntry{
			{Name: "Default", Desc: "Scrollable message window with scroll action button.", GoCode: `@ui.MessageScroller(ui.MessageScrollerProps{}) {
  @ui.MessageScrollerViewport(ui.DOMProps{}) {
    @ui.MessageScrollerContent(ui.DOMProps{}) { Messages }
  }
}`, Preview: examplePreview("Message Scroller")},
		},
	},
}
var blockDocs = []BlockEntry{
	{Slug: "dashboard-01", Title: "A dashboard with sidebar, charts and data table", Description: "A dense app shell that combines a sidebar, KPI cards, charts, and a data table.", Category: "featured", Command: "templcn add dashboard-01", Files: []string{"app/dashboard/page.templ", "components/app-sidebar.templ", "components/chart-area-interactive.templ", "components/data-table.templ", "components/section-cards.templ", "components/site-header.templ"}},
	{Slug: "sidebar-07", Title: "A sidebar that collapses to icons", Description: "A compact application shell with icon-only collapse behavior.", Category: "sidebar", Command: "templcn add sidebar-07", Files: []string{"app/dashboard/page.templ", "components/app-sidebar.templ", "components/team-switcher.templ"}},
	{Slug: "sidebar-03", Title: "A sidebar with submenus", Description: "Sidebar navigation with nested submenu states and a content region.", Category: "sidebar", Command: "templcn add sidebar-03", Files: []string{"app/dashboard/page.templ", "components/app-sidebar.templ"}},
	{Slug: "login-01", Title: "A simple login form", Description: "A centered login form with a compact footprint.", Category: "login", Command: "templcn add login-01", Files: []string{"app/login/page.templ", "components/login-form.templ"}},
	{Slug: "login-03", Title: "A login page with a muted background color", Description: "A basic authentication surface with soft contrast and a small brand lockup.", Category: "login", Command: "templcn add login-03", Files: []string{"app/login/page.templ", "components/login-form.templ"}},
	{Slug: "login-04", Title: "A login page with form and image", Description: "A wider authentication layout that pairs form and image panels.", Category: "login", Command: "templcn add login-04", Files: []string{"app/login/page.templ", "components/login-form.templ"}},
	{Slug: "signup-01", Title: "A signup page with a cover image", Description: "A registration layout with a strong media panel.", Category: "signup", Command: "templcn add signup-01", Files: []string{"app/signup/page.templ", "components/signup-form.templ"}},
	{Slug: "signup-02", Title: "A signup form with social providers", Description: "A signup flow that emphasizes third-party providers and email signup.", Category: "signup", Command: "templcn add signup-02", Files: []string{"app/signup/page.templ", "components/signup-form.templ"}},
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
		if doc.Install == "" || doc.Install == "templcn add button" || doc.Install == "templcn add <component>" {
			doc.Install = "templcn add " + doc.Slug
		}
		if doc.GoUsage == "" {
			doc.GoUsage = defaultGoUsage(doc)
		}
		if len(doc.APIProps) == 0 {
			doc.APIProps = defaultAPIProps(doc)
		}
		if doc.SourceCode == "" {
			doc.SourceCode = defaultSourceCode(doc)
		}
		if len(doc.Composition) == 0 {
			doc.Composition = defaultComposition(doc)
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
	for _, candidate := range componentSourceCandidates(doc.Slug) {
		raw, err := os.ReadFile(candidate)
		if err == nil && len(raw) > 0 {
			return strings.TrimSpace(string(raw))
		}
	}
	return strings.TrimSpace(doc.GoUsage)
}

func componentSourceCandidates(slug string) []string {
	file := strings.ReplaceAll(slug, "-", "_") + ".go"
	if slug == "data-table" {
		file = "datatable.go"
	}
	return []string{
		filepath.Join("..", "ui", file),
		filepath.Join("ui", file),
		filepath.Join("..", "..", "ui", file),
	}
}

var templCallPattern = regexp.MustCompile(`@ui\.([A-Za-z0-9_]+)`)

func defaultComposition(doc *ComponentDocEntry) []string {
	matches := templCallPattern.FindAllStringSubmatch(doc.GoUsage, -1)
	seen := map[string]struct{}{}
	parts := make([]string, 0, len(matches))
	for _, match := range matches {
		if len(match) < 2 {
			continue
		}
		part := match[1]
		if _, ok := seen[part]; ok {
			continue
		}
		seen[part] = struct{}{}
		parts = append(parts, part)
	}
	if len(parts) > 0 {
		return parts
	}
	name := strings.ReplaceAll(doc.Title, " ", "")
	if name == "" {
		name = strings.ReplaceAll(TitleFromSlug(doc.Slug), " ", "")
	}
	return []string{name}
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
	if len(examples) == 0 {
		examples = append(examples, ExampleEntry{
			Name:    "Default",
			Desc:    defaultExampleDesc(doc, "Default"),
			GoCode:  doc.GoUsage,
			Preview: previewForExample(doc.Slug, "Default"),
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

import "templcn/ui"

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
						<span class="flex size-10 items-center justify-center rounded-lg bg-primary text-primary-foreground text-lg font-semibold">T</span>
						<span class="text-2xl font-semibold">templcn/ui</span>
					</div>
					<h2 class="text-3xl font-bold">Beautiful components built with Radix UI and Tailwind CSS.</h2>
					<p class="text-muted-foreground">A set of accessible, customizable, and themeable components that you can copy and paste into your apps.</p>
				</div>
			</div>
			<div class="flex w-full lg:w-1/2 items-center justify-center p-8">
				<div class="w-full max-w-sm">
					<div class="mb-8 lg:hidden">
						<div class="flex items-center gap-2">
							<span class="flex size-8 items-center justify-center rounded-lg bg-primary text-primary-foreground text-sm font-semibold">T</span>
							<span class="text-xl font-semibold">templcn/ui</span>
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
