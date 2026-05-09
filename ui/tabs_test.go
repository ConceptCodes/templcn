package ui

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestTabsRootExposesValue(t *testing.T) {
	html := mustRender(t, Tabs(TabsProps{DefaultValue: "account"}))
	if attrValue(html, "data-value") != "account" {
		t.Fatalf("Tabs data-value = %q, want account", attrValue(html, "data-value"))
	}
}

func TestTabsTriggerAndContentReflectContextValue(t *testing.T) {
	html := renderWithTabsValue(t, "account",
		TabsTrigger(TabsTriggerProps{Value: "account"}),
		TabsContent(TabsContentProps{Value: "account"}),
		TabsTrigger(TabsTriggerProps{Value: "password"}),
		TabsContent(TabsContentProps{Value: "password"}),
	)

	if strings.Count(html, `data-state="active"`) != 2 {
		t.Fatalf("expected active trigger and panel only, got %s", html)
	}
	if strings.Count(html, `data-state="inactive"`) != 2 {
		t.Fatalf("expected inactive trigger and panel only, got %s", html)
	}
	if strings.Count(html, `hidden`) < 1 {
		t.Fatalf("inactive panel should be hidden, got %s", html)
	}
	if strings.Count(html, `role="tab"`) != 2 || strings.Count(html, `role="tabpanel"`) != 2 {
		t.Fatalf("tabs should render tab and tabpanel roles, got %s", html)
	}
}

func TestTabsTriggerButtonTypeAndDisabledState(t *testing.T) {
	html := mustRender(t, TabsTrigger(TabsTriggerProps{Disabled: true}))
	if attrValue(html, "type") != "button" {
		t.Fatalf("TabsTrigger type = %q, want button", attrValue(html, "type"))
	}
	if attrValue(html, "aria-disabled") != "true" {
		t.Fatalf("TabsTrigger aria-disabled = %q, want true", attrValue(html, "aria-disabled"))
	}
}

func renderWithTabsValue(t *testing.T, value string, components ...templ.Component) string {
	t.Helper()
	var buf bytes.Buffer
	ctx := context.WithValue(context.Background(), tabsRenderStateKey{}, tabsRenderState{value: value})
	for _, component := range components {
		if err := component.Render(ctx, &buf); err != nil {
			t.Fatalf("render error: %v", err)
		}
	}
	return buf.String()
}
