package main

import (
	"log"
	"net/http"

	"github.com/kubeden/openssd/templates"

	"github.com/gorilla/mux"
)

func StartServer() error {
	r := mux.NewRouter()

	r.HandleFunc("/", handleIndex)
	r.HandleFunc("/blog", handleBlog)
	r.HandleFunc("/blog/{slug}", handleArticle)
	r.HandleFunc("/info", handleInfo)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	return http.ListenAndServe(":8080", r)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handleIndex)
	r.HandleFunc("/blog", handleBlog)
	r.HandleFunc("/blog/{slug}", handleArticle)
	r.HandleFunc("/info", handleInfo)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Client server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling index request")
	templates.Index().Render(r.Context(), w)
}

func handleBlog(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling blog request")
	// TODO: Implement Blog template
	w.Write([]byte("Blog page"))
}

func handleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]
	log.Printf("Handling article request for slug: %s", slug)
	// TODO: Implement Article template
	w.Write([]byte("Article page for " + slug))
}

func handleInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling info request")
	// TODO: Implement Info template
	w.Write([]byte("Info page"))
}
