package ui

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestPopoverPartsReflectClosedAndOpenState(t *testing.T) {
	closedRoot := mustRender(t, Popover(PopoverProps{}))
	closedContent := renderWithFloatingState(t, false, PopoverContent(DOMProps{}))
	if attrValue(closedRoot, "data-state") != "closed" {
		t.Fatalf("Popover root data-state = %q, want closed", attrValue(closedRoot, "data-state"))
	}
	if !strings.Contains(closedContent, `data-state="closed"`) || !strings.Contains(closedContent, `hidden`) {
		t.Fatalf("closed popover content should be hidden, got %s", closedContent)
	}

	openRoot := mustRender(t, Popover(PopoverProps{Open: true}))
	openContent := renderWithFloatingState(t, true, PopoverContent(DOMProps{}))
	if attrValue(openRoot, "data-state") != "open" || attrValue(openRoot, "data-open") != "true" {
		t.Fatalf("open popover root state mismatch, got %s", openRoot)
	}
	if !strings.Contains(openContent, `data-state="open"`) || strings.Contains(openContent, ` hidden`) {
		t.Fatalf("open popover content should not be hidden, got %s", openContent)
	}
}

func TestFloatingTriggerAccessibilityAttributes(t *testing.T) {
	cases := map[string]templ.Component{
		"PopoverTrigger":   PopoverTrigger(DOMProps{}),
		"TooltipTrigger":   TooltipTrigger(DOMProps{}),
		"HoverCardTrigger": HoverCardTrigger(DOMProps{}),
	}
	for name, component := range cases {
		html := mustRender(t, component)
		if attrValue(html, "type") != "button" {
			t.Fatalf("%s type = %q, want button in %s", name, attrValue(html, "type"), html)
		}
	}

	popover := mustRender(t, PopoverTrigger(DOMProps{}))
	if attrValue(popover, "aria-haspopup") != "dialog" || attrValue(popover, "aria-expanded") != "false" {
		t.Fatalf("PopoverTrigger aria attrs mismatch, got %s", popover)
	}
}

func TestTooltipContentIsTooltipAndClosedByDefault(t *testing.T) {
	html := renderWithFloatingState(t, false, TooltipContent(DOMProps{}))
	if !strings.Contains(html, `role="tooltip"`) || !strings.Contains(html, `data-state="closed"`) || !strings.Contains(html, `hidden`) {
		t.Fatalf("closed tooltip content should be hidden tooltip, got %s", html)
	}
}

func TestHoverCardPartsReflectClosedAndOpenState(t *testing.T) {
	closed := renderWithFloatingState(t, false, HoverCardContent(DOMProps{}))
	if !strings.Contains(closed, `data-state="closed"`) || !strings.Contains(closed, `hidden`) {
		t.Fatalf("closed hover card content should be hidden, got %s", closed)
	}

	openRoot := mustRender(t, HoverCard(HoverCardProps{DefaultOpen: true}))
	open := renderWithFloatingState(t, true, HoverCardContent(DOMProps{}))
	if attrValue(openRoot, "data-state") != "open" || attrValue(openRoot, "data-default-open") != "true" {
		t.Fatalf("open hover card root state mismatch, got %s", openRoot)
	}
	if !strings.Contains(open, `data-state="open"`) || strings.Contains(open, ` hidden`) {
		t.Fatalf("open hover card content should not be hidden, got %s", open)
	}
}

func renderWithFloatingState(t *testing.T, open bool, components ...templ.Component) string {
	t.Helper()
	var buf bytes.Buffer
	ctx := context.WithValue(context.Background(), floatingRenderStateKey{}, floatingRenderState{open: open})
	for _, component := range components {
		if err := component.Render(ctx, &buf); err != nil {
			t.Fatalf("render error: %v", err)
		}
	}
	return buf.String()
}
