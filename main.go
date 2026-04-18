package main

import (
	"net/http"
	"os"

	"shadcn/views"

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
	router.HandleFunc("/docs/components", HandleComponentsIndex).Methods("GET")
	router.HandleFunc("/docs/components/{slug}", HandleComponentDoc).Methods("GET")
	router.HandleFunc("/docs/components/{variant}/{slug}", HandleComponentDoc).Methods("GET")
	router.HandleFunc("/blocks", HandleBlocksIndex).Methods("GET")
	router.HandleFunc("/blocks/{slug}", HandleBlockPage).Methods("GET")
	router.HandleFunc("/charts", HandleCharts).Methods("GET")
	router.HandleFunc("/directory", HandleDirectory).Methods("GET")
	router.HandleFunc("/create", HandleCreate).Methods("GET")

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

func HandleCharts(w http.ResponseWriter, r *http.Request) {
	views.ChartsPage().Render(r.Context(), w)
}

func HandleDirectory(w http.ResponseWriter, r *http.Request) {
	views.DirectoryPage().Render(r.Context(), w)
}

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	views.CreatePage().Render(r.Context(), w)
}
