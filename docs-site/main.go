package main

import (
	"encoding/json"
	"net/http"
	"os"

	"shadcn/docs-site/views"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	router := mux.NewRouter()

	router.HandleFunc("/", HandleHome).Methods("GET")
	router.HandleFunc("/docs", HandleDocs).Methods("GET")
	router.HandleFunc("/docs/installation", HandleInstallation).Methods("GET")
	router.HandleFunc("/docs/components-json", HandleComponentsJSON).Methods("GET")
	router.HandleFunc("/docs/package-imports", HandlePackageImports).Methods("GET")
	router.HandleFunc("/docs/theming", HandleTheming).Methods("GET")
	router.HandleFunc("/docs/dark-mode", HandleDarkMode).Methods("GET")
	router.HandleFunc("/docs/cli", HandleCLI).Methods("GET")
	router.HandleFunc("/docs/monorepo", HandleMonorepo).Methods("GET")
	router.HandleFunc("/docs/rtl", HandleRTL).Methods("GET")
	router.HandleFunc("/docs/javascript", HandleJavaScript).Methods("GET")
	router.HandleFunc("/docs/llms", HandleLLMS).Methods("GET")
	router.HandleFunc("/docs/forms", HandleForms).Methods("GET")
	router.HandleFunc("/docs/changelog", HandleChangelog).Methods("GET")
	router.HandleFunc("/docs/directory", HandleDirectoryDocs).Methods("GET")
	router.HandleFunc("/docs/components", HandleComponentsIndex).Methods("GET")
	router.HandleFunc("/docs/components/{slug}", HandleComponentDoc).Methods("GET")
	router.HandleFunc("/docs/components/{variant}/{slug}", HandleComponentDoc).Methods("GET")
	router.HandleFunc("/blocks", HandleBlocksIndex).Methods("GET")
	router.HandleFunc("/blocks/{slug}", HandleBlockPage).Methods("GET")
	router.HandleFunc("/blocks/{slug}/preview", HandleBlockPreview).Methods("GET")
	router.HandleFunc("/blocks/dashboard-01/preview", HandleDashboard01Preview).Methods("GET")
	router.HandleFunc("/blocks/sidebar-07/preview", HandleSidebar07Preview).Methods("GET")
	router.HandleFunc("/blocks/sidebar-03/preview", HandleSidebar03Preview).Methods("GET")
	router.HandleFunc("/blocks/login-01/preview", HandleLogin01Preview).Methods("GET")
	router.HandleFunc("/blocks/login-03/preview", HandleLogin03Preview).Methods("GET")
	router.HandleFunc("/blocks/login-04/preview", HandleLogin04Preview).Methods("GET")
	router.HandleFunc("/blocks/signup-01/preview", HandleSignup01Preview).Methods("GET")
	router.HandleFunc("/blocks/signup-02/preview", HandleSignup02Preview).Methods("GET")
	router.HandleFunc("/charts", HandleChartsRedirect).Methods("GET")
	router.HandleFunc("/charts/area", HandleAreaChart).Methods("GET")
	router.HandleFunc("/charts/bar", HandleBarChart).Methods("GET")
	router.HandleFunc("/charts/line", HandleLineChart).Methods("GET")
	router.HandleFunc("/charts/pie", HandlePieChart).Methods("GET")
	router.HandleFunc("/charts/radial", HandleRadialChart).Methods("GET")
	router.HandleFunc("/directory", HandleDirectoryRedirect).Methods("GET")
	router.HandleFunc("/create", HandleCreate).Methods("GET")
	router.HandleFunc("/llms.txt", HandleLLMSText).Methods("GET")
	router.HandleFunc("/search-index.json", HandleSearchIndex).Methods("GET")
	router.NotFoundHandler = http.HandlerFunc(HandleNotFound)

	router.PathPrefix("/public/").Handler(
		http.StripPrefix("/public/", http.FileServer(http.Dir("public"))),
	)

	log.Info().Msgf("Listening on port %v\n", port)
	http.ListenAndServe(port, router)
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	views.Index().Render(r.Context(), w)
}

