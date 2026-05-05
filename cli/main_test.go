package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSanitizeProjectName(t *testing.T) {
	got := sanitizeProjectName("My App!")
	if got != "my-app" {
		t.Fatalf("expected my-app, got %q", got)
	}
}

func TestBuildRegistryIndexFindsButton(t *testing.T) {
	index, err := buildRegistryIndex()
	if err != nil {
		t.Fatalf("build registry index: %v", err)
	}
	if _, ok := index.Component["button"]; !ok {
		t.Fatalf("expected button component to be indexed")
	}
}

func TestInitProjectScaffoldsFiles(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("SHADCN_SOURCE_DIR", repoRoot(t))

	if err := InitProject(InitOptions{CWD: dir, Name: "demo-app", Force: true}); err != nil {
		t.Fatalf("init project: %v", err)
	}

	root := filepath.Join(dir, "demo-app")
	mustExist(t, filepath.Join(root, "go.mod"))
	mustExist(t, filepath.Join(root, "main.go"))
	mustExist(t, filepath.Join(root, "app.templ"))
	mustExist(t, filepath.Join(root, "components.json"))
	mustExist(t, filepath.Join(root, "styles", "globals.css"))
	mustExist(t, filepath.Join(root, "ui", "button.go"))
}

func TestAddComponentsCopiesUIPackage(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("SHADCN_SOURCE_DIR", repoRoot(t))

	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/demo\n\ngo 1.23.0\n"), 0644); err != nil {
		t.Fatalf("write go.mod: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "components.json"), []byte(`{"module":"example.com/demo","uiDir":"ui"}`), 0644); err != nil {
		t.Fatalf("write components.json: %v", err)
	}

	if err := AddComponents(AddOptions{CWD: dir, Items: []string{"button"}, Overwrite: true}); err != nil {
		t.Fatalf("add components: %v", err)
	}

	mustExist(t, filepath.Join(dir, "ui", "button.go"))
	mustExist(t, filepath.Join(dir, "ui", "input.go"))
	mustExist(t, filepath.Join(dir, "ui", "textarea.go"))
}

func TestApplyPresetUpdatesConfig(t *testing.T) {
	dir := t.TempDir()

	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/demo\n\ngo 1.23.0\n"), 0644); err != nil {
		t.Fatalf("write go.mod: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "components.json"), []byte(`{"module":"example.com/demo","uiDir":"ui","style":"default"}`), 0644); err != nil {
		t.Fatalf("write components.json: %v", err)
	}

	if err := ApplyPreset(ApplyOptions{CWD: dir, Preset: "new-york", Silent: true}); err != nil {
		t.Fatalf("apply preset: %v", err)
	}

	raw, err := os.ReadFile(filepath.Join(dir, "components.json"))
	if err != nil {
		t.Fatalf("read components.json: %v", err)
	}
	if !strings.Contains(string(raw), `"style": "new-york"`) {
		t.Fatalf("expected style to be updated, got %s", raw)
	}
	mustExist(t, filepath.Join(dir, "assets", "runtime.js"))
	mustExist(t, filepath.Join(dir, "styles", "globals.css"))
}

func repoRoot(t *testing.T) string {
	t.Helper()
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	root := cwd
	for {
		if strings.HasSuffix(root, string(filepath.Separator)+"cli") {
			root = filepath.Dir(root)
			break
		}
		parent := filepath.Dir(root)
		if parent == root {
			t.Fatalf("could not find repo root from %s", cwd)
		}
		root = parent
	}
	return root
}

func mustExist(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected %s to exist: %v", path, err)
	}
}
