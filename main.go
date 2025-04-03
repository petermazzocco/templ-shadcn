package main

import (
	"fmt"
	"net/http"
	"templ-shadcn/internal/component"
	"templ-shadcn/internal/generate"
	"templ-shadcn/internal/middleware"
	"templ-shadcn/internal/template"
	"templ-shadcn/internal/view"

	"github.com/joho/godotenv"
)

func main() {

	err := generate.GenerateMain()
	if err != nil {
		panic(err)
	}

	_ = godotenv.Load()
	mux := http.NewServeMux()

	// Load static assets
	mux.HandleFunc("GET /favicon.ico", view.ServeFavicon)
	mux.HandleFunc("GET /static/", view.ServeStaticFiles)

	mux.HandleFunc("POST /hello-name", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		if name == "" {
			name = "Anonymous"
		}
		middleware.RenderComponent(w, r, component.HelloNameResponse(name))
	})

	// Landing page route
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			middleware.Chain(w, r, template.NotFound())
			return
		}
		middleware.Chain(w, r, template.Home())
	})

	// All components route
	mux.HandleFunc("GET /components", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/components" {
			middleware.Chain(w, r, template.NotFound())
			return
		}
		middleware.Chain(w, r, template.Components())
	})

	// Individual components route
	mux.HandleFunc("GET /components/{name}", func(w http.ResponseWriter, r *http.Request) {
		page := r.PathValue("name")
		fmt.Println(page)
		middleware.Chain(w, r, template.Component(page))
	})

	fmt.Printf("server is running on port 8080")
	err = http.ListenAndServe(":"+"8080", mux)
	if err != nil {
		fmt.Println(err)
	}

}
