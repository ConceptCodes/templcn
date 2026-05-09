package ui

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestRadioGroupPropagatesNameValueAndState(t *testing.T) {
	html := renderRadioGroupWithItems(t)
	if !strings.Contains(html, `role="radiogroup"`) || !strings.Contains(html, `data-value="comfortable"`) {
		t.Fatalf("radio group attrs mismatch, got %s", html)
	}
	if !strings.Contains(html, `name="density"`) || !strings.Contains(html, `value="comfortable"`) || !strings.Contains(html, `checked`) {
		t.Fatalf("radio item should inherit name and checked state, got %s", html)
	}
	if !strings.Contains(html, `data-state="checked"`) || !strings.Contains(html, `aria-checked="true"`) {
		t.Fatalf("radio checked state missing, got %s", html)
	}
}

func TestToggleGroupItemsReflectPressedState(t *testing.T) {
	html := renderToggleGroupWithItems(t)
	if !strings.Contains(html, `role="group"`) || !strings.Contains(html, `data-value="bold"`) {
		t.Fatalf("toggle group attrs mismatch, got %s", html)
	}
	if !strings.Contains(html, `data-state="on"`) || !strings.Contains(html, `aria-pressed="true"`) {
		t.Fatalf("selected toggle attrs missing, got %s", html)
	}
	if !strings.Contains(html, `type="button"`) {
		t.Fatalf("toggle item should render button type, got %s", html)
	}
}

func TestSliderRendersRangeInputWithValueAndAria(t *testing.T) {
	html := mustRender(t, Slider(SliderProps{Name: "volume", Value: []float64{40}, Max: 100, Step: 5}))
	if attrValue(html, "data-slot") != "slider" || attrValue(html, "data-value") != "40" {
		t.Fatalf("slider root attrs mismatch, got %s", html)
	}
	if !strings.Contains(html, `type="range"`) || !strings.Contains(html, `name="volume"`) || !strings.Contains(html, `aria-valuenow="40"`) {
		t.Fatalf("slider input attrs mismatch, got %s", html)
	}
}

func renderRadioGroupWithItems(t *testing.T) string {
	t.Helper()
	children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return RadioGroupItem(RadioGroupItemProps{Value: "comfortable", Label: "Comfortable"}).Render(templ.ClearChildren(ctx), w)
	})
	var buf bytes.Buffer
	if err := RadioGroup(RadioGroupProps{Name: "density", DefaultValue: "comfortable"}).Render(templ.WithChildren(context.Background(), children), &buf); err != nil {
		t.Fatalf("render error: %v", err)
	}
	return buf.String()
}

func renderToggleGroupWithItems(t *testing.T) string {
	t.Helper()
	children := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return ToggleGroupItem(ToggleGroupItemProps{Value: "bold"}).Render(templ.ClearChildren(ctx), w)
	})
	var buf bytes.Buffer
	if err := ToggleGroup(ToggleGroupProps{Type: "single", DefaultValue: []string{"bold"}}).Render(templ.WithChildren(context.Background(), children), &buf); err != nil {
		t.Fatalf("render error: %v", err)
	}
	return buf.String()
}
