package ui

import (
	"strings"
	"testing"
)

func TestCommandAccessibilityAttrs(t *testing.T) {
	root := mustRender(t, Command(CommandProps{DefaultValue: "beta"}))
	if attrValue(root, "role") != "combobox" || attrValue(root, "aria-expanded") != "true" || attrValue(root, "data-value") != "beta" {
		t.Fatalf("command root attrs mismatch, got %s", root)
	}

	input := mustRender(t, CommandInput(InputProps{}))
	if attrValue(input, "role") != "searchbox" || attrValue(input, "aria-autocomplete") != "list" {
		t.Fatalf("command input attrs mismatch, got %s", input)
	}

	list := mustRender(t, CommandList(DOMProps{}))
	if attrValue(list, "role") != "listbox" {
		t.Fatalf("command list role = %q, want listbox", attrValue(list, "role"))
	}

	item := mustRender(t, CommandItem(DropdownMenuItemProps{Value: "beta", Disabled: true}))
	if attrValue(item, "type") != "button" || attrValue(item, "role") != "option" || attrValue(item, "data-value") != "beta" {
		t.Fatalf("command item attrs mismatch, got %s", item)
	}
	if !hasAttr(item, "disabled") || attrValue(item, "aria-disabled") != "true" {
		t.Fatalf("disabled command item attrs mismatch, got %s", item)
	}
}

func TestComboboxRootTriggerContentAndItems(t *testing.T) {
	root := mustRender(t, Combobox(ComboboxProps{Name: "framework", DefaultValue: "go"}))
	if attrValue(root, "data-state") != "closed" || attrValue(root, "data-value") != "go" {
		t.Fatalf("combobox root state mismatch, got %s", root)
	}
	if !strings.Contains(root, `type="hidden"`) || !strings.Contains(root, `name="framework"`) || !strings.Contains(root, `value="go"`) {
		t.Fatalf("combobox should render hidden form input, got %s", root)
	}

	trigger := mustRender(t, ComboboxTrigger(SelectTriggerProps{}))
	if attrValue(trigger, "type") != "button" || attrValue(trigger, "role") != "combobox" || attrValue(trigger, "aria-haspopup") != "listbox" || attrValue(trigger, "aria-expanded") != "false" {
		t.Fatalf("combobox trigger attrs mismatch, got %s", trigger)
	}

	closed := renderWithSelectState(t, false, "go", ComboboxContent(SelectContentProps{}), ComboboxItem(DropdownMenuItemProps{Value: "go"}))
	if !strings.Contains(closed, `role="listbox"`) || !strings.Contains(closed, `hidden`) || !strings.Contains(closed, `aria-selected="true"`) {
		t.Fatalf("closed combobox content/item attrs mismatch, got %s", closed)
	}
}
