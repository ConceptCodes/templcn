package ui

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestPaginationPreviousAndNextRenderWithoutRecursion(t *testing.T) {
	for _, tc := range []struct {
		name      string
		component func() templ.Component
		label     string
	}{
		{name: "previous", component: func() templ.Component { return PaginationPrevious(PaginationLinkProps{}) }, label: "Previous"},
		{name: "next", component: func() templ.Component { return PaginationNext(PaginationLinkProps{}) }, label: "Next"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			html := mustRender(t, tc.component())
			if !strings.Contains(html, tc.label) {
				t.Fatalf("expected default %q label in %q", tc.label, html)
			}
		})
	}
}

func TestPaginationPreviousAndNextRenderCallerChildren(t *testing.T) {
	custom := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, "custom")
		return err
	})

	for _, tc := range []struct {
		name      string
		component func(PaginationLinkProps) templ.Component
	}{
		{name: "previous", component: func(props PaginationLinkProps) templ.Component { return PaginationPrevious(props) }},
		{name: "next", component: func(props PaginationLinkProps) templ.Component { return PaginationNext(props) }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			wrapper := templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
				return tc.component(PaginationLinkProps{}).Render(templ.WithChildren(ctx, custom), w)
			})
			html := mustRender(t, wrapper)
			if !strings.Contains(html, "custom") {
				t.Fatalf("expected caller children in %q", html)
			}
		})
	}
}
