package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const templDependency = "github.com/a-h/templ"
const templVersion = "v0.3.1001"

type ProjectConfig struct {
	Module string `json:"module"`
	UIDir  string `json:"uiDir"`
	Style  string `json:"style,omitempty"`
}

type InitOptions struct {
	CWD       string
	Name      string
	Template  string
	Base      string
	Preset    string
	Yes       bool
	Defaults  bool
	Force     bool
	Silent    bool
	Monorepo  bool
	RTL       bool
	Reinstall bool
}

type AddOptions struct {
	CWD       string
	Path      string
	Items     []string
	Yes       bool
	Overwrite bool
	All       bool
	Silent    bool
	DryRun    bool
	Diff      string
	View      string
}

type ApplyOptions struct {
	CWD    string
	Preset string
	Yes    bool
	Silent bool
}

type ViewOptions struct {
	CWD   string
	Items []string
}

type SearchOptions struct {
	CWD        string
	Query      string
	Limit      int
	Offset     int
	Registries []string
}

type BuildOptions struct {
	CWD       string
	OutputDir string
	Registry  string
}

type DocsOptions struct {
	CWD       string
	Base      string
	JSON      bool
	Component string
}

type InfoOptions struct {
	CWD  string
	JSON bool
}

type registryFile struct {
	RelPath string   `json:"relPath"`
	Symbols []string `json:"symbols"`
}

type registryIndex struct {
	Root      string
	Files     []registryFile
	Component map[string]registryFile
}

func InitProject(opts InitOptions) error {
	root, created, err := projectRootForInit(opts.CWD, opts.Name)
	if err != nil {
		return err
	}

	cfg := ProjectConfig{Module: projectModule(root, opts.Name), UIDir: "ui", Style: "default"}
	if err := ensureGoMod(root, cfg.Module); err != nil {
		return err
	}

	if err := writeProjectConfig(root, cfg, opts.Force); err != nil {
		return err
	}

	if err := scaffoldStarterProject(root, cfg, opts.Force); err != nil {
		return err
	}

	if err := syncUIPackage(context.Background(), root, cfg.UIDir, opts.Force, false); err != nil {
		return err
	}

	if !opts.Silent {
		if created {
			fmt.Printf("created new project in %s\n", root)
		} else {
			fmt.Printf("initialized project in %s\n", root)
		}
	}
	return nil
}

func AddComponents(opts AddOptions) error {
	root, err := detectProjectRoot(opts.CWD)
	if err != nil {
		return err
	}
	if len(opts.Items) == 0 && !opts.All {
		return errors.New("add: specify one or more components or pass --all")
	}

	cfg, err := loadProjectConfig(root)
	if err != nil {
		return err
	}
	targetDir := cfg.UIDir
	if opts.Path != "" {
		targetDir = opts.Path
	}

	index, err := buildRegistryIndex()
	if err != nil {
		return err
	}
	if len(opts.Items) > 0 {
		if err := validateRequestedComponents(index, opts.Items); err != nil {
			return err
		}
	}

	if opts.Diff != "" || opts.View != "" {
		return showFileView(root, targetDir, opts, index)
	}

	return syncUIPackage(context.Background(), root, targetDir, opts.Overwrite, opts.DryRun)
}

func ApplyPreset(opts ApplyOptions) error {
	if opts.Preset == "" {
		return errors.New("apply: missing preset")
	}
	preset := normalizeName(opts.Preset)
	if preset == "newyork" {
		preset = "new-york"
	}
	if !supportedPreset(preset) {
		return fmt.Errorf("apply: unknown preset %q", opts.Preset)
	}

	root, err := detectProjectRoot(opts.CWD)
	if err != nil {
		return err
	}
	cfg, err := loadProjectConfig(root)
	if err != nil {
		return err
	}
	if cfg.Module == "" {
		cfg.Module = projectModule(root, "")
	}
	cfg.Style = preset
	if err := writeProjectConfig(root, cfg, true); err != nil {
		return err
	}

	for rel, content := range map[string]string{
		"styles/globals.css": starterCSS(),
		"assets/runtime.js":  starterRuntimeJS(),
	} {
		target := filepath.Join(root, rel)
		if fileExists(target) {
			continue
		}
		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			return err
		}
		if err := os.WriteFile(target, []byte(content), 0644); err != nil {
			return err
		}
	}

	if !opts.Silent {
		fmt.Printf("applied preset %q\n", preset)
	}
	return nil
}

func supportedPreset(preset string) bool {
	switch preset {
	case "default", "new-york", "nova":
		return true
	default:
		return false
	}
}

