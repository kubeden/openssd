package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday/v2"
)

type Article struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Slug    string    `json:"slug"`
	Date    time.Time `json:"date"`
	Excerpt string    `json:"excerpt"`
}

var articles []Article

func StartServer() error {
	r := mux.NewRouter()
	r.HandleFunc("/api/articles", getArticles).Methods("GET")
	r.HandleFunc("/api/articles/{slug}", getArticle).Methods("GET")

	return http.ListenAndServe(":8081", r)
}

func main() {
	loadArticles()

	r := mux.NewRouter()
	r.HandleFunc("/api/articles", getArticles).Methods("GET")
	r.HandleFunc("/api/articles/{slug}", getArticle).Methods("GET")

	log.Println("API server starting on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func loadArticles() {
	files, err := ioutil.ReadDir("./articles")
	if err != nil {
		log.Fatalf("Error reading articles directory: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			article, err := parseArticle(file.Name())
			if err != nil {
				log.Printf("Error parsing article %s: %v", file.Name(), err)
				continue
			}
			articles = append(articles, article)
			log.Printf("Loaded article: %s", article.Title)
		}
	}

	sort.Slice(articles, func(i, j int) bool {
		return articles[i].Date.After(articles[j].Date)
	})

	log.Printf("Loaded %d articles", len(articles))
}

func parseArticle(filename string) (Article, error) {
	content, err := ioutil.ReadFile(filepath.Join("./articles", filename))
	if err != nil {
		return Article{}, fmt.Errorf("error reading file %s: %v", filename, err)
	}

	var article Article
	rest, err := frontmatter.Parse(strings.NewReader(string(content)), &article)
	if err != nil {
		return Article{}, fmt.Errorf("error parsing frontmatter for %s: %v", filename, err)
	}

	article.Content = string(blackfriday.Run(rest))

	if article.Slug == "" {
		article.Slug = strings.TrimSuffix(filename, filepath.Ext(filename))
	}

	return article, nil
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, article := range articles {
		if article.Slug == slug {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(article)
			return
		}
	}

	http.NotFound(w, r)
}
