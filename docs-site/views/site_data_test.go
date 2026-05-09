package views

import (
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
