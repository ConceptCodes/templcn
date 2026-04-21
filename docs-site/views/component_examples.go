package views

import (
	"context"
	"html"
	"io"
	"strings"

	"github.com/a-h/templ"
	"shadcn/ui"
)

func previewFunc(title string) func() templ.Component {
	return func() templ.Component {
		return genericPreview(title)
	}
}

func previewForExample(slug, name string) templ.Component {
	switch slug {
	case "button":
		switch strings.ToLower(name) {
		case "outline":
			return ui.Button(ui.ButtonProps{Label: "Outline", Variant: ui.ButtonVariantOutline})
		case "secondary":
			return ui.Button(ui.ButtonProps{Label: "Secondary", Variant: ui.ButtonVariantSecondary})
		case "ghost":
			return ui.Button(ui.ButtonProps{Label: "Ghost", Variant: ui.ButtonVariantGhost})
		case "destructive":
			return ui.Button(ui.ButtonProps{Label: "Delete", Variant: ui.ButtonVariantDestructive})
		case "link":
			return ui.Button(ui.ButtonProps{Label: "Link", Variant: ui.ButtonVariantLink, Href: "#"})
		case "icon":
			return ui.Button(ui.ButtonProps{Label: "★", Size: ui.ButtonSizeIcon})
		default:
			return ui.Button(ui.ButtonProps{Label: "Click me"})
		}
	case "badge":
		switch strings.ToLower(name) {
		case "outline":
			return ui.Badge(ui.BadgeProps{Label: "Outline", Variant: ui.BadgeVariantOutline})
		case "secondary":
			return ui.Badge(ui.BadgeProps{Label: "Secondary", Variant: ui.BadgeVariantSecondary})
		case "ghost":
			return ui.Badge(ui.BadgeProps{Label: "Ghost", Variant: ui.BadgeVariantGhost})
		default:
			return ui.Badge(ui.BadgeProps{Label: "Badge"})
		}
	case "input":
		return inputPreview()
	case "checkbox":
		return checkboxPreview()
	case "select":
		return selectPreview()
	case "textarea":
		return textareaPreview()
	case "tabs":
		return tabsPreview()
	case "accordion":
		return accordionPreview()
	case "collapsible":
		return collapsiblePreview()
	case "dialog":
		return dialogPreview()
	case "dropdown-menu":
		return dropdownMenuPreview()
	case "sheet":
		return sheetPreview()
	case "table":
		return tablePreview()
	case "pagination":
		return paginationPreview()
	case "progress":
		return progressPreview()
	case "hover-card":
		return hoverCardPreview()
	case "card":
		return cardPreview()
	case "alert":
		return alertPreview()
	case "avatar":
		return avatarPreview()
	case "button-group":
		return buttonGroupPreview()
	case "native-select":
		return nativeSelectPreview()
	case "spinner":
		return spinnerPreview()
	case "skeleton":
		return skeletonPreview()
	case "separator":
		return separatorPreview()
	case "input-group":
		return inputGroupPreview()
	case "input-otp":
		return inputOTPPreview()
	case "slider":
		return sliderPreview()
	case "toggle-group":
		return toggleGroupPreview()
	case "toggle":
		return togglePreview()
	case "switch":
		return switchPreview()
	case "toast", "sonner":
		return toastPreview()
	case "tooltip":
		return tooltipPreview()
	case "popover":
		return popoverPreview()
	case "command":
		return commandPreview()
	case "empty":
		return emptyPreview()
	case "item":
		return itemPreview()
	case "calendar":
		return calendarPreview()
	case "date-picker":
		return datePickerPreview()
	case "carousel":
		return carouselPreview()
	case "chart":
		return chartPreview()
	case "breadcrumb":
		return breadcrumbPreview()
	case "menubar":
		return menubarPreview()
	case "navigation-menu":
		return navigationMenuPreview()
	case "radio-group":
		return radioGroupPreview()
	case "resizable":
		return resizablePreview()
	case "scroll-area":
		return scrollAreaPreview()
	case "typography":
		return typographyPreview()
	default:
		if preview, ok := componentPreviews[slug]; ok {
			return preview()
		}
		return genericPreview(TitleFromSlug(slug))
	}
}

func ComponentPreview(slug string) templ.Component {
	if preview, ok := componentPreviews[slug]; ok {
		return preview()
	}
	return genericPreview(TitleFromSlug(slug))
}

var componentPreviews = map[string]func() templ.Component{
	"accordion":       previewFunc("Accordion"),
	"alert":           alertPreview,
	"alert-dialog":    previewFunc("Alert Dialog"),
	"aspect-ratio":    previewFunc("Aspect Ratio"),
	"avatar":          avatarPreview,
	"badge":           badgePreview,
	"breadcrumb":      breadcrumbPreview,
	"button":          buttonPreview,
	"button-group":    buttonGroupPreview,
	"calendar":        calendarPreview,
	"card":            cardPreview,
	"carousel":        carouselPreview,
	"chart":           chartPreview,
	"checkbox":        checkboxPreview,
	"collapsible":     collapsiblePreview,
	"combobox":        previewFunc("Combobox"),
	"command":         commandPreview,
	"context-menu":    previewFunc("Context Menu"),
	"data-table":      dataTablePreview,
	"date-picker":     datePickerPreview,
	"dialog":          dialogPreview,
	"direction":       previewFunc("Direction Provider"),
	"drawer":          previewFunc("Drawer"),
	"dropdown-menu":   dropdownMenuPreview,
	"empty":           emptyPreview,
	"field":           previewFunc("Field"),
	"hover-card":      hoverCardPreview,
	"input":           inputPreview,
	"input-group":     inputGroupPreview,
	"input-otp":       inputOTPPreview,
	"item":            itemPreview,
	"kbd":             previewFunc("Keyboard Shortcut"),
	"label":           labelPreview,
	"menubar":         menubarPreview,
	"native-select":   nativeSelectPreview,
	"navigation-menu": navigationMenuPreview,
	"pagination":      paginationPreview,
	"popover":         popoverPreview,
	"progress":        progressPreview,
	"radio-group":     radioGroupPreview,
	"resizable":       resizablePreview,
	"scroll-area":     scrollAreaPreview,
	"select":          selectPreview,
	"separator":       separatorPreview,
	"sheet":           sheetPreview,
	"sidebar":         sidebarPreview,
	"skeleton":        skeletonPreview,
	"slider":          sliderPreview,
	"sonner":          toastPreview,
	"spinner":         spinnerPreview,
	"switch":          switchPreview,
	"table":           tablePreview,
	"tabs":            tabsPreview,
	"textarea":        textareaPreview,
	"toast":           toastPreview,
	"toggle":          togglePreview,
	"toggle-group":    toggleGroupPreview,
	"tooltip":         tooltipPreview,
	"typography":      typographyPreview,
}

