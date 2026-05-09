package ui

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestAccordionItemReflectsDefaultValue(t *testing.T) {
	html := renderWithAccordionValues(t, []string{"billing"}, AccordionItem(AccordionItemProps{Value: "billing"}))
	if !strings.Contains(html, `open`) || !strings.Contains(html, `data-state="open"`) {
		t.Fatalf("accordion item should render open, got %s", html)
	}
}

func TestAccordionTriggerAndContentReflectItemState(t *testing.T) {
	open := renderWithAccordionItemState(t, true, AccordionTrigger(DOMProps{}), AccordionContent(DOMProps{}))
	if !strings.Contains(open, `aria-expanded="true"`) || !strings.Contains(open, `data-state="open"`) {
		t.Fatalf("open accordion parts should reflect state, got %s", open)
	}

	closed := renderWithAccordionItemState(t, false, AccordionTrigger(DOMProps{}), AccordionContent(DOMProps{}))
	if !strings.Contains(closed, `aria-expanded="false"`) || !strings.Contains(closed, `data-state="closed"`) {
		t.Fatalf("closed accordion parts should reflect state, got %s", closed)
	}
}

func renderWithAccordionValues(t *testing.T, values []string, components ...templ.Component) string {
	t.Helper()
	openValues := map[string]struct{}{}
	for _, value := range values {
		openValues[value] = struct{}{}
	}
	ctx := context.WithValue(context.Background(), accordionRenderStateKey{}, accordionRenderState{openValues: openValues})
	return renderComponentsWithContext(t, ctx, components...)
}

func renderWithAccordionItemState(t *testing.T, open bool, components ...templ.Component) string {
	t.Helper()
	ctx := context.WithValue(context.Background(), accordionItemOpenKey{}, open)
	return renderComponentsWithContext(t, ctx, components...)
}

func renderComponentsWithContext(t *testing.T, ctx context.Context, components ...templ.Component) string {
	t.Helper()
	var buf bytes.Buffer
	for _, component := range components {
		if err := component.Render(ctx, &buf); err != nil {
			t.Fatalf("render error: %v", err)
		}
	}
	return buf.String()
}
