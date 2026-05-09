package ui

import (
	"strings"
	"testing"
)

func TestContextMenuAccessibilityAttrs(t *testing.T) {
	trigger := mustRender(t, ContextMenuTrigger(DOMProps{}))
	if attrValue(trigger, "aria-haspopup") != "menu" {
		t.Fatalf("ContextMenuTrigger aria-haspopup = %q, want menu", attrValue(trigger, "aria-haspopup"))
	}

	content := mustRender(t, ContextMenuContent(DOMProps{}))
	if !strings.Contains(content, `role="menu"`) || !strings.Contains(content, `data-state="closed"`) || !strings.Contains(content, `hidden`) {
		t.Fatalf("closed context menu content should be hidden menu, got %s", content)
	}

	item := mustRender(t, ContextMenuItem(DropdownMenuItemProps{Value: "copy", Disabled: true}))
	if attrValue(item, "type") != "button" || attrValue(item, "role") != "menuitem" || attrValue(item, "data-value") != "copy" {
		t.Fatalf("context menu item attrs mismatch, got %s", item)
	}
	if !hasAttr(item, "disabled") || attrValue(item, "aria-disabled") != "true" {
		t.Fatalf("disabled context menu item attrs mismatch, got %s", item)
	}
}

func TestMenubarAccessibilityAttrs(t *testing.T) {
	root := mustRender(t, Menubar(MenubarProps{}))
	if attrValue(root, "role") != "menubar" || attrValue(root, "data-state") != "closed" {
		t.Fatalf("menubar root attrs mismatch, got %s", root)
	}

	trigger := mustRender(t, MenubarTrigger(DOMProps{}))
	if attrValue(trigger, "type") != "button" || attrValue(trigger, "role") != "menuitem" || attrValue(trigger, "aria-haspopup") != "menu" || attrValue(trigger, "aria-expanded") != "false" {
		t.Fatalf("menubar trigger attrs mismatch, got %s", trigger)
	}

	content := mustRender(t, MenubarContent(DOMProps{}))
	if !strings.Contains(content, `role="menu"`) || !strings.Contains(content, `data-state="closed"`) || !strings.Contains(content, `hidden`) {
		t.Fatalf("closed menubar content should be hidden menu, got %s", content)
	}
}

func TestNavigationMenuAccessibilityAttrs(t *testing.T) {
	root := mustRender(t, NavigationMenu(NavigationMenuProps{}))
	if attrValue(root, "data-state") != "closed" {
		t.Fatalf("navigation menu data-state = %q, want closed", attrValue(root, "data-state"))
	}

	trigger := mustRender(t, NavigationMenuTrigger(DOMProps{}))
	if attrValue(trigger, "type") != "button" || attrValue(trigger, "aria-expanded") != "false" {
		t.Fatalf("navigation menu trigger attrs mismatch, got %s", trigger)
	}

	content := mustRender(t, NavigationMenuContent(DOMProps{}))
	if !strings.Contains(content, `data-state="closed"`) || !strings.Contains(content, `hidden`) {
		t.Fatalf("closed navigation menu content should be hidden, got %s", content)
	}
}
