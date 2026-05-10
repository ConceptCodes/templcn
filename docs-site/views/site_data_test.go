package views

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"
)

func TestComponentUsageSnippetsAreConcrete(t *testing.T) {
	for _, doc := range ComponentIndex() {
		if strings.Contains(doc.GoUsage, "{ ... }") || strings.Contains(doc.GoUsage, "{...}") {
			t.Fatalf("%s GoUsage contains placeholder children: %s", doc.Slug, doc.GoUsage)
		}
		if strings.Contains(doc.GoUsage, "ui.SelectItemProps") {
			t.Fatalf("%s GoUsage references a non-existent SelectItemProps type", doc.Slug)
		}
	}
}

func TestGetStartedPagesRender(t *testing.T) {
	pages := map[string]interface {
		Render(context.Context, io.Writer) error
	}{
		"components-json": ComponentsJSONPage(),
		"package-imports": PackageImportsPage(),
		"dark-mode":       DarkModePage(),
		"monorepo":        MonorepoPage(),
		"javascript":      JavaScriptPage(),
		"llms":            LLMSPage(),
	}
	for name, page := range pages {
		var buf bytes.Buffer
		if err := page.Render(context.Background(), &buf); err != nil {
			t.Fatalf("%s render: %v", name, err)
		}
		if !strings.Contains(buf.String(), "<h1") {
			t.Fatalf("%s rendered without h1", name)
		}
	}
}

func TestLLMSTextIncludesProjectContext(t *testing.T) {
	text := LLMSText()
	required := []string{
		"# templcn/ui",
		"/docs/components/button",
		"templcn add button",
		"ui/*.go",
		"assets/runtime.js",
		"styles/globals.css",
		"@theme inline",
	}
	for _, fragment := range required {
		if !strings.Contains(text, fragment) {
			t.Fatalf("llms.txt missing %q", fragment)
		}
	}
}

func TestComponentDocsHaveParitySections(t *testing.T) {
	for _, doc := range ComponentIndex() {
		if doc.Install != "templcn add "+doc.Slug {
			t.Fatalf("%s install command = %q", doc.Slug, doc.Install)
		}
		if len(doc.Composition) == 0 {
			t.Fatalf("%s missing composition parts", doc.Slug)
		}
		if len(doc.Examples) == 0 {
			t.Fatalf("%s missing examples", doc.Slug)
		}
		if strings.TrimSpace(doc.SourceCode) == "" {
			t.Fatalf("%s missing source code", doc.Slug)
		}
		if strings.Contains(doc.SourceCode, "func "+strings.ReplaceAll(doc.Title, " ", "")+"Preview() templ.Component") {
			t.Fatalf("%s still uses generated placeholder source", doc.Slug)
		}
	}
}