func HandleDocs(w http.ResponseWriter, r *http.Request) {
	views.DocsIndexPage().Render(r.Context(), w)
}

func HandleInstallation(w http.ResponseWriter, r *http.Request) {
	views.InstallationPage().Render(r.Context(), w)
}

func HandleComponentsJSON(w http.ResponseWriter, r *http.Request) {
	views.ComponentsJSONPage().Render(r.Context(), w)
}

func HandlePackageImports(w http.ResponseWriter, r *http.Request) {
	views.PackageImportsPage().Render(r.Context(), w)
}

func HandleTheming(w http.ResponseWriter, r *http.Request) {
	views.ThemingPage().Render(r.Context(), w)
}

func HandleDarkMode(w http.ResponseWriter, r *http.Request) {
	views.DarkModePage().Render(r.Context(), w)
}

func HandleCLI(w http.ResponseWriter, r *http.Request) {
	views.CLIPage().Render(r.Context(), w)
}

func HandleMonorepo(w http.ResponseWriter, r *http.Request) {
	views.MonorepoPage().Render(r.Context(), w)
}

func HandleRTL(w http.ResponseWriter, r *http.Request) {
	views.RTLPage().Render(r.Context(), w)
}

func HandleJavaScript(w http.ResponseWriter, r *http.Request) {
	views.JavaScriptPage().Render(r.Context(), w)
}

func HandleLLMS(w http.ResponseWriter, r *http.Request) {
	views.LLMSPage().Render(r.Context(), w)
}

func HandleLLMSText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(views.LLMSText()))
}

func HandleForms(w http.ResponseWriter, r *http.Request) {
	views.FormsPage().Render(r.Context(), w)
}

func HandleChangelog(w http.ResponseWriter, r *http.Request) {
	views.ChangelogPage().Render(r.Context(), w)
}

func HandleDirectoryDocs(w http.ResponseWriter, r *http.Request) {
	views.DirectoryDocsPage().Render(r.Context(), w)
}

func HandleDirectoryRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/docs/directory", http.StatusMovedPermanently)
}

func HandleComponentsIndex(w http.ResponseWriter, r *http.Request) {
	views.ComponentsIndexPage().Render(r.Context(), w)
}

func HandleComponentDoc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]
	if slug == "" {
		http.NotFound(w, r)
		return
	}

	doc, ok := views.FindComponentDoc(slug)
	if !ok {
		http.NotFound(w, r)
		return
	}

	prev, hasPrev := views.ComponentDocBefore(slug)
	next, hasNext := views.ComponentDocAfter(slug)
	views.ComponentDetailPage(doc, prev, hasPrev, next, hasNext).Render(r.Context(), w)
}

func HandleBlocksIndex(w http.ResponseWriter, r *http.Request) {
	views.BlocksIndexPage().Render(r.Context(), w)
}

func HandleBlockPage(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	if slug == "" {
		http.NotFound(w, r)
		return
	}

	if block, ok := views.FindBlockDoc(slug); ok {
		views.BlockDetailPage(block).Render(r.Context(), w)
		return
	}

	categoryBlocks := views.BlocksByCategory(slug)
	if len(categoryBlocks) == 0 {
		http.NotFound(w, r)
		return
	}

	views.BlockCategoryPage(slug).Render(r.Context(), w)
}

func HandleChartsRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/charts/area", http.StatusMovedPermanently)
}

func HandleCharts(w http.ResponseWriter, r *http.Request) {
	views.ChartsPage().Render(r.Context(), w)
}

func HandleAreaChart(w http.ResponseWriter, r *http.Request) {
	views.AreaChartPage().Render(r.Context(), w)
}

func HandleBarChart(w http.ResponseWriter, r *http.Request) {
	views.BarChartPage().Render(r.Context(), w)
}

func HandleLineChart(w http.ResponseWriter, r *http.Request) {
	views.LineChartPage().Render(r.Context(), w)
}

func HandlePieChart(w http.ResponseWriter, r *http.Request) {
	views.PieChartPage().Render(r.Context(), w)
}

