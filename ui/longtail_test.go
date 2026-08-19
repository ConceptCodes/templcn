package ui

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestCheckboxAndSwitchExposeCheckedState(t *testing.T) {
	checkbox := mustRender(t, Checkbox(CheckboxProps{Name: "terms", DefaultChecked: true}))
	if attrValue(checkbox, "data-state") != "checked" || attrValue(checkbox, "aria-checked") != "true" || !hasAttr(checkbox, "checked") {
		t.Fatalf("checkbox checked attrs mismatch, got %s", checkbox)
	}

	sw := mustRender(t, Switch(SwitchProps{Name: "dark", Checked: true}))
	if attrValue(sw, "role") != "switch" || attrValue(sw, "data-state") != "checked" || attrValue(sw, "aria-checked") != "true" {
		t.Fatalf("switch attrs mismatch, got %s", sw)
	}
}

func TestCollapsibleContentReflectsClosedState(t *testing.T) {
	html := mustRender(t, Collapsible(CollapsibleProps{}))
	if attrValue(html, "data-state") != "closed" {
		t.Fatalf("collapsible state = %q, want closed", attrValue(html, "data-state"))
	}
	content := mustRender(t, CollapsibleContent(DOMProps{}))
	if !strings.Contains(content, `data-state="closed"`) || !strings.Contains(content, `hidden`) {
		t.Fatalf("collapsible content should be hidden by default, got %s", content)
	}
}

func TestInputOTPRendersHiddenValueAndAccessibleSlots(t *testing.T) {
	root := mustRender(t, InputOTP(InputOTPProps{Name: "code", DefaultValue: "12"}))
	if !strings.Contains(root, `type="hidden"`) || !strings.Contains(root, `name="code"`) || !strings.Contains(root, `value="12"`) {
		t.Fatalf("input otp hidden value mismatch, got %s", root)
	}
	slot := mustRender(t, InputOTPSlot(InputOTPSlotProps{Index: 0, Value: "1"}))
	if attrValue(slot, "role") != "textbox" || attrValue(slot, "tabindex") != "0" {
		t.Fatalf("input otp slot attrs mismatch, got %s", slot)
	}
}

func TestLongTailInteractiveAttrs(t *testing.T) {
	carousel := mustRender(t, Carousel(CarouselProps{StartIndex: 1}))
	if attrValue(carousel, "role") != "region" || attrValue(carousel, "aria-roledescription") != "carousel" || attrValue(carousel, "data-index") != "1" {
		t.Fatalf("carousel attrs mismatch, got %s", carousel)
	}

	handle := mustRender(t, ResizableHandle(ResizableHandleProps{}))
	if attrValue(handle, "role") != "separator" || attrValue(handle, "tabindex") != "0" {
		t.Fatalf("resizable handle attrs mismatch, got %s", handle)
	}

	toast := mustRender(t, Toast(ToastProps{DefaultOpen: true}))
	if attrValue(toast, "role") != "status" || attrValue(toast, "data-state") != "open" {
		t.Fatalf("toast attrs mismatch, got %s", toast)
	}

	sidebar := mustRender(t, SidebarProvider(SidebarProviderProps{DefaultOpen: true}))
	if attrValue(sidebar, "data-slot") != "sidebar-provider" || attrValue(sidebar, "data-state") != "open" {
		t.Fatalf("sidebar provider attrs mismatch, got %s", sidebar)
	}

	datePicker := mustRender(t, DatePicker(DatePickerProps{Name: "date", Value: "2026-05-08"}))
	if attrValue(datePicker, "data-state") != "closed" || !strings.Contains(datePicker, `type="hidden"`) || !strings.Contains(datePicker, `value="2026-05-08"`) {
		t.Fatalf("date picker attrs mismatch, got %s", datePicker)
	}
}

func TestButtonLikeControlsDoNotSubmitByDefault(t *testing.T) {
	inputGroupButton := mustRender(t, InputGroupButton(InputGroupButtonProps{}))
	if attrValue(inputGroupButton, "type") != "button" {
		t.Fatalf("input group button type mismatch, got %s", inputGroupButton)
	}

	toggle := mustRender(t, Toggle(ToggleProps{DefaultPressed: true}))
	if attrValue(toggle, "type") != "button" || attrValue(toggle, "aria-pressed") != "true" || attrValue(toggle, "data-state") != "on" {
		t.Fatalf("toggle pressed attrs mismatch, got %s", toggle)
	}

	offToggle := mustRender(t, Toggle(ToggleProps{}))
	if attrValue(offToggle, "aria-pressed") != "false" || attrValue(offToggle, "data-state") != "off" {
		t.Fatalf("toggle off attrs mismatch, got %s", offToggle)
	}
}

func TestSkeletonDoesNotRenderAmbientChildren(t *testing.T) {
	html := mustRender(t, templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return Card(DOMProps{}).Render(templ.WithChildren(ctx, Skeleton(DOMProps{Class: "h-4 w-24"})), w)
	}))
	if strings.Count(html, `data-slot="skeleton"`) != 1 {
		t.Fatalf("skeleton should render once without recursive ambient children, got %s", html)
	}
}
