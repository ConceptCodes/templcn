module templcn/docs-site

go 1.25.0

require (
	github.com/a-h/templ v0.3.1020
	github.com/gorilla/mux v1.8.1
	github.com/rs/zerolog v1.32.0
	templcn/ui v0.0.0
)

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.41.0 // indirect
)

replace templcn/ui => ../ui