func ViewItems(opts ViewOptions) error {
	root, err := detectProjectRoot(opts.CWD)
	if err != nil {
		return err
	}
	cfg, err := loadProjectConfig(root)
	if err != nil {
		return err
	}
	index, err := buildRegistryIndex()
	if err != nil {
		return err
	}
	if len(opts.Items) == 0 {
		return errors.New("view: specify one or more component names or file paths")
	}
	for _, item := range opts.Items {
		if err := printItemView(root, cfg.UIDir, index, item); err != nil {
			return err
		}
	}
	return nil
}

func SearchComponents(opts SearchOptions) error {
	index, err := buildRegistryIndex()
	if err != nil {
		return err
	}
	query := normalizeName(opts.Query)
	if query == "" && len(opts.Registries) > 0 {
		query = normalizeName(opts.Registries[0])
	}
	if query == "" {
		return errors.New("search: provide a query or registry name")
	}
	matches := searchRegistry(index, query)
	if len(matches) == 0 {
		fmt.Printf("no components matched %q\n", opts.Query)
		return nil
	}
	for _, match := range matches {
		fmt.Printf("%s\t%s\n", strings.Join(match.Symbols, ", "), match.RelPath)
	}
	return nil
}

func BuildRegistry(opts BuildOptions) error {
	root, err := detectProjectRoot(opts.CWD)
	if err != nil {
		return err
	}
	cfg, err := loadProjectConfig(root)
	if err != nil {
		return err
	}
	index, err := buildRegistryIndex()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(opts.OutputDir, 0755); err != nil {
		return err
	}
	payload := map[string]any{
		"module":     cfg.Module,
		"uiDir":      cfg.UIDir,
		"components": index.Files,
	}
	if opts.Registry != "" {
		payload["registry"] = opts.Registry
	}
	data, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(opts.OutputDir, "registry.json"), data, 0644)
}

func ShowDocs(opts DocsOptions) error {
	index, err := buildRegistryIndex()
	if err != nil {
		return err
	}
	if opts.Component == "" {
		fmt.Println("available components:")
		for _, file := range index.Files {
			fmt.Printf("- %s (%s)\n", strings.Join(file.Symbols, ", "), file.RelPath)
		}
		return nil
	}
	key := normalizeName(opts.Component)
	file, ok := index.Component[key]
	if !ok {
		return fmt.Errorf("docs: unknown component %q", opts.Component)
	}
	if opts.JSON {
		payload := map[string]any{
			"component": opts.Component,
			"file":      file.RelPath,
			"symbols":   file.Symbols,
			"base":      opts.Base,
		}
		data, err := json.MarshalIndent(payload, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(data))
		return nil
	}
	fmt.Printf("%s\nfile: %s\nsymbols: %s\nbase: %s\n", opts.Component, file.RelPath, strings.Join(file.Symbols, ", "), opts.Base)
	return nil
}

func PrintInfo(opts InfoOptions) error {
	root, err := detectProjectRoot(opts.CWD)
	if err != nil {
		return err
	}
	cfg, err := loadProjectConfig(root)
	if err != nil {
		return err
	}
	index, err := buildRegistryIndex()
	if err != nil {
		return err
	}
	if opts.JSON {
		payload := map[string]any{
			"root":       root,
			"module":     cfg.Module,
			"uiDir":      cfg.UIDir,
			"components": len(index.Files),
		}
		data, err := json.MarshalIndent(payload, "", "  ")
		if err != nil {
			return err
		}
		fmt.Println(string(data))
		return nil
	}
	fmt.Printf("root: %s\nmodule: %s\nui dir: %s\ncomponents: %d\n", root, cfg.Module, cfg.UIDir, len(index.Files))
	return nil
}

func projectRootForInit(cwd string, name string) (string, bool, error) {
	base, err := filepath.Abs(cwd)
	if err != nil {
		return "", false, err
	}
	if name == "" {
		return base, false, nil
	}
	root := filepath.Join(base, sanitizeProjectName(name))
	if err := os.MkdirAll(root, 0755); err != nil {
		return "", false, err
	}
	return root, true, nil
}

func projectModule(root string, fallback string) string {
	if module, ok, err := readGoModModule(filepath.Join(root, "go.mod")); err == nil && ok {
		return module
	}
	if fallback != "" {
		return sanitizeModulePath(fallback)
	}
	return sanitizeModulePath(filepath.Base(root))
}

