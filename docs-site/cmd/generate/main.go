package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"shadcn/docs-site/views"
)

type SearchEntry struct {
	Title string `json:"title"`
	URL   string `json:"url"`
	Type  string `json:"type"`
}

func main() {
	outputDir := flag.String("output", "./bin/site", "output directory for static site")
	flag.Parse()

	ctx := context.Background()

	if err := generate(ctx, *outputDir); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("static site generated in %s\n", *outputDir)
}

func generate(ctx context.Context, outputDir string) error {
	if err := cleanOutputDir(outputDir); err != nil {
		return err
	}

	dirs := []string{
		outputDir,
		filepath.Join(outputDir, "docs", "components"),
		filepath.Join(outputDir, "blocks"),
		filepath.Join(outputDir, "charts"),
		filepath.Join(outputDir, "public"),
	}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("create dir %s: %w", dir, err)
		}
	}

	if err := copyDir("public", filepath.Join(outputDir, "public")); err != nil {
		return fmt.Errorf("copy public: %w", err)
	}

	// Generate search index
	if err := generateSearchIndex(outputDir); err != nil {
		return fmt.Errorf("generate search index: %w", err)
	}

	pages := []struct {
		dir  string
		comp interface {
			Render(context.Context, io.Writer) error
		}
	}{
		{outputDir, views.Index()},
		{filepath.Join(outputDir, "docs"), views.DocsIndexPage()},
		{filepath.Join(outputDir, "docs", "components"), views.ComponentsIndexPage()},
		{filepath.Join(outputDir, "docs", "installation"), views.InstallationPage()},
		{filepath.Join(outputDir, "docs", "components-json"), views.ComponentsJSONPage()},
		{filepath.Join(outputDir, "docs", "package-imports"), views.PackageImportsPage()},
		{filepath.Join(outputDir, "docs", "theming"), views.ThemingPage()},
		{filepath.Join(outputDir, "docs", "dark-mode"), views.DarkModePage()},
		{filepath.Join(outputDir, "docs", "cli"), views.CLIPage()},
		{filepath.Join(outputDir, "docs", "monorepo"), views.MonorepoPage()},
		{filepath.Join(outputDir, "docs", "rtl"), views.RTLPage()},
		{filepath.Join(outputDir, "docs", "javascript"), views.JavaScriptPage()},
		{filepath.Join(outputDir, "docs", "llms"), views.LLMSPage()},
		{filepath.Join(outputDir, "docs", "forms"), views.FormsPage()},
		{filepath.Join(outputDir, "docs", "changelog"), views.ChangelogPage()},
		{filepath.Join(outputDir, "docs", "directory"), views.DirectoryDocsPage()},
		{filepath.Join(outputDir, "blocks"), views.BlocksIndexPage()},
		{filepath.Join(outputDir, "charts"), views.ChartsPage()},
		{filepath.Join(outputDir, "charts", "area"), views.AreaChartPage()},
		{filepath.Join(outputDir, "charts", "bar"), views.BarChartPage()},
		{filepath.Join(outputDir, "charts", "line"), views.LineChartPage()},
		{filepath.Join(outputDir, "charts", "pie"), views.PieChartPage()},
		{filepath.Join(outputDir, "charts", "radial"), views.RadialChartPage()},
		{filepath.Join(outputDir, "directory"), views.DirectoryPage()},
		{filepath.Join(outputDir, "create"), views.CreatePage()},
	}

	for _, p := range pages {
		if err := renderPage(ctx, p.dir, "index.html", p.comp); err != nil {
			return err
		}
	}

	for _, doc := range views.ComponentIndex() {
		prev, hasPrev := views.ComponentDocBefore(doc.Slug)
		next, hasNext := views.ComponentDocAfter(doc.Slug)
		dir := filepath.Join(outputDir, "docs", "components", doc.Slug)
		if err := renderPage(ctx, dir, "index.html", views.ComponentDetailPage(doc, prev, hasPrev, next, hasNext)); err != nil {
			return fmt.Errorf("component %s: %w", doc.Slug, err)
		}
	}

	for _, block := range views.Blocks() {
		catDir := filepath.Join(outputDir, "blocks", block.Category)
		if err := renderPage(ctx, catDir, "index.html", views.BlockCategoryPage(block.Category)); err != nil {
			return fmt.Errorf("block category %s: %w", block.Category, err)
		}

		detailDir := filepath.Join(outputDir, "blocks", block.Slug)
		if err := renderPage(ctx, detailDir, "index.html", views.BlockDetailPage(block)); err != nil {
			return fmt.Errorf("block %s: %w", block.Slug, err)
		}

		preview := getBlockPreview(block.Slug)
		if preview != nil {
			previewDir := filepath.Join(outputDir, "blocks", block.Slug)
			if err := renderPage(ctx, previewDir, "preview.html", preview); err != nil {
				return fmt.Errorf("block preview %s: %w", block.Slug, err)
			}
		}
	}

	return nil
}

