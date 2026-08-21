package views

import (
	"context"
	"html"
	"io"
	"strings"

	"github.com/a-h/templ"
	"templcn/ui"
)

func previewFunc(title string) func() templ.Component {
	return func() templ.Component {
		return examplePreview(title)
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
	case "toast":
		return toastPreview()
	case "sonner":
		return sonnerPreview()
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
		return examplePreview(TitleFromSlug(slug))
	}
}

func ComponentPreview(slug string) templ.Component {
	if preview, ok := componentPreviews[slug]; ok {
		return preview()
	}
	return examplePreview(TitleFromSlug(slug))
}

var componentPreviews = map[string]func() templ.Component{
	"accordion":       AccordionPreview,
	"alert":           AlertDefaultPreview,
	"alert-dialog":    alertDialogPreview,
	"aspect-ratio":    AspectRatioPreview,
	"avatar":          avatarPreview,
	"badge":           BadgePreview,
	"breadcrumb":      BreadcrumbPreview,
	"button":          ButtonVariantsPreview,
	"button-group":    ButtonGroupPreview,
	"calendar":        calendarPreview,
	"card":            CardPreview,
	"carousel":        carouselPreview,
	"chart":           chartPreview,
	"checkbox":        CheckboxPreview,
	"collapsible":     CollapsiblePreview,
	"combobox":        commandPreview,
	"command":         commandPreview,
	"context-menu":    previewFunc("Context Menu"),
	"data-table":      dataTablePreview,
	"date-picker":     datePickerPreview,
	"dialog":          dialogPreview,
	"direction":       previewFunc("Direction Provider"),
	"drawer":          drawerPreview,
	"dropdown-menu":   DropdownMenuPreview,
	"empty":           emptyPreview,
	"field":           previewFunc("Field"),
	"hover-card":      hoverCardPreview,
	"input":           InputPreview,
	"input-group":     inputGroupPreview,
	"input-otp":       InputOTPPreview,
	"item":            itemPreview,
	"kbd":             KbdPreview,
	"label":           LabelPreview,
	"menubar":         menubarPreview,
	"native-select":   NativeSelectPreview,
	"navigation-menu": navigationMenuPreview,
	"pagination":      PaginationPreview,
	"popover":         PopoverPreview,
	"progress":        ProgressPreview,
	"radio-group":     RadioGroupPreview,
	"resizable":       resizablePreview,
	"scroll-area":     scrollAreaPreview,
	"select":          SelectPreview,
	"separator":       SeparatorPreview,
	"sheet":           sheetPreview,
	"sidebar":         sidebarPreview,
	"skeleton":        SkeletonPreview,
	"slider":          SliderPreview,
	"sonner":          sonnerPreview,
	"spinner":         SpinnerPreview,
	"switch":          SwitchPreview,
	"table":           TablePreview,
	"tabs":            TabsPreview,
	"textarea":        TextareaPreview,
	"toast":           toastPreview,
	"toggle":          TogglePreview,
	"toggle-group":    toggleGroupPreview,
	"tooltip":         TooltipPreview,
	"typography":      TypographyPreview,
}

func examplePreview(title string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, `<div class="rounded-xl border bg-background p-6">
  <div class="text-sm font-medium">`+html.EscapeString(title)+`</div>
  <div class="mt-4 rounded-md border bg-muted/30 p-4 text-sm text-muted-foreground">`+html.EscapeString(title)+` example</div>
</div>`)
		return err
	})
}

func htmlPreview(markup string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, markup)
		return err
	})
}

func textNode(value string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, templ.EscapeString(value))
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

