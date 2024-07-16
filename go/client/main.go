package client

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/kubeden/openssd/types"

	"github.com/gorilla/mux"
)

var config types.Config

// StartServer initializes and starts the client server
func StartServer(cfg types.Config) error {
	config = cfg

	r := mux.NewRouter()

	r.HandleFunc("/", handleIndex)
	r.HandleFunc("/blog", handleBlog)
	r.HandleFunc("/blog/{slug}", handleArticle)
	r.HandleFunc("/info", handleInfo)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Client server starting on 0.0.0.0:8080")
	return http.ListenAndServe("0.0.0.0:8080", r)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {
	layoutFile := filepath.Join("templates", config.TemplateChoice, "layout.html")
	pageFile := filepath.Join("templates", config.TemplateChoice, tmpl+".html")

	t, err := template.ParseFiles(layoutFile, pageFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Add XUserFullName to the data map
	data["XUserFullName"] = config.XUserFullName
	data["XUserName"] = config.XUserName

	err = t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling index request")
	renderTemplate(w, "index", map[string]interface{}{
		"Title": "Welcome",
	})
}

func handleBlog(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling blog request")
	renderTemplate(w, "blog", map[string]interface{}{
		"Title": "Blog Posts",
	})
}

func handleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]
	log.Printf("Handling article request for slug: %s", slug)
	renderTemplate(w, "article", map[string]interface{}{
		"Title": "Article: " + slug,
		"Slug":  slug,
	})
}

func handleInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling info request")
	renderTemplate(w, "info", map[string]interface{}{
		"Title": "About This Blog",
	})
}