func detectProjectRoot(cwd string) (string, error) {
	root, err := filepath.Abs(cwd)
	if err != nil {
		return "", err
	}
	for {
		if fileExists(filepath.Join(root, "go.mod")) {
			return root, nil
		}
		parent := filepath.Dir(root)
		if parent == root {
			return "", fmt.Errorf("could not find a Go project root from %s", cwd)
		}
		root = parent
	}
}

func loadProjectConfig(root string) (ProjectConfig, error) {
	cfg := ProjectConfig{UIDir: "ui"}
	raw, err := os.ReadFile(filepath.Join(root, "components.json"))
	if errors.Is(err, os.ErrNotExist) {
		return cfg, nil
	}
	if err != nil {
		return cfg, err
	}
	if err := json.Unmarshal(raw, &cfg); err != nil {
		return cfg, err
	}
	if cfg.UIDir == "" {
		cfg.UIDir = "ui"
	}
	return cfg, nil
}

func writeProjectConfig(root string, cfg ProjectConfig, force bool) error {
	path := filepath.Join(root, "components.json")
	if fileExists(path) && !force {
		return nil
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(data, '\n'), 0644)
}

func ensureGoMod(root, modulePath string) error {
	path := filepath.Join(root, "go.mod")
	if fileExists(path) {
		raw, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		updated := string(raw)
		if !strings.Contains(updated, templDependency) {
			updated = strings.TrimRight(updated, "\n") + "\n\nrequire " + templDependency + " " + templVersion + "\n"
			return os.WriteFile(path, []byte(updated), 0644)
		}
		return nil
	}

	content := fmt.Sprintf("module %s\n\ngo 1.23.0\n\nrequire %s %s\n", modulePath, templDependency, templVersion)
	return os.WriteFile(path, []byte(content), 0644)
}

func readGoModModule(path string) (string, bool, error) {
	raw, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return "", false, nil
	}
	if err != nil {
		return "", false, err
	}
	for _, line := range strings.Split(string(raw), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), true, nil
		}
	}
	return "", false, nil
}

func scaffoldStarterProject(root string, cfg ProjectConfig, force bool) error {
	modulePath := cfg.Module
	if modulePath == "" {
		modulePath = sanitizeModulePath(filepath.Base(root))
	}

	files := map[string]string{
		"main.go":            starterMainGo(),
		"app.templ":          starterAppTempl(modulePath),
		"styles/globals.css": starterCSS(),
		"assets/runtime.js":  starterRuntimeJS(),
	}

	for rel, content := range files {
		target := filepath.Join(root, rel)
		if fileExists(target) && !force {
			continue
		}
		if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
			return err
		}
		if err := os.WriteFile(target, []byte(content), 0644); err != nil {
			return err
		}
	}
	return nil
}

func syncUIPackage(ctx context.Context, root, targetDir string, overwrite bool, dryRun bool) error {
	sourceRoot, err := locateRegistryRoot()
	if err != nil {
		return err
	}

	sourceDir := filepath.Join(sourceRoot, "ui")
	if !fileExists(sourceDir) {
		return fmt.Errorf("could not find source ui package at %s", sourceDir)
	}

	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		return err
	}

	dstDir := filepath.Join(root, targetDir)
	if dryRun {
		for _, entry := range entries {
			if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
				continue
			}
			fmt.Println(filepath.Join(targetDir, entry.Name()))
		}
		return nil
	}

	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}
		src := filepath.Join(sourceDir, entry.Name())
		dst := filepath.Join(dstDir, entry.Name())
		if err := copyFile(src, dst, overwrite); err != nil {
			return err
		}
	}
	return nil
}

func copyFile(src, dst string, overwrite bool) error {
	if fileExists(dst) && !overwrite {
		return nil
	}
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}
	return os.WriteFile(dst, data, 0644)
}

func locateRegistryRoot() (string, error) {
	if override := os.Getenv("SHADCN_SOURCE_DIR"); override != "" {
		return filepath.Abs(override)
	}

	candidates := []string{}
	if cwd, err := os.Getwd(); err == nil {
		candidates = append(candidates, cwd)
	}
	if exe, err := os.Executable(); err == nil {
		candidates = append(candidates, filepath.Dir(exe))
	}
	for _, start := range candidates {
		if root, ok := walkForRegistryRoot(start); ok {
			return root, nil
		}
	}
	return "", errors.New("could not locate the source ui package; set SHADCN_SOURCE_DIR")
}

func walkForRegistryRoot(start string) (string, bool) {
	dir := start
	for {
		if fileExists(filepath.Join(dir, "ui", "go.mod")) {
			return dir, true
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", false
		}
		dir = parent
	}
}

