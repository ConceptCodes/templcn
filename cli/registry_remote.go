package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func isRegistryItemRef(item string) bool {
	if strings.HasPrefix(item, "http://") || strings.HasPrefix(item, "https://") {
		return true
	}
	if strings.HasPrefix(item, "@") && strings.Contains(item, "/") {
		return true
	}
	if strings.HasSuffix(strings.ToLower(item), ".json") {
		return true
	}
	return false
}

func splitRegistryItemRefs(items []string) ([]string, []string) {
	refs := make([]string, 0)
	local := make([]string, 0, len(items))
	for _, item := range items {
		if isRegistryItemRef(item) {
			refs = append(refs, item)
			continue
		}
		local = append(local, item)
	}
	return refs, local
}

func loadRegistryItemRefs(root string, cfg ProjectConfig, refs []string) ([]shadcnRegistryItem, error) {
	seen := map[string]struct{}{}
	items := make([]shadcnRegistryItem, 0, len(refs))
	var visit func(string) error
	visit = func(ref string) error {
		resolved, err := resolveRegistryItemRef(root, cfg, ref)
		if err != nil {
			return err
		}
		if _, ok := seen[resolved]; ok {
			return nil
		}
		seen[resolved] = struct{}{}
		item, err := loadRegistryItemRef(resolved)
		if err != nil {
			return err
		}
		items = append(items, item)
		for _, dep := range item.RegistryDependencies {
			if isRegistryItemRef(dep) {
				if err := visit(dep); err != nil {
					return err
				}
			}
		}
		return nil
	}
	for _, ref := range refs {
		if err := visit(ref); err != nil {
			return nil, err
		}
	}
	return items, nil
}

func loadRegistryItemRef(ref string) (shadcnRegistryItem, error) {
	var raw []byte
	var err error
	if strings.HasPrefix(ref, "http://") || strings.HasPrefix(ref, "https://") {
		client := http.Client{Timeout: 20 * time.Second}
		req, err := http.NewRequest(http.MethodGet, ref, nil)
		if err != nil {
			return shadcnRegistryItem{}, err
		}
		req.Header.Set("User-Agent", "templcn")
		req.Header.Set("Accept", "application/vnd.shadcn.v1+json, application/json;q=0.9")
		resp, err := client.Do(req)
		if err != nil {
			return shadcnRegistryItem{}, err
		}
		defer resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			return shadcnRegistryItem{}, fmt.Errorf("registry item %s returned %s", ref, resp.Status)
		}
		raw, err = io.ReadAll(resp.Body)
	} else {
		raw, err = os.ReadFile(ref)
	}
	if err != nil {
		return shadcnRegistryItem{}, err
	}
	var item shadcnRegistryItem
	if err := json.Unmarshal(raw, &item); err != nil {
		return shadcnRegistryItem{}, err
	}
	if item.Name == "" {
		item.Name = strings.TrimSuffix(filepath.Base(ref), filepath.Ext(ref))
	}
	return item, nil
}

func resolveRegistryItemRef(root string, cfg ProjectConfig, ref string) (string, error) {
	if !strings.HasPrefix(ref, "@") {
		return ref, nil
	}
	parts := strings.SplitN(strings.TrimPrefix(ref, "@"), "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", fmt.Errorf("invalid registry item reference %q", ref)
	}
	base := cfg.Registries[parts[0]]
	if base == "" {
		return "", fmt.Errorf("unknown registry namespace @%s", parts[0])
	}
	name := strings.TrimSuffix(parts[1], ".json") + ".json"
	if strings.HasPrefix(base, "http://") || strings.HasPrefix(base, "https://") {
		return strings.TrimRight(base, "/") + "/" + name, nil
	}
	if !filepath.IsAbs(base) {
		base = filepath.Join(root, base)
	}
	if strings.HasSuffix(base, ".json") {
		return filepath.Join(filepath.Dir(base), name), nil
	}
	return filepath.Join(base, name), nil
}

func registryDependenciesForExternalItems(items []shadcnRegistryItem) []string {
	seen := map[string]struct{}{}
	for _, item := range items {
		for _, dep := range item.RegistryDependencies {
			if isRegistryItemRef(dep) {
				continue
			}
			seen[dep] = struct{}{}
		}
	}
	out := make([]string, 0, len(seen))
	for dep := range seen {
		out = append(out, dep)
	}
	sort.Strings(out)
	return out
}

