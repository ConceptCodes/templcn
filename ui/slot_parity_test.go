package ui

import "testing"

func assertSlot(t *testing.T, name string, html string, want string) {
	t.Helper()
	if got := attrValue(html, "data-slot"); got != want {
		t.Fatalf("%s data-slot = %q, want %q", name, got, want)
	}
}

func TestOverlayFamilySlotsAreDistinct(t *testing.T) {
	assertSlot(t, "AlertDialog", mustRender(t, AlertDialog(AlertDialogProps{})), "alert-dialog")
	assertSlot(t, "AlertDialogTrigger", mustRender(t, AlertDialogTrigger(DOMProps{})), "alert-dialog-trigger")
	assertSlot(t, "AlertDialogContent", mustRender(t, AlertDialogContent(DOMProps{})), "alert-dialog-content")
	assertSlot(t, "AlertDialogAction", mustRender(t, AlertDialogAction(DOMProps{})), "alert-dialog-action")
	assertSlot(t, "AlertDialogCancel", mustRender(t, AlertDialogCancel(DOMProps{})), "alert-dialog-cancel")

	assertSlot(t, "SheetTrigger", mustRender(t, SheetTrigger(DOMProps{})), "sheet-trigger")
	assertSlot(t, "SheetOverlay", mustRender(t, SheetOverlay(DOMProps{})), "sheet-overlay")
	assertSlot(t, "SheetContent", mustRender(t, SheetContent(DOMProps{})), "sheet-content")
	assertSlot(t, "SheetClose", mustRender(t, SheetClose(DOMProps{})), "sheet-close")

	assertSlot(t, "DrawerTrigger", mustRender(t, DrawerTrigger(DOMProps{})), "drawer-trigger")
	assertSlot(t, "DrawerOverlay", mustRender(t, DrawerOverlay(DOMProps{})), "drawer-overlay")
	assertSlot(t, "DrawerContent", mustRender(t, DrawerContent(DOMProps{})), "drawer-content")
	assertSlot(t, "DrawerClose", mustRender(t, DrawerClose(DOMProps{})), "drawer-close")
}

func TestFloatingFamilySlotsAreDistinct(t *testing.T) {
	assertSlot(t, "PopoverTrigger", mustRender(t, PopoverTrigger(DOMProps{})), "popover-trigger")
	assertSlot(t, "PopoverHeader", mustRender(t, PopoverHeader(DOMProps{})), "popover-header")
	assertSlot(t, "PopoverTitle", mustRender(t, PopoverTitle(DOMProps{})), "popover-title")
	assertSlot(t, "PopoverDescription", mustRender(t, PopoverDescription(DOMProps{})), "popover-description")

	assertSlot(t, "HoverCardTrigger", mustRender(t, HoverCardTrigger(DOMProps{})), "hover-card-trigger")
	assertSlot(t, "TooltipTrigger", mustRender(t, TooltipTrigger(DOMProps{})), "tooltip-trigger")

	assertSlot(t, "SelectLabel", mustRender(t, SelectLabel(DOMProps{})), "select-label")
	assertSlot(t, "SelectItem", mustRender(t, SelectItem(DropdownMenuItemProps{})), "select-item")
	assertSlot(t, "SelectSeparator", mustRender(t, SelectSeparator(DOMProps{})), "select-separator")
	assertSlot(t, "SelectScrollUpButton", mustRender(t, SelectScrollUpButton(DOMProps{})), "select-scroll-up-button")
	assertSlot(t, "SelectScrollDownButton", mustRender(t, SelectScrollDownButton(DOMProps{})), "select-scroll-down-button")
}

func TestMenuFamilySlotsAreDistinct(t *testing.T) {
	assertSlot(t, "DropdownMenuPortal", mustRender(t, DropdownMenuPortal(DOMProps{})), "dropdown-menu-portal")
	assertSlot(t, "DropdownMenuSubContent", mustRender(t, DropdownMenuSubContent(DOMProps{})), "dropdown-menu-sub-content")

	assertSlot(t, "ContextMenu", mustRender(t, ContextMenu(ContextMenuProps{})), "context-menu")
	assertSlot(t, "ContextMenuTrigger", mustRender(t, ContextMenuTrigger(DOMProps{})), "context-menu-trigger")
	assertSlot(t, "ContextMenuContent", mustRender(t, ContextMenuContent(DOMProps{})), "context-menu-content")
	assertSlot(t, "ContextMenuItem", mustRender(t, ContextMenuItem(DropdownMenuItemProps{})), "context-menu-item")
	assertSlot(t, "ContextMenuLabel", mustRender(t, ContextMenuLabel(DOMProps{})), "context-menu-label")
	assertSlot(t, "ContextMenuSeparator", mustRender(t, ContextMenuSeparator(DOMProps{})), "context-menu-separator")

	assertSlot(t, "Menubar", mustRender(t, Menubar(MenubarProps{})), "menubar")
	assertSlot(t, "MenubarTrigger", mustRender(t, MenubarTrigger(DOMProps{})), "menubar-trigger")
	assertSlot(t, "MenubarContent", mustRender(t, MenubarContent(DOMProps{})), "menubar-content")
	assertSlot(t, "MenubarItem", mustRender(t, MenubarItem(DropdownMenuItemProps{})), "menubar-item")
	assertSlot(t, "MenubarLabel", mustRender(t, MenubarLabel(DOMProps{})), "menubar-label")
	assertSlot(t, "MenubarSeparator", mustRender(t, MenubarSeparator(DOMProps{})), "menubar-separator")
}

