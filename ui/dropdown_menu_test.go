package ui

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestDropdownMenuTriggerAccessibilityAttributes(t *testing.T) {
	html := mustRender(t, DropdownMenuTrigger(DOMProps{}))
	if attrValue(html, "type") != "button" {
		t.Fatalf("DropdownMenuTrigger type = %q, want button", attrValue(html, "type"))
	}
	if attrValue(html, "aria-haspopup") != "menu" {
		t.Fatalf("DropdownMenuTrigger aria-haspopup = %q, want menu", attrValue(html, "aria-haspopup"))
	}
	if attrValue(html, "aria-expanded") != "false" {
		t.Fatalf("DropdownMenuTrigger aria-expanded = %q, want false", attrValue(html, "aria-expanded"))
	}
}

func TestDropdownMenuContentReflectsOpenState(t *testing.T) {
	closed := renderWithMenuState(t, false, DropdownMenuContent(DOMProps{}))
	if !strings.Contains(closed, `role="menu"`) || !strings.Contains(closed, `data-state="closed"`) || !strings.Contains(closed, `hidden`) {
		t.Fatalf("closed DropdownMenuContent should be hidden menu, got %s", closed)
	}

	open := renderWithMenuState(t, true, DropdownMenuContent(DOMProps{}))
	if !strings.Contains(open, `role="menu"`) || !strings.Contains(open, `data-state="open"`) {
		t.Fatalf("open DropdownMenuContent should render open menu, got %s", open)
	}
	if strings.Contains(open, ` hidden`) {
		t.Fatalf("open DropdownMenuContent should not be hidden, got %s", open)
	}
}

func TestDropdownMenuItemAccessibilityAttributes(t *testing.T) {
	html := mustRender(t, DropdownMenuItem(DropdownMenuItemProps{Value: "profile"}))
	if attrValue(html, "type") != "button" {
		t.Fatalf("DropdownMenuItem type = %q, want button", attrValue(html, "type"))
	}
	if attrValue(html, "role") != "menuitem" {
		t.Fatalf("DropdownMenuItem role = %q, want menuitem", attrValue(html, "role"))
	}
	if attrValue(html, "data-value") != "profile" {
		t.Fatalf("DropdownMenuItem data-value = %q, want profile", attrValue(html, "data-value"))
	}
}

func renderWithMenuState(t *testing.T, open bool, components ...templ.Component) string {
	t.Helper()
	var buf bytes.Buffer
	ctx := context.WithValue(context.Background(), menuRenderStateKey{}, menuRenderState{open: open})
	for _, component := range components {
		if err := component.Render(ctx, &buf); err != nil {
			t.Fatalf("render error: %v", err)
		}
	}
	return buf.String()
}