func badgePreview() templ.Component {
	return htmlPreview(`<div class="flex gap-2 rounded-xl border bg-background p-6"><span class="inline-flex items-center rounded-full bg-primary px-2.5 py-0.5 text-xs font-medium text-primary-foreground">Stable</span><span class="inline-flex items-center rounded-full border px-2.5 py-0.5 text-xs font-medium">Preview</span></div>`)
}
func cardPreview() templ.Component {
	return htmlPreview(`<div class="max-w-sm rounded-xl border bg-card p-6 text-card-foreground shadow-sm"><div class="font-semibold">Project velocity</div><p class="mt-1 text-sm text-muted-foreground">12 components installed this week.</p></div>`)
}
func alertPreview() templ.Component {
	return htmlPreview(`<div role="alert" class="rounded-lg border p-4"><div class="font-medium">Heads up</div><div class="text-sm text-muted-foreground">This action updates copied source files.</div></div>`)
}
func inputPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><input class="h-9 w-full max-w-sm rounded-md border bg-background px-3 text-sm" placeholder="email@example.com"/></div>`)
}
func checkboxPreview() templ.Component {
	return htmlPreview(`<label class="flex items-center gap-2 rounded-xl border bg-background p-6 text-sm"><input data-slot="checkbox" type="checkbox" checked aria-checked="true" data-state="checked"/>Accept terms</label>`)
}
func labelPreview() templ.Component {
	return htmlPreview(`<div class="grid max-w-sm gap-2 rounded-xl border bg-background p-6"><label class="text-sm font-medium" for="preview-email">Email</label><input id="preview-email" class="h-9 rounded-md border px-3 text-sm"/></div>`)
}
func selectPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><button type="button" role="combobox" aria-expanded="false" class="h-9 min-w-48 rounded-md border px-3 text-left text-sm">Starter</button></div>`)
}
func textareaPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><textarea class="min-h-24 w-full max-w-sm rounded-md border bg-background p-3 text-sm">Ship the release notes.</textarea></div>`)
}
func dialogPreview() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return ui.Dialog(ui.DialogProps{Modal: true}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := ui.DialogTrigger(ui.DOMProps{Class: "inline-flex h-9 items-center justify-center rounded-md bg-primary px-4 py-2 text-sm font-medium text-primary-foreground shadow-xs hover:bg-primary/90"}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				_, err := io.WriteString(w, "Open dialog")
				return err
			})), w); err != nil {
				return err
			}
			return ui.DialogContent(ui.DOMProps{}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				if err := ui.DialogHeader(ui.DOMProps{}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
					if err := ui.DialogTitle(ui.DOMProps{}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
						_, err := io.WriteString(w, "Edit profile")
						return err
					})), w); err != nil {
						return err
					}
					return ui.DialogDescription(ui.DOMProps{}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
						_, err := io.WriteString(w, "Make changes to your profile here. Click save when you're done.")
						return err
					})), w)
				})), w); err != nil {
					return err
				}
				return ui.DialogFooter(ui.DOMProps{}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
					return ui.DialogClose(ui.DOMProps{Class: "inline-flex h-9 items-center justify-center rounded-md border px-4 py-2 text-sm font-medium"}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
						_, err := io.WriteString(w, "Close")
						return err
					})), w)
				})), w)
			})), w)
		})), w)
	})
}

func alertDialogPreview() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return ui.AlertDialog(ui.AlertDialogProps{}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := ui.AlertDialogTrigger(ui.DOMProps{Class: "inline-flex h-9 items-center justify-center rounded-md bg-destructive px-4 py-2 text-sm font-medium text-destructive-foreground shadow-xs hover:bg-destructive/90"}).Render(templ.WithChildren(ctx, textNode("Delete account")), w); err != nil {
				return err
			}
			return ui.AlertDialogContent(ui.DOMProps{}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				if err := ui.AlertDialogHeader(ui.DOMProps{}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
					if err := ui.AlertDialogTitle(ui.DOMProps{}).Render(templ.WithChildren(ctx, textNode("Are you absolutely sure?")), w); err != nil {
						return err
					}
					return ui.AlertDialogDescription(ui.DOMProps{}).Render(templ.WithChildren(ctx, textNode("This action cannot be undone.")), w)
				})), w); err != nil {
					return err
				}
				return ui.AlertDialogFooter(ui.DOMProps{}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
					if err := ui.AlertDialogCancel(ui.DOMProps{}).Render(templ.WithChildren(ctx, textNode("Cancel")), w); err != nil {
						return err
					}
					return ui.AlertDialogAction(ui.DOMProps{}).Render(templ.WithChildren(ctx, textNode("Continue")), w)
				})), w)
			})), w)
		})), w)
	})
}

func dropdownMenuPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><button class="rounded-md border px-3 py-2 text-sm">Open</button><div class="mt-2 w-40 rounded-md border bg-popover p-1 text-sm shadow"><div class="rounded-sm px-2 py-1.5">Profile</div><div class="rounded-sm px-2 py-1.5">Billing</div></div></div>`)
}
func tabsPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><div role="tablist" class="inline-flex rounded-md bg-muted p-1"><button class="rounded bg-background px-3 py-1 text-sm">Account</button><button class="px-3 py-1 text-sm">Password</button></div><div class="mt-4 text-sm">Account settings panel</div></div>`)
}
func sheetPreview() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return ui.Sheet(ui.SheetProps{Side: "right"}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := ui.SheetTrigger(ui.DOMProps{Class: "inline-flex h-9 items-center justify-center rounded-md border px-4 py-2 text-sm font-medium"}).Render(templ.WithChildren(ctx, textNode("Open sheet")), w); err != nil {
				return err
			}
			return ui.SheetContent(ui.DOMProps{}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				return ui.SheetHeader(ui.DOMProps{}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
					if err := ui.SheetTitle(ui.DOMProps{}).Render(templ.WithChildren(ctx, textNode("Edit profile")), w); err != nil {
						return err
					}
					return ui.SheetDescription(ui.DOMProps{}).Render(templ.WithChildren(ctx, textNode("Make changes and save.")), w)
				})), w)
			})), w)
		})), w)
	})
}

func drawerPreview() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return ui.Drawer(ui.DrawerProps{SheetProps: ui.SheetProps{Side: "bottom"}}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
			if err := ui.DrawerTrigger(ui.DOMProps{Class: "inline-flex h-9 items-center justify-center rounded-md border px-4 py-2 text-sm font-medium"}).Render(templ.WithChildren(ctx, textNode("Open drawer")), w); err != nil {
				return err
			}
			return ui.DrawerContent(ui.DOMProps{}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				if err := ui.DrawerHeader(ui.DOMProps{}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
					if err := ui.DrawerTitle(ui.DOMProps{}).Render(templ.WithChildren(ctx, textNode("Move goal")), w); err != nil {
						return err
					}
					return ui.DrawerDescription(ui.DOMProps{}).Render(templ.WithChildren(ctx, textNode("Set your daily activity target.")), w)
				})), w); err != nil {
					return err
				}
				return ui.DrawerFooter(ui.DOMProps{}).Render(templ.WithChildren(ctx, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
					return ui.DrawerClose(ui.DOMProps{Class: "inline-flex h-9 items-center justify-center rounded-md border px-4 py-2 text-sm font-medium"}).Render(templ.WithChildren(ctx, textNode("Close")), w)
				})), w)
			})), w)
		})), w)
	})
}
func tablePreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><table class="w-full text-sm"><thead><tr class="border-b"><th class="py-2 text-left">Name</th><th class="py-2 text-left">Status</th></tr></thead><tbody><tr><td class="py-2">Components</td><td class="py-2">Ready</td></tr></tbody></table></div>`)
}
func dataTablePreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><div class="mb-3 flex justify-between"><input class="h-8 rounded-md border px-2 text-sm" placeholder="Filter"/><button class="h-8 rounded-md border px-3 text-sm">Columns</button></div><table class="w-full text-sm"><tbody><tr class="border-t"><td class="py-2">button</td><td class="py-2">installed</td></tr></tbody></table></div>`)
}
func paginationPreview() templ.Component {
	return htmlPreview(`<nav class="flex items-center gap-1 rounded-xl border bg-background p-6"><button class="rounded-md border px-3 py-1 text-sm">Previous</button><button class="rounded-md bg-primary px-3 py-1 text-sm text-primary-foreground">1</button><button class="rounded-md border px-3 py-1 text-sm">Next</button></nav>`)
}
func progressPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><div class="h-2 w-full max-w-sm overflow-hidden rounded-full bg-primary/20"><div class="h-full w-2/3 bg-primary"></div></div></div>`)
}
func accordionPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><details open class="border-b py-2"><summary class="font-medium">Is it accessible?</summary><p class="pt-2 text-sm text-muted-foreground">Keyboard behavior is covered by browser tests.</p></details></div>`)
}
func collapsiblePreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><details open><summary class="cursor-pointer font-medium">Runtime files</summary><div class="pt-2 text-sm text-muted-foreground">assets/runtime.js</div></details></div>`)
}
func hoverCardPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><button class="underline">Hover David</button><div class="mt-2 w-56 rounded-md border bg-popover p-3 text-sm shadow">Maintainer preview</div></div>`)
}
func avatarPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><div class="flex size-10 items-center justify-center rounded-full bg-muted text-sm font-medium">DO</div></div>`)
}
func buttonGroupPreview() templ.Component {
	return htmlPreview(`<div class="inline-flex rounded-xl border bg-background p-6"><button class="rounded-l-md border px-3 py-2 text-sm">Left</button><button class="border-y px-3 py-2 text-sm">Center</button><button class="rounded-r-md border px-3 py-2 text-sm">Right</button></div>`)
}
func nativeSelectPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><select class="h-9 rounded-md border bg-background px-3 text-sm"><option>Go</option><option>templ</option></select></div>`)
}
func spinnerPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><div role="status" class="size-6 animate-spin rounded-full border-2 border-muted border-t-primary"></div></div>`)
}
func skeletonPreview() templ.Component {
	return htmlPreview(`<div class="space-y-3 rounded-xl border bg-background p-6"><div class="h-4 w-48 rounded bg-muted"></div><div class="h-4 w-32 rounded bg-muted"></div></div>`)
}
func separatorPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6 text-sm">Top<hr class="my-4 border-border"/>Bottom</div>`)
}
func inputGroupPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><div class="flex max-w-sm rounded-md border"><span class="px-3 py-2 text-sm text-muted-foreground">https://</span><input class="min-w-0 flex-1 px-3 text-sm" value="example.com"/></div></div>`)
}
func inputOTPPreview() templ.Component {
	return htmlPreview(`<div class="flex gap-2 rounded-xl border bg-background p-6"><div class="grid size-10 place-items-center rounded-md border">1</div><div class="grid size-10 place-items-center rounded-md border">2</div><div class="grid size-10 place-items-center rounded-md border">3</div></div>`)
}
func sliderPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><input type="range" min="0" max="100" value="40" class="w-full max-w-sm"/></div>`)
}
func toggleGroupPreview() templ.Component {
	return htmlPreview(`<div class="inline-flex gap-1 rounded-xl border bg-background p-6"><button aria-pressed="true" class="rounded-md bg-primary px-3 py-2 text-sm text-primary-foreground">B</button><button class="rounded-md border px-3 py-2 text-sm">I</button></div>`)
}
func togglePreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><button aria-pressed="true" class="rounded-md bg-primary px-3 py-2 text-sm text-primary-foreground">Bold</button></div>`)
}
func switchPreview() templ.Component {
	return htmlPreview(`<label class="flex items-center gap-3 rounded-xl border bg-background p-6 text-sm"><input role="switch" type="checkbox" checked aria-checked="true"/>Notifications</label>`)
}
func toastPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><div role="status" class="max-w-sm rounded-lg border bg-background p-4 shadow"><div class="font-medium">Saved</div><p class="text-sm text-muted-foreground">Your changes were stored.</p></div></div>`)
}
func sonnerPreview() templ.Component {
	return htmlPreview(`<div class="flex flex-wrap gap-3 rounded-xl border bg-background p-6">
  <button type="button" class="inline-flex items-center justify-center rounded-md border border-input bg-background px-4 py-2 text-sm font-medium shadow-sm transition-colors hover:bg-accent hover:text-accent-foreground" onclick="toast('Event has been created', { description: 'Sunday, December 03, 2023 at 9:00 AM', action: { label: 'Undo', onClick: () => console.log('Undo') } })">Default</button>
  <button type="button" class="inline-flex items-center justify-center rounded-md border border-input bg-background px-4 py-2 text-sm font-medium shadow-sm transition-colors hover:bg-accent hover:text-accent-foreground" onclick="toast.success('Event has been created')">Success</button>
  <button type="button" class="inline-flex items-center justify-center rounded-md border border-input bg-background px-4 py-2 text-sm font-medium shadow-sm transition-colors hover:bg-accent hover:text-accent-foreground" onclick="toast.info('Be at the area 10 minutes before start.')">Info</button>
  <button type="button" class="inline-flex items-center justify-center rounded-md border border-input bg-background px-4 py-2 text-sm font-medium shadow-sm transition-colors hover:bg-accent hover:text-accent-foreground" onclick="toast.warning('Event start time has been updated.')">Warning</button>
  <button type="button" class="inline-flex items-center justify-center rounded-md border border-input bg-background px-4 py-2 text-sm font-medium shadow-sm transition-colors hover:bg-accent hover:text-accent-foreground" onclick="toast.error('Event has not been created')">Error</button>
  <button type="button" class="inline-flex items-center justify-center rounded-md border border-input bg-background px-4 py-2 text-sm font-medium shadow-sm transition-colors hover:bg-accent hover:text-accent-foreground" onclick="toast('Event created', { action: { label: 'Undo', onClick: () => console.log('Undo') } })">Action</button>
