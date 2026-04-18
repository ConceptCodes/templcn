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

	router.PathPrefix("/public/").Handler(
		http.StripPrefix("/public/", http.FileServer(http.Dir("public"))),
	)

	log.Info().Msgf("Listening on port %v\n", port)
	http.ListenAndServe(port, router)
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	views.Index().Render(r.Context(), w)
}