func cleanOutputDir(outputDir string) error {
	if outputDir == "" || outputDir == "." || outputDir == string(filepath.Separator) {
		return fmt.Errorf("refusing to clear unsafe output directory %q", outputDir)
	}
	abs, err := filepath.Abs(outputDir)
	if err != nil {
		return fmt.Errorf("resolve output dir: %w", err)
	}
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("resolve working directory: %w", err)
	}
	rel, err := filepath.Rel(cwd, abs)
	if err != nil {
		return fmt.Errorf("resolve relative output dir: %w", err)
	}
	if rel == "." || rel == ".." || rel == filepath.Join("..") || relHasParentPrefix(rel) {
		return fmt.Errorf("refusing to clear output directory outside docs site: %s", outputDir)
	}
	if err := os.RemoveAll(abs); err != nil {
		return fmt.Errorf("clear output dir %s: %w", outputDir, err)
	}
	return nil
}

func relHasParentPrefix(rel string) bool {
	return len(rel) >= 3 && rel[:3] == ".."+string(filepath.Separator)
}

func renderPage(ctx context.Context, dir, filename string, comp interface {
	Render(context.Context, io.Writer) error
}) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(dir, filename))
	if err != nil {
		return err
	}
	defer f.Close()
	return comp.Render(ctx, f)
}

func getBlockPreview(slug string) interface {
	Render(context.Context, io.Writer) error
} {
	switch slug {
	case "dashboard-01":
		return views.Dashboard01Preview()
	case "sidebar-07":
		return views.Sidebar07Preview()
	case "sidebar-03":
		return views.Sidebar03Preview()
	case "login-01":
		return views.Login01Preview()
	case "login-03":
		return views.Login03Preview()
	case "login-04":
		return views.Login04Preview()
	default:
		return nil
	}
}

func copyDir(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)

		if d.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, 0644)
	})
}

func generateSearchIndex(outputDir string) error {
	entries := []SearchEntry{}

	// Add all components
	for _, doc := range views.ComponentIndex() {
		entries = append(entries, SearchEntry{
			Title: doc.Title,
			URL:   "/docs/components/" + doc.Slug,
			Type:  "component",
		})
	}

	// Add doc section pages
	entries = append(entries, SearchEntry{Title: "Introduction", URL: "/docs", Type: "docs"})
	entries = append(entries, SearchEntry{Title: "Installation", URL: "/docs/installation", Type: "docs"})
	entries = append(entries, SearchEntry{Title: "components.json", URL: "/docs/components-json", Type: "docs"})
	entries = append(entries, SearchEntry{Title: "Package Imports", URL: "/docs/package-imports", Type: "docs"})
	entries = append(entries, SearchEntry{Title: "Theming", URL: "/docs/theming", Type: "docs"})
	entries = append(entries, SearchEntry{Title: "CLI", URL: "/docs/cli", Type: "docs"})
	entries = append(entries, SearchEntry{Title: "Dark Mode", URL: "/docs/dark-mode", Type: "docs"})
	entries = append(entries, SearchEntry{Title: "Monorepo", URL: "/docs/monorepo", Type: "docs"})
	entries = append(entries, SearchEntry{Title: "RTL", URL: "/docs/rtl", Type: "docs"})
	entries = append(entries, SearchEntry{Title: "JavaScript", URL: "/docs/javascript", Type: "docs"})
	entries = append(entries, SearchEntry{Title: "llms.txt", URL: "/docs/llms", Type: "docs"})
	entries = append(entries, SearchEntry{Title: "Forms", URL: "/docs/forms", Type: "docs"})
	entries = append(entries, SearchEntry{Title: "Changelog", URL: "/docs/changelog", Type: "docs"})
	entries = append(entries, SearchEntry{Title: "Directory", URL: "/docs/directory", Type: "docs"})

	// Add chart pages
	entries = append(entries, SearchEntry{Title: "Area Chart", URL: "/charts/area", Type: "chart"})
	entries = append(entries, SearchEntry{Title: "Bar Chart", URL: "/charts/bar", Type: "chart"})
	entries = append(entries, SearchEntry{Title: "Line Chart", URL: "/charts/line", Type: "chart"})
	entries = append(entries, SearchEntry{Title: "Pie Chart", URL: "/charts/pie", Type: "chart"})
	entries = append(entries, SearchEntry{Title: "Radial Chart", URL: "/charts/radial", Type: "chart"})

	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal search index: %w", err)
	}

	searchIndexPath := filepath.Join(outputDir, "search-index.json")
	if err := os.WriteFile(searchIndexPath, data, 0644); err != nil {
		return fmt.Errorf("write search index: %w", err)
	}

	return nil
}
