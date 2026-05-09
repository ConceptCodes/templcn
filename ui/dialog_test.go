package ui

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestDialogPartsRenderClosedStateByDefault(t *testing.T) {
	html := mustRender(t, Dialog(DialogProps{}))
	childHTML := renderWithDialogState(t, false, DialogOverlay(DOMProps{}), DialogContent(DOMProps{}))

	if attrValue(html, "role") != "" {
		t.Fatalf("Dialog root role = %q, want empty wrapper role", attrValue(html, "role"))
	}
	if !strings.Contains(childHTML, `data-slot="dialog-overlay"`) || !strings.Contains(childHTML, `data-slot="dialog-content"`) {
		t.Fatalf("expected overlay and content in %s", childHTML)
	}
	if strings.Count(html+childHTML, `data-state="closed"`) < 3 {
		t.Fatalf("expected root, overlay, and content to render closed state, got %s %s", html, childHTML)
	}
	if strings.Contains(html+childHTML, `data-state="open"`) {
		t.Fatalf("closed dialog should not render open child state, got %s %s", html, childHTML)
	}
}

func TestDialogPartsRenderOpenState(t *testing.T) {
	html := mustRender(t, Dialog(DialogProps{Open: true}))
	childHTML := renderWithDialogState(t, true, DialogOverlay(DOMProps{}), DialogContent(DOMProps{}))

	if strings.Contains(html, ` open`) {
		t.Fatal("dialog root wrapper should not render native open attribute")
	}
	if !strings.Contains(childHTML, `data-slot="dialog-content"`) || !strings.Contains(childHTML, ` open`) {
		t.Fatalf("open dialog content should render native open attribute, got %s", childHTML)
	}
	if strings.Count(html+childHTML, `data-state="open"`) < 3 {
		t.Fatalf("expected root, overlay, and content to render open state, got %s %s", html, childHTML)
	}
	if strings.Contains(html+childHTML, `data-state="closed"`) {
		t.Fatalf("open dialog should not render closed child state, got %s %s", html, childHTML)
	}
}

func TestDialogTriggerAccessibilityAttributes(t *testing.T) {
	html := mustRender(t, DialogTrigger(DOMProps{}))
	if attrValue(html, "type") != "button" {
		t.Fatalf("DialogTrigger type = %q, want button", attrValue(html, "type"))
	}
	if attrValue(html, "aria-haspopup") != "dialog" {
		t.Fatalf("DialogTrigger aria-haspopup = %q, want dialog", attrValue(html, "aria-haspopup"))
	}
	if attrValue(html, "aria-expanded") != "false" {
		t.Fatalf("DialogTrigger aria-expanded = %q, want false", attrValue(html, "aria-expanded"))
	}
}

func TestOverlayFamilyAccessibilityAttributes(t *testing.T) {
	for name, component := range map[string]templ.Component{
		"AlertDialogTrigger": AlertDialogTrigger(DOMProps{}),
		"SheetTrigger":       SheetTrigger(DOMProps{}),
		"DrawerTrigger":      DrawerTrigger(DOMProps{}),
		"AlertDialogAction":  AlertDialogAction(DOMProps{}),
		"AlertDialogCancel":  AlertDialogCancel(DOMProps{}),
		"SheetClose":         SheetClose(DOMProps{}),
		"DrawerClose":        DrawerClose(DOMProps{}),
	} {
		html := mustRender(t, component)
		if attrValue(html, "type") != "button" {
			t.Fatalf("%s type = %q, want button in %s", name, attrValue(html, "type"), html)
		}
	}
}

func TestDialogFamilyContentOwnsNativeDialogElement(t *testing.T) {
	cases := map[string]templ.Component{
		"AlertDialogContent": renderWithAlertDialogStateComponent(true, AlertDialogContent(DOMProps{})),
		"SheetContent":       renderWithSheetStateComponent(true, SheetContent(DOMProps{})),
		"DrawerContent":      renderWithDrawerStateComponent(true, DrawerContent(DOMProps{})),
	}
	for name, component := range cases {
		html := mustRender(t, component)
		if !strings.Contains(html, `<dialog`) || !strings.Contains(html, ` open`) {
			t.Fatalf("%s should render open native dialog content, got %s", name, html)
		}
		wantRole := map[string]string{
			"AlertDialogContent": "alertdialog",
			"SheetContent":       "dialog",
			"DrawerContent":      "dialog",
		}[name]
		if attrValue(html, "role") != wantRole {
			t.Fatalf("%s role = %q, want %q in %s", name, attrValue(html, "role"), wantRole, html)
		}
	}
}

func renderWithDialogState(t *testing.T, open bool, components ...templ.Component) string {
	t.Helper()
	var buf bytes.Buffer
	ctx := context.WithValue(context.Background(), dialogRenderStateKey{}, dialogRenderState{open: open, slot: "dialog"})
	for _, component := range components {
		if err := component.Render(ctx, &buf); err != nil {
			t.Fatalf("render error: %v", err)
		}
	}
	return buf.String()
}

func renderWithAlertDialogStateComponent(open bool, component templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		ctx = context.WithValue(ctx, dialogRenderStateKey{}, dialogRenderState{open: open, slot: "alert-dialog", modal: true})
		return component.Render(ctx, w)
	})
}

func renderWithSheetStateComponent(open bool, component templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		ctx = context.WithValue(ctx, dialogRenderStateKey{}, dialogRenderState{open: open, slot: "sheet", modal: true})
		return component.Render(ctx, w)
	})
}

func renderWithDrawerStateComponent(open bool, component templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		ctx = context.WithValue(ctx, dialogRenderStateKey{}, dialogRenderState{open: open, slot: "drawer", modal: true})
		return component.Render(ctx, w)
	})
}