func buildRegistryIndex() (registryIndex, error) {
	root, err := locateRegistryRoot()
	if err != nil {
		return registryIndex{}, err
	}
	uiDir := filepath.Join(root, "ui")
	entries, err := os.ReadDir(uiDir)
	if err != nil {
		return registryIndex{}, err
	}
	index := registryIndex{
		Root:      root,
		Files:     make([]registryFile, 0, len(entries)),
		Component: make(map[string]registryFile),
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}
		rel := filepath.Join("ui", entry.Name())
		symbols, err := exportedSymbols(filepath.Join(uiDir, entry.Name()))
		if err != nil {
			return registryIndex{}, err
		}
		file := registryFile{RelPath: rel, Symbols: symbols}
		index.Files = append(index.Files, file)
		for _, symbol := range symbols {
			index.Component[normalizeName(symbol)] = file
		}
	}
	return index, nil
}

func exportedSymbols(path string) ([]string, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	seen := map[string]struct{}{}
	lines := bytes.Split(raw, []byte{'\n'})
	for _, line := range lines {
		line = bytes.TrimSpace(line)
		if !bytes.HasPrefix(line, []byte("func ")) {
			continue
		}
		name := line[len("func "):]
		idx := bytes.IndexByte(name, '(')
		if idx < 0 {
			continue
		}
		symbol := string(bytes.TrimSpace(name[:idx]))
		if symbol == "" || symbol[0] < 'A' || symbol[0] > 'Z' {
			continue
		}
		seen[symbol] = struct{}{}
	}
	out := make([]string, 0, len(seen))
	for symbol := range seen {
		out = append(out, symbol)
	}
	sort.Strings(out)
	return out, nil
}

func searchRegistry(index registryIndex, query string) []registryFile {
	matches := make([]registryFile, 0)
	for _, file := range index.Files {
		if strings.Contains(normalizeName(strings.Join(file.Symbols, " ")), query) {
			matches = append(matches, file)
			continue
		}
		for _, symbol := range file.Symbols {
			if strings.Contains(normalizeName(symbol), query) {
				matches = append(matches, file)
				break
			}
		}
	}
	return matches
}

func validateRequestedComponents(index registryIndex, items []string) error {
	missing := make([]string, 0)
	for _, item := range items {
		if _, ok := index.Component[normalizeName(item)]; !ok {
			missing = append(missing, item)
		}
	}
	if len(missing) > 0 {
		return fmt.Errorf("unknown components: %s", strings.Join(missing, ", "))
	}
	return nil
}

func printItemView(root, uiDir string, index registryIndex, item string) error {
	if filepath.IsAbs(item) || strings.Contains(item, string(os.PathSeparator)) {
		raw, err := os.ReadFile(filepath.Join(root, item))
		if err != nil {
			return err
		}
		fmt.Printf("== %s ==\n%s\n", item, string(raw))
		return nil
	}
	file, ok := index.Component[normalizeName(item)]
	if !ok {
		return fmt.Errorf("view: unknown component or file %q", item)
	}
	raw, err := os.ReadFile(filepath.Join(root, file.RelPath))
	if err != nil {
		return err
	}
	fmt.Printf("== %s ==\nfile: %s\nsymbols: %s\n%s\n", item, file.RelPath, strings.Join(file.Symbols, ", "), string(raw))
	return nil
}

func showFileView(root, uiDir string, opts AddOptions, index registryIndex) error {
	if opts.View != "" {
		return printItemView(root, uiDir, index, opts.View)
	}
	if opts.Diff != "" {
		raw, err := os.ReadFile(filepath.Join(root, opts.Diff))
		if err != nil {
			return err
		}
		fmt.Printf("== %s ==\n%s\n", opts.Diff, string(raw))
		return nil
	}
	return nil
}

func sanitizeProjectName(name string) string {
	cleaned := strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z':
			return r
		case r >= 'A' && r <= 'Z':
			return r + ('a' - 'A')
		case r >= '0' && r <= '9':
			return r
		case r == '-' || r == '_' || r == '.':
			return r
		default:
			return '-'
		}
	}, strings.TrimSpace(name))
	cleaned = strings.Trim(cleaned, "-_.")
	if cleaned == "" {
		return "app"
	}
	return cleaned
}

func sanitizeModulePath(name string) string {
	name = sanitizeProjectName(name)
	return strings.ReplaceAll(name, " ", "-")
}

func normalizeName(name string) string {
	return strings.ToLower(strings.Map(func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			return r
		case r >= 'A' && r <= 'Z':
			return r + ('a' - 'A')
		default:
			return -1
		}
	}, name))
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
