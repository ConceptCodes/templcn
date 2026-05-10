package ui

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestSelectRootRendersFormValueAndState(t *testing.T) {
	html := mustRender(t, Select(SelectProps{Name: "plan", DefaultValue: "pro"}))
	if attrValue(html, "data-slot") != "select" {
		t.Fatalf("Select data-slot = %q", attrValue(html, "data-slot"))
	}
	if attrValue(html, "data-state") != "closed" {
		t.Fatalf("Select data-state = %q, want closed", attrValue(html, "data-state"))
	}
	if attrValue(html, "data-value") != "pro" {
		t.Fatalf("Select data-value = %q, want pro", attrValue(html, "data-value"))
	}
	if !strings.Contains(html, `type="hidden"`) || !strings.Contains(html, `name="plan"`) || !strings.Contains(html, `value="pro"`) {
		t.Fatalf("Select should render hidden form input, got %s", html)
	}
}

func TestSelectTriggerAccessibilityAttributes(t *testing.T) {
	html := mustRender(t, SelectTrigger(SelectTriggerProps{}))
	if attrValue(html, "type") != "button" {
		t.Fatalf("SelectTrigger type = %q, want button", attrValue(html, "type"))
	}
	if attrValue(html, "role") != "combobox" {
		t.Fatalf("SelectTrigger role = %q, want combobox", attrValue(html, "role"))
	}
	if attrValue(html, "aria-haspopup") != "listbox" {
		t.Fatalf("SelectTrigger aria-haspopup = %q, want listbox", attrValue(html, "aria-haspopup"))
	}
	if attrValue(html, "aria-expanded") != "false" {
		t.Fatalf("SelectTrigger aria-expanded = %q, want false", attrValue(html, "aria-expanded"))
	}
}

func TestSelectContentAndItemsReflectState(t *testing.T) {
	closed := renderWithSelectState(t, false, "pro", SelectContent(SelectContentProps{}), SelectItem(DropdownMenuItemProps{Value: "pro"}))
	if !strings.Contains(closed, `role="listbox"`) || !strings.Contains(closed, `hidden`) {
		t.Fatalf("closed select content should be hidden listbox, got %s", closed)
	}
	if !strings.Contains(closed, `aria-selected="true"`) || !strings.Contains(closed, `data-state="checked"`) {
		t.Fatalf("selected item should render checked state, got %s", closed)
	}

	open := renderWithSelectState(t, true, "free", SelectContent(SelectContentProps{}), SelectItem(DropdownMenuItemProps{Value: "pro"}))
	if strings.Contains(open, ` hidden`) {
		t.Fatalf("open select content should not be hidden, got %s", open)
	}
	if !strings.Contains(open, `aria-selected="false"`) || !strings.Contains(open, `data-state="unchecked"`) {
		t.Fatalf("unselected item should render unchecked state, got %s", open)
	}
}

func renderWithSelectState(t *testing.T, open bool, value string, components ...templ.Component) string {
	t.Helper()
	var buf bytes.Buffer
	ctx := context.WithValue(context.Background(), selectRenderStateKey{}, selectRenderState{open: open, value: value})
	for _, component := range components {
		if err := component.Render(ctx, &buf); err != nil {
			t.Fatalf("render error: %v", err)
		}
	}
	return buf.String()
}