func HandleRadialChart(w http.ResponseWriter, r *http.Request) {
	views.RadialChartPage().Render(r.Context(), w)
}

func HandleDirectory(w http.ResponseWriter, r *http.Request) {
	views.DirectoryPage().Render(r.Context(), w)
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	views.CreatePage().Render(r.Context(), w)
}

func HandleDashboard01Preview(w http.ResponseWriter, r *http.Request) {
	views.Dashboard01Preview().Render(r.Context(), w)
}

func HandleSidebar07Preview(w http.ResponseWriter, r *http.Request) {
	views.Sidebar07Preview().Render(r.Context(), w)
}

func HandleSidebar03Preview(w http.ResponseWriter, r *http.Request) {
	views.Sidebar03Preview().Render(r.Context(), w)
}

func HandleLogin01Preview(w http.ResponseWriter, r *http.Request) {
	views.Login01Preview().Render(r.Context(), w)
}

func HandleLogin03Preview(w http.ResponseWriter, r *http.Request) {
	views.Login03Preview().Render(r.Context(), w)
}

func HandleLogin04Preview(w http.ResponseWriter, r *http.Request) {
	views.Login04Preview().Render(r.Context(), w)
}

func HandleSignup01Preview(w http.ResponseWriter, r *http.Request) {
	views.Signup01Preview().Render(r.Context(), w)
}

func HandleSignup02Preview(w http.ResponseWriter, r *http.Request) {
	views.Signup02Preview().Render(r.Context(), w)
}

func HandleSearchIndex(w http.ResponseWriter, r *http.Request) {
	type entry struct {
		Title string `json:"title"`
		URL   string `json:"url"`
		Type  string `json:"type"`
	}

	index := make([]entry, 0, len(views.ComponentIndex())+len(views.Blocks())+16)
	for _, doc := range views.ComponentIndex() {
		index = append(index, entry{
			Title: doc.Title,
			URL:   "/docs/components/" + doc.Slug,
			Type:  "component",
		})
	}
	for _, block := range views.Blocks() {
		index = append(index, entry{
			Title: block.Title,
			URL:   "/blocks/" + block.Slug,
			Type:  "block",
		})
	}
	docs := []entry{
		{Title: "Introduction", URL: "/docs", Type: "doc"},
		{Title: "Installation", URL: "/docs/installation", Type: "doc"},
		{Title: "components.json", URL: "/docs/components-json", Type: "doc"},
		{Title: "Package Imports", URL: "/docs/package-imports", Type: "doc"},
		{Title: "Theming", URL: "/docs/theming", Type: "doc"},
		{Title: "Dark Mode", URL: "/docs/dark-mode", Type: "doc"},
		{Title: "CLI", URL: "/docs/cli", Type: "doc"},
		{Title: "Monorepo", URL: "/docs/monorepo", Type: "doc"},
		{Title: "RTL", URL: "/docs/rtl", Type: "doc"},
		{Title: "JavaScript", URL: "/docs/javascript", Type: "doc"},
		{Title: "llms.txt", URL: "/docs/llms", Type: "doc"},
		{Title: "Forms", URL: "/docs/forms", Type: "doc"},
		{Title: "Changelog", URL: "/docs/changelog", Type: "doc"},
		{Title: "Directory", URL: "/docs/directory", Type: "doc"},
		{Title: "Create", URL: "/create", Type: "doc"},
		{Title: "Components", URL: "/docs/components", Type: "doc"},
		{Title: "Blocks", URL: "/blocks", Type: "doc"},
		{Title: "Charts", URL: "/charts/area", Type: "doc"},
	}
	index = append(index, docs...)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(index); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HandleBlockPreview(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	if slug == "" {
		http.NotFound(w, r)
		return
	}

	if preview, ok := views.BlockPreviewForSlug(slug); ok {
		preview.Render(r.Context(), w)
		return
	}

	http.NotFound(w, r)
}

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	views.NotFoundPage().Render(r.Context(), w)
}
