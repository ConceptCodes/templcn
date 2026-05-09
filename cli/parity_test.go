package main

import (
	"os"
	"testing"
)

func TestBuildParityReportFlagsMissingUpstreamComponents(t *testing.T) {
	t.Setenv("SHADCN_SOURCE_DIR", repoRoot(t))

	index, err := buildRegistryIndex()
	if err != nil {
		t.Fatalf("build registry index: %v", err)
	}
	report := buildParityReport("testdata/upstream.json", []upstreamRegistryItem{
		{Name: "button", Type: "registry:ui"},
		{Name: "missing-upstream-component", Type: "registry:ui"},
	}, index)

	if report.LocalCount == 0 {
		t.Fatal("expected local registry items")
	}
	if !containsString(report.Missing, "missing-upstream-component") {
		t.Fatalf("expected missing component in report, got %#v", report.Missing)
	}
	if containsString(report.Missing, "button") {
		t.Fatalf("button should be present locally, missing list: %#v", report.Missing)
	}
}

func TestUpstreamParitySnapshot(t *testing.T) {
	if os.Getenv("SHADCN_CHECK_UPSTREAM") != "1" {
		t.Skip("set SHADCN_CHECK_UPSTREAM=1 to compare against upstream shadcn/ui registry")
	}
	t.Setenv("SHADCN_SOURCE_DIR", repoRoot(t))

	report, err := checkParity("")
	if err != nil {
		t.Fatalf("check upstream parity: %v", err)
	}
	if len(report.Missing) > 0 {
		t.Fatalf("missing upstream components: %v", report.Missing)
	}
}

func containsString(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}