</div>`)
}
func tooltipPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><button class="rounded-md border px-3 py-2 text-sm">Hover</button><span role="tooltip" class="ml-3 rounded bg-primary px-2 py-1 text-xs text-primary-foreground">Tooltip</span></div>`)
}
func popoverPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><button class="rounded-md border px-3 py-2 text-sm">Open</button><div class="mt-2 max-w-xs rounded-md border bg-popover p-3 text-sm shadow">Popover content</div></div>`)
}
func commandPreview() templ.Component {
	return htmlPreview(`<div class="max-w-sm rounded-xl border bg-popover p-2"><input class="mb-2 h-9 w-full rounded-md border px-3 text-sm" placeholder="Search commands"/><div class="rounded-sm px-2 py-1.5 text-sm">Add component</div></div>`)
}
func emptyPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-10 text-center"><div class="font-medium">No components installed</div><p class="text-sm text-muted-foreground">Run templcn add button.</p></div>`)
}
func itemPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><div class="flex items-center justify-between rounded-md border p-3"><div><div class="font-medium">Component item</div><div class="text-sm text-muted-foreground">button.go</div></div><button class="rounded-md border px-3 py-1 text-sm">View</button></div></div>`)
}
func calendarPreview() templ.Component {
	return ui.Calendar(ui.CalendarProps{
		DOMProps:       ui.DOMProps{Class: "rounded-lg border"},
		Mode:           "single",
		DefaultMonth:   "2026-05",
		Selected:       "2026-05-07",
		CaptionLayout:  "dropdown",
		ShowWeekNumber: true,
		FromYear:       2024,
		ToYear:         2028,
	})
}
func datePickerPreview() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return ui.DatePicker(ui.DatePickerProps{Name: "date", Value: "2026-05-08"}).Render(templ.WithChildren(ctx, ui.Calendar(ui.CalendarProps{
			DOMProps:     ui.DOMProps{Class: "rounded-lg border"},
			Mode:         "single",
			DefaultMonth: "2026-05",
			Selected:     "2026-05-08",
		})), w)
	})
}
func carouselPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><div class="grid h-28 place-items-center rounded-md border bg-muted/30">Slide 1</div><div class="mt-3 flex gap-2"><button class="rounded-md border px-3 py-1 text-sm">Previous</button><button class="rounded-md border px-3 py-1 text-sm">Next</button></div></div>`)
}
func chartPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><div class="flex h-32 items-end gap-2"><div class="h-12 w-10 rounded-t bg-primary"></div><div class="h-24 w-10 rounded-t bg-primary"></div><div class="h-16 w-10 rounded-t bg-primary"></div></div></div>`)
}
func breadcrumbPreview() templ.Component {
	return htmlPreview(`<nav class="rounded-xl border bg-background p-6 text-sm"><a class="text-muted-foreground">Docs</a><span class="mx-2">/</span><span>Components</span></nav>`)
}
func menubarPreview() templ.Component {
	return htmlPreview(`<div class="rounded-xl border bg-background p-6"><div role="menubar" class="inline-flex gap-1 rounded-md border p-1"><button class="px-3 py-1 text-sm">File</button><button class="px-3 py-1 text-sm">Edit</button></div></div>`)
}
func navigationMenuPreview() templ.Component {
	return htmlPreview(`<nav class="rounded-xl border bg-background p-6"><button class="rounded-md px-4 py-2 text-sm hover:bg-accent">Products</button><button class="rounded-md px-4 py-2 text-sm hover:bg-accent">Docs</button></nav>`)
}
func radioGroupPreview() templ.Component {
	return htmlPreview(`<div role="radiogroup" class="grid gap-2 rounded-xl border bg-background p-6 text-sm"><label><input type="radio" name="density" checked/> Comfortable</label><label><input type="radio" name="density"/> Compact</label></div>`)
}
func resizablePreview() templ.Component {
	return htmlPreview(`<div class="flex h-24 rounded-xl border bg-background p-6"><div class="basis-1/2 rounded bg-muted"></div><div role="separator" class="mx-2 w-1 bg-border"></div><div class="basis-1/2 rounded bg-muted"></div></div>`)
}
func scrollAreaPreview() templ.Component {
	return htmlPreview(`<div class="h-32 overflow-auto rounded-xl border bg-background p-6 text-sm"><p>Scrollable content</p><p class="mt-16">End</p></div>`)
}
func sidebarPreview() templ.Component {
	return htmlPreview(`<div class="flex h-48 rounded-xl border bg-background"><aside class="w-48 border-r p-3"><div class="font-medium">Dashboard</div><div class="mt-3 text-sm text-muted-foreground">Projects</div></aside><main class="flex-1 p-3 text-sm">Content</main></div>`)
}
func typographyPreview() templ.Component {
	return htmlPreview(`<article class="rounded-xl border bg-background p-6"><h3 class="text-xl font-semibold">Typography</h3><p class="mt-2 text-sm leading-6 text-muted-foreground">Readable defaults for prose and interface copy.</p></article>`)
}