func genericPreview(title string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, `<div class="flex min-h-40 items-center justify-center rounded-xl border bg-background p-6">
  <div class="max-w-sm text-center">
    <div class="text-sm font-medium">`+html.EscapeString(title)+`</div>
    <div class="mt-1 text-sm text-muted-foreground">Generated preview for `+html.EscapeString(title)+`.</div>
  </div>
</div>`)
		return err
	})
}

func buttonPreview() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, `<div class="flex flex-wrap items-center gap-2 rounded-xl border bg-background p-6">`)
		if err != nil {
			return err
		}
		components := []templ.Component{
			ui.Button(ui.ButtonProps{Label: "Default"}),
			ui.Button(ui.ButtonProps{Label: "Outline", Variant: ui.ButtonVariantOutline}),
			ui.Button(ui.ButtonProps{Label: "Ghost", Variant: ui.ButtonVariantGhost}),
			ui.Button(ui.ButtonProps{Label: "Delete", Variant: ui.ButtonVariantDestructive}),
		}
		for _, component := range components {
			if err := component.Render(ctx, w); err != nil {
				return err
			}
		}
		_, err = io.WriteString(w, `</div>`)
		return err
	})
}

func badgePreview() templ.Component { return genericPreview("Badge") }
func cardPreview() templ.Component { return genericPreview("Card") }
func alertPreview() templ.Component { return genericPreview("Alert") }
func inputPreview() templ.Component { return genericPreview("Input") }
func checkboxPreview() templ.Component { return genericPreview("Checkbox") }
func labelPreview() templ.Component { return genericPreview("Label") }
func selectPreview() templ.Component { return genericPreview("Select") }
func textareaPreview() templ.Component { return genericPreview("Textarea") }
func dialogPreview() templ.Component { return genericPreview("Dialog") }
func dropdownMenuPreview() templ.Component { return genericPreview("Dropdown Menu") }
func tabsPreview() templ.Component { return genericPreview("Tabs") }
func sheetPreview() templ.Component { return genericPreview("Sheet") }
func tablePreview() templ.Component { return genericPreview("Table") }
func dataTablePreview() templ.Component { return genericPreview("Data Table") }
func paginationPreview() templ.Component { return genericPreview("Pagination") }
func progressPreview() templ.Component { return genericPreview("Progress") }
func accordionPreview() templ.Component { return genericPreview("Accordion") }
func collapsiblePreview() templ.Component { return genericPreview("Collapsible") }
func hoverCardPreview() templ.Component { return genericPreview("Hover Card") }
func avatarPreview() templ.Component { return genericPreview("Avatar") }
func buttonGroupPreview() templ.Component { return genericPreview("Button Group") }
func nativeSelectPreview() templ.Component { return genericPreview("Native Select") }
func spinnerPreview() templ.Component { return genericPreview("Spinner") }
func skeletonPreview() templ.Component { return genericPreview("Skeleton") }
func separatorPreview() templ.Component { return genericPreview("Separator") }
func inputGroupPreview() templ.Component { return genericPreview("Input Group") }
func inputOTPPreview() templ.Component { return genericPreview("Input OTP") }
func sliderPreview() templ.Component { return genericPreview("Slider") }
func toggleGroupPreview() templ.Component { return genericPreview("Toggle Group") }
func togglePreview() templ.Component { return genericPreview("Toggle") }
func switchPreview() templ.Component { return genericPreview("Switch") }
func toastPreview() templ.Component { return genericPreview("Toast") }
func tooltipPreview() templ.Component { return genericPreview("Tooltip") }
func popoverPreview() templ.Component { return genericPreview("Popover") }
func commandPreview() templ.Component { return genericPreview("Command") }
func emptyPreview() templ.Component { return genericPreview("Empty") }
func itemPreview() templ.Component { return genericPreview("Item") }
func calendarPreview() templ.Component { return genericPreview("Calendar") }
func datePickerPreview() templ.Component { return genericPreview("Date Picker") }
func carouselPreview() templ.Component { return genericPreview("Carousel") }
func chartPreview() templ.Component { return genericPreview("Chart") }
func breadcrumbPreview() templ.Component { return genericPreview("Breadcrumb") }
func menubarPreview() templ.Component { return genericPreview("Menubar") }
func navigationMenuPreview() templ.Component { return genericPreview("Navigation Menu") }
func radioGroupPreview() templ.Component { return genericPreview("Radio Group") }
func resizablePreview() templ.Component { return genericPreview("Resizable") }
func scrollAreaPreview() templ.Component { return genericPreview("Scroll Area") }
func sidebarPreview() templ.Component { return genericPreview("Sidebar") }
func typographyPreview() templ.Component { return genericPreview("Typography") }