func TestCompositeSlotsAreDistinct(t *testing.T) {
	assertSlot(t, "NavigationMenuList", mustRender(t, NavigationMenuList(DOMProps{})), "navigation-menu-list")
	assertSlot(t, "NavigationMenuItem", mustRender(t, NavigationMenuItem(DOMProps{})), "navigation-menu-item")
	assertSlot(t, "NavigationMenuTrigger", mustRender(t, NavigationMenuTrigger(DOMProps{})), "navigation-menu-trigger")
	assertSlot(t, "NavigationMenuContent", mustRender(t, NavigationMenuContent(DOMProps{})), "navigation-menu-content")
	assertSlot(t, "NavigationMenuLink", mustRender(t, NavigationMenuLink(BreadcrumbLinkProps{})), "navigation-menu-link")

	assertSlot(t, "ComboboxValue", mustRender(t, ComboboxValue(DOMProps{})), "combobox-value")
	assertSlot(t, "ComboboxTrigger", mustRender(t, ComboboxTrigger(SelectTriggerProps{})), "combobox-trigger")
	assertSlot(t, "ComboboxClear", mustRender(t, ComboboxClear(DOMProps{})), "combobox-clear")
	assertSlot(t, "ComboboxInput", mustRender(t, ComboboxInput(InputProps{})), "combobox-input")
	assertSlot(t, "ComboboxContent", mustRender(t, ComboboxContent(SelectContentProps{})), "combobox-content")
	assertSlot(t, "ComboboxItem", mustRender(t, ComboboxItem(DropdownMenuItemProps{})), "combobox-item")
	assertSlot(t, "ComboboxSeparator", mustRender(t, ComboboxSeparator(DOMProps{})), "combobox-separator")

	assertSlot(t, "CommandDialog", mustRender(t, CommandDialog(CommandProps{})), "command-dialog")
	assertSlot(t, "CommandInput", mustRender(t, CommandInput(InputProps{})), "command-input-wrapper")
	assertSlot(t, "CommandList", mustRender(t, CommandList(DOMProps{})), "command-list")
	assertSlot(t, "CommandEmpty", mustRender(t, CommandEmpty(DOMProps{})), "command-empty")
	assertSlot(t, "CommandGroup", mustRender(t, CommandGroup(DOMProps{})), "command-group")
	assertSlot(t, "CommandItem", mustRender(t, CommandItem(DropdownMenuItemProps{})), "command-item")
	assertSlot(t, "CommandShortcut", mustRender(t, CommandShortcut(DOMProps{})), "command-shortcut")
	assertSlot(t, "CommandSeparator", mustRender(t, CommandSeparator(DOMProps{})), "command-separator")

	assertSlot(t, "InputGroupInput", mustRender(t, InputGroupInput(InputGroupInputProps{})), "input-group-control")
	assertSlot(t, "InputGroupTextarea", mustRender(t, InputGroupTextarea(InputGroupTextareaProps{})), "input-group-control")

	assertSlot(t, "SidebarInput", mustRender(t, SidebarInput(InputProps{})), "sidebar-input")
	assertSlot(t, "SidebarSeparator", mustRender(t, SidebarSeparator(DOMProps{})), "sidebar-separator")
	assertSlot(t, "SidebarMenuBadge", mustRender(t, SidebarMenuBadge(DOMProps{})), "sidebar-menu-badge")
	assertSlot(t, "SidebarMenuSkeleton", mustRender(t, SidebarMenuSkeleton(DOMProps{})), "sidebar-menu-skeleton")
	assertSlot(t, "SidebarMenuSubButton", mustRender(t, SidebarMenuSubButton(SidebarMenuButtonProps{})), "sidebar-menu-sub-button")
}