func syncExternalRegistryItems(root string, items []shadcnRegistryItem, modulePath string, overwrite bool, dryRun bool) error {
	for _, item := range items {
		for _, file := range item.Files {
			if file.Path == "" {
				continue
			}
			if dryRun {
				fmt.Println(targetPathForExternalRegistryFile(file))
				continue
			}
			content := strings.ReplaceAll(file.Content, "{{module}}", modulePath)
			if err := writeFileIfAllowed(filepath.Join(root, targetPathForExternalRegistryFile(file)), []byte(content), overwrite); err != nil {
				return err
			}
		}
	}
	return nil
}

func targetPathForExternalRegistryFile(file shadcnRegistryFile) string {
	if file.Target != "" {
		return file.Target
	}
	return file.Path
}

func printRegistryItemView(root string, cfg ProjectConfig, ref string) error {
	resolved, err := resolveRegistryItemRef(root, cfg, ref)
	if err != nil {
		return err
	}
	item, err := loadRegistryItemRef(resolved)
	if err != nil {
		return err
	}
	fmt.Printf("== %s ==\n", ref)
	fmt.Printf("name: %s\ntype: %s\n", item.Name, item.Type)
	if len(item.RegistryDependencies) > 0 {
		fmt.Printf("registryDependencies: %s\n", strings.Join(item.RegistryDependencies, ", "))
	}
	for _, file := range item.Files {
		fmt.Printf("-- %s --\n%s\n", file.Path, file.Content)
	}
	return nil
}

type namespaceRegistryIndex struct {
	Name     string               `json:"name"`
	Homepage string               `json:"homepage"`
	Items    []shadcnRegistryItem `json:"items"`
}

func SearchNamespaceRegistry(root string, cfg ProjectConfig, opts SearchOptions) error {
	if root == "" {
		return fmt.Errorf("search: namespaced registries require a project with components.json")
	}
	namespace := strings.TrimPrefix(opts.Registries[0], "@")
	base := cfg.Registries[namespace]
	if base == "" {
		return fmt.Errorf("unknown registry namespace @%s", namespace)
	}
	index, err := loadNamespaceRegistryIndex(root, base)
	if err != nil {
		return err
	}
	query := normalizeName(opts.Query)
	start := opts.Offset
	if start < 0 {
		start = 0
	}
	limit := opts.Limit
	if limit <= 0 {
		limit = 100
	}
	count := 0
	for _, item := range index.Items {
		if query != "" && !strings.Contains(normalizeName(item.Name+" "+item.Description), query) {
			continue
		}
		if start > 0 {
			start--
			continue
		}
		if count >= limit {
			break
		}
		fmt.Printf("@%s/%s\t%s\n", namespace, item.Name, item.Description)
		count++
	}
	if count == 0 {
		fmt.Printf("no registry items matched %q\n", opts.Query)
	}
	return nil
}

func loadNamespaceRegistryIndex(root string, base string) (namespaceRegistryIndex, error) {
	ref := base
	if strings.HasPrefix(base, "http://") || strings.HasPrefix(base, "https://") {
		ref = strings.TrimRight(base, "/") + "/registry.json"
	} else {
		if !filepath.IsAbs(ref) {
			ref = filepath.Join(root, ref)
		}
		if !strings.HasSuffix(ref, ".json") {
			ref = filepath.Join(ref, "registry.json")
		}
	}
	raw, err := readRegistryPayload(ref)
	if err != nil {
		return namespaceRegistryIndex{}, err
	}
	var index namespaceRegistryIndex
	if err := json.Unmarshal(raw, &index); err != nil {
		return namespaceRegistryIndex{}, err
	}
	return index, nil
}

func readRegistryPayload(ref string) ([]byte, error) {
	if strings.HasPrefix(ref, "http://") || strings.HasPrefix(ref, "https://") {
		client := http.Client{Timeout: 20 * time.Second}
		req, err := http.NewRequest(http.MethodGet, ref, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("User-Agent", "templcn")
		req.Header.Set("Accept", "application/vnd.shadcn.v1+json, application/json;q=0.9")
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			return nil, fmt.Errorf("registry %s returned %s", ref, resp.Status)
		}
		return io.ReadAll(resp.Body)
	}
	return os.ReadFile(ref)
}
