package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type blockRegistryFile struct {
	RelPath string
	Target  string
	Type    string
	Content string
}

type blockRegistryItem struct {
	Name         string
	Files        []blockRegistryFile
	Dependencies []string
	Runtime      []string
}

func isBlockName(name string) bool {
	_, ok := blockRegistry()[normalizeName(name)]
	return ok
}

func resolveBlockItems(requested []string) []blockRegistryItem {
	blocks := blockRegistry()
	seen := map[string]struct{}{}
	out := make([]blockRegistryItem, 0)
	for _, name := range requested {
		key := normalizeName(name)
		block, ok := blocks[key]
		if !ok {
			continue
		}
		if _, ok := seen[block.Name]; ok {
			continue
		}
		seen[block.Name] = struct{}{}
		out = append(out, block)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Name < out[j].Name })
	return out
}

func componentItemsForRequests(items []string) []string {
	out := make([]string, 0, len(items))
	for _, item := range items {
		if isBlockName(item) || isRegistryItemRef(item) {
			continue
		}
		out = append(out, item)
	}
	return out
}

func blockDependencies(blocks []blockRegistryItem) []string {
	seen := map[string]struct{}{}
	for _, block := range blocks {
		for _, dep := range block.Dependencies {
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

func blockRuntimeFiles(blocks []blockRegistryItem) []string {
	seen := map[string]struct{}{}
	for _, block := range blocks {
		for _, runtime := range block.Runtime {
			seen[runtime] = struct{}{}
		}
	}
	out := make([]string, 0, len(seen))
	for runtime := range seen {
		out = append(out, runtime)
	}
	sort.Strings(out)
	return out
}

func syncBlockFiles(root string, blocks []blockRegistryItem, modulePath string, overwrite bool, dryRun bool) error {
	for _, block := range blocks {
		for _, file := range block.Files {
			if dryRun {
				fmt.Println(file.RelPath)
				continue
			}
			content := strings.ReplaceAll(file.Content, "{{module}}", modulePath)
			target := filepath.Join(root, targetPathForBlockFile(file))
			if err := writeFileIfAllowed(target, []byte(content), overwrite); err != nil {
				return err
			}
		}
	}
	return nil
}

func writeFileIfAllowed(path string, data []byte, overwrite bool) error {
	if fileExists(path) && !overwrite {
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func blockRegistry() map[string]blockRegistryItem {
	blocks := map[string]blockRegistryItem{
		normalizeName("dashboard-01"): {
			Name:         "dashboard-01",
			Dependencies: []string{"badge", "button", "card", "chart", "data-table", "separator", "sidebar", "table"},
			Runtime:      []string{"assets/runtime.js"},
			Files: []blockRegistryFile{
				blockPageFile("dashboard-01", "app/dashboard/page.templ", dashboardPageTempl()),
				blockComponentFile("dashboard-01", "components/app-sidebar.templ", appSidebarTempl()),
				blockComponentFile("dashboard-01", "components/chart-area-interactive.templ", chartAreaInteractiveTempl()),
				blockComponentFile("dashboard-01", "components/data-table.templ", dataTableBlockTempl()),
				blockComponentFile("dashboard-01", "components/section-cards.templ", sectionCardsTempl()),
				blockComponentFile("dashboard-01", "components/site-header.templ", siteHeaderTempl()),
			},
		},
	}
	for _, name := range []string{
		"sidebar-01", "sidebar-02", "sidebar-03", "sidebar-04", "sidebar-05",
		"sidebar-06", "sidebar-07", "sidebar-08", "sidebar-09", "sidebar-10",
		"sidebar-11", "sidebar-12", "sidebar-13", "sidebar-14", "sidebar-15",
	} {
		variant := "default"
		if name == "sidebar-03" {
			variant = "submenu"
		}
		if name == "sidebar-07" {
			variant = "icon"
		}
		blocks[normalizeName(name)] = sidebarBlock(name, variant)
	}
	for _, block := range []struct {
		name        string
		route       string
		title       string
		description string
	}{
		{"login-01", "login", "Sign in", "Enter your credentials to access your account"},
		{"login-02", "login", "Welcome back", "Sign in with your email and password"},
		{"login-03", "login", "Welcome back", "Sign in to continue to your account"},
		{"login-04", "login", "Sign in", "Enter your email and password to sign in"},
		{"login-05", "login", "Sign in", "Use your account to continue"},
		{"signup-01", "signup", "Create an account", "Enter your information to get started"},
		{"signup-02", "signup", "Create an account", "Sign up with email or a provider"},
		{"signup-03", "signup", "Create an account", "Start building with your new workspace"},
	} {
		if block.name == "login-01" {
			blocks[normalizeName(block.name)] = login01Block()
			continue
		}
		blocks[normalizeName(block.name)] = authBlock(block.name, block.route, block.title, block.description)
	}
	return blocks
}

func blockPageFile(block string, target string, content string) blockRegistryFile {
	return blockRegistryFile{
		RelPath: filepath.Join("blocks", block, filepath.Base(target)),
		Target:  target,
		Type:    "registry:page",
		Content: content,
	}
}

func blockComponentFile(block string, target string, content string) blockRegistryFile {
	return blockRegistryFile{
		RelPath: filepath.Join("blocks", block, target),
		Target:  target,
		Type:    "registry:component",
		Content: content,
	}
}

func sidebarBlock(name string, variant string) blockRegistryItem {
	return blockRegistryItem{
		Name:         name,
		Dependencies: []string{"breadcrumb", "button", "card", "separator", "sidebar"},
		Runtime:      []string{"assets/runtime.js"},
		Files: []blockRegistryFile{
			blockPageFile(name, "app/dashboard/page.templ", sidebarPageTempl(variant)),
			blockComponentFile(name, "components/app-sidebar.templ", appSidebarTempl()),
			blockComponentFile(name, "components/team-switcher.templ", teamSwitcherTempl()),
		},
	}
}

func authBlock(name string, route string, title string, description string) blockRegistryItem {
	return blockRegistryItem{
		Name:         name,
		Dependencies: []string{"button", "card", "checkbox", "input", "label"},
		Runtime:      []string{"assets/runtime.js"},
		Files: []blockRegistryFile{
			blockPageFile(name, filepath.Join("app", route, "page.templ"), authPageTempl(route)),
			blockComponentFile(name, filepath.Join("components", route+"-form.templ"), authFormTempl(route, title, description)),
		},
	}
}

func login01Block() blockRegistryItem {
	return blockRegistryItem{
		Name:         "login-01",
		Dependencies: []string{"button", "card", "field", "input"},
		Runtime:      []string{"assets/runtime.js"},
		Files: []blockRegistryFile{
			blockPageFile("login-01", "app/login/page.templ", login01PageTempl()),
			blockComponentFile("login-01", "components/login-form.templ", login01FormTempl()),
		},
	}
}

func targetPathForBlockFile(file blockRegistryFile) string {
	if file.Target != "" {
		return file.Target
	}
	return file.RelPath
}

func dashboardPageTempl() string {
	return `package dashboard

import "{{module}}/components"

templ Page() {
	<div class="flex min-h-svh bg-background">
		@components.AppSidebar()
		<main class="flex-1 space-y-6 p-6">
			@components.SiteHeader("Dashboard")
			@components.SectionCards()
			@components.ChartAreaInteractive()
			@components.DataTable()
		</main>
	</div>
}
`
}

func login01PageTempl() string {
	return `package login

import "{{module}}/components"

templ Page() {
	<div class="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
		<div class="w-full max-w-sm">
			@components.LoginForm()
		</div>
	</div>
}
`
}

func login01FormTempl() string {
	return `package components

import (
	ui "{{module}}/ui"
	t "github.com/a-h/templ"
)

templ LoginForm() {
	<div class="flex flex-col gap-6">
		@ui.Card(ui.DOMProps{}) {
			@ui.CardHeader(ui.DOMProps{}) {
				@ui.CardTitle(ui.DOMProps{}) { <span>Login to your account</span> }
				@ui.CardDescription(ui.DOMProps{}) { <span>Enter your email below to login to your account</span> }
			}
			@ui.CardContent(ui.DOMProps{}) {
				<form>
					@ui.FieldGroup(ui.DOMProps{}) {
						@ui.Field(ui.DOMProps{}) {
							@ui.FieldLabel(ui.DOMProps{Attrs: t.Attributes{"for": "email"}}) { <span>Email</span> }
							@ui.Input(ui.InputProps{DOMProps: ui.DOMProps{ID: "email", Attrs: t.Attributes{"required": true}}, Type: "email", Placeholder: "m@example.com"})
						}
						@ui.Field(ui.DOMProps{}) {
							<div class="flex items-center">
								@ui.FieldLabel(ui.DOMProps{Attrs: t.Attributes{"for": "password"}}) { <span>Password</span> }
								<a href="#" class="ml-auto inline-block text-sm underline-offset-4 hover:underline">Forgot your password?</a>
							</div>
							@ui.Input(ui.InputProps{DOMProps: ui.DOMProps{ID: "password", Attrs: t.Attributes{"required": true}}, Type: "password"})
						}
						@ui.Field(ui.DOMProps{}) {
							@ui.Button(ui.ButtonProps{Label: "Login", Type: "submit"})
							@ui.Button(ui.ButtonProps{Label: "Login with Google", Variant: ui.ButtonVariantOutline})
							@ui.FieldDescription(ui.DOMProps{Class: "text-center"}) {
								<span>Don't have an account? </span><a href="#">Sign up</a>
							}
						}
					}
				</form>
			}
		}
	</div>
}
`
}

func sidebarPageTempl(variant string) string {
	return `package dashboard

import "{{module}}/components"

templ Page() {
	<div class="flex min-h-svh bg-background">
		@components.AppSidebar()
		<main class="flex-1 space-y-6 p-6">
			@components.SiteHeader("Dashboard")
			<div class="grid gap-4 md:grid-cols-2">
				@components.SummaryCard("Navigation", "Sidebar ` + variant + ` block copied into your Go/templ app.")
				@components.SummaryCard("Content", "Replace this region with your server-rendered page content.")
			</div>
		</main>
	</div>
}
`
}

func authPageTempl(route string) string {
	component := "LoginForm"
	if route == "signup" {
		component = "SignupForm"
	}
	return `package ` + route + `

import "{{module}}/components"

templ Page() {
	<div class="flex min-h-svh w-full items-center justify-center bg-background p-6 md:p-10">
		<div class="w-full max-w-sm">
			@components.` + component + `()
		</div>
	</div>
}
`
}

func authFormTempl(route string, title string, description string) string {
	name := "LoginForm"
	action := "Sign in"
	if route == "signup" {
		name = "SignupForm"
		action = "Create account"
	}
	return `package components

import ui "{{module}}/ui"

templ ` + name + `() {
	@ui.Card(ui.DOMProps{}) {
		@ui.CardHeader(ui.DOMProps{Class: "space-y-1 text-center"}) {
			@ui.CardTitle(ui.DOMProps{}) { <span>` + title + `</span> }
			@ui.CardDescription(ui.DOMProps{}) { <span>` + description + `</span> }
		}
		@ui.CardContent(ui.DOMProps{Class: "space-y-4"}) {
			<div class="space-y-2">
				@ui.Label(ui.LabelProps{For: "` + route + `-email"}) { <span>Email</span> }
				@ui.Input(ui.InputProps{DOMProps: ui.DOMProps{ID: "` + route + `-email"}, Type: "email", Placeholder: "name@example.com"})
			</div>
			<div class="space-y-2">
				@ui.Label(ui.LabelProps{For: "` + route + `-password"}) { <span>Password</span> }
				@ui.Input(ui.InputProps{DOMProps: ui.DOMProps{ID: "` + route + `-password"}, Type: "password"})
			</div>
			@ui.Button(ui.ButtonProps{Label: "` + action + `", DOMProps: ui.DOMProps{Class: "w-full"}})
		}
	}
}
`
}

func appSidebarTempl() string {
	return `package components

import ui "{{module}}/ui"

templ AppSidebar() {
	@ui.SidebarProvider(ui.SidebarProviderProps{DefaultOpen: true}) {
		@ui.Sidebar(ui.SidebarProps{Side: "left", Variant: "sidebar", Collapsible: "icon"}) {
			@ui.SidebarHeader(ui.DOMProps{Class: "p-4"}) {
				<div class="flex items-center gap-2 font-semibold">
					<span class="flex size-7 items-center justify-center rounded-md bg-primary text-xs text-primary-foreground">A</span>
					<span>App</span>
				</div>
			}
			@ui.SidebarContent(ui.DOMProps{}) {
				@ui.SidebarGroup(ui.DOMProps{}) {
					@ui.SidebarGroupLabel(ui.DOMProps{}) { <span>Platform</span> }
					@ui.SidebarMenu(ui.DOMProps{}) {
						@ui.SidebarMenuItem(ui.DOMProps{}) { @ui.SidebarMenuButton(ui.SidebarMenuButtonProps{IsActive: true}) { <span>Dashboard</span> } }
						@ui.SidebarMenuItem(ui.DOMProps{}) { @ui.SidebarMenuButton(ui.SidebarMenuButtonProps{}) { <span>Projects</span> } }
						@ui.SidebarMenuItem(ui.DOMProps{}) { @ui.SidebarMenuButton(ui.SidebarMenuButtonProps{}) { <span>Settings</span> } }
					}
				}
			}
		}
	}
}
`
}

func siteHeaderTempl() string {
	return `package components

import ui "{{module}}/ui"

templ SiteHeader(title string) {
	<header class="flex items-center justify-between border-b border-border/70 pb-4">
		<h1 class="text-2xl font-semibold tracking-tight">{ title }</h1>
		@ui.Button(ui.ButtonProps{Label: "New", Size: ui.ButtonSizeSM})
	</header>
}

templ SummaryCard(title string, description string) {
	@ui.Card(ui.DOMProps{}) {
		@ui.CardHeader(ui.DOMProps{}) {
			@ui.CardTitle(ui.DOMProps{}) { <span>{ title }</span> }
			@ui.CardDescription(ui.DOMProps{}) { <span>{ description }</span> }
		}
	}
}
`
}

func sectionCardsTempl() string {
	return `package components

import ui "{{module}}/ui"

templ SectionCards() {
	<div class="grid gap-4 md:grid-cols-3">
		@SummaryCard("Total Revenue", "+20.1% from last month")
		@SummaryCard("Subscriptions", "+180.1% from last month")
		@SummaryCard("Sales", "+19% from last month")
	</div>
}
`
}

func chartAreaInteractiveTempl() string {
	return `package components

import ui "{{module}}/ui"

templ ChartAreaInteractive() {
	@ui.Card(ui.DOMProps{}) {
		@ui.CardHeader(ui.DOMProps{}) {
			@ui.CardTitle(ui.DOMProps{}) { <span>Area Chart</span> }
			@ui.CardDescription(ui.DOMProps{}) { <span>Server-rendered chart shell ready for your data.</span> }
		}
		@ui.CardContent(ui.DOMProps{}) {
			<div class="flex h-64 items-center justify-center rounded-md border border-dashed text-sm text-muted-foreground">Chart data goes here</div>
		}
	}
}
`
}

func dataTableBlockTempl() string {
	return `package components

import ui "{{module}}/ui"

templ DataTable() {
	@ui.Card(ui.DOMProps{}) {
		@ui.CardHeader(ui.DOMProps{}) {
			@ui.CardTitle(ui.DOMProps{}) { <span>Recent Activity</span> }
			@ui.CardDescription(ui.DOMProps{}) { <span>Latest transactions and updates.</span> }
		}
		@ui.CardContent(ui.DOMProps{}) {
			@ui.Table(ui.DOMProps{}) {
				@ui.TableHeader(ui.DOMProps{}) {
					@ui.TableRow(ui.DOMProps{}) {
						@ui.TableHead(ui.DOMProps{}) { <span>Invoice</span> }
						@ui.TableHead(ui.DOMProps{}) { <span>Status</span> }
						@ui.TableHead(ui.DOMProps{}) { <span>Amount</span> }
					}
				}
				@ui.TableBody(ui.DOMProps{}) {
					@ui.TableRow(ui.DOMProps{}) {
						@ui.TableCell(ui.DOMProps{}) { <span>INV001</span> }
						@ui.TableCell(ui.DOMProps{}) { @ui.Badge(ui.BadgeProps{Label: "Paid", Variant: ui.BadgeVariantSecondary}) }
						@ui.TableCell(ui.DOMProps{}) { <span>$250.00</span> }
					}
				}
			}
		}
	}
}
`
}

func teamSwitcherTempl() string {
	return `package components

import ui "{{module}}/ui"

templ TeamSwitcher() {
	@ui.Button(ui.ButtonProps{Label: "Acme Inc.", Variant: ui.ButtonVariantOutline, Size: ui.ButtonSizeSM})
}
`
}
