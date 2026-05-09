package main

import (
	"path/filepath"
	"sort"
	"strings"
)

var sharedRegistryFiles = []string{
	"ui/cn.go",
	"ui/floating.go",
	"ui/html.go",
	"ui/render.go",
	"ui/types.go",
}

var internalRegistryFiles = map[string]struct{}{
	"cn":       {},
	"floating": {},
	"html":     {},
	"render":   {},
	"types":    {},
}

var componentFileNames = map[string]string{
	"data-table": "datatable",
}

var componentNamesByFile = map[string]string{
	"datatable": "data-table",
}

var componentFileDependencies = map[string][]string{
	"alert-dialog": {"dialog"},
	"combobox":     {"command", "dropdown-menu", "input", "select"},
	"command":      {"dialog", "dropdown-menu", "input"},
	"context-menu": {"dropdown-menu"},
	"data-table":   {"pagination"},
	"date-picker":  {"button", "calendar", "popover"},
	"drawer":       {"sheet"},
	"input-group":  {"button", "input"},
	"input-otp":    {"input"},
	"menubar":      {"dropdown-menu"},
	"select":       {"dropdown-menu"},
	"sheet":        {"dialog"},
	"sidebar":      {"input", "separator"},
	"sonner":       {"toast"},
}

var runtimeComponents = map[string][]string{
	"accordion":       {"assets/runtime.js"},
	"alert-dialog":    {"assets/runtime.js"},
	"combobox":        {"assets/runtime.js"},
	"command":         {"assets/runtime.js"},
	"context-menu":    {"assets/runtime.js"},
	"calendar":        {"assets/runtime.js"},
	"carousel":        {"assets/runtime.js"},
	"checkbox":        {"assets/runtime.js"},
	"collapsible":     {"assets/runtime.js"},
	"date-picker":     {"assets/runtime.js"},
	"dialog":          {"assets/runtime.js"},
	"drawer":          {"assets/runtime.js"},
	"dropdown-menu":   {"assets/runtime.js"},
	"hover-card":      {"assets/runtime.js"},
	"input-otp":       {"assets/runtime.js"},
	"menubar":         {"assets/runtime.js"},
	"navigation-menu": {"assets/runtime.js"},
	"popover":         {"assets/runtime.js"},
	"radio-group":     {"assets/runtime.js"},
	"resizable":       {"assets/runtime.js"},
	"select":          {"assets/runtime.js"},
	"sheet":           {"assets/runtime.js"},
	"sidebar":         {"assets/runtime.js"},
	"slider":          {"assets/runtime.js"},
	"sonner":          {"assets/runtime.js"},
	"switch":          {"assets/runtime.js"},
	"tabs":            {"assets/runtime.js"},
	"toggle-group":    {"assets/runtime.js"},
	"tooltip":         {"assets/runtime.js"},
	"toast":           {"assets/runtime.js"},
}

var upstreamPrimitiveComponents = map[string]struct{}{
	"accordion":       {},
	"alert-dialog":    {},
	"combobox":        {},
	"command":         {},
	"context-menu":    {},
	"dialog":          {},
	"drawer":          {},
	"dropdown-menu":   {},
	"hover-card":      {},
	"menubar":         {},
	"navigation-menu": {},
	"popover":         {},
	"radio-group":     {},
	"select":          {},
	"sheet":           {},
	"slider":          {},
	"tabs":            {},
	"toggle-group":    {},
	"tooltip":         {},
}

func componentNameFromFile(rel string) string {
	base := strings.TrimSuffix(filepath.Base(rel), ".go")
	if name, ok := componentNamesByFile[base]; ok {
		return name
	}
	return strings.ReplaceAll(base, "_", "-")
}

func componentSourcePath(name string) string {
	if fileName, ok := componentFileNames[name]; ok {
		return filepath.Join("ui", fileName+".go")
	}
	return filepath.Join("ui", strings.ReplaceAll(name, "-", "_")+".go")
}

func buildRegistryItem(file registryFile) registryItem {
	name := componentNameFromFile(file.RelPath)
	item := registryItem{
		Name:        name,
		Type:        "registry:ui",
		Description: componentDescription(name),
		Files:       []registryFile{file},
		CSS:         []string{"styles/globals.css"},
		DocsURL:     "https://ui.shadcn.com/docs/components/" + name,
		Examples:    []string{"docs-site/views/component_examples.go"},
	}
	item.Dependencies = registryDependencies(name)
	item.Runtime = append([]string(nil), runtimeComponents[name]...)
	sort.Strings(item.Runtime)
	return item
}

func registryDependencies(name string) []string {
	seen := map[string]struct{}{}
	for _, rel := range sharedRegistryFiles {
		if rel != componentSourcePath(name) {
			seen[rel] = struct{}{}
		}
	}
	for _, dep := range componentFileDependencies[name] {
		seen[componentSourcePath(dep)] = struct{}{}
	}
	out := make([]string, 0, len(seen))
	for rel := range seen {
		out = append(out, rel)
	}
	sort.Strings(out)
	return out
}

func componentDescription(name string) string {
	if _, ok := upstreamPrimitiveComponents[name]; ok {
		return "Templ-native shadcn/ui primitive with runtime behavior metadata."
	}
	return "Templ-native shadcn/ui component."
}

func resolveRegistryItems(index registryIndex, requested []string, all bool) ([]registryItem, error) {
	names := make([]string, 0)
	if all {
		for name := range index.Items {
			names = append(names, name)
		}
	} else {
		for _, item := range requested {
			file, ok := index.Component[normalizeName(item)]
			if !ok {
				return nil, nil
			}
			names = append(names, componentNameFromFile(file.RelPath))
		}
	}
	sort.Strings(names)

	seen := map[string]struct{}{}
	var out []registryItem
	var visit func(string)
	visit = func(name string) {
		if _, ok := seen[name]; ok {
			return
		}
		seen[name] = struct{}{}
		for _, dep := range componentFileDependencies[name] {
			visit(dep)
		}
		if item, ok := index.Items[name]; ok {
			out = append(out, item)
		}
	}
	for _, name := range names {
		visit(name)
	}
	sort.SliceStable(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out, nil
}

func filesForRegistryItems(items []registryItem) []string {
	seen := map[string]struct{}{}
	for _, item := range items {
		for _, file := range item.Files {
			seen[file.RelPath] = struct{}{}
		}
		for _, dep := range item.Dependencies {
			seen[dep] = struct{}{}
		}
		for _, runtime := range item.Runtime {
			seen[runtime] = struct{}{}
		}
	}
	out := make([]string, 0, len(seen))
	for rel := range seen {
		out = append(out, rel)
	}
	sort.Strings(out)
	return out
}
