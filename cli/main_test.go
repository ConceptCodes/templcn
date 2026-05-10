package main

import (
	"os"
	"os/exec"
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
	if _, ok := index.Items["data-table"]; !ok {
		t.Fatalf("expected data-table component to use public slug")
	}
	for _, internal := range []string{"cn", "html", "render", "types", "floating", "dialog-test"} {
		if _, ok := index.Items[internal]; ok {
			t.Fatalf("internal/test file %q should not be a registry item", internal)
		}
	}
	for name, item := range index.Items {
		if len(item.Files) == 0 {
			t.Fatalf("%s missing source files", name)
		}
		if len(item.Dependencies) == 0 {
			t.Fatalf("%s missing shared dependencies", name)
		}
		if len(item.CSS) == 0 {
			t.Fatalf("%s missing css metadata", name)
		}
		if item.DocsURL == "" {
			t.Fatalf("%s missing docs url", name)
		}
		if len(item.Examples) == 0 {
			t.Fatalf("%s missing example metadata", name)
		}
	}
}

func TestInitProjectScaffoldsFiles(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("TEMPLCN_SOURCE_DIR", repoRoot(t))

	if err := InitProject(InitOptions{CWD: dir, Name: "demo-app", Force: true}); err != nil {
		t.Fatalf("init project: %v", err)
	}

	root := filepath.Join(dir, "demo-app")
	mustExist(t, filepath.Join(root, "go.mod"))
	mustExist(t, filepath.Join(root, "main.go"))
	mustExist(t, filepath.Join(root, "app.templ"))
	mustExist(t, filepath.Join(root, "app_templ.go"))
	mustExist(t, filepath.Join(root, "components.json"))
	mustExist(t, filepath.Join(root, "styles", "globals.css"))
	mustExist(t, filepath.Join(root, "ui", "button.go"))
}

func TestInitProjectBuildsGeneratedStarter(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("TEMPLCN_SOURCE_DIR", repoRoot(t))

	if err := InitProject(InitOptions{CWD: dir, Name: "demo-build", Force: true, Silent: true}); err != nil {
		t.Fatalf("init project: %v", err)
	}

	root := filepath.Join(dir, "demo-build")
	tidy := exec.Command("go", "mod", "tidy")
	tidy.Dir = root
	if out, err := tidy.CombinedOutput(); err != nil {
		t.Fatalf("generated starter should tidy dependencies: %v\n%s", err, out)
	}
	cmd := exec.Command("go", "build", "./...")
	cmd.Dir = root
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("generated starter should build: %v\n%s", err, out)
	}
}

func TestAddComponentsCopiesOnlyRequestedComponentDependencies(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("TEMPLCN_SOURCE_DIR", repoRoot(t))

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
	mustExist(t, filepath.Join(dir, "ui", "cn.go"))
	mustExist(t, filepath.Join(dir, "ui", "html.go"))
	mustExist(t, filepath.Join(dir, "ui", "render.go"))
	mustExist(t, filepath.Join(dir, "ui", "types.go"))
	mustNotExist(t, filepath.Join(dir, "ui", "input.go"))
	mustNotExist(t, filepath.Join(dir, "ui", "textarea.go"))
	mustNotExist(t, filepath.Join(dir, "assets", "runtime.js"))
}

func TestAddSelectCopiesRuntimeAndDeclaredDependencies(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("TEMPLCN_SOURCE_DIR", repoRoot(t))

	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/demo\n\ngo 1.23.0\n"), 0644); err != nil {
		t.Fatalf("write go.mod: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "components.json"), []byte(`{"module":"example.com/demo","uiDir":"ui"}`), 0644); err != nil {
		t.Fatalf("write components.json: %v", err)
	}

	if err := AddComponents(AddOptions{CWD: dir, Items: []string{"select"}, Overwrite: true}); err != nil {
		t.Fatalf("add select: %v", err)
	}

	mustExist(t, filepath.Join(dir, "ui", "select.go"))
	mustExist(t, filepath.Join(dir, "ui", "dropdown_menu.go"))
	mustExist(t, filepath.Join(dir, "assets", "runtime.js"))
	mustNotExist(t, filepath.Join(dir, "ui", "button.go"))
}

func TestAddAllCopiesComponentsWithoutTestsOrInternalComponents(t *testing.T) {
	dir := newConfiguredProject(t)
	t.Setenv("TEMPLCN_SOURCE_DIR", repoRoot(t))

	if err := AddComponents(AddOptions{CWD: dir, All: true, Overwrite: true}); err != nil {
		t.Fatalf("add all: %v", err)
	}

	mustExist(t, filepath.Join(dir, "ui", "dialog.go"))
	mustExist(t, filepath.Join(dir, "ui", "calendar.go"))
	mustExist(t, filepath.Join(dir, "ui", "floating.go"))
	mustExist(t, filepath.Join(dir, "assets", "runtime.js"))
	mustNotExist(t, filepath.Join(dir, "ui", "dialog_test.go"))
	mustNotExist(t, filepath.Join(dir, "ui", "floating_test.go"))
}

func TestAddDryRunDoesNotWriteFiles(t *testing.T) {
	dir := newConfiguredProject(t)
	t.Setenv("TEMPLCN_SOURCE_DIR", repoRoot(t))

	if err := AddComponents(AddOptions{CWD: dir, Items: []string{"button"}, DryRun: true}); err != nil {
		t.Fatalf("dry run add: %v", err)
	}

	mustNotExist(t, filepath.Join(dir, "ui", "button.go"))
}

func TestViewAndDiffResolveRegistrySource(t *testing.T) {
	dir := newConfiguredProject(t)
	t.Setenv("TEMPLCN_SOURCE_DIR", repoRoot(t))

	if err := ViewItems(ViewOptions{CWD: dir, Items: []string{"button"}}); err != nil {
		t.Fatalf("view button: %v", err)
	}
	if err := DiffItems(DiffOptions{CWD: dir, Items: []string{"button"}}); err != nil {
		t.Fatalf("diff button: %v", err)
	}
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

func TestStarterRuntimeJavaScriptParses(t *testing.T) {
	node, err := exec.LookPath("node")
	if err != nil {
		t.Skip("node is not installed")
	}
	path := filepath.Join(t.TempDir(), "runtime.js")
	if err := os.WriteFile(path, []byte(starterRuntimeJS()), 0644); err != nil {
		t.Fatalf("write runtime: %v", err)
	}
	cmd := exec.Command(node, "--check", path)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("starter runtime should parse: %v\n%s", err, out)
	}
}

func newConfiguredProject(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module example.com/demo\n\ngo 1.23.0\n"), 0644); err != nil {
		t.Fatalf("write go.mod: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "components.json"), []byte(`{"module":"example.com/demo","uiDir":"ui"}`), 0644); err != nil {
		t.Fatalf("write components.json: %v", err)
	}
	return dir
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

func mustNotExist(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); err == nil {
		t.Fatalf("expected %s not to exist", path)
	} else if !os.IsNotExist(err) {
		t.Fatalf("stat %s: %v", path, err)
	}
}
