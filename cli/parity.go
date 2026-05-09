package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"
)

const upstreamBaselineCommit = "dd3567c39da374346d2daa07f926a20d5036c492"
const defaultUpstreamRegistryIndex = "https://raw.githubusercontent.com/shadcn-ui/ui/" + upstreamBaselineCommit + "/apps/v4/public/r/index.json"

type upstreamRegistryItem struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Files []struct {
		Path string `json:"path"`
		Type string `json:"type"`
	} `json:"files"`
	Meta map[string]any `json:"meta"`
}

type parityComponent struct {
	Name         string   `json:"name"`
	Local        bool     `json:"local"`
	Upstream     bool     `json:"upstream"`
	LocalFiles   []string `json:"localFiles,omitempty"`
	Runtime      []string `json:"runtime,omitempty"`
	Dependencies []string `json:"dependencies,omitempty"`
	Status       string   `json:"status"`
}

type parityReport struct {
	GeneratedAt string            `json:"generatedAt"`
	Upstream    string            `json:"upstream"`
	LocalCount  int               `json:"localCount"`
	RemoteCount int               `json:"remoteCount"`
	Missing     []string          `json:"missing"`
	Extra       []string          `json:"extra"`
	Components  []parityComponent `json:"components"`
}

func checkParity(source string) (parityReport, error) {
	if source == "" {
		source = defaultUpstreamRegistryIndex
	}
	upstreamItems, err := loadUpstreamRegistry(source)
	if err != nil {
		return parityReport{}, err
	}
	index, err := buildRegistryIndex()
	if err != nil {
		return parityReport{}, err
	}
	return buildParityReport(source, upstreamItems, index), nil
}

func loadUpstreamRegistry(source string) ([]upstreamRegistryItem, error) {
	var raw []byte
	var err error
	if strings.HasPrefix(source, "http://") || strings.HasPrefix(source, "https://") {
		client := http.Client{Timeout: 20 * time.Second}
		resp, err := client.Get(source)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			return nil, fmt.Errorf("upstream registry returned %s", resp.Status)
		}
		raw, err = io.ReadAll(resp.Body)
	} else {
		raw, err = os.ReadFile(source)
	}
	if err != nil {
		return nil, err
	}
	var items []upstreamRegistryItem
	if err := json.Unmarshal(raw, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func buildParityReport(source string, upstream []upstreamRegistryItem, index registryIndex) parityReport {
	local := map[string]registryItem{}
	for name, item := range index.Items {
		local[name] = item
	}
	remote := map[string]upstreamRegistryItem{}
	for _, item := range upstream {
		if item.Type != "" && item.Type != "registry:ui" {
			continue
		}
		if item.Name != "" {
			remote[item.Name] = item
		}
	}
	namesSeen := map[string]struct{}{}
	for name := range local {
		namesSeen[name] = struct{}{}
	}
	for name := range remote {
		namesSeen[name] = struct{}{}
	}
	names := make([]string, 0, len(namesSeen))
	for name := range namesSeen {
		names = append(names, name)
	}
	sort.Strings(names)

	report := parityReport{
		GeneratedAt: time.Now().UTC().Format(time.RFC3339),
		Upstream:    source,
		LocalCount:  len(local),
		RemoteCount: len(remote),
	}
	for _, name := range names {
		localItem, hasLocal := local[name]
		_, hasRemote := remote[name]
		status := "present"
		if !hasLocal && hasRemote {
			status = "missing"
			report.Missing = append(report.Missing, name)
		}
		if hasLocal && !hasRemote {
			status = "extra"
			report.Extra = append(report.Extra, name)
		}
		component := parityComponent{Name: name, Local: hasLocal, Upstream: hasRemote, Status: status}
		if hasLocal {
			for _, file := range localItem.Files {
				component.LocalFiles = append(component.LocalFiles, file.RelPath)
			}
			component.Dependencies = append(component.Dependencies, localItem.Dependencies...)
			component.Runtime = append(component.Runtime, localItem.Runtime...)
		}
		report.Components = append(report.Components, component)
	}
	return report
}

func printParityReport(report parityReport) {
	fmt.Printf("upstream: %s\n", report.Upstream)
	fmt.Printf("local components: %d\n", report.LocalCount)
	fmt.Printf("upstream components: %d\n", report.RemoteCount)
	fmt.Printf("missing: %d\n", len(report.Missing))
	for _, name := range report.Missing {
		fmt.Printf("- %s\n", name)
	}
	if len(report.Extra) > 0 {
		fmt.Printf("extra: %d\n", len(report.Extra))
		for _, name := range report.Extra {
			fmt.Printf("+ %s\n", name)
		}
	}
}
